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
  <div class="review-container" style="max-height: none;">
    <div class="lister-item-content">
      <div class="ipl-ratings-bar">
            <span class="rating-other-user-rating">
            <svg class="ipl-icon ipl-star-icon  " xmlns="http://www.w3.org/2000/svg" fill="#000000" height="24" viewBox="0 0 24 24" width="24">
                <path d="M0 0h24v24H0z" fill="none"></path>
                <path d="M12 17.27L18.18 21l-1.64-7.03L22 9.24l-7.19-.61L12 2 9.19 8.63 2 9.24l5.46 4.73L5.82 21z"></path>
                <path d="M0 0h24v24H0z" fill="none"></path>
            </svg>
                <span>9</span><span class="point-scale">/10</span>
            </span>
      </div>
      <a href="/review/rw1488015/?ref_=tt_urv" class="title"> An excellent show about really nasty people and the heinous things they do for money. Big money.
      </a>            <div class="display-name-date">
      <span class="display-name-link"><a href="/user/ur12122603/?ref_=tt_urv">swetepete</a></span><span class="review-date">28 September 2006</span>
    </div>
      <div class="content">
        <div class="text show-more__control">This is one of the best crime dramas on TV, maybe almost as good as early 'Sopranos' or 'Shield.' It has a wicked soundtrack, it's beautifully shot and choreographed, and sports some of the best acting I've seen on a network show. The characters are, basically, incorrigible scum; but, they pull off their heists with such style, I find myself very much hoping they stick around for a couple dozen episodes. No doubt most or all of them will meet nasty ends, but in the meantime, 'Smith' is a good place to indulge in some cathartic mayhem and villainy. The soundtrack is phenomenal--Rancid, Dirty Pretty Things, Imogen Heap, Thom Yorke, and the Ramones have all popped up early on in the series. I expect the music to stay high-end as the first season progresses. 'Smith' is not for an audience who needs to cheer for the forces of law and order. It's all about the bad guys, and not saying sorry.</div>
        <div class="actions text-muted">
          46 out of 51 found this helpful.
          <span>
                            Was this review helpful? <a href="/registration/signin?ref_=urv"> Sign in</a> to vote.
                        </span>
          <br>
          <a href="/review/rw1488015/?ref_=tt_urv">Permalink</a>
        </div>
      </div>
    </div>
    <div class="clear"></div>
    <div class="lister-item-content">
      <div class="ipl-ratings-bar">
            <span class="rating-other-user-rating">
            <svg class="ipl-icon ipl-star-icon  " xmlns="http://www.w3.org/2000/svg" fill="#000000" height="24" viewBox="0 0 24 24" width="24">
                <path d="M0 0h24v24H0z" fill="none"></path>
                <path d="M12 17.27L18.18 21l-1.64-7.03L22 9.24l-7.19-.61L12 2 9.19 8.63 2 9.24l5.46 4.73L5.82 21z"></path>
                <path d="M0 0h24v24H0z" fill="none"></path>
            </svg>
                <span>9</span><span class="point-scale">/10</span>
            </span>
      </div>
      <a href="/review/rw2138436/?ref_=tt_urv" class="title"> strange that CBS canceled this after only one season
      </a>            <div class="display-name-date">
      <span class="display-name-link"><a href="/user/ur9298705/?ref_=tt_urv">tom-rusch</a></span><span class="review-date">9 October 2009</span>
    </div>
      <div class="content">
        <div class="text show-more__control clickable">after watching many of the balance of the episodes on DirecTV's 101, i'm puzzled why this didn't get more that a showing of the pilot.<br><br>interesting characters. definite possibilities for the direction the show might have gone. <br><br>all the characters might have great back stories that would give the writers an option to move beyond the mere caper plot of other similar series.<br><br>the second to the last episode on 101 suggested the background for simon baker's character and suggested a little redemptive option.<br><br>and the same episode hints at the history of ray liotta's and virginia madsen's troubled relationship.<br><br>ah well, another example of 'if i like it, they'll cancel it.'</div>

      </div>
    </div>
    <div class="clear"></div>
  </div>
</template>

<script>
export default {
  name: "Comments",
  methods: {
    open() {
      this.$alert('这是一段内容', '标题名称', {
        confirmButtonText: '确定',
        callback: action => {
          this.$message({
            type: 'info',
            message: `action: ${ action }`
          });
        }
      });
    }
  }
}
</script>

<style scoped>

</style>
