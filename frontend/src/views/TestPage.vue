<template>
  <div class="TestPage">
    <SimpleBar class="SimpleBar">
      <h1>{{ title }}</h1>
      <div class="hint">若此处没时间显示，说明与后端通信有问题！</div>
      <div class="input-box">
        <input class="n-input" v-model="math.n1" type="text" autocomplete="off" />
        <input class="n-input" v-model="math.n2" type="text" autocomplete="off" />
        <button class="btn" @click="doAdd">Add</button>
        <div class="result">
          Result： {{ math.result }} 。<span class="hint">如果计算的值不正确则表示有问题</span>
        </div>
      </div>

      <div>计数：{{ count }} 双倍计数：{{ doubleCount }}</div>

      <CounterOptions />

      <div @contextmenu.prevent="showMenu($event)">
        <p>前端页面级右键菜单 (右键点我)</p>
        <ul
          v-if="menuInfo.visible"
          class="context-menu"
          :style="{ top: menuInfo.y + 'px', left: menuInfo.x + 'px' }"
        >
          <li @click="refresh">刷新页面</li>
          <li @click="quit">退出应用</li>
        </ul>
      </div>
    </SimpleBar>
  </div>
</template>

<script lang="ts">
import CounterOptions from '@src/stores/CounterOptions.vue';
import { useCounterOptionsStore } from '@src/stores/counterOptions';
import { mapState } from 'pinia';
import SimpleBar from 'simplebar-vue';

export default {
  name: 'TestPage',
  components: {
    CounterOptions,
    SimpleBar,
  },
  data() {
    return {
      title: '',
      math: {
        n1: '',
        n2: '',
        result: '',
      },
      menuInfo: {
        visible: false,
        x: 0,
        y: 0,
      },
    };
  },
  computed: {
    ...mapState(useCounterOptionsStore, ['count', 'doubleCount']),
  },
  mounted() {
    this.title = `This is ${this.$options.name} ! `;
    // 点击其他地方时关闭菜单
    window.addEventListener('click', this.hideMenu);
  },
  beforeUnmount() {
    window.removeEventListener('click', this.hideMenu);
  },
  methods: {
    doAdd() {
      const _this = this;
      const n1 = Number(_this.math.n1);
      const n2 = Number(_this.math.n2);
      _this.math.result = String(n1 + n2);
    },
    showMenu(e: MouseEvent) {
      this.menuInfo.x = e.clientX; // 相对于整个文档左上角
      this.menuInfo.y = e.clientY;
      this.menuInfo.visible = true;
    },
    refresh() {
      console.info('刷新逻辑');
      this.menuInfo.visible = false;
    },
    quit() {
      console.info('退出逻辑');
      this.menuInfo.visible = false;
    },
    hideMenu() {
      this.menuInfo.visible = false;
    },
  },
};
</script>

<style scoped lang="scss">
.TestPage {
  border: 1px solid #000;
  padding: 2px;
  border-radius: 10px;
  box-sizing: border-box;
  height: 100%;
  background-color: rgba($color: #fff, $alpha: 0.8);

  .SimpleBar {
    padding: 15px;
    padding-bottom: 40px;
    box-sizing: border-box;
    height: 100%;
  }
}

.hint {
  color: #888;
  text-shadow: 1px 1px 2px #ccc;
  font-size: 12px;
}

.n-input {
  width: 80px;
  margin-right: 10px;
  padding: 4px 8px;
  font-size: 14px;
}

.context-menu {
  position: fixed;
  background: #fff;
  border: 1px solid #ccc;
  list-style: none;
  padding: 4px 0;
  margin: 0;
  z-index: 9999;
}
.context-menu li {
  padding: 6px 12px;
  cursor: pointer;
}
.context-menu li:hover {
  background: #eee;
}
</style>
