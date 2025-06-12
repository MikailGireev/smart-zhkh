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
        <div class="card-topbar"></div>
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

<style scoped>
.charges-container {
  position: relative;
  overflow: hidden;
  margin: 3rem auto;
  padding: 3rem 2rem;
  background: rgba(255, 255, 255, 0.8);
  backdrop-filter: blur(12px);
  border-radius: 1.75rem;
  box-shadow:
    0 12px 30px -8px rgba(0, 0, 0, 0.2),
    inset 0 2px 4px rgba(255, 255, 255, 0.6);
  border: 1px solid rgba(255, 255, 255, 0.4);
  animation: cardEntrance 0.8s cubic-bezier(0.22, 1, 0.36, 1);
  color: var(--color-text-dark);
}
.charges-container::before {
  content: '';
  position: absolute;
  top: -40%;
  left: -40%;
  width: 180%;
  height: 180%;
  background: radial-gradient(circle at center, rgba(37,99,235,0.1) 0%, transparent 70%);
  animation: rotate 30s linear infinite;
  z-index: -1;
}

.charges-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 2rem;
}
.charges-title {
  font-size: 2rem;
  font-weight: 800;
  position: relative;
  display: inline-block;
}
.charges-title::after {
  content: '';
  position: absolute;
  bottom: -6px;
  left: 0;
  width: 60px;
  height: 4px;
  background: var(--color-primary-light);
  border-radius: 2px;
  box-shadow: 0 0 8px var(--color-primary-light);
}

.add-button {
  display: inline-flex;
  align-items: center;
  gap: 0.5rem;
}
.btn-icon {
  width: 1.25rem;
  height: 1.25rem;
  stroke: currentColor;
  stroke-width: 2;
  fill: none;
  transition: transform 0.3s;
}
.add-button:hover .btn-icon {
  transform: scale(1.2);
}

.loading,
.empty-state {
  text-align: center;
  padding: 2rem 0;
  font-size: 1rem;
  color: var(--color-text-dark);
}

.card-grid {
  display: grid;
  gap: 2rem;
  grid-template-columns: repeat(auto-fit, minmax(240px, 1fr));
}

.charge-card {
  position: relative;
  background: rgba(255, 255, 255, 0.9);
  backdrop-filter: blur(8px);
  border-radius: 1.25rem;
  box-shadow:
    0 6px 16px -4px rgba(0, 0, 0, 0.1),
    inset 0 1px 3px rgba(255, 255, 255, 0.5);
  padding: 2rem 1.5rem 1.5rem;
  transition: transform 0.4s cubic-bezier(0.22, 1, 0.36, 1),
              box-shadow 0.4s;
  overflow: hidden;
  color: var(--color-text-dark);
  text-align: center;
}
.charge-card:hover {
  transform: translateY(-6px) scale(1.02);
  box-shadow:
    0 12px 24px -6px rgba(0, 0, 0, 0.15),
    inset 0 1px 4px rgba(255, 255, 255, 0.6);
}

.card-topbar {
  position: absolute;
  top: 0; left: 0; right: 0;
  height: 4px;
  background: linear-gradient(
    90deg,
    var(--color-primary-light),
    var(--color-primary-dark)
  );
}

.charge-icon {
  width: 2rem;
  height: 2rem;
  margin-bottom: 1rem;
  stroke: var(--color-primary);
  stroke-width: 2;
  fill: none;
  transition: transform 0.3s, filter 0.3s;
  filter: drop-shadow(0 2px 4px rgba(37, 99, 235, 0.2));
}
.charge-card:hover .charge-icon {
  transform: scale(1.2) translateY(-4px);
  filter: drop-shadow(0 4px 8px rgba(37, 99, 235, 0.3));
}

.category {
  font-size: 1.25rem;
  font-weight: 700;
  color: var(--color-primary-dark);
  margin-bottom: 0.5rem;
}

.amount {
  font-size: 1.125rem;
  font-weight: 600;
  color: var(--color-success);
  margin-bottom: 0.5rem;
}

.date {
  font-size: 0.875rem;
  color: var(--color-text-dark);
}

@keyframes cardEntrance {
  from { opacity: 0; transform: translateY(30px) scale(0.95); }
  to   { opacity: 1; transform: translateY(0) scale(1); }
}
@keyframes rotate {
  from { transform: rotate(0deg); }
  to   { transform: rotate(360deg); }
}

/* Responsive */
@media (max-width: 768px) {
  .charges-container {
    padding: 2rem 1rem;
  }
  .charges-title {
    font-size: 1.5rem;
  }
  .card-grid {
    gap: 1.5rem;
  }
}
@media (max-width: 480px) {
  .charge-card {
    padding: 1.5rem 1rem 1rem;
  }
  .category {
    font-size: 1.125rem;
  }
  .amount {
    font-size: 1rem;
  }
  .date {
    font-size: 0.8rem;
  }
}
</style>
