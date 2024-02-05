import Vue from "vue";
import VueRouter from "vue-router";
import Home from "@/views/Home.vue";

import Dashboard from "@/views/DashBoard.vue";
import Report from "@/views/Report.vue";
// import Configure from "@/views/Configure.vue";
import IpoInfo from "@/views/IpoInfo.vue";
import Sgb from "@/views/Sgb.vue";
import GSec from "@/views/PreGsec.vue";
import CorporateBond from "@/views/PreCorprate.vue"
import MutalFunds from "@/views/PreMutal.vue"

import Credential from "@/views/MemberDetail.vue";
import MemberUser from "@/views/MemberUser.vue";
import TaskMaster from "@/views/TaskMaster.vue"
import RoleMaster from "@/views/RoleMaster.vue"
import PageNotFound from "@/views/PageNotFound.vue"
import DomainMaster from "@/views/DomainMaster.vue"
import LookUp from "@/views/LookUp.vue"

import Fetchmaster from "@/components/ReportComp/ManualFetch/manualFetch.vue";

import Login from "@/views/PreDash.vue";
import LandingIpo from "../views/PreIpo.vue"
import LandingSgb from "@/views/PreSgb.vue"
import LandingGsec from "@/views/PreGsec.vue"
import LandingCorporateBond from "@/views/PreCorprate.vue"
import LandingMutalFunds from "@/views/PreMutal.vue"

import UnderConstruction from "@/views/UnderConstruction.vue"

import DashboardSetup from "@/components/Dash/DashboardSetup.vue"
import logdownload from "@/views/logfiledownload.vue"
import  versionControl  from "@/views/versioncontrol.vue";
import MasterControl from "@/views/MasterControl.vue"

import Ncb from "@/views/Ncb.vue";

Vue.use(VueRouter);

const routes = [
  {
    path: "/",
    name: "login",
    component: Login,
  },
  {
    path: "/pre/ipo",
    name: "Ipo",
    component: LandingIpo,
  },
  {
    path: "/pre/sgb",
    name: "Sgb",
    component: LandingSgb,
  },
  {
    path: "/pre/gsec",
    name: "G-Sec",
    component: LandingGsec,
  },
  {
    path: "/pre/corporateBond",
    name: "CorporateBond",
    component: LandingCorporateBond,
  },
  {
    path: "/pre/mutalFunds",
    name: "MutalFund",
    component: LandingMutalFunds,
  },
  {
    path: "/dashboard",
    name: "Dashboard",
    component: Dashboard,
  },
  {
    path: "/ipo",
    name: "ipo",
    component: Home,
    meta: {
      needsAuth: true,
      // title: 'IPO',
      metaTags: [
        {
          name: 'description',
          content: 'This is the main page of my appication.'
        }
      ]
    },
  },
  {
    path: "/gsec",
    name: "G-Secs",
    component: GSec,
  },
  {
    path: "/corporatebonds",
    name: "CorporateBonds",
    component: CorporateBond,
  },
  {
    path: "/mutalfunds",
    name: "MutalFunds",
    component: MutalFunds,
  },

  {
    path: "/report",
    name: "report",
    component: Report,
    meta: {
      needsAuth: true,
      // title: 'Report',
      metaTags: [
        {
          name: 'description',
          content: 'This Page is allowed for admin only.'
        }
      ]
    },
  },
  // {
  //   path: "/config",
  //   name: "configure",
  //   component: Configure,
  //   meta: {
  //     needsAuth: true,
  //   },
  // },
  {
    path: "/assigntask",
    name: "roletask",
    component: TaskMaster,
    meta: {
      needsAuth: true,
    },
  },
  {
    path: "/createtask",
    name: "roleMaster",
    component: RoleMaster,
    meta: {
      needsAuth: true,
    },
  },
  {
    path: "/ipoInfo",
    name: "ipoInfo",
    component: IpoInfo,
    meta: {
      needsAuth: true,
    },
  },
  {
    path: "/sgb",
    name: "sgb",
    component: Sgb,
    meta: {
      needsAuth: true,
    },
  },
  {
    path: "/credential",
    name: "credential",
    component: Credential,
    meta: {
      needsAuth: true,
    },
  },
  {
    path: "/managerole",
    name: "memberUser",
    component: MemberUser,
    meta: {
      needsAuth: true,
    },
  },
  {
    path: "/fetchmaster",
    name: "fetchmaster",
    component: Fetchmaster,
    meta: {
      needsAuth: true,
    }
  },
  {
    path: "/domainsetup",
    name: "domainsetup",
    component: DomainMaster,
    meta: {
      needsAuth: true,
    }
  },
  {
    path: "/lookup",
    name: "lookup",
    component: LookUp,
    meta: {
      needsAuth: true,
    }
  },
  {
    path: "*",
    name: "pagenotfound",
    component: PageNotFound,
    meta: { hideAppBar: true }, // Add this meta field
  },
  {
    path: "/underConstruction",
    name: "underConstruction",
    component: UnderConstruction,
    meta: { hideAppBar: true }, // Add this meta field
  },

  {
    path:"/dashboard/setup",
    name: "dashboardSetup",
    component:DashboardSetup
  },

  {
    path: "/logdownload",
    name: "logdownload",
    component: logdownload,
  },
  {
    path: "/versioncontrol",
    name: "versioncontrol",
    component: versionControl,
    meta: {
      needsAuth: true,
    }
  },
   {
    path: "/ncb",
    name: "ncb",
    component: Ncb,
    meta: {
      needsAuth: true,
    },
    },
    {
      path: "/mastercontrol",
      name: "MasterControl",
      component: MasterControl,
      meta: {
        needsAuth: true,
      }
    },
];

const router = new VueRouter({
  routes,
  mode: "history",
  base: process.env.BASE_URL,
  scrollBehavior(to, from, savedPosition) {
    // If there is a savedPosition (e.g., user pressed back/forward buttons), use it
    if (savedPosition) {
      return savedPosition;
    } else {
      // Otherwise, scroll to the top of the page
      return { x: 0, y: 0 };
    }
  },
});


export default router;
