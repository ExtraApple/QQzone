<template>
  <div class="admin-page">
    <el-card shadow="hover">
      <template #header>
        <span>管理后台</span>
      </template>

      <el-result
        v-if="adminData"
        icon="success"
        title="管理员身份验证通过"
        :sub-title="`欢迎，${adminData.user?.Username || adminData.user?.username || 'admin'}`"
      >
        <template #extra>
          <el-tag type="success">{{ adminData.msg }}</el-tag>
        </template>
      </el-result>

      <el-result
        v-else
        icon="warning"
        title="管理员验证中..."
        sub-title="正在确认您的管理员身份"
      />
    </el-card>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { adminCheck } from '@/api/user'

const adminData = ref<any>(null)

onMounted(async () => {
  try {
    const res = await adminCheck()
    adminData.value = res.data
  } catch {}
})
</script>

<style scoped>
.admin-page {
  max-width: 720px;
  margin: 0 auto;
}
</style>
