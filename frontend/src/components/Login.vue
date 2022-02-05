<template>
  <!--  <div class="vue-tempalte">-->
  <!--    <form>-->
  <!--      <h3>Sign In</h3>-->

  <!--      <div class="form-group" name="">-->
  <!--        <label>Email address</label>-->
  <!--        <input type="text" class="form-control form-control-lg" name="username"/>-->
  <!--      </div>-->

  <!--      <div class="form-group">-->
  <!--        <label>Password</label>-->
  <!--        <input type="password" class="form-control form-control-lg"name="password"/>-->
  <!--      </div>-->

  <!--      <button type="submit" class="btn btn-dark btn-lg btn-block">Sign In</button>-->

  <!--      <p class="forgot-password text-right mt-2 mb-4">-->
  <!--        <router-link to="/forgot-password">Forgot password ?</router-link>-->
  <!--      </p>-->

  <!--      <div class="social-icons">-->
  <!--        <ul>-->
  <!--          <li><a href="#"><i class="fa fa-google"></i></a></li>-->
  <!--          <li><a href="#"><i class="fa fa-facebook"></i></a></li>-->
  <!--          <li><a href="#"><i class="fa fa-twitter"></i></a></li>-->
  <!--        </ul>-->
  <!--      </div>-->

  <!--    </form>-->
  <!--  </div>-->
  <div>
    <el-form ref="loginForm" :model="form" :rules="rules" label-width="80px" class="login-box">
      <h3 class="login-title">欢迎登录</h3>
      <el-form-item label="账号" prop="username">
        <el-input type="text" placeholder="请输入账号" v-model="form.username"/>
      </el-form-item>
      <el-form-item label="密码" prop="password">
        <el-input type="password" placeholder="请输入密码" v-model="form.password"/>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" v-on:click="onSubmit('loginForm')">登录</el-button>
      </el-form-item>
    </el-form>

    <el-dialog
      title="温馨提示"
      :visible.sync="dialogVisible"
      width="30%"
      >
      <span>请输入账号和密码</span>
      <span slot="footer" class="dialog-footer">
        <el-button type="primary" @click="dialogVisible = false">确 定</el-button>
      </span>
    </el-dialog>
  </div>
</template>

<script>
export default {
  data() {
    return {
      form: {
        username: '',
        password: ''
      },
      //表单验证，需要再el-form-item 元素中增加prop属性
      rules: {
        username: [
          {required: true, message: '账号不能为空', trigger: 'blur'}
        ],
        password: [
          {required: true, message: '密码不能为空', trigger: 'blur'}
        ]
      },
      //对话框显示和隐藏
      dialogVisible: false
    }
  },
  methods: {
    // testAxios() {
    //   // 由于 main.js 里全局定义的 axios,此处直接使用 $axios 即可。
    //   // 由于 main.js 里定义了每个请求前缀，此处的 / 即为 /api/，
    //   // 经过 vue.config.js 配置文件的代理设置，会自动转为 https://www.baidu.com/，从而解决跨域问题
    //   this.$axios.get("v1/signin").then(response => {
    //     if (response.data) {
    //       console.log(response.data)
    //     }
    //   }).catch(err => {
    //     alert('请求失败')
    //   })
    // },
    onSubmit(formName) {
      //为表单绑定验证功能
      this.$refs[formName].validate((valid) => {
        if (valid) {
          //使用 vue-router路由到指定页面，该方式称之为编程式导航
          console.log("fasd"+this.form.username);
          let params={
            username:this.form.username,
            password:this.form.password
          }
          this.$axios.post("/v1/signin",this.$qs.stringify(params))
            .then(res => {
              console.log("resdate:"+res.data.code);
            if (res.data.code==200){
              this.$router.push('/hello')
            }
            })
        } else {
          this.dialogVisible = true;
          return false;
        }
      });
    }
  }
}
</script>
<!--<style lang="scss" scoped>-->
<!--.login-box{-->
<!--  border: 1px solid #DCDFE6;-->
<!--  width: 350px;-->
<!--  margin:180px auto;-->
<!--  padding:35px 35px 15px 35px;-->
<!--  border-radius: 5px;-->
<!--  -webkit-border-radius: 5px;-->
<!--  -moz-border-radius: 5px;-->
<!--  box-shadow:0 0 25px #909399;-->
<!--}-->

<!--.login-title{-->
<!--  text-align:center;-->
<!--  margin:0 auto 40px auto;-->
<!--  color:#303133;-->
<!--}-->
<!--\\</style>-->
