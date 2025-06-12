<template>
  <div class="dashboard-container container">
    <h2 class="dashboard-title">👋 Привет, {{ auth.username }}</h2>
    <p class="dashboard-subtitle">Добро пожаловать в личный кабинет</p>

    <div class="card-grid">
      <RouterLink to="/charges" class="info-card">
        <svg viewBox="0 0 24 24" class="card-icon" aria-hidden="true">
          <path
            d="M20 7H4a1 1 0 0 0-1 1v8a1 1 0 0 0 1 1h16a1 1 0 0 0 1-1V8a1 1 0 0 0-1-1zm-1 2v2H5V9h14zm0 4v2H5v-2h14z"
          />
        </svg>
        <span>Мои начисления</span>
      </RouterLink>

      <div class="info-card disabled">
        <svg viewBox="0 0 24 24" class="card-icon" aria-hidden="true">
          <path d="M12 2a5 5 0 0 0-5 5v4H5l7 7 7-7h-2V7a5 5 0 0 0-5-5z" />
        </svg>
        <span>Устройства (в разработке)</span>
      </div>

      <RouterLink to="/profile" class="info-card">
        <svg viewBox="0 0 24 24" class="card-icon" aria-hidden="true">
          <path d="M12 12a5 5 0 1 0-5-5 5 5 0 0 0 5 5zm0 2c-4 0-7 2-7 4v2h14v-2c0-2-3-4-7-4z" />
        </svg>
        <span>Профиль</span>
      </RouterLink>
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
  padding: 2.5rem 1.5rem;
  background: var(--color-bg-light);
  background: linear-gradient(145deg, var(--color-bg-light) 0%, #ffffff 100%);
  border-radius: 1.5rem;
  box-shadow: var(--shadow-lg);
  animation: fadeIn 0.6s ease;
}

/* Заголовки */
.dashboard-title {
  font-size: 1.75rem;
  font-weight: 700;
  color: var(--color-primary-dark);
  margin-bottom: 0.5rem;
}

.dashboard-subtitle {
  font-size: 1rem;
  color: var(--color-text-dark);
  margin-bottom: 2rem;
}

/* Карточки */
.card-grid {
  display: grid;
  gap: 1.5rem;
  grid-template-columns: repeat(auto-fit, minmax(240px, 1fr));
}

.info-card {
  display: flex;
  flex-direction: column;
  align-items: center;
  padding: 1.75rem;
  background: var(--color-bg-light);
  border-radius: 1rem;
  box-shadow: var(--shadow-md);
  text-align: center;
  font-size: 1.125rem;
  font-weight: 600;
  color: var(--color-text-dark);
  transition: var(--transition-default);
  text-decoration: none;
  border: 2px solid transparent;
}

.info-card:hover {
  transform: translateY(-4px);
  background: var(--color-bg-light);
  border-color: var(--color-primary-light);
  box-shadow: var(--shadow-lg);
  color: var(--color-primary-dark);
}

.info-card.disabled {
  opacity: 0.6;
  cursor: not-allowed;
}
.info-card.disabled:hover {
  transform: none;
  border-color: transparent;
  box-shadow: var(--shadow-md);
  color: var(--color-text-dark);
}

/* Иконки в карточках */
.card-icon {
  width: 2.5rem;
  height: 2.5rem;
  margin-bottom: 0.75rem;
  fill: var(--color-primary);
  transition: var(--transition-default);
}
.info-card:hover .card-icon {
  transform: scale(1.1);
}

/* Анимации */
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

/* Адаптивность */
@media (max-width: 768px) {
  .dashboard-container {
    margin: 1.5rem 0.75rem;
    padding: 2rem 1rem;
  }
  .dashboard-title {
    font-size: 1.5rem;
  }
  .dashboard-subtitle {
    font-size: 0.95rem;
  }
}

@media (max-width: 480px) {
  .dashboard-container {
    padding: 1.5rem 1rem;
  }
  .card-grid {
    gap: 1rem;
  }
  .info-card {
    padding: 1.25rem;
    font-size: 1rem;
  }
  .card-icon {
    width: 2rem;
    height: 2rem;
  }
}
</style>
