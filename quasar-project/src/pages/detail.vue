<template>
  <div class="q-pa-md row items-start q-gutter-md">

    <q-card dark bordered class="bg-grey-9 my-card">
      <q-card-section>
        <div class="text-h4">{{ title }}<q-btn
        size = 'lg'
        round
        flat
        dense
         :icon="expanded ? 'star_outline' : 'star'"
        @click="expanded = !expanded"
        />
        </div>
        <div>
        <q-chip icon="today">Data: {{ date }}</q-chip>
        <q-chip icon="person">Author: {{ author }}</q-chip>
        <q-chip icon="update">Version: {{ version }}</q-chip>
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
        {{ content }}
      </q-card-section>
    </q-card>

  </div>
</template>

<script>
import { ref } from 'vue'
import { useRoute } from 'vue-router';
import axios from 'axios';
export default {
  setup () {
    const route = useRoute()
    console.log(route.query.docID);

    const title = ref('');
    const content = ref('');
    const author = ref('');
    const date = ref('');
    const version = ref('First');

    function getDoc (id) {
      axios
        .get(`http://localhost:8000/TSMC/docs/getDoc/${id}`)
        .then(response => {
          console.log(response);
          console.log(response.data);
          title.value = response.data.Title
          content.value = response.data.Content
          author.value = response.data.AuthorName
          date.value = response.data.UpdatedAt
        })
        .catch(error => {
          console.error(error);
        });
    };
    getDoc(route.query.docID);

    return {
      expanded: ref(false),
      title,
      content,
      author,
      date,
      version
    }
  }
}
</script>
<style lang="sass" scoped>
.my-card
  width: 100%
  height: 100%
  max-width: 100%
</style>
