<template>
  <div class="dashboard-editor-container">
    <el-row :gutter="12">
      <el-col
        :sm="24"
        :xs="24"
        :md="6"
        :xl="6"
        :lg="6"
        :style="{ marginBottom: '12px' }"
      >
        <chart-card title="总发送条数" :total="count">
          <el-tooltip
            slot="action"
            class="item"
            effect="dark"
            content="指标说明"
            placement="top-start"
          >
            <i class="el-icon-warning-outline" />
          </el-tooltip>
          <div>
            <trend flag="top" style="margin-right: 16px" rate="12">
              <span slot="term">周同比</span>
            </trend>
            <trend flag="bottom" rate="11">
              <span slot="term">日同比</span>
            </trend>
          </div>
          <template slot="footer">日均短信条数: <span>10 条</span></template>
        </chart-card>
      </el-col>
      <el-col
        :sm="24"
        :xs="24"
        :md="6"
        :xl="6"
        :lg="6"
        :style="{ marginBottom: '12px' }"
      >
        <chart-card title="成功条数" :total="successCount">
          <el-tooltip
            slot="action"
            class="item"
            effect="dark"
            content="指标说明"
            placement="top-start"
          >
            <i class="el-icon-warning-outline" />
          </el-tooltip>
          <div>
            <mini-area />
          </div>
          <template
            slot="footer"
          >日成功<span> {{ "200条" }}</span></template>
        </chart-card>
      </el-col>
      <el-col
        :sm="24"
        :xs="24"
        :md="6"
        :xl="6"
        :lg="6"
        :style="{ marginBottom: '12px' }"
      >
        <chart-card title="接口请求次数" :total="requestCount">
          <el-tooltip
            slot="action"
            class="item"
            effect="dark"
            content="指标说明"
            placement="top-start"
          >
            <i class="el-icon-warning-outline" />
          </el-tooltip>
          <div>
            <mini-bar />
          </div>
          <template slot="footer">成功率 <span>100%</span></template>
        </chart-card>
      </el-col>
      <el-col
        :sm="24"
        :xs="24"
        :md="6"
        :xl="6"
        :lg="6"
        :style="{ marginBottom: '12px' }"
      >
        <chart-card title="运营活动效果" total="78%">
          <el-tooltip
            slot="action"
            class="item"
            effect="dark"
            content="指标说明"
            placement="top-start"
          >
            <i class="el-icon-warning-outline" />
          </el-tooltip>
          <div>
            <mini-progress
              color="rgb(19, 194, 194)"
              :target="80"
              :percentage="78"
              height="8px"
            />
          </div>
          <template slot="footer">
            <trend flag="top" style="margin-right: 16px" rate="12">
              <span slot="term">同周比</span>
            </trend>
            <trend flag="bottom" rate="80">
              <span slot="term">日环比</span>
            </trend>
          </template>
        </chart-card>
      </el-col>
    </el-row>

    <el-card :bordered="false" :body-style="{ padding: '0' }">
      <div class="salesCard">
        <el-tabs>
          <el-tab-pane label="发送量">
            <el-row>
              <el-col :xl="8" :lg="8" :md="8" :sm="24" :xs="24">
                <rank-list title="应用剩余额度" :list="rankList" />
              </el-col>
              <el-col :xl="16" :lg="8" :md="8" :sm="24" :xs="24">
                <bar :list="barData" title="应用发送量排行榜" />
              </el-col>
              <el-col :xl="8" :lg="8" :md="8" :sm="24" :xs="24">
                <rank-list title="业务发送量排行榜" :list="rankList" />
              </el-col>
            </el-row>
          </el-tab-pane>
        </el-tabs>
      </div>
    </el-card>
  </div>
</template>

<script>
import ChartCard from '@/components/ChartCard'
import Trend from '@/components/Trend'
import MiniArea from '@/components/MiniArea'
import MiniBar from '@/components/MiniBar'
import MiniProgress from '@/components/MiniProgress'
import RankList from '@/components/RankList/index'
import Bar from '@/components/Bar.vue'
import { getAllCount, getSuccessCount, getRequestCount } from '@/api/dashboard'
import { listSmsAppConfig } from '@/api/admin/sms-app-config'

const barData = []
const barData2 = []
for (let i = 0; i < 12; i += 1) {
  barData.push({
    x: `${i + 1}月`,
    y: Math.floor(Math.random() * 1000) + 200
  })
  barData2.push({
    x: `${i + 1}月`,
    y: Math.floor(Math.random() * 1000) + 200
  })
}

export default {
  name: 'DashboardAdmin',
  components: {
    ChartCard,
    Trend,
    MiniArea,
    MiniBar,
    MiniProgress,
    RankList,
    Bar
  },
  data() {
    return {
      barData,
      barData2,
      rankList: [],
      count: 0,
      successCount: 0,
      requestCount: 0
    }
  },
  created() {
    this.getAllCount()
    this.getSuccessCount()
    this.getRequestCount()
    this.getAppList()
  },
  methods: {
    getAllCount() {
      getAllCount().then((res) => {
        this.count = res.data
      })
    },

    getSuccessCount() {
      getSuccessCount().then((res) => {
        this.successCount = res.data
      })
    },
    getRequestCount() {
      getRequestCount().then((res) => {
        this.requestCount = res.data
      })
    },

    getAppList() {
      listSmsAppConfig({ page: 1, size: 10 }).then((resp) => {
        resp.data.list.forEach((item) => {
          this.rankList.push({
            name: item.appName,
            total: item.useNumber
          })
        })
      })
    }
  }
}
</script>

<style lang="scss" scoped>
.dashboard-editor-container {
  padding: 12px;
  background-color: rgb(240, 242, 245);
  position: relative;

  .github-corner {
    position: absolute;
    top: 0;
    border: 0;
    right: 0;
  }

  .chart-wrapper {
    background: #fff;
    padding: 16px 16px 0;
    margin-bottom: 32px;
  }
}

::v-deep .el-tabs__item {
  padding-left: 16px !important;
  height: 50px;
  line-height: 50px;
}

@media (max-width: 1024px) {
  .chart-wrapper {
    padding: 8px;
  }
}
</style>
