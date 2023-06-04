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
        <q-chip v-show="watchflag" icon="today">Create Data: {{ date }}</q-chip>
        <q-chip v-show="watchflag" icon="person">Author: {{ author }}</q-chip>
        <q-chip icon="today">Update Data: {{ updateDate }}</q-chip>
        <q-chip icon="person">Updater: {{ updater }}</q-chip>
        <q-chip icon="update">Version: {{ version }}</q-chip>
        </div><br />

        <router-link :to="{path: '/MyHistory', query: {docID: thisDoc }}">
          <q-btn
          unelevated rounded
          glossy
          icon="menu"
          class="right-btn"
          color="dark"
          label="history"
          align="right"
          />
        </router-link>

      </q-card-section>
      <q-separator dark inset />
      <q-card-section class="my-doc-content">
        <div class="col-5 ">
          <div class="row">
            <q-editor v-show="editflag" v-model="dContent" class="section-card "
                toolbar-text-color="grey-6"
                toolbar-toggle-color="white"
                content-class="bg-grey-8 "
                toolbar-bg="grey-10"
                flat
                :toolbar="[
                  ['bold', 'italic', 'strike', 'underline'],
                  ['fullscreen'],
                  ['save']
                ]"
                :definitions="{
                  save: {
                    tip: 'Save your document',
                    icon: 'save',
                    label: 'Save',
                    handler: updateDoc
                  }
                }"
            />
          <div v-show="!editflag" class="row">
            {{ content }}
          </div>
          </div>
          <div v-show="watchflag" class="row justify-center">
            <q-btn v-show="!editflag" label="編輯" color="black" type="submit" @click="edit" />
            <q-btn v-show="editflag" label="取消編輯" color="black" type="submit" @click="edit" />
          </div>
        </div>
      </q-card-section>
    </q-card>

  </div>
</template>

<script>
import { ref } from 'vue'
import { useRoute } from 'vue-router';
import axios from 'axios';
import { useQuasar } from 'quasar'
export default {
  setup () {
    const route = useRoute()
    console.log(route.query.docID);
    console.log(route.query.ver);
    const $q = useQuasar()
    const title = ref('');
    const content = ref('');
    const author = ref('');
    const updater = ref('');
    const date = ref('');
    const updateDate = ref('');
    const belong = ref('');
    const version = ref('First');
    const editflag = ref(false);
    const watchflag = ref(false); // 從版控頁來的不能編輯

    const dContent = ref(content.value);

    function getDoc (id) {
      axios
        .get(`http://localhost:8000/TSMC/docs/getDoc/${id}`)
        .then(response => {
          console.log(response);
          console.log(response.data);
          title.value = response.data.Title
          content.value = response.data.Content
          author.value = response.data.AuthorName
          date.value = response.data.CreatedAt
          updateDate.value = response.data.UpdatedAt
          belong.value = response.data.Belong
          dContent.value = response.data.Content
        })
        .catch(error => {
          console.error(error);
        });
    };
    function getDocVersion (id) {
      axios
        .get(`http://localhost:8000/TSMC/docs/getDocAllVers/${id}`)
        .then(response => {
          console.log(response);
          console.log(response.data);

          const index = response.data.Vers.length
          version.value = index
          updater.value = response.data.Vers[index - 1].UpdaterName
          // updateDate.value = response.data.Vers[index - 1].UpdatedAt
          // console.log(updater.value);
        })
        .catch(error => {
          console.error(error);
        });
    };
    function getDocWithVersion (id, ver) {
      axios
        .get(`http://localhost:8000/TSMC/docs/getDocAllVers/${id}`)
        .then(response => {
          console.log(response);
          console.log(response.data);

          version.value = ver
          updater.value = response.data.Vers[ver - 1].UpdaterName
          title.value = response.data.Vers[ver - 1].Title
          content.value = response.data.Vers[ver - 1].Content
          author.value = response.data.Vers[ver - 1].AuthorName
          updateDate.value = response.data.Vers[ver - 1].UpdatedAt
          dContent.value = response.data.Vers[ver - 1].Content
        })
        .catch(error => {
          console.error(error);
        });
    };

    function edit () {
      editflag.value = !editflag.value;
    };
    function updateDoc () {
      const doc = {
        Title: title.value,
        Content: dContent.value,
        Belong: belong.value
      };
      axios
        .put(`http://localhost:8000/TSMC/docs/update/${thisDoc.value}`, doc)
        .then(response => {
          console.log(response);
          console.log(response.data);

          getDoc(thisDoc.value);

          if (route.query.ver === undefined) {
            getDocVersion(thisDoc.value);
          } else {
            getDocWithVersion(thisDoc.value, version.value)
          }

          edit();
          $q.notify({
            message: 'Saved Successfully!',
            color: 'green-4',
            textColor: 'white',
            icon: 'cloud_done'
          })
        })
        .catch(error => {
          console.error(error);
        });
    };

    const thisDoc = ref(route.query.docID);
    getDoc(thisDoc.value);
    if (route.query.ver === undefined) {
      watchflag.value = true;
      getDocVersion(thisDoc.value);
      console.log('PPPPPPPPPPPPPPPPPPPPPPPPPP')
    } else {
      version.value = route.query.ver;
      watchflag.value = false;
      getDocWithVersion(thisDoc.value, version.value);
      console.log('HEREEEEEEEEEEEEEEEEEEEEEEEE')
    }

    return {
      expanded: ref(false),
      title,
      content,
      author,
      updater,
      date,
      updateDate,
      belong,
      version,
      editflag,
      edit,
      dContent,
      thisDoc,
      updateDoc,
      watchflag
    }
  }
}
</script>
<style lang="sass" scoped>
.my-card
  width: 100%
  height: 100%
  max-width: 100%
.my-doc-content
  width: 100%
  height: 100%
  full-height: 100%
.section-card

  width: 100%
  height: 100%
  max-width: 100%
</style>
