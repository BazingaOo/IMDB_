<template>
  <Layout :has_share="false" title="My Account">
    <div class="wrap">
      <mu-list>
        <mu-list-item title="Watch History" to="user/history">
          <mu-icon slot="left" value="history"/>
          <mu-icon slot="right" value="keyboard_arrow_right"/>
        </mu-list-item>
        <mu-list-item title="Watchlist" to="user/favorite">
          <mu-icon slot="left" value="favorite"/>
          <mu-icon slot="right" value="keyboard_arrow_right"/>
        </mu-list-item>
        <mu-list-item title="Setting" to="user/settings">
          <mu-icon slot="left" value="settings"/>
          <mu-icon slot="right" value="keyboard_arrow_right"/>
        </mu-list-item>
        <mu-list-item title="About" to="user/about">
          <mu-icon slot="left" value="timeline"/>
          <mu-icon slot="right" value="keyboard_arrow_right"/>
        </mu-list-item>
      </mu-list>
    </div>
  </Layout>
</template>

<script>
let _self;


export default {
  data: function() {
    return {

    };
  },
  created() {
    _self = this;
  },
  mounted() {

  },
  methods: {

  },
  computed: {

  },
  components: {
    Layout,
  }
};

</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="less">

</style>
