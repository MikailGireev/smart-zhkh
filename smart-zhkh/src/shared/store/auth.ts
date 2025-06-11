import { defineStore } from 'pinia';

export const useAuthStore = defineStore('auth', {
  state: () => ({
    username: localStorage.getItem('username') || '',
    token: localStorage.getItem('token') || '',
    userId: Number(localStorage.getItem('userId')) || 0,
  }),

  actions: {
    login(username: string, token: string, userId: number) {
      this.username = username;
      this.token = token;
      this.userId = userId;

      localStorage.setItem('username', username);
      localStorage.setItem('token', token);
      localStorage.setItem('userId', userId.toString());
    },

    logout() {
      this.username = '';
      this.token = '';
      this.userId = 0;

      localStorage.removeItem('username');
      localStorage.removeItem('token');
      localStorage.removeItem('userId');
    },
  },

  getters: {
    isLoggedIn: (state) => !!state.token,
  },
});
