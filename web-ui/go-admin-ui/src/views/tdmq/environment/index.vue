<template>
  <BasicLayout>
    <template #wrapper>
      <el-card class="box-card">
        <el-form ref="queryForm" :model="queryParams" :inline="true">
          <el-form-item label="集群ID" prop="status">
            <el-select
              v-model="queryParams.status"
              placeholder="集群ID"
              clearable
              size="small"
              style="width: 260px"
              @change="selectChange"
            >
              <el-option
                v-for="dict in statusOptions"
                :key="dict.value"
                :label="dict.label"
                :value="dict.value"
              />
            </el-select>
          </el-form-item>
        </el-form>

        <el-table
          v-loading="loading"
          :data="roleList"
          border
          @selection-change="handleSelectionChange"
          @sort-change="handleSortChang"
        >
          <el-table-column type="selection" width="55" align="center" />
          <el-table-column
            label="命名空间名称"
            sortable="custom"
            prop="EnvironmentId"
            width="200"
          />
          <el-table-column
            label="说明"
            sortable="custom"
            prop="Remark"
            :show-overflow-tooltip="true"
          />
          <el-table-column
            label="命名空间ID"
            sortable="custom"
            prop="NamespaceId"
            :show-overflow-tooltip="true"
          />

          <el-table-column
            label="命名空间名称"
            sortable="custom"
            prop="NamespaceName"
            :show-overflow-tooltip="true"
          />
          <el-table-column
            label="Topic数量"
            sortable="custom"
            prop="TopicNum"
            :show-overflow-tooltip="true"
          />
          <el-table-column
            label="创建时间"
            sortable="custom"
            prop="CreateTime"
            width="160"
          >
            <template slot-scope="scope">
              <span>{{ parseTime(scope.row.CreateTime) }}</span>
            </template>
          </el-table-column>
          <el-table-column
            label="操作"
            align="left"
            class-name="small-padding fixed-width"
            width="80"
          >
            <template slot-scope="scope">
              <el-button
                v-if="scope.row.roleKey !== 'admin'"
                v-permisaction="['admin:sysRole:remove']"
                size="mini"
                type="text"
                icon="el-icon-delete"
                @click="handleDelete(scope.row)"
                >删除</el-button
              >
            </template>
          </el-table-column>
        </el-table>

        <pagination
          v-show="total > 0"
          :total="total"
          :page.sync="queryParams.pageIndex"
          :limit.sync="queryParams.pageSize"
          @pagination="getList"
        />
      </el-card>
    </template>
  </BasicLayout>
</template>

<script>
import {
  listRole,
  getRole,
  delRole,
  addRole,
  updateRole,
  dataScope,
  changeRoleStatus,
} from "@/api/admin/sys-role";
import { listEnvironments } from "@/api/tdmq/environments";
import { describeSelectClusters } from "@/api/tdmq/clusters";
import {
  treeselect as deptTreeselect,
  roleDeptTreeselect,
} from "@/api/admin/sys-dept";
import { formatJson } from "@/utils";

export default {
  name: "Environment",
  components: {},
  data() {
    return {
      // 遮罩层
      loading: true,
      // 选中数组
      ids: [],
      // 非单个禁用
      single: true,
      // 非多个禁用
      multiple: true,
      // 总条数
      total: 0,
      // 角色表格数据
      roleList: [],
      menuIdsChecked: [],
      // 弹出层标题
      title: "",
      // 是否显示弹出层
      open: false,
      // 是否显示弹出层（数据权限）
      openDataScope: false,
      isEdit: false,
      // 日期范围
      dateRange: [],
      // 状态数据字典
      statusOptions: [],
      // 数据范围选项
      dataScopeOptions: [
        {
          value: "1",
          label: "全部数据权限",
        },
        {
          value: "2",
          label: "自定数据权限",
        },
        {
          value: "3",
          label: "本部门数据权限",
        },
        {
          value: "4",
          label: "本部门及以下数据权限",
        },
        {
          value: "5",
          label: "仅本人数据权限",
        },
      ],
      // 菜单列表
      menuOptions: [],
      menuList: [],
      // 部门列表
      deptOptions: [],
      menuOptionsAlert: "加载中，请稍后",
      // 查询参数
      queryParams: {
        pageIndex: 1,
        pageSize: 10,
        roleName: undefined,
        roleKey: undefined,
        status: undefined,
      },
      // 表单参数
      form: {
        sysMenu: [],
      },
      defaultProps: {
        children: "children",
        label: "label",
      },
      // 表单校验
      rules: {
        roleName: [
          { required: true, message: "角色名称不能为空", trigger: "blur" },
        ],
        roleKey: [
          { required: true, message: "权限字符不能为空", trigger: "blur" },
        ],
        roleSort: [
          { required: true, message: "角色顺序不能为空", trigger: "blur" },
        ],
      },
    };
  },
  created() {
    this.getList("pulsar-44or73egbj4g");
    this.describeSelectClusters();
  },
  methods: {
    /** 查询角色列表 */
    getList(clustersId) {
      this.loading = true;
      listEnvironments(clustersId).then((response) => {
        this.roleList = response.data.list;
        this.total = response.data.count;
        this.loading = false;
      });
    },
    describeSelectClusters() {
      describeSelectClusters().then((res) => {
        this.statusOptions = res.data;
      });
    },
    //selectChange
    selectChange(e) {
      this.getList(e);
    },

    // 取消按钮
    cancel() {
      this.open = false;
      this.reset();
    },

    // 表单重置
    reset() {
      this.menuOptions = this.menuList;
      if (this.$refs.menuTree !== undefined) {
        this.$refs.menuTree.setCheckedKeys([]);
      }
      this.form = {
        roleId: undefined,
        roleName: undefined,
        roleKey: undefined,
        roleSort: 0,
        status: "2",
        menuIds: [],
        deptIds: [],
        sysMenu: [],
        remark: undefined,
      };
      this.resetForm("form");
    },
    /** 搜索按钮操作 */
    handleQuery() {
      this.queryParams.pageIndex = 1;
      this.getList();
    },
    /** 重置按钮操作 */
    resetQuery() {
      this.dateRange = [];
      this.resetForm("queryForm");
      this.handleQuery();
    },
    // 多选框选中数据
    handleSelectionChange(selection) {
      this.ids = selection.map((item) => item.roleId);
      this.single = selection.length !== 1;
      this.multiple = !selection.length;
    },
    /** 新增按钮操作 */
    handleAdd() {
      this.reset();

      this.open = true;
      this.title = "添加角色";
      this.isEdit = false;
    },
    handleSortChang(column, prop, order) {
      prop = column.prop;
      order = column.order;
      if (order === "descending") {
        this.queryParams[prop + "Order"] = "desc";
      } else if (order === "ascending") {
        this.queryParams[prop + "Order"] = "asc";
      } else {
        this.queryParams[prop + "Order"] = undefined;
      }
      this.getList();
    },

    /** 提交按钮 */
    submitForm: function () {
      this.$refs["form"].validate((valid) => {
        if (valid) {
          if (this.form.roleId !== undefined) {
            this.form.menuIds = this.getMenuAllCheckedKeys();
            updateRole(this.form, this.form.roleId).then((response) => {
              if (response.code === 200) {
                this.msgSuccess(response.msg);
                this.open = false;
                this.getList();
              } else {
                this.msgError(response.msg);
              }
            });
          } else {
            this.form.menuIds = this.getMenuAllCheckedKeys();
            addRole(this.form).then((response) => {
              if (response.code === 200) {
                this.msgSuccess(response.msg);
                this.open = false;
                this.getList();
              } else {
                this.msgError(response.msg);
              }
            });
          }
        }
      });
    },

    /** 删除按钮操作 */
    handleDelete(row) {
      const roleIds = (row.roleId && [row.roleId]) || this.ids;
      this.$confirm(
        '是否确认删除角色编号为"' + roleIds + '"的数据项?',
        "警告",
        {
          confirmButtonText: "确定",
          cancelButtonText: "取消",
          type: "warning",
        }
      )
        .then(function () {
          return delRole({ ids: roleIds });
        })
        .then((response) => {
          this.getList();
          this.msgSuccess(response.msg);
        })
        .catch(function () {});
    },
  },
};
</script>
