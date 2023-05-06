<template>
  <div class="q-pa-md row items-start q-gutter-md">

    <q-card dark bordered class="bg-grey-9 my-card">
      <q-card-section>
        <div class="text-h4">Title<q-btn
        size = 'lg'
        round
        flat
        dense
         :icon="expanded ? 'star_outline' : 'star'"
        @click="expanded = !expanded"
        />
        </div>
        <div>
        <q-chip icon="today">Data:</q-chip>
        <q-chip icon="person">Author:</q-chip>
        <q-chip icon="update">Version:</q-chip>
        </div><br />
        <q-btn
        unelevated rounded
        glossy
        icon="menu"
        class="right-btn"
        color="dark"
        label="history"
        align="right"
        />
      </q-card-section>
      <q-separator dark inset />
      <q-card-section>
        {{ lorem }}
      </q-card-section>
    </q-card>

  </div>
</template>

<script>
import { ref } from 'vue'
export default {
  setup () {
    return {
      expanded: ref(false),
      lorem: 'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.'
    }
  }
}
</script>
