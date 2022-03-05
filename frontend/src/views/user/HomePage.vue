<template>
  <div class="block">
    <el-image
      style="width: 315px; height: 110px"
      :src="require('@/assets/IMDB-CLone-FinalLogo.png')"
      :preview-src-list="require('@/assets/IMDB-CLone-FinalLogo.png')">
    </el-image>
    <el-dropdown>
      <el-button type="primary">
        MENU<i class="el-icon-arrow-down el-icon--right"></i>
      </el-button>
      <el-dropdown-menu slot="dropdown">
        <el-dropdown-item></el-dropdown-item>
        <el-dropdown-item>Top Movies</el-dropdown-item>
        <el-dropdown-items>Top Cast</el-dropdown-items>
        <el-dropdown-item>Genre</el-dropdown-item>
        <el-dropdown-item>Awards</el-dropdown-item>
      </el-dropdown-menu>
    </el-dropdown>
    <el-row>
      <a href="http://localhost:8080/#/UserLogIn">
        <el-button>Sign in</el-button>
      </a>
    </el-row>

    <el-form ref="form" :model="searchForm" label-width="40px">
      <el-form-item label="IMDB">
        <el-input v-model="searchForm.searchName" placeholder="Search movies"></el-input>
      </el-form-item>
      <a href="http://localhost:8080/#/Movies">
        <el-button type="primary" @click="onSubmit">Search</el-button>
      </a>
    </el-form>
    <!--    <a href="http://localhost:8080/#/Movies"><el-button type="primary" icon="el-icon-search">Search</el-button></a>-->


    <!--    <div class="line"></div>-->
    <el-menu
      :default-active="activeIndex2"
      class="el-menu-demo"
      mode="horizontal"
      @select="handleSelect"
      background-color="#545c64"
      text-color="#fff"
      active-text-color="#ffd04b">
      <el-menu-item index="1">Movie Ratings</el-menu-item>
      <el-submenu index="2">
        <template slot="title">Settings</template>
        <el-menu-item index="2-1">Personal Information</el-menu-item>
        <el-menu-item index="2-2">Language</el-menu-item>
        <el-menu-item index="2-3">item3</el-menu-item>
        <el-submenu index="2-4">
          <template slot="title">item4</template>
          <el-menu-item index="2-4-1">subitem1</el-menu-item>
          <el-menu-item index="2-4-2">subitem2</el-menu-item>
          <el-menu-item index="2-4-3">subitem3</el-menu-item>
        </el-submenu>
      </el-submenu>
      <el-menu-item index="3" disabled>Information Updates</el-menu-item>
      <el-menu-item index="4"><a href="https://www.ele.me" target="_blank">Watchlist</a></el-menu-item>
    </el-menu>
    <span class="demonstration">Click to Swtich Moive Poster</span>
    <el-carousel :interval="4000" type="card" height="300px">
      <el-carousel-item>
        <div class="item">
          <div class="item__content">
            The first movie
          </div>
          <img class="item__image" src="https://picsum.photos/300?random=1" style="width: 900px; height: 300px" alt=""/>
        </div>
      </el-carousel-item>
      <el-carousel-item>
        <div class="item">
          <div class="item__content">
            The second movie
          </div>
          <img class="item__image" src="https://picsum.photos/300?random=2" style="width: 900px; height: 300px" alt=""/>
        </div>
      </el-carousel-item>
      <el-carousel-item>
        <div class="item">
          <div class="item__content">
            Another one movie
          </div>
          <img class="item__image" src="https://picsum.photos/300?random=3" style="width: 900px; height: 300px" alt=""/>
        </div>
      </el-carousel-item>
    </el-carousel>

    <el-popover
      placement="top-start"
      title="标题"
      width="200"
      trigger="hover"
      content="这是一段内容,这是一段内容,这是一段内容,这是一段内容。">
      <el-button slot="reference">hover activate</el-button>
    </el-popover>

    <el-popover
      placement="bottom"
      title="标题"
      width="200"
      trigger="click"
      content="这是一段内容,这是一段内容,这是一段内容,这是一段内容。">
      <el-button slot="reference">click activate</el-button>
    </el-popover>

    <el-popover
      ref="popover"
      placement="right"
      title="标题"
      width="200"
      trigger="focus"
      content="这是一段内容,这是一段内容,这是一段内容,这是一段内容。">
    </el-popover>
    <el-button v-popover:popover>focus activate</el-button>

    <el-popover
      placement="bottom"
      title="标题"
      width="200"
      trigger="manual"
      content="这是一段内容,这是一段内容,这是一段内容,这是一段内容。"
      v-model="visible">
      <el-button slot="reference" @click="visible = !visible">mannually activate</el-button>
    </el-popover>

  </div>

</template>

<script>
export default {
  data() {
    return {
      searchForm: {
        searchName: ''
      },
      loginForm: {
        username: 'Joker',
      },
      url: "require('@/assets/IMDB-CLone-FinalLogo.png')",
      srcList: [
        "require('@/assets/IMDB-CLone-FinalLogo.png')"
      ],
      input: '',
      activeIndex: '1',
      activeIndex2: '1'
    }
  },
  methods: {
    onSubmit() {
      let params = {
        searchName: this.searchForm.searchName,
      }
      this.$axios.post("/api/user/movie/searchMovieByName", this.$qs.stringify(params))
        .then(res => {
          if (res.status == 200) {
            localStorage.setItem('token', res.data.token);
            this.$router.push({path:'/Movies',query:{movie:res.data.movie,msg:res.data.message}});
            this.$message({
              title: '查询提示',
              message: '查询成功',
              showClose: true,
              center: true,
              type: 'success'
            });
          } else {
            this.$message({
              title: '查询提示',
              message: '电影不存在',
              center: true,
              type: 'error'
            });
            this.load = false;
          }
        }).catch(error => {
          this.load = false;
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


    handleSelect(key, keyPath) {
      console.log(key, keyPath);
    },
    handleClick() {
      alert('button click');
    }
  }
}

</script>

<style scoped>
.el-carousel__item h3 {
  color: #475669;
  font-size: 14px;
  opacity: 0.75;
  line-height: 300px;

  margin: 0;
}

.el-carousel__item:nth-child(2n) {
  background-color: #99a9bf;
}

.el-carousel__item:nth-child(2n+1) {
  background-color: #d3dce6;
}

.el-dropdown {
  vertical-align: top;
}

.el-dropdown + .el-dropdown {
  margin-left: 15px;
}

.el-icon-arrow-down {
  font-size: 12px;
}

.el-input {
  vertical-align: baseline;
}

.el-input + .el-input {
  margin-left: 15px;
}


</style>
