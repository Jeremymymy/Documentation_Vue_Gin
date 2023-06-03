<template>
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
  <div class=" q-pa-md row section-card items-start q-gutter-md"  align="center">
    <q-card class="my-card" flat bordered>
      <q-img
        src="https://cdn.quasar.dev/img/parallax2.jpg"
      />
      <q-card-section>
        <div class="text-overline text-orange-9">Overline</div>
        <div class="text-h5 q-mt-sm q-mb-xs">Title</div>
        <div class="text-caption text-grey">
          Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.
        </div>
      </q-card-section>
      <q-card-actions align="right">
        <q-btn flat round color="red" icon="favorite" />
        <q-btn flat round color="teal" icon="bookmark" />
        <q-btn flat round color="primary" icon="share" />
      </q-card-actions>
    </q-card>
     <q-card class="my-card" flat bordered>
      <q-img
        src="https://cdn.quasar.dev/img/parallax2.jpg"
      />

      <q-card-section>
        <div class="text-overline text-orange-9">Overline</div>
        <div class="text-h5 q-mt-sm q-mb-xs">Title</div>
        <div class="text-caption text-grey">
          Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.
        </div>
      </q-card-section>

      <q-card-actions align="right">
        <q-btn flat round color="red" icon="favorite" />
        <q-btn flat round color="teal" icon="bookmark" />
        <q-btn flat round color="primary" icon="share" />
      </q-card-actions>
    </q-card>
     <q-card class="my-card" flat bordered>
      <q-img
        src="https://cdn.quasar.dev/img/parallax2.jpg"
      />

      <q-card-section>
        <div class="text-overline text-orange-9">Overline</div>
        <div class="text-h5 q-mt-sm q-mb-xs">Title</div>
        <div class="text-caption text-grey">
          Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.
        </div>
      </q-card-section>

      <q-card-actions align="right">
        <q-btn flat round color="red" icon="favorite" />
        <q-btn flat round color="teal" icon="bookmark" />
        <q-btn flat round color="primary" icon="share" />
      </q-card-actions>
    </q-card>
     <q-card class="my-card" flat bordered>
      <q-img
        src="https://cdn.quasar.dev/img/parallax2.jpg"
      />

      <q-card-section>
        <div class="text-overline text-orange-9">Overline</div>
        <div class="text-h5 q-mt-sm q-mb-xs">Title</div>
        <div class="text-caption text-grey">
          Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.
        </div>
      </q-card-section>

      <q-card-actions align="right">
        <q-btn flat round color="red" icon="favorite" />
        <q-btn flat round color="teal" icon="bookmark" />
        <q-btn flat round color="primary" icon="share" />
      </q-card-actions>
    </q-card>
  </div>
  <!-- <div  class="text-h6 q-pa-sm " align="center"><b class="title">發文紀錄</b></div>
      <div class=" q-pa-md row section-card items-start q-gutter-md col justify-center"  align="center">
      <q-card class="my-card" flat bordered  v-for="item in paginatedDoc()" :key="item.ID">

        <q-card-actions align="left">
           <q-btn flat round icon="favorite" :color="item.favorite ? 'red' : 'gray'" @click="createCollect(item)"/>
        </q-card-actions>

        <q-card-section>

          <div class="text-h5 q-mt-sm q-mb-xs">{{ item.Title }}</div>
          <div class="text-overline text-orange-10">Author: {{ item.AuthorName }}</div>
          <div class="text-caption text-grey">
            {{ item.Content }}
          </div>
        </q-card-section>
        <q-card-actions align="right">

            <router-link :to="{path: '/detail', query: {docID: item.ID }}">
              <q-btn flat round color="primary" icon="edit" />
            </router-link>

            <q-btn flat round color="teal" icon="delete" @click="dialogDelete = !dialogDelete" >
              <q-dialog v-model="dialogDelete">
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
          v-model="currentDoc"
          :min="1"
          :max="Math.ceil(allDoc.length/4)"
          :input="true"
          input-class="text-orange-10"
        />
      </div> -->
</template>

<script>
import { ref } from 'vue'
import axios from 'axios';
export default {
  name: 'Index2',
  props: ['data'],
  setup () {
    const allDepDoc = ref('');
    const pageSize = 4;
    const currentDoc = ref(1);
    function paginatedDoc () {
      const startIndex = (currentDoc.value - 1) * pageSize;
      return allDepDoc.value.slice(startIndex, startIndex + pageSize);
    };
    function getAllDepDoc () {
      axios
        .get('http://localhost:8000/TSMC/users/getMyDetail')
        .then(response => {
          console.log(response);

          // allDepDoc.value = response.data.PostDocs;
          // allCollect.value = response.data.CollectDocs;
          // console.log(allCollect.value.length);
          // allDoc.value.forEach((elem) => {
          //   // elem.page =
          //   allCollect.value.forEach((elemC) => {
          //     // console.log(elem);
          //     // console.log(elemC.Title);
          //     if (elemC.Title !== elem.Title && elem.favorite !== true) {
          //       elem.favorite = false;
          //       // console.log('yes');
          //     } else {
          //       elem.favorite = true;
          //       // console.log('no');
          //     }
          //   });
          // });

          // console.log(allDoc.value);
          // console.log(allCollect.value);
        })
        .catch(error => {
          console.error(error);
          // Handle registration error
          // You can display an error message or perform other actions
        });
    };
    return {
      expanded: ref(false),
      slide: ref('style'),
      autoplay: ref(true),
      lorem: 'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.',
      paginatedDoc,
      getAllDepDoc,
      allDepDoc,
      pageSize,
      currentDoc
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
  max-width: 350px
</style>
