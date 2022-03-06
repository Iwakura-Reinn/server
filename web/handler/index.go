// Copyright (c) 2022 Trim21 <trim21.me@gmail.com>
//
// SPDX-License-Identifier: AGPL-3.0-only
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published
// by the Free Software Foundation, version 3.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.
// See the GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program. If not, see <https://www.gnu.org/licenses/>

package handler

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"

	"github.com/bangumi/server/domain"
	"github.com/bangumi/server/internal/errgo"
	"github.com/bangumi/server/internal/strparse"
	"github.com/bangumi/server/web/handler/cachekey"
	"github.com/bangumi/server/web/res"
)

func (h Handler) getIndex(c *fiber.Ctx, id uint32) error {
	i, err := h.i.Get(c.Context(), id)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			return fiber.ErrNotFound
		}

		return errgo.Wrap(err, "Index.Get")
	}

	u, err := h.u.GetByID(c.Context(), i.CreatorID)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			h.log.Error("index missing creator",
				zap.Uint32("index_id", id), zap.Uint32("user_id", i.CreatorID))
			return fiber.ErrInternalServerError
		}

		return errgo.Wrap(err, "Index.Get")
	}

	return c.JSON(res.Index{
		CreatedAt: i.CreatedAt,
		Creator: res.Creator{
			Username: u.UserName,
			Nickname: u.NickName,
		},
		Title:       i.Title,
		Description: i.Description,
		Total:       i.Total,
		ID:          id,
		Stat: res.Stat{
			Comments: i.Comments,
			Collects: i.Collects,
		},
		Ban: i.Ban,
	})
}

func (h Handler) GetIndex(c *fiber.Ctx) error {
	v := h.getUser(c)

	id, err := strparse.Uint32(c.Params("id"))
	if err != nil || id == 0 {
		return fiber.NewError(http.StatusBadRequest, "bad index_id")
	}

	if !v.AllowNSFW() {
		nsfw, err := h.getIndexNsfwWithCache(c.Context(), id)
		if err != nil {
			return err
		}

		if nsfw {
			return fiber.ErrNotFound
		}
	}

	return h.getIndex(c, id)
}

func (h Handler) getIndexNsfwWithCache(ctx context.Context, id uint32) (bool, error) {
	cacheKey := cachekey.IndexNSFW(id)
	var nsfw bool
	cached, err := h.cache.Get(ctx, cacheKey, &nsfw)
	if err != nil {
		h.log.Error("can't read cache", zap.String("key", cacheKey), zap.Error(err))
		return false, errgo.Wrap(err, "Cache.Get")
	}

	if cached {
		return nsfw, nil
	}

	nsfw, err = h.i.IsNsfw(ctx, id)
	if err != nil {
		return false, errgo.Wrap(err, "Index.IsNsfw")
	}

	if err = h.cache.Set(ctx, cacheKey, nsfw, time.Hour); err != nil {
		h.log.Error("can't read cache", zap.String("key", cacheKey), zap.Error(err))
		return false, errgo.Wrap(err, "Cache.Get")
	}

	return nsfw, nil
}
