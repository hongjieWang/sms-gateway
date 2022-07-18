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
          <el-form-item label="命名空间ID" prop="status">
            <el-select
              v-model="selectedEnvironments"
              placeholder="命名空间ID"
              clearable
              size="small"
              style="width: 260px"
              @change="selectChangeEnvironments"
            >
              <el-option
                v-for="dict in environmentsOptions"
                :key="dict.value"
                :label="dict.label"
                :value="dict.value"
              />
            </el-select>
          </el-form-item>
        </el-form>

        <el-row :gutter="10" class="mb8">
          <el-col :span="1.5">
            <el-button
              type="primary"
              icon="el-icon-plus"
              size="mini"
              @click="handleAdd"
              >新增</el-button
            >
          </el-col>
        </el-row>

        <el-table
          v-loading="loading"
          :data="roleList"
          border
          @selection-change="handleSelectionChange"
          @sort-change="handleSortChang"
        >
          <el-table-column type="selection" width="55" align="center" />
          <el-table-column
            label="队列名称"
            sortable="custom"
            prop="TopicName"
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
            prop="EnvironmentId"
            :show-overflow-tooltip="true"
          />
          <el-table-column
            label="队列类型"
            sortable="custom"
            prop="TopicType"
            :show-overflow-tooltip="true"
            width="80"
          />
          <el-table-column
            label="消费者数量"
            sortable="custom"
            prop="ConsumerCount"
            :show-overflow-tooltip="true"
          />
          <el-table-column
            label="生产者数量"
            sortable="custom"
            prop="ProducerCount"
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

        <!-- 添加或修改角色配置对话框 -->
        <el-dialog
          v-if="open"
          :title="title"
          :visible.sync="open"
          width="600px"
        >
          <el-form ref="form" :model="form" :rules="rules" label-width="80px">
            <el-form-item label="集群ID">
              <el-select
                v-model="form.clusterId"
                placeholder="请选择"
                @change="selectChange"
              >
                <el-option
                  v-for="item in statusOptions"
                  :key="item.value"
                  :label="item.label"
                  :value="item.value"
                  :disabled="item.status == 1"
                />
              </el-select>
            </el-form-item>
            <el-form-item label="命名空间">
              <el-select v-model="form.environmentId" placeholder="请选择">
                <el-option
                  v-for="item in environmentsOptions"
                  :key="item.value"
                  :label="item.label"
                  :value="item.value"
                  :disabled="item.status == 1"
                />
              </el-select>
            </el-form-item>

            <el-form-item label="Topic名称" prop="topicName">
              <el-input
                v-model="form.topicName"
                placeholder="请输入Topic名称"
                :disabled="isEdit"
              />
            </el-form-item>
            <el-form-item label="备注">
              <el-input
                v-model="form.remark"
                :disabled="isEdit"
                placeholder="请输入内容"
              />
            </el-form-item>
            <el-form-item label="类型">
              <el-select v-model="form.topicType" placeholder="请选择">
                <el-option
                  v-for="item in topicTypeOptions"
                  :key="item.id"
                  :label="item.value"
                  :value="item.id"
                />
              </el-select>
            </el-form-item>
            <el-form-item label="分区数量" prop="partitions">
              <el-input-number
                v-model="form.partitions"
                controls-position="right"
                :min="0"
              />
            </el-form-item>
          </el-form>
          <div slot="footer" class="dialog-footer">
            <el-button type="primary" @click="submitForm">确 定</el-button>
            <el-button @click="cancel">取 消</el-button>
          </div>
        </el-dialog>
      </el-card>
    </template>
  </BasicLayout>
</template>

<script>
import { selectEnvironments } from "@/api/tdmq/environments";
import { describeSelectClusters } from "@/api/tdmq/clusters";
import { listTopic, addTopic, delTopic } from "@/api/tdmq/topic";

export default {
  name: "Topic",
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
      //命名空间选择
      environmentsOptions: [],
      //topic类型
      topicTypeOptions: [
        { id: 0, value: "普通消息" },
        { id: 1, value: "全局顺序消息" },
        { id: 2, value: "局部顺序消息" },
        { id: 3, value: "重试队列" },
        { id: 4, value: "死信队列" },
      ],
      menuOptionsAlert: "加载中，请稍后",
      selectedClustersId: "",
      selectedEnvironments: "",
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
        clusterId: "",
        environmentId: "",
        topicName: "",
        remark: "",
        topicType: "",
        partitions: 1,
      },
      defaultProps: {
        children: "children",
        label: "label",
      },
      // 表单校验
      rules: {
        topicName: [
          { required: true, message: "队列名称不能为空", trigger: "blur" },
        ],
        clusterId: [
          { required: true, message: "集群ID不能为空", trigger: "blur" },
        ],
        environmentId: [
          { required: true, message: "命名空间不能为空", trigger: "blur" },
        ],
        topicType: [
          { required: true, message: "队列类型不能为空", trigger: "blur" },
        ],
      },
    };
  },
  created() {
    this.getList();
    this.selectEnvironments("pulsar-44or73egbj4g");
    this.describeSelectClusters();
  },
  methods: {
    /** 查询角色列表 */
    getList() {
      this.loading = true;
      listTopic({
        environmentId: this.selectedEnvironments,
        clusterId: this.selectedClustersId,
      }).then((response) => {
        this.loading = false;
        this.roleList = response.data.list;
        this.total = response.data.count;
      });
    },
    describeSelectClusters() {
      describeSelectClusters().then((res) => {
        this.statusOptions = res.data;
      });
    },
    selectEnvironments(clustersId) {
      selectEnvironments(clustersId).then((res) => {
        this.environmentsOptions = res.data;
      });
    },
    //selectChange集群ID
    selectChange(e) {
      this.selectedClustersId = e;
      this.selectEnvironments(e);
    },
    //change命名空间
    selectChangeEnvironments(e) {
      this.selectedEnvironments = e;
      this.getList();
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
        partitions: 0,
        remark: "",
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
      this.title = "添加Topic";
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
          addTopic(this.form).then((res) => {
            if (res.code === 200) {
              this.msgSuccess(res.msg);
              this.open = false;
              this.getList();
            } else {
              this.msgError(res.msg);
            }
          });
        }
      });
    },

    /** 删除按钮操作 */
    handleDelete(row) {
      row.ClustersId = this.selectedClustersId;
      this.$confirm(
        '是否确认删除Topic为"' + row.TopicName + '"的数据项?',
        "警告",
        {
          confirmButtonText: "确定",
          cancelButtonText: "取消",
          type: "warning",
        }
      )
        .then(function () {
          return delTopic({
            topicName: row.TopicName,
            environmentId: row.EnvironmentId,
            clusterId: row.ClustersId,
          });
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
