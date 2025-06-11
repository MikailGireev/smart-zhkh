<template>
  <div class="profile-container">
    <h2>👤 Профиль</h2>

    <div class="user-info">
      <p><strong>Пользователь:</strong> {{ auth.username }}</p>
    </div>

    <h3>Ваши счета</h3>

    <div v-if="isLoading" class="loading">Загрузка...</div>

    <div v-else-if="filteredAccounts.length === 0" class="empty-state">
      <p>У вас пока нет заполненного профиля.</p>
      <RouterLink to="/accounts/add" class="add-profile-button">📝 Заполнить профиль</RouterLink>
    </div>

    <div v-else class="account-list">
      <div v-for="acc in filteredAccounts" :key="acc.id" class="account-card">
        <p><strong>Номер счёта:</strong> {{ acc.account_num }}</p>
        <p><strong>ФИО:</strong> {{ acc.full_name }}</p>
        <p><strong>Адрес:</strong> {{ acc.address }}</p>
        <p><strong>Площадь:</strong> {{ acc.area }} м²</p>
        <RouterLink :to="`/accounts/${acc.id}/edit`" class="edit-link">✏️ Редактировать</RouterLink>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, computed } from 'vue';
import { useAuthStore } from '@/shared/store/auth';
import { fetchAccounts } from '@/shared/api/accounts';

const auth = useAuthStore();
const accounts = ref<any[]>([]);
const isLoading = ref(true);

onMounted(async () => {
  try {
    const data = await fetchAccounts();
    accounts.value = data;
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

<style scoped>
.profile-container {
  max-width: 800px;
  margin: 3rem auto;
  padding: 2.5rem;
  background: linear-gradient(to bottom right, #f9fafb, #e0f2fe);
  border-radius: 16px;
  box-shadow: 0 15px 40px rgba(0, 0, 0, 0.05);
  animation: fadeIn 0.4s ease;
}

h2 {
  font-size: 28px;
  color: #0f172a;
  margin-bottom: 1.5rem;
}

h3 {
  font-size: 22px;
  margin-bottom: 1rem;
  color: #1e293b;
}

.user-info {
  margin-bottom: 2rem;
  background: #f1f5f9;
  padding: 1rem;
  border-radius: 8px;
  color: #334155;
  border-left: 5px solid #2563eb;
}

.loading {
  text-align: center;
  font-weight: 500;
  color: #64748b;
}

.empty-state {
  text-align: center;
  padding: 2rem;
  background: #fff1f2;
  border: 1px dashed #fca5a5;
  border-radius: 12px;
  color: #b91c1c;
}

.add-profile-button {
  display: inline-block;
  margin-top: 1rem;
  padding: 0.6rem 1.2rem;
  background: #10b981;
  color: white;
  border-radius: 8px;
  text-decoration: none;
  font-weight: 600;
  transition: background 0.3s;
}

.add-profile-button:hover {
  background: #059669;
}

.account-list {
  display: grid;
  grid-template-columns: 1fr;
  gap: 1.5rem;
}

.account-card {
  background: #f8fafc;
  padding: 1.5rem;
  border-radius: 10px;
  box-shadow: 0 3px 10px rgba(0, 0, 0, 0.03);
  border: 1px solid #e2e8f0;
  transition: transform 0.2s;
}

.account-card:hover {
  transform: translateY(-4px);
}

.edit-link {
  display: inline-block;
  margin-top: 0.75rem;
  color: #3b82f6;
  font-weight: 500;
  font-size: 0.95rem;
  transition: color 0.3s;
}

.edit-link:hover {
  color: #1d4ed8;
  text-decoration: underline;
}

@keyframes fadeIn {
  from {
    opacity: 0;
    transform: translateY(15px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}
</style>
