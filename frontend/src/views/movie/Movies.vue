<template>
  <div class="block">
    <el-container>
      <el-header>
        <template >
          {{ movieName }}
        </template>
      </el-header>
      <el-container>
        <el-aside>Movie Poster
          <template  v-slot="props">
          <el-image
            style="width: 259px; height: 375px"
            :src="require('@/assets/'+$data.image)" fit="cover">
          </el-image></template>
        </el-aside>
        <el-container>
          <el-main>Movie Information
            <el-table
              :data="tableData"
              style="width: 100%">
              <el-table-column
                prop="genre"
                label="Genre"
              >
              </el-table-column>
              <el-table-column
                prop="abstract"
                label="Abstract"
                width="200">
              </el-table-column>
              <el-table-column

                label="Cast">
<!--                <el-table :data="castList">-->
<!--                  <el-table-column>-->
<!--                    <template v-slot="props">-->
<!--                      {{props.row.CastName}}-->
<!--                    </template>-->
<!--                  </el-table-column>-->
<!--                </el-table>-->
                <span v-for="(item, i) in castList">
                  {{item.CastName }} <br/>
                </span>
              </el-table-column>
              <el-table-column
                label="Rating">
                <template>
                  {{grade}}
                </template>

              </el-table-column>
            </el-table>

          </el-main>
          <el-footer>
            Storyline

          </el-footer>
          <hgroup class="ipc-title ipc-title--section-title ipc-title--base ipc-title--on-textPrimary">
            <h3 class="ipc-title__text">
              Comments
            </h3>
          </hgroup>
        </el-container>
      </el-container>
    </el-container>

  </div>


</template>

<script>
export default {
  data() {
    return {
      tags: [
        {name: 'Action', type: ''},
        {name: 'Comedy', type: 'success'},
        {name: 'Drama', type: 'info'},
        {name: 'Si-Fi', type: 'warning'},
        {name: 'Romance', type: 'danger'}
      ],
      movieName: '',
      year: 0,
      image:'',
      grade:0,
      tableData: [{
        abstract: '',
        cast: 'Joaquin Phoenix',
      }],
      castList: []


    };

  },
  mounted() {
    this.fetchData()
    this.searchMovieByTitle()
    this.showCast()
  }, methods: {
    fetchData() {
      this.movieId = this.$route.query.movieId
    },
    showCast(){
      let params = {
        movieId: this.movieId,
      }
      this.$axios.post("/api/user/cast/searchCastByMovieId", this.$qs.stringify(params))
        .then(res => {
          if (res.data.code === 200) {
            this.castList = res.data.cast
            this.tableData[0].abstract=res.data.movie.Description
            this.$message({
              title: '查询提示',
              message: 'This is what we found by movie titles',
              showClose: true,
              center: true,
              type: 'success'
            });
          } else {
            this.$message({
              title: '查询提示',
              message: 'Sorry, we cannot find any result by movie titles',
              center: true,
              type: 'error'
            });
          }
        }).catch(error => {
          console.log(error);
          this.$message({//这里采用element ui的一个错误显示效果模板
            title: '系统提示',
            message: error.message,
            center: true,
            type: 'warning'
          });
        }
      )
      ;
    },

    //根据movie_id来找movie信息
    searchMovieByTitle() {
      console.log("movieId:" + this.movieId)
      let params = {
        movieId: this.movieId,
      }
      this.$axios.post("/api/user/movie/searchMovieByMovieId", this.$qs.stringify(params))
        .then(res => {
          if (res.data.code === 200) {
            this.grade = res.data.movie.Grade
            this.image=res.data.movie.Image
            this.movieName = res.data.movie.Movie_name
            this.year = res.data.movie.Year
            //this.cast = res.data.movie.
            this.tableData[0].abstract=res.data.movie.Description
            this.$message({
              title: '查询提示',
              message: 'This is what we found by movie titles',
              showClose: true,
              center: true,
              type: 'success'
            });
          } else {
            this.$message({
              title: '查询提示',
              message: 'Sorry, we cannot find any result by movie titles',
              center: true,
              type: 'error'
            });
          }
        }).catch(error => {
          console.log(error);
          this.$message({//这里采用element ui的一个错误显示效果模板
            title: '系统提示',
            message: error.message,
            center: true,
            type: 'warning'
          });
        }
      )
      ;
    },
  }
}
</script>

export default {
data() {
return {
value1: null,
value2: null,
colors: ['#99A9BF', '#F7BA2A', '#FF9900']  // 等同于 { 2: '#99A9BF', 4: { value: '#F7BA2A', excluded: true }, 5: '#FF9900' }
}
}
}
<style scoped>
.el-header, .el-footer {
  background-color: #B3C0D1;
  color: #333;
  text-align: center;
  line-height: 60px;
}

.el-aside {
  background-color: #D3DCE6;
  color: #333;
  text-align: center;
  line-height: 150px;
  width: fit-content;
}

.el-main {
  background-color: #E9EEF3;
  color: #333;
  text-align: center;
  line-height: 150px;
  width: 100%;
  margin-left: auto;
  margin-right: auto;
}

body > .el-container {
  margin-bottom: 40px;
}

.el-container:nth-child(5) .el-aside,
.el-container:nth-child(6) .el-aside {
  line-height: 260px;
}

.el-container:nth-child(7) .el-aside {
  line-height: 320px;
}

</style>
