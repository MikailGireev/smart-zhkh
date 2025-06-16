<template>
  <div class="admin-layout">
    <aside class="sidebar">
      <h2 class="logo" @click="goHome">📦 Admin Panel</h2>
      <nav>
        <RouterLink to="/admin/tasks" class="nav-link" active-class="active">Заявки</RouterLink>
        <RouterLink to="/admin/settings" class="nav-link" active-class="active"
          >Настройки</RouterLink
        >
      </nav>
    </aside>

    <main class="main-content">
      <RouterView />
    </main>
  </div>
</template>

<script setup lang="ts">
import { useRouter } from 'vue-router';
const router = useRouter();

function goHome() {
  router.push('/admin/tasks');
}
</script>

<style scoped>
.admin-layout {
  display: flex;
  min-height: 100vh;
  background-color: #f8fafc;
}

.sidebar {
  width: 250px;
  background-color: #1e293b;
  color: white;
  padding: 1rem;
  display: flex;
  flex-direction: column;
}

.logo {
  font-size: 1.4rem;
  font-weight: bold;
  margin-bottom: 2rem;
  cursor: pointer;
}

.nav-link {
  color: white;
  text-decoration: none;
  margin-bottom: 1rem;
  display: block;
  padding: 0.5rem 0.75rem;
  border-radius: 8px;
  transition: background 0.2s;
}

.nav-link:hover {
  background-color: #334155;
}

.nav-link.active {
  background-color: #475569;
}

.main-content {
  flex: 1;
  padding: 2rem;
}
</style>
