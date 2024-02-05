<template>
  <div>
    <v-container class="my-5">
      <v-layout row wrap class="mt-5 ml-3">
        <v-flex class="d-flex justify-left" lg2>
          <v-icon left color="blue darken-4" medium>mdi-download</v-icon>
          <b>Log File Download</b>
        </v-flex>
        
        
          <v-flex>
            <v-btn  color="primary"  small @click="logfileDownload">
              Download
            </v-btn>
          </v-flex>
      </v-layout>
        
    
    </v-container>
  </div>
</template>

<script>
import EventService from "@/services/EventServices.js";
export default {
  components: {},
  //data() {},
  methods: {
    logfileDownload() {
      // console.log("logfileDownload+");
      EventService.LogDownload()
        .then((response) => {
          // console.log(response.data.content);
          // console.log(response.data.fileName);
          if (response.data.status == "S") {
            // ---------------------------------------
            const fileData = new Blob([response.data.content], {
              type: "text/plain",
            });
            const fileUrl = URL.createObjectURL(fileData);

            const link = document.createElement("a");
            link.href = fileUrl;
            link.download = response.data.fileName; // Replace with the desired file name
            link.click();
            
            // Cleanup
            URL.revokeObjectURL(fileUrl);

            // ---------------------------------------

            this.MessageBar("S", "Successfully Downloaded");
          } else if (response.data.status == "E") {
            //this.errMsg = response.data.ErrMsg;
            this.MessageBar("E", response.data.errMsg);
          }
        })
        .catch((error) => {
          this.MessageBar("E", error.response);
        });
      console.log("logfileDownload-");
    },

    // logfileDownload() {
    //   EventService.LogDownload()
    //     .then((response) => {
    //       console.log("response.data.content", response.data);
    //       console.log("response.data.fileName", response.data.fileName);
    //       //if (response.headers["content-type"] == "application/json") {
    //         // console.log(response.data);
    //         this.blobToJson(response.data)
    //           .then((json) => {

    //             // Handle the JSON data here
    //             console.log(json);
    //             if (json.status == "E") {
    //               alert(json.statusCode + "/" + json.msg);
    //             }
    //           })
    //           .catch((error) => {
    //             console.error("Error converting Blob to JSON:", error);
    //           });
    //       //} else {
    //         const blob = new Blob([response.data.content], {
    //           type: response.headers["content-type"],
    //         });
    //         // const blob = new Blob([response], { type: "application/pdf" });
    //         if (window.navigator.msSaveBlob) {
    //           // For IE and Edge browsers
    //           window.navigator.msSaveBlob(blob, response.data.fileName);
    //         } else {
    //           // For other browsers
    //           const downloadLink = document.createElement("a");
    //           downloadLink.href = URL.createObjectURL(blob);
    //           downloadLink.download = response.data.fileName;
    //           downloadLink.click();
    //         }
    //       //}
    //     })
    //     .catch((error) => {
    //       console.error("Error downloading file:", error);
    //     });
    // },
    // blobToJson(blob) {
    //   console.log("blobToJson +")
    //   return new Promise((resolve, reject) => {
    //     const reader = new FileReader();
    //     reader.onload = () => {
    //       try {
    //         const json = JSON.parse(reader.result);
    //         resolve(json);
    //       } catch (error) {
    //         reject(error);
    //       }
    //     };
    //     reader.onerror = () => {
    //       reject(new Error("Error reading blob data"));
    //     };
    //     reader.readAsText(blob);
    //     console.log(reader.readAsText(blob))
    //   });
    // },
  },
};
</script>
