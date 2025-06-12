<template>
  <div class="app-layout">
    <header class="app-header">
      <div class="header-container container">
        <h1 class="app-logo" @click="goHome">
          <span class="logo-icon">🧠</span>
          <span class="logo-text">IoT Dashboard</span>
        </h1>

        <button
          class="mobile-menu-button"
          @click.stop="isMobileMenuOpen = !isMobileMenuOpen"
          :aria-expanded="isMobileMenuOpen"
          aria-label="Toggle navigation"
        >
          <svg viewBox="0 0 24 24" class="menu-icon" v-if="!isMobileMenuOpen">
            <path d="M3 12h18M3 6h18M3 18h18" />
          </svg>
          <svg viewBox="0 0 24 24" class="menu-icon" v-else>
            <path d="M6 6l12 12M6 18L18 6" />
          </svg>
        </button>

        <nav
          v-if="auth.isLoggedIn"
          ref="navRef"
          class="main-navigation"
          :class="{ 'mobile-open': isMobileMenuOpen }"
        >
          <RouterLink to="/dashboard" class="nav-item" active-class="router-link-active">
            <svg viewBox="0 0 24 24" class="nav-icon">
              <path d="M3 9l9-7 9 7v11a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2z" />
              <polyline points="9 22 9 12 15 12 15 22" />
            </svg>
            <span class="nav-text" title="Главная">Главная</span>
          </RouterLink>
          <RouterLink to="/charges" class="nav-item" active-class="router-link-active">
            <svg viewBox="0 0 24 24" class="nav-icon">
              <path d="M14 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8z" />
              <polyline points="14 2 14 8 20 8" />
              <line x1="16" y1="13" x2="8" y2="13" />
              <line x1="16" y1="17" x2="8" y2="17" />
              <polyline points="10 9 9 9 8 9" />
            </svg>
            <span class="nav-text" title="Начисления">Начисления</span>
          </RouterLink>
          <RouterLink to="/settings" class="nav-item" active-class="router-link-active">
            <svg viewBox="0 0 24 24" class="nav-icon">
              <circle cx="12" cy="12" r="3" />
              <path
                d="M19.4 15a1.65 1.65 0 0 0 .33 1.82l.06.06a2 2 0 0 1 0
                  2.83 2 2 0 0 1-2.83 0l-.06-.06a1.65
                  1.65 0 0 0-1.82-.33 1.65 1.65 0 0
                  0-1 1.51V21a2 2 0 0 1-2 2 2 2 0 0
                  1-2-2v-.09A1.65 1.65 0 0 0 9
                  19.4a1.65 1.65 0 0 0-1.82.33l-.06.06a2
                  2 0 0 1-2.83 0 2 2 0 0 1 0-2.83l.06-.06a1.65
                  1.65 0 0 0 .33-1.82 1.65 1.65 0 0
                  0-1.51-1H3a2 2 0 0 1-2-2 2 2 0 0
                  1 2-2h.09A1.65 1.65 0 0 0 4.6
                  9a1.65 1.65 0 0 0-.33-1.82l-.06-.06a2
                  2 0 0 1 0-2.83 2 2 0 0 1 2.83 0l.06.06a1.65
                  1.65 0 0 0 1.82.33H9a1.65 1.65 0 0 0
                  1-1.51V3a2 2 0 0 1 2-2 2 2 0 0
                  1 2 2v.09a1.65 1.65 0 0 0 1
                  1.51 1.65 1.65 0 0 0 1.82-.33l.06-.06a2
                  2 0 0 1 2.83 0 2 2 0 0
                  1 0 2.83l-.06.06a1.65 1.65 0 0
                  0-.33 1.82V9a1.65 1.65 0 0 0 1.51
                  1H21a2 2 0 0 1 2 2 2 2 0 0
                  1-2 2h-.09a1.65 1.65 0 0 0-1.51
                  1z"
              />
            </svg>
            <span class="nav-text" title="Настройки">Настройки</span>
          </RouterLink>
          <RouterLink to="/profile" class="nav-item" active-class="router-link-active">
            <svg viewBox="0 0 24 24" class="nav-icon">
              <path d="M20 21v-2a4 4 0 0 0-4-4H8a4 4 0 0 0-4 4v2" />
              <circle cx="12" cy="7" r="4" />
            </svg>
            <span class="nav-text" title="Профиль">Профиль</span>
          </RouterLink>
          <button @click="logout" class="logout-button">
            <svg viewBox="0 0 24 24" class="logout-icon">
              <path d="M9 21H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h4" />
              <polyline points="16 17 21 12 16 7" />
              <line x1="21" y1="12" x2="9" y2="12" />
            </svg>
            <span class="logout-text">Выйти</span>
          </button>
        </nav>

        <nav
          v-else
          ref="navRef"
          class="auth-navigation"
          :class="{ 'mobile-open': isMobileMenuOpen }"
        >
          <RouterLink to="/login" class="auth-button login-button">
            <span>Вход</span>
          </RouterLink>
          <RouterLink to="/register" class="auth-button register-button">
            <span>Регистрация</span>
          </RouterLink>
        </nav>
      </div>
    </header>

    <main class="app-main">
      <RouterView />
    </main>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, onBeforeUnmount } from 'vue';
import { useRouter } from 'vue-router';
import { useAuthStore } from '@/shared/store/auth';

const auth = useAuthStore();
const router = useRouter();
const isMobileMenuOpen = ref(false);
const navRef = ref<HTMLElement | null>(null);

function logout() {
  auth.logout();
  router.push('/login');
  isMobileMenuOpen.value = false;
}

function goHome() {
  router.push(auth.isLoggedIn ? '/' : '/login');
  isMobileMenuOpen.value = false;
}

// Закрытие мобильного меню при клике вне
function handleClickOutside(event: MouseEvent) {
  if (navRef.value && !navRef.value.contains(event.target as Node)) {
    isMobileMenuOpen.value = false;
  }
}

onMounted(() => {
  document.addEventListener('click', handleClickOutside);
});

onBeforeUnmount(() => {
  document.removeEventListener('click', handleClickOutside);
});
</script>

<style scoped>
/* Layout */
.app-layout {
  min-height: 100vh;
  display: flex;
  flex-direction: column;
}

/* Header */
.app-header {
  background: linear-gradient(135deg, var(--color-primary-dark) 0%, var(--color-primary) 100%);
  color: var(--color-text-light);
  padding: 0.75rem 0;
  box-shadow: var(--shadow-lg);
  position: sticky;
  top: 0;
  z-index: 50;
  backdrop-filter: blur(8px);
  border-bottom: 1px solid rgba(255, 255, 255, 0.1);
}

/* Container override */
.header-container.container {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

/* Logo */
.app-logo {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  font-size: 1.5rem;
  font-weight: 700;
  cursor: pointer;
  transition: var(--transition-default);
  color: var(--color-text-light);
}
.app-logo:hover {
  opacity: 0.9;
}
.logo-icon {
  font-size: 1.8rem;
}
.logo-text {
  display: inline;
}

/* Navigation */
.main-navigation,
.auth-navigation {
  display: flex;
  align-items: center;
  gap: 0.75rem;
}

.nav-item {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  color: var(--color-text-light);
  padding: 0.5rem 1rem;
  border-radius: 8px;
  font-weight: 500;
  transition: var(--transition-default);
  position: relative;
}
.nav-item:hover {
  background: rgba(255, 255, 255, 0.1);
}

/* Active link indicator */
.nav-item.router-link-active {
  background: rgba(255, 255, 255, 0.15);
}
.nav-item.router-link-active::before {
  content: '';
  position: absolute;
  left: 0;
  top: 0.5rem;
  bottom: 0.5rem;
  width: 4px;
  background-color: var(--color-text-light);
  border-radius: 2px;
  box-shadow: 0 0 4px rgba(255, 255, 255, 0.6);
  transition: var(--transition-default);
}

/* SVG icons */
.nav-icon,
.logout-icon {
  width: 1.25rem;
  height: 1.25rem;
  stroke: currentColor;
  stroke-width: 2;
  stroke-linecap: round;
  stroke-linejoin: round;
  fill: none;
}

/* Logout button */
.logout-button {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  color: var(--color-text-light);
  background: rgba(255, 255, 255, 0.1);
  padding: 0.5rem 1rem;
  border-radius: 8px;
  font-weight: 500;
  transition: var(--transition-default);
}
.logout-button:hover {
  background: var(--color-error);
}

/* Auth buttons */
.auth-button {
  padding: 0.5rem 1.25rem;
  border-radius: 8px;
  font-weight: 500;
  transition: var(--transition-default);
}
.login-button {
  color: var(--color-text-light);
  border: 1px solid rgba(255, 255, 255, 0.3);
}
.login-button:hover {
  background: rgba(255, 255, 255, 0.1);
}
.register-button {
  background: var(--color-text-light);
  color: var(--color-primary-dark);
}
.register-button:hover {
  background: rgba(255, 255, 255, 0.9);
  box-shadow: 0 2px 8px rgba(255, 255, 255, 0.2);
}

/* Main content */
.app-main {
  flex: 1;
  padding: 2rem;
  background: var(--color-bg-light);
  min-height: calc(100vh - 60px);
}

/* Mobile menu button */
.mobile-menu-button {
  display: none;
  background: none;
  border: none;
  color: var(--color-text-light);
  padding: 0.5rem;
  cursor: pointer;
}
.menu-icon {
  width: 1.5rem;
  height: 1.5rem;
  stroke: currentColor;
  stroke-width: 2;
  stroke-linecap: round;
  stroke-linejoin: round;
}

/* Responsive */
@media (max-width: 768px) {
  .main-navigation,
  .auth-navigation {
    display: none;
    position: absolute;
    top: 100%;
    left: 0;
    right: 0;
    background: var(--color-primary-dark);
    flex-direction: column;
    gap: 0;
    padding: 0.5rem 1rem;
    box-shadow: var(--shadow-md);
    z-index: 40;
  }
  .main-navigation.mobile-open,
  .auth-navigation.mobile-open {
    display: flex;
  }
  .nav-item,
  .logout-button,
  .auth-button {
    width: 100%;
    padding: 0.75rem 1rem;
    border-radius: 0;
    justify-content: center;
  }
  .mobile-menu-button {
    display: block;
  }
  .app-main {
    padding: 1.5rem 1rem;
  }
}

@media (max-width: 480px) {
  .logo-text {
    display: none;
  }
  .logo-icon {
    font-size: 2rem;
  }
}
</style>
