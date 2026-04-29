<template>
  <div class="friends-page">
    <el-card shadow="hover">
      <template #header>
        <span>我的好友</span>
      </template>

      <div class="add-friend">
        <el-input v-model.number="addId" placeholder="输入好友ID" style="width: 200px" size="small" />
        <el-button type="primary" size="small" :loading="adding" @click="handleAdd">
          添加好友
        </el-button>
      </div>

      <el-table v-if="friends.length" :data="friends" style="width: 100%; margin-top: 16px">
        <el-table-column prop="id" label="ID" width="80" />
        <el-table-column prop="username" label="用户名" />
        <el-table-column prop="role" label="角色" width="100" />
        <el-table-column label="操作" width="100">
          <template #default="{ row }">
            <el-button type="danger" size="small" text @click="handleDelete(row.id)">
              删除
            </el-button>
          </template>
        </el-table-column>
      </el-table>

      <el-empty v-else description="还没有好友" :image-size="80" />
    </el-card>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { listFriends, addFriend, deleteFriend } from '@/api/user'
import { ElMessage, ElMessageBox } from 'element-plus'
import type { UserInfo } from '@/api/user'

const friends = ref<UserInfo[]>([])
const addId = ref<number | null>(null)
const adding = ref(false)

async function fetchFriends() {
  try {
    const res = await listFriends()
    friends.value = res.data.friends || []
  } catch {}
}

async function handleAdd() {
  if (!addId.value) {
    ElMessage.warning('请输入好友ID')
    return
  }
  adding.value = true
  try {
    await addFriend(addId.value)
    ElMessage.success('添加成功')
    addId.value = null
    await fetchFriends()
  } finally {
    adding.value = false
  }
}

async function handleDelete(id: number) {
  await ElMessageBox.confirm('确定要删除该好友吗？', '确认删除', { type: 'warning' })
  try {
    await deleteFriend(id)
    ElMessage.success('删除成功')
    await fetchFriends()
  } catch {}
}

onMounted(fetchFriends)
</script>

<style scoped>
.friends-page {
  max-width: 720px;
  margin: 0 auto;
}

.add-friend {
  display: flex;
  gap: 8px;
  align-items: center;
}
</style>
