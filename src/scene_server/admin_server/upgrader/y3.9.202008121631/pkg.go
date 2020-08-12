/*
 * Tencent is pleased to support the open source community by making 蓝鲸 available.
 * Copyright (C) 2017-2018 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 */

package y3_9_202008121631

import (
	"context"

	"configcenter/src/common/blog"
	"configcenter/src/scene_server/admin_server/upgrader"
	"configcenter/src/storage/dal"
)

const currentPackageName = "y3.9.202008121631"

func init() {
	upgrader.RegistUpgrader(currentPackageName, upgrade)
}

// upgrade  不在前端界面展示指定模型,将其添加ishide=true字段
func upgrade(ctx context.Context, db dal.RDB, conf *upgrader.Config) (err error) {
	blog.Infof("start execute %s", currentPackageName)

	// 在cc_ObjDes表中将process和plat模型实例增加ishide=true字段
	if err := addModelFieldIshide(ctx, db, conf); err != nil {
		blog.Errorf("[%s] failed to model add field ishide, error: %s", currentPackageName, err.Error())
		return err
	}

	return nil
}
