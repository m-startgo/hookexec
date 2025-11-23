<template>
  <TitleBar />
  <div class="page">
    <RouterView />
  </div>
</template>

<script lang="ts">
import { RouterView } from 'vue-router';

import TitleBar from '@src/components/TitleBar.vue';

export default {
  name: 'App',
  components: {
    TitleBar,
    RouterView,
  },
  data() {
    return {};
  },
  computed: {},

  mounted() {},
  methods: {
    IsShow_TitleBar() {},
  },
};
</script>

<style lang="scss" scoped>
.page {
  height: 100%;
  width: 100%;
}
</style>
