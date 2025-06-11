<template>
  <div class="dashboard-container">
    <h2>Привет, {{ auth.username }} 👋</h2>
    <p>Добро пожаловать в личный кабинет</p>

    <div class="card-grid">
      <RouterLink to="/charges" class="info-card">💳 Мои начисления</RouterLink>
      <div class="info-card">🏠 Устройства (в разработке)</div>
      <RouterLink to="/profile" class="info-card">👤 Профиль</RouterLink>
    </div>
  </div>
</template>

<script setup lang="ts">
import { onMounted } from 'vue';
import { useRouter } from 'vue-router';
import { useAuthStore } from '@/shared/store/auth';

const router = useRouter();
const auth = useAuthStore();

onMounted(() => {
  if (!auth.isLoggedIn) {
    router.push('/login');
  }
});
</script>

<style scoped>
.dashboard-container {
  max-width: 960px;
  margin: 2rem auto;
  padding: 2.5rem;
  background: linear-gradient(145deg, #ffffff, #f8fafc);
  border-radius: 20px;
  box-shadow: 0 8px 30px rgba(0, 0, 0, 0.05);
  animation: fadeIn 0.6s ease;
}

@keyframes fadeIn {
  from {
    opacity: 0;
    transform: translateY(8px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

h2 {
  font-size: 28px;
  font-weight: 700;
  color: #1e3a8a;
  margin-bottom: 0.5rem;
}

p {
  font-size: 16px;
  color: #64748b;
  margin-bottom: 2rem;
}

.card-grid {
  display: grid;
  gap: 1.5rem;
  grid-template-columns: repeat(auto-fit, minmax(240px, 1fr));
}

.info-card {
  background: white;
  padding: 1.75rem;
  border-radius: 16px;
  box-shadow: 0 6px 16px rgba(0, 0, 0, 0.06);
  text-align: center;
  font-size: 18px;
  font-weight: 600;
  color: #1e293b;
  transition:
    transform 0.25s ease,
    background 0.25s ease;
  border: 2px solid transparent;
  text-decoration: none;
}

.info-card:hover {
  transform: translateY(-4px);
  background: #f0f9ff;
  border-color: #38bdf8;
  color: #0c4a6e;
  box-shadow: 0 8px 20px rgba(14, 165, 233, 0.15);
}
</style>
