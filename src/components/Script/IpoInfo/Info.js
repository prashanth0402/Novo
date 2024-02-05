import EventServices from "@/services/EventServices";
export default {
  data: () => ({
    search: "",
    dialog: false,
    dialogDelete: false,
    checkBox: false,
    rules: [(value) => !!value || "Required."],
    headers1: [
      {
        text: "Symbol",
        align: "left",
        sortable: false,
        value: "symbol",
      },
      {
        text: "ISIN",
        align: "start",
        sortable: false,
        value: "isin",
      },
      { text: "BLOG Link", value: "blogLink", sortable: false },
      { text: "DRHP Link", value: "drhpLink", sortable: false },
      { text: "SME", align: "center", value: "sme", sortable: false },
      {
        text: "CreatedDate",
        align: "center",
        value: "createdDate",
        sortable: true,
      },
      {
        text: "CreatedBy",
        align: "center",
        value: "createdBy",
        sortable: false,
      },
      { text: "", align: "center", value: "actions", sortable: false },
    ],
    headers2: [
      {
        text: "ISIN",
        align: "start",
        sortable: false,
        value: "symbol",
      },
    ],
    linksArr: [],
    editedIndex: -1,
    editedItem: {
      id: 0,
      symbol: "",
      isin: "",
      blogLink: "",
      drhpLink: "",
      sme: true,
      allotmentFinal: "",
      refundInitiate: "",
      dematTransfer: "",
      listingDate: "",
    },
    defaultItem: {
      id: 0,
      symbol: "",
      isin: "",
      blogLink: "",
      drhpLink: "",
      sme: true,
      allotmentFinal: "",
      refundInitiate: "",
      dematTransfer: "",
      listingDate: "",
    },
    copy: [],
    loading: false,
    menu1: false,
    menu2: false,
    menu3: false,
    menu4: false,
  }),

  computed: {
    formTitle() {
      return this.editedIndex === -1 ? "Add Link" : "Edit Link";
    },
  },
  created() {
    this.initialize();
  },

  methods: {
    initialize() {
      this.loading = true;
      EventServices.GetBlogLink()
        .then((response) => {
          this.loading = false;
          if (response.data.status == "S") {
            this.linksArr = response.data.blogLinkArr;
          }
        })
        .catch((error) => {
          this.loading = false;
          this.MessageBar("E", error);
        });
    },

    editItem(item) {
      this.editedIndex = this.linksArr.indexOf(item);
      this.editedItem = Object.assign({}, item);
      this.dialog = true;
    },
    add() {
      this.dialog = true;
    },
    close() {
      // console.log("Closed")
      this.dialog = false;
      this.$nextTick(() => {
        this.editedItem = Object.assign({}, this.defaultItem);
        this.editedIndex = -1;
      });
      this.$refs.form.resetValidation();
    },
    save() {
      this.$refs.form.validate();
      if (
        this.editedItem.isin != "" &&
        this.editedItem.drhpLink != "" &&
        this.editedItem.allotmentFinal != "" &&
        this.editedItem.refundInitiate != "" &&
        this.editedItem.dematTransfer != "" &&
        this.editedItem.listingDate != ""
      ) {
        this.$globalData.overlay = true;
        EventServices.AddBlogLink(this.editedItem)
          .then((response) => {
            this.$globalData.overlay = false;
            if (response.data.status == "S") {
              this.initialize();
              this.MessageBar("S", "");
            } else {
              this.MessageBar("E", response.data.errMsg);
            }
          })
          .catch((error) => {
            this.$globalData.overlay = false;
            this.MessageBar("E", error);
          });
        this.close();
      } else {
        this.MessageBar("E", "Fill all the details");
      }
    },
    // validateDate(value) {
    //   const regex = /^\d{4}-\d{2}-\d{2}$/;

    //   if (!regex.test(value)) {
    //     console.log(value);
    //     // Handle invalid date format
    //     // You can display an error message or take other actions
    //   } else {
    //     console.log(value);
    //   }
    // },
  },
  watch: {
    enteredDate() {
      if (this.enteredDate.length == 4 || this.enteredDate.length == 7) {
        this.enteredDate += "-";
      }
    },
  },
};
