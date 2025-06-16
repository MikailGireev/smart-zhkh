<template>
  <div class="task-detail-page">
    <h2>🛠 Заявка №{{ task?.id }}</h2>

    <div v-if="loading">Загрузка...</div>
    <div v-if="error">{{ error }}</div>

    <div v-if="task" class="task-info">
      <p><strong>Заголовок:</strong> {{ task.title }}</p>
      <p><strong>Категория:</strong> {{ task.category }}</p>
      <p><strong>Описание:</strong> {{ task.description }}</p>

      <p>
        <strong>Статус:</strong>
        <select v-model="task.status">
          <option value="created">Создана</option>
          <option value="in_progress">В работе</option>
          <option value="done">Завершена</option>
        </select>
      </p>

      <p>
        <strong>Исполнитель:</strong>
        <select v-model="task.executor_id">
          <option disabled value="">Выбери исполнителя</option>
          <option v-for="user in users" :key="user.id" :value="user.id">
            {{ user.username }}
          </option>
        </select>
      </p>

      <button @click="saveChanges">💾 Сохранить изменения</button>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue';
import { useRoute, useRouter } from 'vue-router';

const route = useRoute();
const router = useRouter();
const taskId = route.params.id as string;

const task = ref<any>(null);
const users = ref<any[]>([]);
const loading = ref(true);
const error = ref('');

async function loadTask() {
  try {
    const res = await fetch(`http://localhost:8081/api/v1/tasks/${taskId}`);
    if (!res.ok) throw new Error('Ошибка при загрузке');
    task.value = await res.json();

    if (!('executor_id' in task.value)) {
      task.value.executor_id = '';
    }
  } catch (err) {
    error.value = '❌ Не удалось загрузить задачу';
  } finally {
    loading.value = false;
  }
}

async function loadUsers() {
  users.value = [
    {
      id: 'admin-001',
      username: 'Админ Иванов',
    },
  ];
  // try {
  //   const res = await fetch('http://localhost:8080/api/users');
  //   if (!res.ok) throw new Error();
  //   users.value = await res.json();
  // } catch {
  //   users.value = [];
  // }
}

async function saveChanges() {
  if (!task.value.executor_id) {
    alert('❗ Выберите исполнителя');
    return;
  }

  const payload = {
    status: task.value.status,
    executor_id: task.value.executor_id,
  };

  try {
    const res = await fetch(`http://localhost:8081/api/v1/tasks/${taskId}`, {
      method: 'PUT',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(payload),
    });
    if (!res.ok) throw new Error();
    alert('✅ Обновлено!');
    router.push('/admin/tasks');
  } catch {
    alert('❌ Ошибка обновления');
  }
}

onMounted(() => {
  loadTask();
  loadUsers();
});
</script>

<style scoped>
.task-detail-page {
  max-width: 600px;
  margin: 2rem auto;
  padding: 2rem;
  background: #ffffff;
  border-radius: 1rem;
  box-shadow: 0 10px 25px rgba(0, 0, 0, 0.08);
  font-family: 'Segoe UI', sans-serif;
  color: #1f2937;
}

.task-detail-page h2 {
  font-size: 1.75rem;
  font-weight: 700;
  margin-bottom: 1.5rem;
  display: flex;
  align-items: center;
  gap: 0.5rem;
}

.task-info p {
  margin-bottom: 1.2rem;
  font-size: 1rem;
  line-height: 1.5;
}

.task-info strong {
  display: inline-block;
  min-width: 120px;
  color: #374151;
}

.task-info select {
  padding: 0.5rem 0.75rem;
  font-size: 1rem;
  border-radius: 0.5rem;
  border: 1px solid #d1d5db;
  background-color: #f9fafb;
  transition: border-color 0.2s ease;
}

.task-info select:focus {
  border-color: #3b82f6;
  outline: none;
  background-color: #fff;
}

button {
  margin-top: 1.5rem;
  padding: 0.75rem 1.5rem;
  background-color: #2563eb;
  color: white;
  font-weight: 600;
  font-size: 1rem;
  border: none;
  border-radius: 0.75rem;
  cursor: pointer;
  transition: background-color 0.2s ease;
}

button:hover {
  background-color: #1e40af;
}

button:disabled {
  opacity: 0.6;
  cursor: not-allowed;
}

.loading,
.error {
  margin-top: 1rem;
  font-weight: 600;
  font-size: 1rem;
}

.error {
  color: #dc2626;
}
</style>
