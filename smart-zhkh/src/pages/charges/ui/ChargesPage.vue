<template>
  <div class="charges-container">
    <div class="charges-header">
      <h2>Ваши начисления</h2>
      <RouterLink to="/charges/add" class="add-button">➕ Добавить</RouterLink>
    </div>

    <div v-if="isLoading" class="loading">⏳ Загрузка...</div>
    <div v-else-if="filteredCharges.length === 0" class="empty-state">
      <p>Пока нет начислений</p>
    </div>

    <div class="card-grid" v-else>
      <div v-for="charge in filteredCharges" :key="charge.id" class="charge-card">
        <div class="charge-icon">📂</div>
        <h3 class="category">{{ charge.category }}</h3>
        <p class="amount">💰 {{ charge.amount.toLocaleString() }} ₽</p>
        <p class="date">📅 {{ formatDate(charge.date) }}</p>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, computed } from 'vue';
import { useAuthStore } from '@/shared/store/auth';
import { fetchCharges } from '@/shared/api/charges';

const auth = useAuthStore();
const charges = ref<any[]>([]);
const isLoading = ref(true);

onMounted(async () => {
  try {
    const data = await fetchCharges(auth.userId);
    charges.value = data;
    console.log('Загружены начисления:', data);
  } catch (e) {
    console.error('Ошибка загрузки:', e);
  } finally {
    isLoading.value = false;
  }
});

const filteredCharges = computed(() => charges.value);

function formatDate(date: string) {
  return new Date(date).toLocaleDateString('ru-RU');
}
</script>

<style scoped>
.charges-container {
  max-width: 900px;
  margin: 3rem auto;
  padding: 2.5rem;
  background: linear-gradient(to bottom right, #f9fafb, #e0f2fe);
  border-radius: 16px;
  box-shadow: 0 10px 35px rgba(0, 0, 0, 0.05);
  animation: fadeIn 0.3s ease-in-out;
}

.charges-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 2rem;
}

h2 {
  font-size: 26px;
  color: #0f172a;
}

.add-button {
  padding: 0.6rem 1.2rem;
  background: #3b82f6;
  color: white;
  border-radius: 8px;
  text-decoration: none;
  font-weight: 600;
  transition: background 0.3s;
}
.add-button:hover {
  background: #2563eb;
}

.loading,
.empty-state {
  text-align: center;
  color: #64748b;
  font-size: 16px;
  margin-top: 2rem;
}

.card-grid {
  display: grid;
  gap: 1.5rem;
  grid-template-columns: repeat(auto-fit, minmax(240px, 1fr));
}

.charge-card {
  background: #f1f5f9;
  padding: 1.5rem;
  border-radius: 12px;
  box-shadow: 0 4px 15px rgba(0, 0, 0, 0.04);
  display: flex;
  flex-direction: column;
  align-items: flex-start;
  transition: transform 0.2s ease;
}
.charge-card:hover {
  transform: translateY(-4px);
}

.charge-icon {
  font-size: 24px;
  margin-bottom: 0.5rem;
}

.category {
  font-size: 18px;
  color: #1e293b;
  margin-bottom: 0.5rem;
}

.amount {
  font-weight: bold;
  color: #0f766e;
  margin-bottom: 0.3rem;
}

.date {
  font-size: 14px;
  color: #64748b;
}

@keyframes fadeIn {
  from {
    opacity: 0;
    transform: translateY(20px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}
</style>
