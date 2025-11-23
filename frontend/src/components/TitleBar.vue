<template>
  <div class="TitleBar" @click.stop>
    <button class="TitleShowBtn" @click.stop="isHide = !isHide">
      {{ isHide ? '显示菜单' : '隐藏菜单' }}
    </button>
    <img v-if="!isHide" alt="AppLogo" class="AppLogo" src="/img/appicon.png" draggable="false" />
    <SimpleBar v-if="!isHide" class="Content">
      <div class="AppInfo">
        <div class="url" @dblclick="SelectAll">
          {{ fullUrl }}
        </div>
        <div class="nav">
          <RouterLink v-for="item in routes" :key="item.name" class="nav-item" :to="item.path">
            {{ item.name }}
          </RouterLink>
        </div>
      </div>
    </SimpleBar>
  </div>
</template>

<script lang="ts">
import router from '@src/router';
import { RouterLink } from 'vue-router';
import SimpleBar from 'simplebar-vue';

const routes = router.getRoutes();

export default {
  name: 'TitleBar',
  components: {
    RouterLink,
    SimpleBar,
  },
  data() {
    return {
      isHide: true,
      routes,
      fullUrl: window.location.href,
      isBut: true,
    };
  },
  computed: {},
  watch: {
    $route() {
      this.fullUrl = window.location.href;
    },
  },
  mounted() {
    document.onclick = () => {
      if (!this.isHide) {
        this.isHide = true;
      }
    };
  },
  beforeUnmount() {},
  methods: {
    SelectAll(event: Event) {
      const range = document.createRange();
      range.selectNodeContents(event?.target as Node);
      const selection = window.getSelection();
      selection?.removeAllRanges();
      selection?.addRange(range);
    },
  },
};
</script>

<style scoped lang="scss">
@use '@src/assets/global.scss';

.TitleBar {
  border: 1px solid #000;
  border-radius: 4px;
  background-color: rgba($color: #e09bc7, $alpha: 0.5);
  position: fixed;
  z-index: 999;
  left: 10px;
  top: 10px;
}
.AppLogo {
  position: absolute;
  right: 0;
  top: 0;
  z-index: 3;
  width: 60px;
  opacity: 0.9;
  will-change: opacity;
}

.TitleShowBtn {
  margin: 5px;
}

.Content {
  width: 300px;
  height: 300px;
  padding: 5px;
  position: relative;
}

.AppInfo {
  background-color: rgba($color: #ffff, $alpha: 0.7);
  text-align: center;
  font-size: 14px;
  .name {
    font-weight: bold;
    text-shadow: 1px 1px 2px #fff;
  }
  .url {
    color: #888;
    word-break: break-all;
    user-select: text;
    text-shadow: 1px 1px 2px #fff;
  }

  .nav {
    display: flex;
    align-items: center;
    flex-wrap: wrap;
    padding: 5px;
    .nav-item {
      font-size: 12px;
      margin: 4px;
      padding: 4px 6px;
      cursor: pointer;
      background-color: #fff;
      border-radius: 4px;
      @include global.no-select;
    }
    .router-link-exact-active {
      // 精确匹配（只有当前路由和链接路径完全一致时才会加上）
      padding: 6px 8px;
      background-color: rgba($color: #409eff, $alpha: 0.8);
      color: #fff;
    }
    .router-link-active {
      // 模糊匹配（当前路由包含该链接的路径时就会加上）
      @extend .router-link-exact-active;
    }
  }
}

.BtnWrapper {
  margin-top: 5px;
  margin-bottom: 5px;
  .btn {
    cursor: pointer;
  }
}

.SysInfo {
  display: block;
  background-color: #000;
  color: #fff;
  font-size: 12px;
}
</style>
<style>
.TitleBar {
  .JsonBlock {
    padding: 5px;
    .jv-code {
      padding: 0;
    }
    .jv-node.toggle {
      margin-left: 0 !important;
    }
  }
}
</style>
