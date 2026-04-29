<template>
  <el-container class="layout">
    <el-header class="header">
      <div class="header-left">
        <router-link to="/" class="logo">QQZone</router-link>
        <router-link to="/friends" v-if="userStore.isLoggedIn" class="nav-link">好友</router-link>
      </div>
      <div class="header-right">
        <template v-if="userStore.isLoggedIn">
          <el-button type="primary" @click="$router.push('/create')">
            <el-icon><Plus /></el-icon>
            发文
          </el-button>
          <el-dropdown>
            <span class="user-dropdown">
              {{ userStore.userInfo?.username }}
              <el-icon><ArrowDown /></el-icon>
            </span>
            <template #dropdown>
              <el-dropdown-menu>
                <el-dropdown-item v-if="userStore.isAdmin" @click="$router.push('/admin')">
                  <el-icon><Setting /></el-icon>管理后台
                </el-dropdown-item>
                <el-dropdown-item @click="handleLogout">
                  <el-icon><SwitchButton /></el-icon>退出登录
                </el-dropdown-item>
              </el-dropdown-menu>
            </template>
          </el-dropdown>
        </template>
        <template v-else>
          <el-button text @click="$router.push('/login')">登录</el-button>
          <el-button type="primary" @click="$router.push('/register')">注册</el-button>
        </template>
      </div>
    </el-header>

    <el-main class="main">
      <router-view />
    </el-main>

    <el-footer class="footer">
      QQZone © 2025
    </el-footer>
  </el-container>
</template>

<script setup lang="ts">
import { useUserStore } from '@/stores/user'
import { useRouter } from 'vue-router'
import { logout as apiLogout } from '@/api/user'

const userStore = useUserStore()
const router = useRouter()

async function handleLogout() {
  try {
    await apiLogout()
  } catch {
  }
  userStore.logout()
  router.push('/login')
}
</script>

<style scoped>
.layout {
  min-height: 100vh;
}

.header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  background: #fff;
  box-shadow: 0 1px 4px rgba(0, 0, 0, 0.08);
  padding: 0 24px;
  position: sticky;
  top: 0;
  z-index: 100;
}

.logo {
  font-size: 22px;
  font-weight: 700;
  color: #409eff;
}

.header-left {
  display: flex;
  align-items: center;
  gap: 32px;
}

.nav-link {
  font-size: 15px;
  color: #555;
  transition: color 0.2s;
}

.nav-link:hover {
  color: #409eff;
}

.header-right {
  display: flex;
  align-items: center;
  gap: 16px;
}

.user-dropdown {
  cursor: pointer;
  display: flex;
  align-items: center;
  gap: 4px;
  font-size: 14px;
}

.main {
  min-height: calc(100vh - 120px);
  padding: 24px;
}

.footer {
  text-align: center;
  color: #999;
  font-size: 13px;
  padding: 16px;
  border-top: 1px solid #eee;
}
</style>
