<template>
  <div>
    <header class="site-header">
      <div class="container">
        <h1 class="logo" @click="goHome">🧠 IoT Dashboard</h1>

        <nav v-if="auth.isLoggedIn">
          <RouterLink to="/dashboard">Главная</RouterLink>
          <RouterLink to="/charges">Начисления</RouterLink>
          <RouterLink to="/settings">Настройки</RouterLink
          ><RouterLink to="/profile">Профиль</RouterLink>
          <button @click="logout">Выйти</button>
        </nav>

        <nav v-else>
          <RouterLink to="/login">Вход</RouterLink>
          <RouterLink to="/register">Регистрация</RouterLink>
        </nav>
      </div>
    </header>

    <main class="main-content">
      <RouterView />
    </main>
  </div>
</template>

<script setup lang="ts">
import { useRouter } from 'vue-router';
import { useAuthStore } from '@/shared/store/auth';

const auth = useAuthStore();
const router = useRouter();

function logout() {
  auth.logout();
  router.push('/login');
}

function goHome() {
  if (auth.isLoggedIn) {
    router.push('/');
  } else {
    router.push('/login');
  }
}
</script>

<style scoped>
.site-header {
  background: linear-gradient(90deg, #1d4ed8 0%, #2563eb 100%);
  color: white;
  padding: 1rem 0;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15);
  position: sticky;
  top: 0;
  z-index: 1000;
}

.container {
  max-width: 1200px;
  margin: 0 auto;
  padding: 0 2rem;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.logo {
  font-size: 1.5rem;
  font-weight: 700;
  cursor: pointer;
  letter-spacing: 0.5px;
  transition: transform 0.2s ease-in-out;
}
.logo:hover {
  transform: scale(1.05);
}

nav {
  display: flex;
  align-items: center;
  gap: 1.25rem;
}

nav a,
nav button {
  color: white;
  font-weight: 500;
  text-decoration: none;
  background: rgba(255, 255, 255, 0.1);
  padding: 0.4rem 0.9rem;
  border-radius: 8px;
  transition: background 0.2s, transform 0.2s;
  font-size: 0.95rem;
  border: none;
  cursor: pointer;
}

nav a:hover,
nav button:hover {
  background: rgba(255, 255, 255, 0.2);
  transform: translateY(-1px);
}

.main-content {
  padding: 2.5rem 2rem;
  background: linear-gradient(to bottom, #f1f5f9 0%, #f9fafb 100%);
  min-height: 85vh;
  box-shadow: inset 0 1px 3px rgba(0, 0, 0, 0.03);
  border-top: 1px solid #e2e8f0;
  transition: background 0.3s ease-in-out;
}

</style>
