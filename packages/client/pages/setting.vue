<template>
  <div>
    <PageTitle title="設定" />

    <v-divider></v-divider>

    <v-data-table :items="menuItems" hide-default-footer hide-default-header>
      <template v-slot:item="{ item }">
        <tr :class="{ highlight: item.type === 'header' }">
          <td>{{ item.name }}</td>
          <td>
            <div v-if="item.key">
              <v-btn v-if="item.key === 'board-history-delete'" small>削除</v-btn>
              <v-btn v-else-if="item.key === 'thread-history-delete'" small>削除</v-btn>
              <v-switch v-else-if="item.key === 'display-latest-50-replies-only'" flat small></v-switch>
            </div>
          </td>
        </tr>
      </template>
    </v-data-table>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import PageTitle from '~/components/PageTitle.vue';

const menuItems = [
  { name: '板', type: 'header' },
  { name: '閲覧履歴', key: 'board-history-delete' },
  { name: 'スレッド', type: 'header' },
  { name: '閲覧履歴', key: 'thread-history-delete' },
  { name: '最新レス50レスのみ表示', key: 'display-latest-50-replies-only' },
  { name: 'コメント', type: 'header' },
  { name: '昇順/降順', key: 'display-latest-50-replies-only' },
];
</script>

<style scoped>
.highlight {
  background-color: #f0f0f0;
}
</style>
