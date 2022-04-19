<template>
<!--  <el-table-->
<!--    :data="tableData"-->
<!--    style="width: 100%">-->
<!--    <el-table-column-->
<!--      type="index"-->
<!--      :index="indexMethod">-->
<!--    </el-table-column>-->
<!--    <el-table-column-->
<!--      prop="date"-->
<!--      label="日期"-->
<!--      width="180">-->
<!--    </el-table-column>-->
<!--    <el-table-column-->
<!--      prop="name"-->
<!--      label="姓名"-->
<!--      width="180">-->
<!--    </el-table-column>-->
<!--    <el-table-column-->
<!--      prop="address"-->
<!--      label="地址">-->
<!--    </el-table-column>-->
<!--  </el-table>-->
  <div style="height: 400px" ref="chartEducation">User Rating</div>



</template>

<script>
import * as echarts from "echarts";

export default {
  mounted() {
    this.$nextTick(() => {
      this.initCharts4sEducation();
    });
  },
  methods: {
    initCharts4sEducation() {
      let myChart = echarts.init(this.$refs.chartEducation);
      myChart.setOption({
        title: "",
        tooltip: {},
        legend: {
          orient: "horizontal",
          x: "right",
          y: "bottom",
          itemWidth: 18, // 设置图例图形的宽
          itemHeight: 30, // 设置图例图形的高
          itemGap: 30,
          backgroundColor: "#eee", // 设置整个图例区域背景颜色
          data: ["5 Stars", "4 Stars", "3 Stars", "2 Stars", "1 Star"],
        },
        series: [
          {
            name: "User Ratings",
            type: "pie", // 设置图表类型为饼图
            radius: ["55%", "65%"], // 饼图的半径，外半径为可视区尺寸（容器高宽中较小一项）的 55% 长度。
            data: [
              // 数据数组，name 为数据项名称，value 为数据项值
              { value: 30, name: "5 Stars" },
              { value: 40, name: "4 Stars" },
              { value: 20, name: "3 Stars" },
              { value: 10, name: "2 Stars" },
              { value: 10, name: "1 Star" },
            ],
            color: ["green", "blue", "Cyan", "red", "yellow"], //各个区域颜色
            itemStyle: {
              normal: {
                label: {
                  show: true,
                  formatter: "{b} {d}%",
                }, //饼图图形上的文本标签，可用于说明图形的一些数据信息，比如值，名称等。可以与itemStyle属性同级，具体看文档
                labelLine: {
                  show: true,
                },
              },
            },
          },
        ],
      });
    },
  },
  data() {
    return {
      // initCharts4sEducation() {
      //   let myChart = echarts.init(this.$refs.chartEducation);
      //   myChart.setOption({
      //     title: "",
      //     tooltip: {},
      //     legend: {
      //       orient: "horizontal",
      //       x: "right",
      //       y: "top",
      //       itemWidth: 18, // 设置图例图形的宽
      //       itemHeight: 30, // 设置图例图形的高
      //       itemGap: 30,
      //       backgroundColor: "#eee", // 设置整个图例区域背景颜色
      //       data: ["大学及以上", "高中及以下", "初中及以下", "小学及以下"],
      //     },
      //     series: [
      //       {
      //         name: "教育程度",
      //         type: "pie", // 设置图表类型为饼图
      //         radius: ["55%", "65%"], // 饼图的半径，外半径为可视区尺寸（容器高宽中较小一项）的 55% 长度。
      //         data: [
      //           // 数据数组，name 为数据项名称，value 为数据项值
      //           { value: 30, name: "大学及以上" },
      //           { value: 40, name: "高中及以下" },
      //           { value: 20, name: "初中及以下" },
      //           { value: 10, name: "小学及以下" },
      //         ],
      //         color: ["green", "blue", "Cyan", "red"], //各个区域颜色
      //         itemStyle: {
      //           normal: {
      //             label: {
      //               show: true,
      //               formatter: "{b} {d}%",
      //             }, //饼图图形上的文本标签，可用于说明图形的一些数据信息，比如值，名称等。可以与itemStyle属性同级，具体看文档
      //             labelLine: {
      //               show: true,
      //             },
      //           },
      //         },
      //       },
      //     ],
      //   });
      // },
      tableData: [{
        date: '2016-05-02',
        name: '王小虎',
        province: '上海',
        city: '普陀区',
        address: '上海市普陀区金沙江路 1518 弄',
        zip: 200333,
        tag: '家'
      }, {
        date: '2016-05-04',
        name: '王小虎',
        province: '上海',
        city: '普陀区',
        address: '上海市普陀区金沙江路 1517 弄',
        zip: 200333,
        tag: '公司'
      }, {
        date: '2016-05-01',
        name: '王小虎',
        province: '上海',
        city: '普陀区',
        address: '上海市普陀区金沙江路 1519 弄',
        zip: 200333,
        tag: '家'
      }, {
        date: '2016-05-03',
        name: '王小虎',
        province: '上海',
        city: '普陀区',
        address: '上海市普陀区金沙江路 1516 弄',
        zip: 200333,
        tag: '公司'
      },{
        date: '2016-05-03',
        name: '王小虎',
        province: '上海',
        city: '普陀区',
        address: '上海市普陀区金沙江路 1516 弄',
        zip: 200333,
        tag: '公司'
      }],
    }
  },
  // methods: {
  //   indexMethod(index) {
  //     return index +1 ;
  //   }
  // }
};
</script>

<style scoped>
.chartEducation h3 {
  color: #475669;
  font-size: 18px;
  opacity: 0.75;
  line-height: 300px;
  margin: 0;
}


</style>
