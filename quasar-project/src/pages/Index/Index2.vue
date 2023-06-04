<template>
<!-- <Suspense> -->
  <div  class="text-h3 q-pa-sm " align="center"><b class="title"><br />TSMC最新公告</b></div>
  <div class="q-pa-md">

    <q-carousel
      v-model="slide"
      :autoplay="autoplay"
      infinite
      swipeable
      animated
      padding
      arrows
      navigation
      navigation-icon="radio_button_unchecked"
      height="500px"
      class="bg-dark glossy text-white rounded-borders"
    >

      <q-carousel-slide name="style" class="text-center">
        <q-scroll-area class="fit">
          <q-icon name="style" size="56px" />
          <div class="q-mt-md">
            {{ lorem }}
          </div>
          <div class="q-mt-md">
            {{ lorem }}
          </div>
          <div class="q-mt-md">
            {{ lorem }}
          </div>
          <div class="q-mt-md">
            {{ lorem }}
          </div>
        </q-scroll-area>
      </q-carousel-slide>

      <q-carousel-slide name="tv" class="text-center">
        <q-scroll-area class="fit">
          <q-icon name="live_tv" size="56px" />
          <div class="q-mt-md">
            {{ lorem }}
          </div>
          <div class="q-mt-md">
            {{ lorem }}
          </div>
          <div class="q-mt-md">
            {{ lorem }}
          </div>
          <div class="q-mt-md">
            {{ lorem }}
          </div>
        </q-scroll-area>
      </q-carousel-slide>

      <q-carousel-slide name="layers" class="text-center">
        <q-scroll-area class="fit">
          <q-icon name="layers" size="56px" />
          <div class="q-mt-md">
            {{ lorem }}
          </div>
          <div class="q-mt-md">
            {{ lorem }}
          </div>
          <div class="q-mt-md">
            {{ lorem }}
          </div>
          <div class="q-mt-md">
            {{ lorem }}
          </div>
        </q-scroll-area>
      </q-carousel-slide>

      <q-carousel-slide name="map" class="text-center">
        <q-scroll-area class="fit">
          <q-icon name="terrain" size="56px" />
          <div class="q-mt-md">
            {{ lorem }}
          </div>
          <div class="q-mt-md">
            {{ lorem }}
          </div>
          <div class="q-mt-md">
            {{ lorem }}
          </div>
          <div class="q-mt-md">
            {{ lorem }}
          </div>
        </q-scroll-area>
      </q-carousel-slide>
    </q-carousel>
  </div>

  <div  class="text-h3 q-pa-sm " align="center"><b class="title">公開文件</b></div>
      <div class=" q-pa-md section-card items-start q-gutter-md row justify-center"  align="center">
      <q-card class="my-card col" flat bordered  v-for="item in paginatedPublicDoc()" :key="item.ID">

        <q-card-actions align="left">
           <q-btn flat round icon="favorite" :color="item.favorite ? 'red' : 'gray'" @click="createCollect(item)"/>
        </q-card-actions>

        <q-card-section>

          <div class="text-h5 q-mt-sm q-mb-xs">{{ item.Title }}</div>
          <div class="text-overline text-orange-10">Author: {{ item.AuthorName }}</div>
          <div class="text-overline text-orange-10">Department: {{ item.Belong }}</div>
          <div class="text-caption text-grey">
            {{ item.Content }}
          </div>
        </q-card-section>
        <q-card-actions align="center">

            <router-link :to="{path: '/detail', query: {docID: item.ID }}">
              <q-btn unelevated rounded glossy color="primary" icon="description" label = "文章內容"/>
            </router-link>

            <q-btn v-show="item.mine" flat round color="teal" icon="delete" @click="item.deleteDialog = !item.deleteDialog" >
              <q-dialog v-model="item.deleteDialog">
                <q-card class="my-card-info-fix q-pa-md" align="center">
                  <q-card-section>
                    <div class="text-h6">確認要刪除 {{ item.Title }} ?</div>
                  </q-card-section>

                <q-btn flat round color="black" label="確認" type="submit" @click="deleteDoc(item)" ></q-btn>
                <q-btn flat round v-close-popup label="取消" color="black"></q-btn>
                </q-card>
              </q-dialog>
            </q-btn>
        </q-card-actions>
      </q-card>
      </div>
      <div class="q-pa-lg flex flex-center justify-center">
        <q-pagination
          v-model="currentPublicDoc"
          :min="1"
          :max="totalPages(allPublicDoc)"
          :input="true"
          input-class="text-orange-10"
        />
      </div>

  <div  class="text-h3 q-pa-sm " align="center"><b class="title">部門文件</b></div>
      <div class=" q-pa-md section-card items-start q-gutter-md row justify-center"  align="center">
      <q-card class="my-card col" flat bordered  v-for="item in paginatedDoc()" :key="item.ID">

        <q-card-actions align="between">
           <q-btn flat round icon="favorite" :color="item.favorite ? 'red' : 'gray'" @click="createCollect(item)"/>
           <q-btn v-show="item.mine" flat round color="teal" icon="delete" @click="item.deleteDialog = !item.deleteDialog" ></q-btn>
        </q-card-actions>

        <q-card-section>

          <div class="text-h5 q-mt-sm q-mb-xs">{{ item.Title }}</div>
          <div class="text-overline text-orange-10">Author: {{ item.AuthorName }}</div>
          <div class="text-overline text-orange-10">Department: {{ item.Belong }}</div>
          <div class="text-caption text-grey">
            {{ item.Content }}
          </div>
        </q-card-section>
        <q-card-actions align="center">

            <router-link :to="{path: '/detail', query: {docID: item.ID }}">
              <q-btn unelevated rounded glossy color="primary" icon="description" label = "文章內容"/>
            </router-link>

              <q-dialog v-model="item.deleteDialog">
                <q-card class="my-card-info-fix q-pa-md" align="center">
                  <q-card-section>
                    <div class="text-h6">確認要刪除 {{ item.Title }} ?</div>
                  </q-card-section>

                <q-btn flat round color="black" label="確認" type="submit" @click="deleteDoc(item)" ></q-btn>
                <q-btn flat round v-close-popup label="取消" color="black"></q-btn>
                </q-card>
              </q-dialog>
        </q-card-actions>
      </q-card>
      </div>
      <div class="q-pa-lg flex flex-center justify-center">
        <q-pagination
          v-model="currentDoc"
          :min="1"
          :max="totalPages(allDepDoc)"
          :input="true"
          input-class="text-orange-10"
        />
      </div>
    <!-- </Suspense> -->
</template>

<script>
import { ref } from 'vue'
import axios from 'axios';
import { LocalStorage } from 'quasar';
export default {
  name: 'Index2',
  props: ['data'],
  setup () {
    const userInfo = LocalStorage.getItem('userInfo');
    const allMyCollect = ref(Object([{}])); // LocalStorage.getItem('userCollect')
    const allDepDoc = ref(Object([{}]));
    const allPublicDoc = ref(Object([{}]));
    const pageSize = 4;
    const currentDoc = ref(1);
    const currentPublicDoc = ref(1);

    function totalPages (item) {
      console.log(item)
      if (item.length === 0) {
        return 1;
      } else {
        return Math.ceil(item.length / pageSize);
      }
    };
    function paginatedDoc () {
      const startIndex = (currentDoc.value - 1) * pageSize;
      return allDepDoc.value.slice(startIndex, startIndex + pageSize);
    };
    function paginatedPublicDoc () {
      const startIndex = (currentPublicDoc.value - 1) * pageSize;
      return allPublicDoc.value.slice(startIndex, startIndex + pageSize);
    };

    function getAllPublicDoc () {
      axios
        .get('http://localhost:8000/TSMC/docs/getDepartmentDocs/Public')
        .then(response => {
          console.log(response);
          console.log(response.data);
          console.log('length ' + response.data.length);

          allPublicDoc.value = response.data;
          console.log(allMyCollect.value.length);
          allPublicDoc.value.forEach((elem) => {
            allMyCollect.value.forEach((elemC) => {
              if (elemC.DocId !== elem.ID && elem.favorite !== true) {
                elem.favorite = false;
              } else {
                elem.favorite = true;
              }
              elem.deleteDialog = false;
            });
            if (elem.AuthorId === userInfo.EmployeeId) {
              elem.mine = true;
            } else {
              elem.mine = false;
            }
          });
          console.log(allPublicDoc.value);
        })
        .catch(error => {
          console.error(error);
        });
    };
    function getAllDepDoc () {
      axios
        .get(`http://localhost:8000/TSMC/docs/getDepartmentDocs/${userInfo.Department}`)
        .then(response => {
          console.log(response);
          console.log(response.data);

          allDepDoc.value = response.data;
          console.log(allMyCollect.value.length);
          allDepDoc.value.forEach((elem) => {
            allMyCollect.value.forEach((elemC) => {
              if (elemC.DocId !== elem.ID && elem.favorite !== true) {
                elem.favorite = false;
              } else {
                elem.favorite = true;
              }
              elem.deleteDialog = false;
            });
            if (elem.AuthorId === userInfo.EmployeeId) {
              elem.mine = true;
            } else {
              elem.mine = false;
            }
          });
          console.log(allDepDoc.value);
        })
        .catch(error => {
          console.error(error);
        });
    };
    function getAllUserInfo () {
      axios
        .get('http://localhost:8000/TSMC/users/getMyDetail')
        .then(response => {
          console.log(response);

          allMyCollect.value = response.data.CollectDocs;
          console.log(allMyCollect.value.length);
          // LocalStorage.set('userCollect', allMyCollect.value);

          console.log(allMyCollect.value);
        })
        .catch(error => {
          console.error(error);
        });
    };
    function createCollect (ff) {
      ff.favorite = !ff.favorite;
      if (ff.favorite === true) {
        axios
          .get(`http://localhost:8000/TSMC/docs/collectDoc/${ff.ID}`)
          .then(response => {
            console.log(response);
            allMyCollect.value = response.data.Collect
            console.log(allMyCollect.value);
            getAllUserInfo();
          })
          .catch(error => {
            console.error(error);
            // Handle registration error
          });
      } else {
        console.log('Delete');
        let colID = null;
        allMyCollect.value.forEach((elem) => {
          if (ff.ID === elem.DocId) {
            colID = elem.ID
          }
        });
        if (colID !== null) {
          axios
            .delete(`http://localhost:8000/TSMC/docs/deleteCollect/${colID}`)
            .then(response => {
              console.log(response);
              getAllUserInfo();
            })
            .catch(error => {
              console.error(error);
            });
        }
      }
    };
    function deleteCollect (ff) {
      // ff.favorite = !ff.favorite;
      console.log('Delete');
      axios
        .delete(`http://localhost:8000/TSMC/docs/deleteCollect/${ff.ID}`)
        .then(response => {
          console.log(response);
          getAllUserInfo();
        })
        .catch(error => {
          console.error(error);
        });
    };
    function deleteDoc (ff) {
      axios
        .delete(`http://localhost:8000/TSMC/docs/deleteDoc/${ff.ID}`)
        .then(response => {
          console.log(response);
          console.log(response.data);
          allMyCollect.value.forEach((elemC) => {
            if (ff.AuthorId === elemC.AuthorId && ff.ID === elemC.DocId) {
              deleteCollect(elemC);
            }
          });
          getAllUserInfo();
          getAllPublicDoc();
          getAllDepDoc();
        })
        .catch(error => {
          console.error(error);
        });
    };

    getAllUserInfo();
    getAllPublicDoc();
    getAllDepDoc();

    return {
      expanded: ref(false),
      slide: ref('style'),
      autoplay: ref(true),
      lorem: 'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.',
      paginatedDoc,
      paginatedPublicDoc,
      allPublicDoc,
      getAllDepDoc,
      allDepDoc,
      pageSize,
      currentDoc,
      currentPublicDoc,
      allMyCollect,
      createCollect,
      getAllUserInfo,
      deleteDoc,
      deleteCollect,
      totalPages
    }
  }
}
</script>

<style lang="sass" scoped>

.section-card
  margin-right: 10%
  margin-left: 10%

.my-card
  width: 100%
  height: 100%
  max-width: 300px
</style>
