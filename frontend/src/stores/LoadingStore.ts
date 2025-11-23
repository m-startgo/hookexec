import { defineStore } from 'pinia';

type LoadingState = {
  status: boolean;
  info: string;
  icon: string;
};

type LoadingShowOpt = {
  info?: string;
  icon?: string;
};

export const useLoadingStore = defineStore('LoadingStore', {
  // 状态：存储数据
  state(): LoadingState {
    return {
      status: false,
      info: '',
      icon: '',
    };
  },

  // Actions ：类似 methods，用于修改状态（支持同步和异步）
  actions: {
    Show(opt?: LoadingShowOpt) {
      this.info = opt?.info || '加载中...';
      this.icon = opt?.icon || 'el-icon-loading';
      this.status = true;
      console.log('显示loading');
    },
    Hide() {
      this.status = false;
      this.info = '';
      this.icon = '';
      console.log('关闭loading');
    },
  },
});
