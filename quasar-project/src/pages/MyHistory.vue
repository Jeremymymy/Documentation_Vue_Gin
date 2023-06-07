<template>
  <q-page>
    <div class="text-h3" align="center">
       <b class="title"><br /><q-icon name="update" /> History</b><br />
    </div>
    <q-card dark bordered class="bg-grey-9 my-card">
      <q-card-section>
        <div class="text-h4 q-mb-sm text-center">{{ title }}</div>
        <hr class="h1" />
        <div v-for="item in version" :key="item.ID">
          <div class="q-ma-md show-row-equal-width">
            <div class="row" >
              <div class="text-h6 col q-ma-md text-center">Version {{ item.Index }}</div>
              <div class="text-h6 col q-ma-md text-center">Updater: {{ item.UpdaterName }}</div>
              <div class="text-h6 col q-ma-md text-center">{{ item.UpdatedAt }}</div>
              <!-- <div class="text-h5 col q-ma-md text-center">{{ item[3] }}</div> -->
              <div class="col q-ma-md text-center">
                <router-link :to="{path: '/detail', query: {docID: item.DocId, ver: item.ID }}">
                  <q-btn color="white" text-color="black" label="View" />
                </router-link>
              </div>
            </div>
          </div>
          <hr class="h1" />
        </div>
      </q-card-section>
    </q-card>
  </q-page>
</template>

<script>
import axios from 'axios';
import { defineComponent, ref } from 'vue';
import { useRoute } from 'vue-router';
export default defineComponent({
  setup () {
    const route = useRoute()
    console.log(route.query.docID);
    const thisDoc = ref(route.query.docID);
    const title = ref('');
    const version = ref(Object([{}]));

    // const fetchData = async () => {
    //   try {
    //     const response = await axios.get('');
    //     const data = response.data;
    //     router.push({ name: '/detail', params: { data } });
    //   } catch (error) {
    //     console.error(error);
    //   }
    // };
    const fetchData = async () => {
      try {
        const response = await axios.get(`http://localhost:8000/TSMC/docs/getDocAllVers/${thisDoc.value}`);
        // const data = response.data;
        version.value = response.data.Vers
        let index = 0;
        version.value.forEach((elem) => {
          index++;
          elem.Index = index;
        });
        console.log(version.value);
        title.value = response.data.Title
        // console.log(data);
      } catch (error) {
        console.error(error);
      }
    };

    const historyInfo = [
      ['Version1', 'Author1', 'Date1'],
      ['Version2', 'Author2', 'Date2'],
      ['Version3', 'Author3', 'Date3']
    ];
    const documentTitle = ref('Title');
    fetchData();
    return {
      fetchData,
      historyInfo,
      documentTitle,
      title,
      version,
      thisDoc
    };
  }
});
</script>

<style lang="sass">
.show-row-equal-width
  .row > div
    padding: 5px 10px
  .row + .row
    margin-top: 1rem

hr.h1
  border: 1px solid black
</style>
