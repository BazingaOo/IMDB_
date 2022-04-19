<template>
  <!--  <div id="app">-->
  <!--    <h2>Todos:</h2>-->
  <!--    <ol>-->
  <!--      <li v-for="todo in todos">-->
  <!--        <label>-->
  <!--          <input type="checkbox"-->
  <!--                 v-on:change="toggle(todo)"-->
  <!--                 v-bind:checked="todo.done">-->

  <!--          &lt;!&ndash;-->
  <!--            {{ todo.text }}-->
  <!--          &ndash;&gt;-->
  <!--        </label>-->
  <!--      </li>-->
  <!--    </ol>-->
  <!--  </div>-->



  <!--  <div class="review-container" style="max-height: none;">-->
  <!--    <div class="lister-item-content">-->
  <!--      <div class="ipl-ratings-bar">-->
  <!--            <span class="rating-other-user-rating">-->
  <!--            <svg class="ipl-icon ipl-star-icon  " xmlns="http://www.w3.org/2000/svg" fill="#000000" height="24" viewBox="0 0 24 24" width="24">-->
  <!--                <path d="M0 0h24v24H0z" fill="none"></path>-->
  <!--                <path d="M12 17.27L18.18 21l-1.64-7.03L22 9.24l-7.19-.61L12 2 9.19 8.63 2 9.24l5.46 4.73L5.82 21z"></path>-->
  <!--                <path d="M0 0h24v24H0z" fill="none"></path>-->
  <!--            </svg>-->
  <!--                <span>9</span><span class="point-scale">/10</span>-->
  <!--            </span>-->
  <!--      </div>-->
  <!--      <a href="/review/rw1488015/?ref_=tt_urv" class="title"> An excellent show about really nasty people and the heinous things they do for money. Big money.-->
  <!--      </a>            <div class="display-name-date">-->
  <!--      <span class="display-name-link"><a href="/user/ur12122603/?ref_=tt_urv">swetepete</a></span><span class="review-date">28 September 2006</span>-->
  <!--    </div>-->
  <!--      <div class="content">-->
  <!--        <div class="text show-more__control">This is one of the best crime dramas on TV, maybe almost as good as early 'Sopranos' or 'Shield.' It has a wicked soundtrack, it's beautifully shot and choreographed, and sports some of the best acting I've seen on a network show. The characters are, basically, incorrigible scum; but, they pull off their heists with such style, I find myself very much hoping they stick around for a couple dozen episodes. No doubt most or all of them will meet nasty ends, but in the meantime, 'Smith' is a good place to indulge in some cathartic mayhem and villainy. The soundtrack is phenomenal&#45;&#45;Rancid, Dirty Pretty Things, Imogen Heap, Thom Yorke, and the Ramones have all popped up early on in the series. I expect the music to stay high-end as the first season progresses. 'Smith' is not for an audience who needs to cheer for the forces of law and order. It's all about the bad guys, and not saying sorry.</div>-->
  <!--        <div class="actions text-muted">-->
  <!--          46 out of 51 found this helpful.-->
  <!--          <span>-->
  <!--                            Was this review helpful? <a href="/registration/signin?ref_=urv"> Sign in</a> to vote.-->
  <!--                        </span>-->
  <!--          <br>-->
  <!--          <a href="/review/rw1488015/?ref_=tt_urv">Permalink</a>-->
  <!--        </div>-->
  <!--      </div>-->
  <!--    </div>-->
  <!--    <div class="clear"></div>-->
  <!--    <div class="lister-item-content">-->
  <!--      <div class="ipl-ratings-bar">-->
  <!--            <span class="rating-other-user-rating">-->
  <!--            <svg class="ipl-icon ipl-star-icon  " xmlns="http://www.w3.org/2000/svg" fill="#000000" height="24" viewBox="0 0 24 24" width="24">-->
  <!--                <path d="M0 0h24v24H0z" fill="none"></path>-->
  <!--                <path d="M12 17.27L18.18 21l-1.64-7.03L22 9.24l-7.19-.61L12 2 9.19 8.63 2 9.24l5.46 4.73L5.82 21z"></path>-->
  <!--                <path d="M0 0h24v24H0z" fill="none"></path>-->
  <!--            </svg>-->
  <!--                <span>9</span><span class="point-scale">/10</span>-->
  <!--            </span>-->
  <!--      </div>-->
  <!--      <a href="/review/rw2138436/?ref_=tt_urv" class="title"> strange that CBS canceled this after only one season-->
  <!--      </a>            <div class="display-name-date">-->
  <!--      <span class="display-name-link"><a href="/user/ur9298705/?ref_=tt_urv">tom-rusch</a></span><span class="review-date">9 October 2009</span>-->
  <!--    </div>-->
  <!--      <div class="content">-->
  <!--        <div class="text show-more__control clickable">after watching many of the balance of the episodes on DirecTV's 101, i'm puzzled why this didn't get more that a showing of the pilot.<br><br>interesting characters. definite possibilities for the direction the show might have gone. <br><br>all the characters might have great back stories that would give the writers an option to move beyond the mere caper plot of other similar series.<br><br>the second to the last episode on 101 suggested the background for simon baker's character and suggested a little redemptive option.<br><br>and the same episode hints at the history of ray liotta's and virginia madsen's troubled relationship.<br><br>ah well, another example of 'if i like it, they'll cancel it.'</div>-->

  <!--      </div>-->
  <!--    </div>-->
  <!--    <div class="clear"></div>-->
  <!--  </div>-->


  <div class="bgView">
<!--        <head-bar :title="newTitle" :set="newSet" :is-back="isBack" :is-set="isSet"></head-bar>-->
    <div class="bgScored">
      <div>
        <img  class="headerImg"  src="http://yijiao.oss-cn-qingdao.aliyuncs.com/web/1472127058468.jpg">
        <div class="grade-box">

        </div>
      </div>
      <div class="scoreInfo">{{scoreInfo}}</div>
    </div>
    <div class="line"></div>
    <div class="inputText">
      <textarea v-bind:maxlength="Surplus" @input="descArea" v-model="inputText" name="abstract" id="abstract" placeholder="Please leave your comment about this movie."></textarea>
      <div class="selectPic">
        <up-Load :multiple="true" :max=6 :list="imgList" ref="upload"></up-Load>
      </div>
    </div>
    <div class="bgName">
      <img src="../../../static/Images/star02@2x.png">
      <div class="unName">Anonymous</div>
    </div>
    <div class="line2"></div>
    <div class="severce">
      <div class="attitudeOfService"> Rate this movie </div>
      <el-rate
        v-model="value" :texts="texts"
        show-text>
      </el-rate>
    </div>
    <div class="publish" @click="publicComment">Submit</div>
  </div>
</template>

<script>
import headerBar from '../../components/headerBar'
import uploadImages from '../../components/uploadImages'
var starOffImg = '../../../static/Images/star02@2x.png'
var starOnImg = '../../../static/Images/star@2x.png'
export default {
  data: function () {
    return {
      texts:['Terrible','Bad','Regular','Good','Fantastic'],
      inputText: '',
      Surplus: 140,
      newTitle: '商品评价',
      scoreStartNum: 5, // 商品评分
      severStartNum: 5, // 服务评分
      newSet: '     ',
      isBack: true,
      isSet: true,
      stars: [{
        src: starOnImg,
        active: true
      }, {
        src: starOnImg,
        active: true
      }, {
        src: starOnImg,
        active: true
      },
        {
          src: starOnImg,
          active: true
        }, {
          src: starOnImg,
          active: true
        }
      ],
      severceStarts: [{
        src: starOnImg,
        active: true
      }, {
        src: starOnImg,
        active: true
      }, {
        src: starOnImg,
        active: true
      },
        {
          src: starOnImg,
          active: true
        }, {
          src: starOnImg,
          active: true
        }]
    }
  },
  components: {
    'head-bar': headerBar,
    'up-Load': uploadImages
  },
  methods: {
    // 评分
    rating: function (index, string) {
      if (string === 'starts') {
        var total = this.stars.length // 星星总数
        var idx = index + 1 // 这代表选的第idx颗星-也代表应该显示的星星数量
        // 进入if说明页面为初始状态
        if (this.scoreStartNum === 0) {
          this.scoreStartNum = idx
          for (var i = 0; i < idx; i++) {
            this.stars[i].src = starOnImg
            this.stars[i].active = true
          }
        } else {
          // 如果再次点击当前选中的星级-仅取消掉当前星级，保留之前的。
          if (idx == this.scoreStartNum) {
            for (var i = index; i < total; i++) {
              if (i != 0) {
                this.stars[i].src = starOffImg
                this.stars[i].active = false
              }
            }
          }
          // 如果小于当前最高星级，则直接保留当前星级
          if (idx < this.scoreStartNum) {
            for (var i = idx; i < this.scoreStartNum; i++) {
              if (i != 0) {
                this.stars[i].src = starOffImg
                this.stars[i].active = false
              }
            }
          }
          // 如果大于当前星级，则直接选到该星级
          if (idx > this.scoreStartNum) {
            for (var i = 0; i < idx; i++) {
              this.stars[i].src = starOnImg
              this.stars[i].active = true
            }
          }

          var count = 0 // 计数器-统计当前有几颗星
          for (var i = 0; i < total; i++) {
            if (this.stars[i].active) {
              count++
            }
          }
          this.scoreStartNum = count
        }
        if (this.scoreStartNum === 1) {
          this.scoreInfo = 'Terrible'
        } else if (this.scoreStartNum === 2) {
          this.scoreInfo = 'Bad'
        } else if (this.scoreStartNum === 3) {
          this.scoreInfo = 'Regular'
        } else if (this.scoreStartNum === 4) {
          this.scoreInfo = 'Good'
        } else if (this.scoreStartNum === 5) {
          this.scoreInfo = 'Fantastic'
        }
      } else {
        var total = this.severceStarts.length // 星星总数
        var idx = index + 1 // 这代表选的第idx颗星-也代表应该显示的星星数量
        // 进入if说明页面为初始状态
        if (this.severStartNum == 0) {
          this.severStartNum = idx
          for (var i = 0; i < idx; i++) {
            this.severceStarts[i].src = starOnImg
            this.severceStarts[i].active = true
          }
        } else {
          // 如果再次点击当前选中的星级-仅取消掉当前星级，保留之前的。
          if (idx === this.severStartNum) {
            for (var i = index; i < total; i++) {
              if (i != 0) {
                this.severceStarts[i].src = starOffImg
                this.severceStarts[i].active = false
              }
            }
          }
          // 如果小于当前最高星级，则直接保留当前星级
          if (idx < this.severStartNum) {
            for (var i = idx; i < this.severStartNum; i++) {
              if (i != 0) {
                this.severceStarts[i].src = starOffImg
                this.severceStarts[i].active = false
              }
            }
          }
          // 如果大于当前星级，则直接选到该星级
          if (idx > this.severStartNum) {
            for (var i = 0; i < idx; i++) {
              this.severceStarts[i].src = starOnImg
              this.severceStarts[i].active = true
            }
          }

          var count = 0 // 计数器-统计当前有几颗星
          for (var i = 0; i < total; i++) {
            if (this.severceStarts[i].active) {
              count++
            }
          }
          this.severStartNum = count
        }
      }
    },
    descArea () {
      var textVal = this.introduct.length
      this.Surplus = 200 - textVal
    },
    // 发布评论
    publicComment () {
      // this.scoreStartNum 代表商品评分
      // this.severStartNum 代表服务态度
      // this.inputText  代表心得
      // base64的图片流直接传给服务端，服务端传给阿里云，也可前端直接上传，拿到url链接上传到服务端。

    }
  },
  name: 'gotoComment'
}
</script>

<style scoped>
.bgView{
  padding: 15px 15px;
}
.headerImg{
  width:4.8125rem;
  height: 4.8125rem;
}
.score{
  color: #333333;
  font-size: .875rem;
  margin-left: 10px;
  display: inline-block;
}
.bgScored{
  position: relative;
  margin-top: -15px;
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 15px;
}

.bgImg img{
  margin-left: 5px;
  width: 1.09375rem;
  height:1.03125rem ;
}
.grade-box{
  position: absolute;
  top: 1.6875rem;
  left: 5.0275rem;
  display: inline-block;
}
.bgImg{
  display: inline-block;
  margin-left: 5px;
}
.line{
  margin-top: -5px;
  margin-left: -23px;
  margin-right: -23px;
  background: #c4c3c3;
  height: 1px;
  opacity: 0.3;
}
.scoreInfo{
  color:#999999 ;
  font-size: .75rem;
}
.inputText{
  width: 100%;
  margin-top: 10px;
  border-bottom: 1px solid #eef1f5;
}
.inputText textarea{
  width: 100%;
  height: 5.5rem;
  /*outline: none;*/
  /*overflow:auto;*/
  border: none;
  font-size: .75rem;
  /*background-attachment:fixed;*/
  /*background-repeat:no-repeat;*/
  /*border-style:solid;*/
  /*border-color:#FFFFFF*/
}
.selectPic{
  margin-top: 10px;
  width: 100%;
}
.unName{
  display: inline-block;
  color:#333333 ;
  margin-left: 5px;
  vertical-align: middle;
  font-size: .75rem;
}
.bgName{
  width: 100%;
  height: 50px;
  line-height: 50px;
  display: flex;
  align-items: center;
}
.bgName img{
  width: 17px;
  height: 17px;
}
.line2{
  margin-top: -5px;
  margin-left: -23px;
  margin-right: -23px;
  background: #c4c3c3;
  height: 8px;
  opacity: 0.3;
}
.attitudeOfService{
  color: #333333;
  font-size: .875rem;
  display: inline-block;
}
.severce{
  display: flex;
  height: 60px;
  align-items: center;
}
.publish{
  width: 100%;
  height: 40px;
  background: #3fc9a5;
  color: white;
  font-size:.875rem ;
  border-radius: 20px;
  line-height: 40px;
}
</style>
