<template>
  <div>
    <q-bar class="bg-secondary text-white shadow-2">
      <div>[[index .Vue.I18n "recursiveListTitle"]]</div>
      <q-space />
      <q-btn v-if = '!isDeleted' dense flat icon="delete" @click="isDeleted = !isDeleted"><q-tooltip>показать список удаленных</q-tooltip></q-btn>
      <q-btn v-else dense round outline icon="delete" @click="isDeleted = !isDeleted"><q-tooltip>показать список активных</q-tooltip></q-btn>
      <q-btn v-if="!readonly" dense flat icon="add" @click="$router.push(`/[[.Vue.RouteName]]/${id}/new`)"/>
    </q-bar>

    <q-list bordered separator>
      <q-item v-for="item in list" :key="item.id">
          [[.PrintListRowAvatar]]
          [[.PrintListRowLabel]]
        <comp-delete-btn-in-list v-if="!readonly" update-method="[[.Name]]_update" :item="item" @success="onChangeList"/>
      </q-item>
    </q-list>
  </div>
</template>

<script>
    export default {
        props: ['id', 'readonly'],
        computed: {
          currentUrl: function () {
            return `/[[.Vue.RouteName]]/${this.id}/`
          },
        },
        data() {
            return {
                list: [],
                isDeleted: false,
            }
        },
        watch: {
            isDeleted() {
                this.reload()
            }
        },
        methods: {
            onChangeList() {
                this.reload()
                this.$emit('update')
            },
            reload() {
                this.$utils.postCallPgMethod({method: '[[.Name]]_list', params: {parent_id: +this.id, deleted: this.isDeleted}}).subscribe(res => {
                    if (res.ok) {
                        this.list = res.result
                    }
                })
            }
        },
        mounted() {
           this.reload()
        }
    }
</script>
