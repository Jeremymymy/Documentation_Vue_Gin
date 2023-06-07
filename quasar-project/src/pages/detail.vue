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
        <!--<q-chip >{{ watchflag }}</q-chip> -->
        <!-- <q-chip >{{ thisMine }}</q-chip> -->
        <!-- <q-chip v-if= "watchflag === 'true'" icon="today">Create Data: {{ date }}</q-chip> -->
        <!-- <q-chip >{{ watchflag }}</q-chip> -->
        <q-chip v-show= "watchflag" icon="today">Create Data: {{ date }}</q-chip>
        <q-chip v-show= "watchflag" icon="person">Author: {{ author }}</q-chip>
        <q-chip icon="today">Update Data: {{ updateDate }}</q-chip>
        <q-chip icon="person">Updater: {{ updater }}</q-chip>
        <!-- <q-chip icon="update">Version: {{ version }}</q-chip> -->
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
          <div v-show= "watchflag"  class="row justify-center">
          <!-- <p v-show="thisMine"> -->
            <q-btn v-show="!editflag" unelevated rounded glossy  icon="edit" label="編輯" color="black" type="submit" @click="edit" />
          <!-- </p> -->
            <q-btn v-show="editflag" unelevated rounded glossy label="取消編輯" color="black" type="submit" @click="edit" />
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
import { useQuasar, LocalStorage } from 'quasar'
export default {
  setup () {
    const route = useRoute()
    const userInfo = LocalStorage.getItem('userInfo');
    console.log(route.query.docID);
    console.log(route.query.ver);
    const $q = useQuasar()
    const title = ref('');
    const content = ref('');
    const author = ref('');
    const authorid = ref('');
    const updater = ref('');
    const thisver = ref('');
    const date = ref('');
    const updateDate = ref('');
    const belong = ref('');
    const version = ref('First');
    const editflag = ref(false);
    const watchflag = ref(false); // 從版控頁來的不能編輯

    const dContent = ref(content.value);
    console.log(watchflag.value)
    function getDoc (id, ver) {
      // axios
      console.log(id);
      console.log(ver.value);
      if (ver.value !== undefined) {
        axios.get(`http://localhost:8000/TSMC/docs/getDocVer/${ver.value}`)
          .then(response => {
            console.log(response);
            console.log(response.data);
            title.value = response.data.Title
            content.value = response.data.Content
            author.value = response.data.AuthorName
            // console.log(response.data.AuthorId);
            authorid.value = response.data.AuthorId
            console.log(authorid.value);
            date.value = response.data.CreatedAt
            updateDate.value = response.data.UpdatedAt
            belong.value = response.data.Belong
            dContent.value = response.data.Content
            console.log(userInfo.EmployeeId);
            if (authorid.value === userInfo.EmployeeId) {
            // if (route.query.ver === undefined) {
              watchflag.value = true;
              getDocVersion(thisDoc.value);
              console.log('PPPPPPPPPPPPPPPPPPPPPPPPPP')
            } else {
            // version.value = route.query.ver;
              watchflag.value = false;
              getDocWithVersion(thisDoc.value, version.value);
              console.log('HEREEEEEEEEEEEEEEEEEEEEEEEE')
            }
          })
          .catch(error => {
            console.error(error);
          });
      } else {
        axios.get(`http://localhost:8000/TSMC/docs/getDoc/${id}`)
          .then(response => {
            console.log(response);
            console.log(response.data);
            title.value = response.data.Title
            content.value = response.data.Content
            author.value = response.data.AuthorName
            // console.log(response.data.AuthorId);
            authorid.value = response.data.AuthorId
            console.log(authorid.value);
            date.value = response.data.CreatedAt
            updateDate.value = response.data.UpdatedAt
            belong.value = response.data.Belong
            dContent.value = response.data.Content
            console.log(userInfo.EmployeeId);
            if (authorid.value === userInfo.EmployeeId) {
            // if (route.query.ver === undefined) {
              watchflag.value = true;
              getDocVersion(thisDoc.value);
              console.log('PPPPPPPPPPPPPPPPPPPPPPPPPP')
            } else {
            // version.value = route.query.ver;
              watchflag.value = false;
              getDocWithVersion(thisDoc.value, version.value);
              console.log('HEREEEEEEEEEEEEEEEEEEEEEEEE')
            }
          })
          .catch(error => {
            console.error(error);
          });
      }
    };
    function getDocVersion (id) {
      axios
        .get(`http://localhost:8000/TSMC/docs/getDocAllVers/${id}`)
        .then(response => {
          console.log(response);
          console.log(response.data);

          const index = response.data.Vers.length
          // version.value = index
          updater.value = response.data.Vers[index - 1].UpdaterName
          version.value = response.data.Vers[index - 1].ID
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
    function updateDoc () { // 從版控頁來的不能編輯
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

          getDoc(thisDoc.value, thisver.value);
          getDocVersion(thisDoc.value);
          // if (route.query.ver === undefined) {
          //   getDocVersion(thisDoc.value);
          // } else {
          //   getDocWithVersion(thisDoc.value, version.value)
          // }

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
    // const thisMine = ref(route.query.docmine);
    console.log(watchflag.value)
    const thisDoc = ref(route.query.docID);
    if (route.query.ver === undefined) {
      thisver.value = version.value
    } else {
      thisver.value = ref(route.query.ver);
    }
    console.log(thisDoc)
    getDoc(thisDoc.value, thisver.value);
    // console.log(thisMine)
    // if (authorid.value === userInfo.EmployeeId) {
    // // if (route.query.ver === undefined) {
    //   watchflag.value = true;
    //   getDocVersion(thisDoc.value);
    //   console.log('PPPPPPPPPPPPPPPPPPPPPPPPPP')
    // } else {
    //   // version.value = route.query.ver;
    //   watchflag.value = false;
    //   getDocWithVersion(thisDoc.value, version.value);
    //   console.log('HEREEEEEEEEEEEEEEEEEEEEEEEE')
    // }
    console.log(watchflag.value)
    console.log(editflag.value)
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
      authorid,
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
