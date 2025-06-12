<template>
  <div class="charges-container container">
    <div class="charges-header">
      <h2 class="charges-title">Ваши начисления</h2>
      <RouterLink to="/charges/add" class="btn btn-primary add-button">
        <svg viewBox="0 0 24 24" class="btn-icon" aria-hidden="true">
          <line x1="12" y1="5" x2="12" y2="19" />
          <line x1="5" y1="12" x2="19" y2="12" />
        </svg>
        <span>Добавить</span>
      </RouterLink>
    </div>

    <div v-if="isLoading" class="loading">⏳ Загрузка...</div>

    <div v-else-if="filteredCharges.length === 0" class="empty-state">
      <p>Пока нет начислений</p>
    </div>

    <div v-else class="card-grid">
      <div v-for="charge in filteredCharges" :key="charge.id" class="charge-card">
        <svg viewBox="0 0 24 24" class="charge-icon" aria-hidden="true">
          <path d="M3 3h18v18H3z" />
          <path d="M3 7h18" />
        </svg>
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
    charges.value = await fetchCharges(auth.userId);
  } catch (e) {
    console.error('Ошибка загрузки начислений:', e);
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
  margin: 3rem auto;
  padding: 2.5rem 1.5rem;
  background: var(--color-bg-light);
  border-radius: 1.5rem;
  box-shadow: var(--shadow-lg);
  animation: fadeIn 0.4s ease;
}

.charges-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 1.75rem;
}
.charges-title {
  font-size: 1.75rem;
  color: var(--color-primary-dark);
}

/* Add-button uses btn and btn-primary */
.btn-icon {
  width: 1rem;
  height: 1rem;
  stroke: currentColor;
  stroke-width: 2;
  stroke-linecap: round;
  stroke-linejoin: round;
  margin-right: 0.5rem;
  fill: none;
}

/* Loading & Empty State */
.loading,
.empty-state {
  text-align: center;
  color: var(--color-text-dark);
  font-size: 1rem;
  padding: 2rem 0;
}

/* Cards Grid */
.card-grid {
  display: grid;
  gap: 1.5rem;
  grid-template-columns: repeat(auto-fit, minmax(240px, 1fr));
}

.charge-card {
  background: var(--color-text-light);
  padding: 1.5rem;
  border-radius: 1rem;
  box-shadow: var(--shadow-md);
  transition: var(--transition-default);
  display: flex;
  flex-direction: column;
  gap: 0.75rem;
}
.charge-card:hover {
  transform: translateY(-4px);
  box-shadow: var(--shadow-lg);
}

/* Charge Icon */
.charge-icon {
  width: 1.5rem;
  height: 1.5rem;
  stroke: var(--color-primary);
  stroke-width: 2;
  fill: none;
  transition: var(--transition-default);
}
.charge-card:hover .charge-icon {
  transform: scale(1.1);
}

/* Text Elements */
.category {
  font-size: 1.125rem;
  color: var(--color-text-dark);
  font-weight: 600;
}
.amount {
  font-size: 1rem;
  color: var(--color-success);
  font-weight: 700;
}
.date {
  font-size: 0.875rem;
  color: var(--color-text-dark);
}

/* Animations */
@keyframes fadeIn {
  from {
    opacity: 0;
    transform: translateY(16px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

/* Responsive */
@media (max-width: 768px) {
  .charges-container {
    padding: 2rem 1rem;
  }
  .charges-title {
    font-size: 1.5rem;
  }
}

@media (max-width: 480px) {
  .charge-card {
    padding: 1.25rem;
  }
  .category {
    font-size: 1rem;
  }
  .amount {
    font-size: 0.95rem;
  }
  .date {
    font-size: 0.8rem;
  }
}
</style>
