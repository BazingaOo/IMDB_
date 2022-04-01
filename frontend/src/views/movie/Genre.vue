<template>
  <div class="types">
    <h2>Search based on types</h2>
    <div class="content">
      <ul class="type-list">
        <li v-for="item in items">
          <a :href="baseUri + item.href">
            {{item.title}}<span></span>
          </a>
        </li>
        <li></li>
      </ul>
    </div>
  </div>
</template>

<script>
export default {
  name: 'types',
  data () {
    return {
      baseUri: '',
      items: [
        {
          title: 'CLASSIC',
          href: 'movie/classic'
        },
        {
          title: 'UNDERRATED',
          href: 'movie/underrated'
        },
        {
          title: 'TOPRATINGS',
          href: 'movie/doubantop'
        },
        {
          title: 'ACTION',
          href: 'movie/action'
        },
        {
          title: 'COMEDY',
          href: 'movie/comedy'
        },
        {
          title: 'ROMANCE',
          href: 'movie/love'
        },
        {
          title: 'MYSTERY',
          href: 'movie/mystery'
        },
        {
          title: 'HORROR',
          href: 'movie/horror'
        },
        {
          title: 'SCIFI',
          href: 'movie/scifi'
        },
        {
          title: 'WHOLESOME',
          href: 'movie/sweet'
        },
        {
          title: 'ART',
          href: 'movie/artfilm'
        },
        {
          title: 'YOUTH',
          href: 'movie/youth'
        },
        {
          title: 'ANIMATION',
          href: 'movie/animation'
        },
        {
          title: 'CHINESE',
          href: 'movie/chinese'
        },
        {
          title: 'WESTERN',
          href: 'movie/western'
        },
        {
          title: 'KOREAN',
          href: 'movie/korean'
        },
        {
          title: 'JAPANESE',
          href: 'movie/japanese'
        }
      ]
    }
  }
}
</script>

<style lang="scss" scoped>
.types {
  margin-top: 1rem;

  h2 {
    height: 2.6rem;
    line-height: 2.6rem;
    padding: 0 1.6rem;
    display: inline-block;
  }
}

.type-list {
  margin: 1.6rem 0 0 1.6rem;
  overflow: hidden;

  li {
    float: left;
    box-sizing: border-box;
    width: 50%;
    padding-right: 1.8rem;
    height: 4.2rem;
    line-height: 4.2rem;
    font-size: 1.6rem;
    border-top: solid 0.1rem #eee;
    border-right: solid 0.1rem #eee;

    a {
      color: #42bd56;
    }
  }

  li:nth-child(2n) {
    padding-left: 1.8rem;
  }

  span {
    color: #ccc;
    float: right;
    font-weight: bold;
    display: inline-block;
    border-right: solid 0.1rem #ccc;
    border-bottom: solid 0.1rem #ccc;
    width: 0.8rem;
    height: 0.8rem;
    transform: rotate(-45deg);
    margin-top: 1rem;
  }
}
</style>
