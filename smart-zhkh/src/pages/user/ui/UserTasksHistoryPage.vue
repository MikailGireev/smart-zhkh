<template>
  <div class="user-history-page">
    <h2>📜 История заявок</h2>

    <div v-if="loading" class="status">Загрузка...</div>
    <div v-if="error" class="status error">{{ error }}</div>

    <table v-if="tasks.length" class="tasks-table">
      <thead>
        <tr>
          <th>ID</th>
          <th>Категория</th>
          <th>Заголовок</th>
          <th>Статус</th>
          <th>Создана</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="task in tasks" :key="task.id">
          <td>{{ task.id }}</td>
          <td>{{ task.category }}</td>
          <td>{{ task.title }}</td>
          <td>{{ formatStatus(task.status) }}</td>
          <td>{{ formatDate(task.created_at) }}</td>
        </tr>
      </tbody>
    </table>

    <p v-if="!loading && !tasks.length" class="status">Заявок не найдено</p>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue';
import { useAuthStore } from '@/shared/store/auth';

const auth = useAuthStore();
const tasks = ref<any[]>([]);
const loading = ref(true);
const error = ref('');

async function loadMyTasks() {
  try {
    const res = await fetch(`http://localhost:8081/api/v1/tasks?account_id=${auth.userId}`);
    if (!res.ok) throw new Error();
    tasks.value = await res.json();
  } catch {
    error.value = '❌ Не удалось загрузить заявки';
  } finally {
    loading.value = false;
  }
}

function formatDate(str: string) {
  return new Date(str).toLocaleString();
}

function formatStatus(status: string) {
  switch (status) {
    case 'created':
      return 'Создана';
    case 'in_progress':
      return 'В работе';
    case 'done':
      return 'Завершена';
    default:
      return status;
  }
}

onMounted(loadMyTasks);
</script>

<style scoped>
.user-history-page {
  background: white;
  padding: 2rem;
  border-radius: 1rem;
  box-shadow: 0 10px 30px rgba(0, 0, 0, 0.05);
  font-family: 'Segoe UI', sans-serif;
}

h2 {
  font-size: 1.5rem;
  font-weight: bold;
  margin-bottom: 1.5rem;
}

.tasks-table {
  width: 100%;
  border-collapse: collapse;
  margin-top: 1rem;
}

.tasks-table th,
.tasks-table td {
  border: 1px solid #ddd;
  padding: 0.75rem;
  text-align: left;
}

.status {
  margin-top: 1rem;
  font-weight: 500;
}

.status.error {
  color: red;
}
</style>
