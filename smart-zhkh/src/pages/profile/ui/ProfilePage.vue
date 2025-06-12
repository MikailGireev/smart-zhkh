<script setup lang="ts">
import { ref, onMounted, computed } from 'vue';
import { useAuthStore } from '@/shared/store/auth';
import { fetchAccounts } from '@/shared/api/accounts';

const auth = useAuthStore();
const accounts = ref<any[]>([]);
const isLoading = ref(true);

onMounted(async () => {
  try {
    accounts.value = await fetchAccounts();
  } catch (e) {
    console.error('Ошибка загрузки счетов:', e);
  } finally {
    isLoading.value = false;
  }
});

const filteredAccounts = computed(() =>
  accounts.value.filter((acc) => String(acc.user_id) === String(auth.userId)),
);
</script>

<template>
  <div class="profile-container container">
    <h2 class="profile-title">
      <svg viewBox="0 0 24 24" class="profile-icon" aria-hidden="true">
        <path d="M12 12a5 5 0 1 0-5-5 5 5 0 0 0 5 5zm0 2c-4 0-7 3-7 5v3h14v-3c0-2-3-5-7-5z" />
      </svg>
      Профиль
    </h2>

    <div class="user-info">
      <p><strong>Пользователь:</strong> {{ auth.username }}</p>
    </div>

    <h3 class="section-title">Ваши счета</h3>

    <div v-if="isLoading" class="loading">Загрузка...</div>

    <div v-else-if="filteredAccounts.length === 0" class="empty-state">
      <p>У вас пока нет заполненного профиля.</p>
      <RouterLink to="/accounts/add" class="btn btn-primary mt-4">
        📝 Заполнить профиль
      </RouterLink>
    </div>

    <div v-else class="account-list">
      <div v-for="acc in filteredAccounts" :key="acc.id" class="account-card">
        <svg viewBox="0 0 24 24" class="card-icon" aria-hidden="true">
          <path d="M3 6h18v12H3z" />
          <path d="M3 10h18" />
        </svg>
        <p><strong>Номер счёта:</strong> {{ acc.account_num }}</p>
        <p><strong>ФИО:</strong> {{ acc.full_name }}</p>
        <p><strong>Адрес:</strong> {{ acc.address }}</p>
        <p><strong>Площадь:</strong> {{ acc.area }} м²</p>
      </div>
    </div>
  </div>
</template>

<style scoped>
.profile-container {
  margin: 2rem auto;
  padding: 2.5rem 1.5rem;
  background: var(--color-bg-light);
  border-radius: 1.5rem;
  box-shadow: var(--shadow-lg);
  animation: fadeIn 0.5s ease;
}

.profile-title {
  display: flex;
  align-items: center;
  gap: 0.75rem;
  font-size: 1.75rem;
  color: var(--color-primary-dark);
  margin-bottom: 1rem;
}

.profile-icon {
  width: 2.5rem;
  height: 2.5rem;
  fill: var(--color-primary);
  transition: var(--transition-default);
}
.profile-icon:hover {
  transform: scale(1.1);
}

.user-info {
  background: var(--color-bg-dark);
  color: var(--color-text-light);
  padding: 1rem;
  border-radius: 0.75rem;
  border-left: 4px solid var(--color-primary);
  margin-bottom: 2rem;
}

.section-title {
  font-size: 1.25rem;
  color: var(--color-text-dark);
  margin-bottom: 1rem;
}

.loading {
  text-align: center;
  color: var(--color-text-dark);
  font-weight: 500;
}

.empty-state {
  text-align: center;
  padding: 2rem;
  background: var(--color-bg-light);
  border: 2px dashed var(--color-warning);
  border-radius: 1rem;
  color: var(--color-warning);
}

.account-list {
  display: grid;
  gap: 1.5rem;
  grid-template-columns: repeat(auto-fit, minmax(260px, 1fr));
}

.account-card {
  background: var(--color-text-light);
  padding: 1.5rem;
  border-radius: 1rem;
  box-shadow: var(--shadow-md);
  transition: var(--transition-default);
  display: flex;
  flex-direction: column;
  gap: 0.5rem;
}
.account-card:hover {
  transform: translateY(-4px);
  box-shadow: var(--shadow-lg);
}

.card-icon {
  width: 1.75rem;
  height: 1.75rem;
  margin-bottom: 0.75rem;
  stroke: var(--color-primary);
  stroke-width: 2;
  fill: none;
  transition: var(--transition-default);
}
.account-card:hover .card-icon {
  transform: scale(1.1);
}

.btn {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  padding: 0.5rem 1rem;
  border-radius: 0.75rem;
  font-weight: 500;
}

.btn-primary {
  background-color: var(--color-primary);
  color: var(--color-text-light);
  transition: var(--transition-default);
}
.btn-primary:hover {
  background-color: var(--color-primary-dark);
}

.mt-4 {
  margin-top: 1rem;
}

/* Анимации */
@keyframes fadeIn {
  from {
    opacity: 0;
    transform: translateY(12px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

/* Адаптивность */
@media (max-width: 768px) {
  .profile-container {
    padding: 2rem 1rem;
  }
  .profile-title {
    font-size: 1.5rem;
  }
  .account-list {
    gap: 1rem;
  }
}

@media (max-width: 480px) {
  .profile-title {
    font-size: 1.25rem;
  }
  .profile-icon {
    width: 2rem;
    height: 2rem;
  }
  .account-card {
    padding: 1rem;
  }
}
</style>
