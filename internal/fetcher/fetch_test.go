package fetcher

import (
	"strings"
	"testing"

	"github.com/Nikkely/GLSite/internal/model"
	"github.com/google/go-cmp/cmp"
)

func TestFetch(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want *[]model.Work
	}{
		{
			name: "ok",
			arg:  dummyHTML,
			want: &[]model.Work{
				{
					Name: "【アズールレーン】指揮官を癒やし隊！・綾波とゆっくり過ごす約1日間【ASMR】",
				},
				{
					Name: "絶対に寝かしつける!からかい上手な後輩彼女のあまあま安眠耳かき",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parse(strings.NewReader(tt.arg))
			if err != nil {
				t.Errorf(err.Error())
				return
			}
			// fmt.Printf("%v", got) // for develop
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("Hogefunc differs: (-got +want)\n%s", diff)
				return
			}
		})
	}
}

const dummyHTML = `<!DOCTYPE html>
<html lang="ja-jp">
<body class="style_home t_male">
    <div data-vue-component="header-banner" data-vue-async="true" data-section_name="campaign_header_banner">
    <div v-if="loading" class="hd_cp_banner type_1bn">
    <ul class="cp_bn_list">
            <li class="cp_bn_item" style="background-color: #FF22DA;">
        <a href="https://www.dlsite.com/home/campaign/sale2022gw" >
          <div class="cp_bn_inner">
            <div class="cp_bn_reminder " style="background-color: ; color: ;">
              <div class="cp_bn_reminder_content"><i class="cp_bn_reminder_period type_date">5/16</i><i class="cp_bn_reminder_period type_time">13:59</i>まで</div>
            </div>
            <div class="cp_bn"><img src="/modpub/images/campaign/super_sale_2204/90per/bn_hd_cp_90per_ja_jp.png" alt="2022年同人GWセール"></div>
                        <div class="cp_bn_work" style="background-color: ;"><div class="cp_bn_work_inner">対象作品<i class="cp_bn_work_count">16000</i>作品以上！</div></div>
                      </div>
        </a>
      </li>
          </ul>
  </div>
    
  <div v-else-if="is_show_frame" :class="style" class="hd_cp_banner" v-cloak>
    <ul class="cp_bn_list">
      <li v-for="campaign in campaigns" class="cp_bn_item" :data-ab-test="campaign.abTest" :style="{'background-color': campaign.mainColor}">
        <a :href="campaign.url" :target="campaign.option.includes('target_blank') ? 'blank' : null">
          <div class="cp_bn_inner">
            <div class="cp_bn_reminder" v-if="campaign.isShowLimit" :class="{blank: !campaign.isShowLimit}" :style="{'background-color': campaign.periodBgColor, 'color': campaign.periodTextColor}">
              <div class="cp_bn_reminder_content" v-html="campaign.period"></div>
            </div>
            <div class="cp_bn"><img :src="campaign.img" :alt="campaign.alt"></div>
            <div
              class="cp_bn_work"
              v-if="campaign.outline"
              v-html="campaign.outline"
              :style="{'background-color': campaign.periodBgColor}"
            >
            </div>
          </div>
        </a>
      </li>
    </ul>
  </div>
  
</div>

  <div data-vue-component="coupon-modal" v-cloak></div>

<div data-vue-component="header-locale-suggest-dialog" data-vue-async="true"></div>

<div data-vue-component="ie-alert" v-cloak></div>
  <!-- container -->
  <div id="container">

    <div id="top_header">
      
    </div>

    <!-- header -->
    <div id="header">
            <header class="l-header" data-vue-component="header" data-search-category="">
  <div class="headerCore">
    <div class="headerCore_top">
      <div class="headerCore_top_item">
        <p class="header_description">
                                    同人誌・同人ゲーム・同人ボイス・ASMRならDLsite                              </p>
      </div>
      <div class="headerCore_top_item">
                <ul v-if="isMemberLogin" v-cloak class="login_information">
          <li class="login_information_item type_point">
            <a href="https://www.dlsite.com/home/mypage/point" class="coupon_text">
            ポイント<span class="number" v-text="point_str">-</span>pt
            </a>
          </li>
          <li class="login_information_item type_coupon" :class="{ 'is-active': noticeCoupons.length, 'singular-number': coupons.length === 1 }">
            <a href="https://www.dlsite.com/home/mypage/coupon/list" class="coupon_text">クーポン<span class="number" v-text="coupons.length">-</span>枚</a>
          </li>
        </ul>
                <header_locale class="header_dropdown_nav type_language">
          <div class="header_dropdown_nav_Link">Language</div>
        </header_locale>
                        <header_service class="header_dropdown_nav type_service">
          <div class="header_dropdown_nav_Link">関連サービス</div>
        </header_service>
              </div>
    </div>

    <div class="headerCore-sub">
            <a href="https://www.dlsite.com/home/" class="logo">
                <img src="/images/web/common/logo/pc/logo-dlsite.png" alt="DLsite" width="89" height="25">
              </a>
            <!-- 検索 -->
              <div class="globalSearch" :class="{active: isFocus}">
          <form class="globalSearch-form" action="https://www.dlsite.com/home/fs" method="post">
            <input name="_qf__fulltext_search" type="hidden" value="">
            <input name="_layout" type="hidden" value="fs">
            <input name="_site" type="hidden" value="home">
            <input name="_form_id" type="hidden" value="FulltextSearchProductForm">
            <input name="_view" type="hidden" value="input">
            <input name="from" type="hidden" value="fs.header" id="header_search_from">

                                        <input type="hidden" name="site_category" :value="selected" v-if="!/drama$/.test(selected)">
                            <input type="hidden" name="work_category" :value="'drama'" v-if="/drama$/.test(selected)">
                            <input type="hidden" name="work_category" :value="'all'" v-if="site.gender === 'female' && !selected">
                            <input type="hidden" name="is_tl"  :value="1" v-if="/^girls/.test(selected)">
              <input type="hidden" name="is_bl"  :value="1" v-if="/^bl/.test(selected)">
              <input type="hidden" name="is_gay" :value="1" v-if="/^bl/.test(selected)">
            
            <div class="globalSearchSelect">
              <div class="globalSearchSelect-lable">
                <span v-text="searchCategories[selected]">すべて</span>
              </div>
              <select v-model="selected" class="globalSearchSelect-list is-active">
                <option value="">すべて</option>
                                <option value="home">同人</option>
                <option value="soft">PCソフト</option>
                <option value="app">アプリ</option>
                              </select>
            </div>
            <div class="globalSearchForm">
              <input
                name="keyword"
                type="search"
                id="search_text"
                ref="keyword"
                maxlength="255"
                autocomplete="off"
                aria-autocomplete="list"
                                placeholder="キーワードから探す（作品名、サークル名など）"                @focus="isFocus = true"
                @blur="isFocus = false"
              >
            </div>
            <button id="search_button" class="globalSearchBtn" type="submit" @click="saveSelected()"><i>検索</i></button>
          </form>
        </div>
        <div
          class="globalSearchBg"
          :class="{active: isFocus}"
        ></div>
        <div class="detailedSearch">
          <a href="https://www.dlsite.com/home/fs">こだわり条件</a>
        </div>
            <!-- /検索 -->
      <!-- アイコンメニュー -->
      <ul class="globalNav">
        <li v-if="!isMemberLogin && !isCircleLogin" class="globalNav-item type-register" v-cloak>
          <a href="https://www.dlsite.com/home/regist/user"><i>新規登録</i></a>
        </li>
        <li v-if="!isMemberLogin && !isCircleLogin" class="globalNav-item type-login" v-cloak>
          <a href="https://www.dlsite.com/home/login/=/skip_register/1" referrerpolicy="no-referrer-when-downgrade"><i>ログイン</i></a>
        </li>
        <li v-if="isMemberLogin" class="globalNav-item type-favorite">
          <a  :href="hasUnboughtFavorites ? 'https://www.dlsite.com/home/mypage/wishlist' : 'https://www.dlsite.com/home/mypage/wishlist'"><i>お気に入り</i></a><template v-if="hasUnboughtFavorites" v-cloak><a href="https://www.dlsite.com/home/mypage/wishlist/=/discount/1" class="notificationBadge" >割引作品あり</a></template>
        </li>
        <li class="globalNav-item type-cart"><a href="https://www.dlsite.com/home/cart"><i>カート</i></a><span v-if="cartActives.length" v-cloak class="cartBadge" v-text="Math.min(cartActives.length, 100)"></span></li>
        <li v-if="isMemberLogin" class="globalNav-item type-play">
          <a rel="noopener" href="https://play.dlsite.com/" target="_blank" v-cloak><i>購入済作品</i></a>
        </li>
        <!-- メンバーアカウント登録の場合　-->
        <li v-if="isMemberLogin" class="globalNav-item type-circle">
          <a v-cloak><i>アカウント</i></a>
          <div class="dropdown_list type_account" v-cloak>
            <div class="dropdown_list_inner">
              <ul class="globalNav">
                <li class="globalNav-item type-review">
                  <a href="https://www.dlsite.com/home/mypage/short-review">
                    <i>評価・レビュー</i>
                  </a>
                </li>
                <li class="globalNav-item type-coupon">
                  <a href="https://www.dlsite.com/home/mypage/coupon/list">
                    <i>クーポン管理</i>
                  </a>
                </li>
                <li class="globalNav-item type-userbuy">
                  <a href="https://www.dlsite.com/home/mypage/userbuy">
                    <i>購入履歴</i>
                  </a>
                </li>
              </ul>
              <ul class="menu_list">
                <li class="menu_list_item type-mypage">
                  <a href="https://www.dlsite.com/home/mypage" class="link">
                    <i>マイページ</i>
                    <span v-if="mypageNotices > 0" class="count" v-cloak>{{ noticesBadge(mypageNotices) }}</span>
                  </a>
                  <div v-if="mypageNotices > 0" class="notice">
                                        重要なお知らせが<span>{{ mypageNotices }}件</span>あります。                  </div>
                </li>
                <li v-if="isCircleLogin" class="menu_list_item type-circle" v-cloak>
                  <a href="https://www.dlsite.com/circle/" class="link">
                    <i>サークル管理</i>
                    <span v-if="circleNotices > 0" class="count" v-cloak>{{ noticesBadge(circleNotices) }}</span>
                  </a>
                  <div v-if="circleNotices > 0" class="notice">
                                        重要なお知らせが<span>{{ circleNotices }}件</span>あります。                  </div>
                </li>
                <li class="menu_list_item">
                  <a href="https://login.dlsite.com/user/self?lang=ja&redirect_uri=https%3A%2F%2Fwww.dlsite.com%2Fhome%2F&cancel_uri=https://www.dlsite.com/home/" class="link" target="_blank">
                    <i>アカウント管理</i>
                  </a>
                </li>
              </ul>
              <div class="login_account">
                <div class="login_account_item">
                  <p v-if="login_id" v-cloak class="account_name">
                                                            <span>{{ login_id.substr(0, 30) }}{{ login_id.length > 30 ? '...' : '' }}</span>
                    さん                  </p>
                </div>
                <div class="login_account_item">
                  <a href="https://www.dlsite.com/home/logout" class="logout">ログアウト</a>
                </div>
              </div>
            </div>
          </div>
          <span v-if="(mypageNotices + circleNotices) > 0 " class="accountBadge" v-text="noticesBadge(mypageNotices + circleNotices)" v-cloak></span>
        </li>
        <!-- サークルアカウントのみ登録の場合　-->
        <li v-if="isCircleLogin && !isMemberLogin" class="globalNav-item type-circle">
          <a v-cloak><i>アカウント</i></a>
          <div class="dropdown_list type_account" v-cloak>
            <div class="dropdown_list_inner">
              <ul class="menu_list">
                <li v-if="isCircleLogin" class="menu_list_item type-circle" v-cloak>
                  <a href="https://www.dlsite.com/circle/" class="link">
                    <i>サークル管理</i>
                    <span v-if="circleNotices > 0" class="count" v-cloak>{{ noticesBadge(circleNotices) }}</span>
                  </a>
                  <div v-if="circleNotices > 0" class="notice">
                                        重要なお知らせが<span>{{ circleNotices }}件</span>あります。                  </div>
                </li>
                <li class="menu_list_item">
                  <a href="https://login.dlsite.com/user/self?lang=ja&redirect_uri=https%3A%2F%2Fwww.dlsite.com%2Fhome%2F&cancel_uri=https://www.dlsite.com/home/" class="link" target="_blank">
                    <i>アカウント管理</i>
                  </a>
                </li>
              </ul>
              <div class="login_account">
                <div class="login_account_item">
                  <p v-if="login_id" v-cloak class="account_name">
                                                            <span>{{ login_id.substr(0, 30) }}{{ login_id.length > 30 ? '...' : '' }}</span>
                    さん                  </p>
                </div>
                <div class="login_account_item">
                  <a href="https://www.dlsite.com/home/logout" class="logout">ログアウト</a>
                </div>
              </div>
            </div>
          </div>
          <span v-if="(circleNotices) > 0 " class="accountBadge" v-text="noticesBadge(circleNotices)" v-cloak></span>
        </li>
        <li class="globalNav-item type-guide">
          <a href="#"><i>ガイド</i></a>
          <div class="dropdown_list type_guide">
            <div class="dropdown_list_inner">
              <ul class="menu_list">
                <li class="menu_list_item"><a rel="noopener" href="https://www.dlsite.com/home/faq/=/type/user" target="_blank" class="link">ヘルプ</a></li>
                <li class="menu_list_item"><a href="https://www.dlsite.com/home/welcome" class="link">初めての方へ</a></li>
                                <li class="menu_list_item"><a href="https://www.dlsite.com/home/circle/invite" class="link">サークル登録について</a></li>
                                <li class="menu_list_item"><a href="https://www.dlsite.com/home/guide/payment" class="link">お支払方法について</a></li>
                <li class="menu_list_item"><a href="https://www.dlsite.com/home/mypage/aboutpoint" class="link">ポイントについて</a></li>
              </ul>
            </div>
          </div>
        </li>
      </ul>
      <!-- /ガイドメニュー -->
    </div>

        <div class="headerCore-main">
        <div class="headerCore-mainInner " :class="{ couponShow: isAccountLogin }">
                <ul class="floorTab">
          <li class="floorTab-item type-doujin is-active"><a href="https://www.dlsite.com/home/">同人</a>     </li>
          <li class="floorTab-item type-com   "><a href="https://www.dlsite.com/soft/">PCソフト</a> </li>
          <li class="floorTab-item type-app   "><a href="https://www.dlsite.com/app/">アプリ</a></li>
          <li class="floorTab-item type-asmr "><a href="https://www.dlsite.com/home/asmr">ボイス・ASMR</a></li>
          <li class="floorTab-item type-tool "><a href="https://www.dlsite.com/home/tool">制作ソフト・素材</a></li>
          <li class="floorTab-item type-nijiyome"><a href="https://www.nijiyome.com/app?en=dlsite&em=tab&et=home">オンラインゲーム</a></li>
        </ul>
        
                <div class="floorNavLink">
                    <div class="floorNavLink-item type-general"><a rel="noopener" href="https://www.dlsite.com/comic/" target="_blank">一般コミックへ</a></div>
                              <div class="floorNavLink-item type-adult"><a href="https://www.dlsite.com/maniax/">男性向け R18へ</a></div>
                              <div class="floorNavLink-item type-female"><a href="https://www.dlsite.com/girls/">女性向けへ</a></div>
                  </div>
              </div>

            <div class="floorSubNav">
        <div class="floorSubNav-item">
          <ul class="headerNav">

      
                  <li class="headerNav-item"><a href="https://www.dlsite.com/home/" class="top">TOP</a></li>
                        <li class="headerNav-item">
              <a href="https://www.dlsite.com/home/works/type/=/work_type_category/game" class="game" :class="{'is-active': currentWorkTypeCategoryEquals('game')}">ゲーム</a>
            </li>
            <li class="headerNav-item">
              <a href="https://www.dlsite.com/home/works/type/=/work_type_category/movie" class="video" :class="{'is-active': currentWorkTypeCategoryEquals('movie')}">動画</a>
            </li>
            <li class="headerNav-item">
              <a href="https://www.dlsite.com/home/works/type/=/work_type_category/audio" class="voice" :class="{'is-active': currentWorkTypeCategoryEquals('audio')}">ボイス・ASMR</a>
            </li>
            <li class="headerNav-item">
              <a href="https://www.dlsite.com/home/works/type/=/work_type_category/comic" class="manga" :class="{'is-active': currentWorkTypeCategoryEquals('comic')}">マンガ</a>
            </li>
            <li class="headerNav-item">
              <a href="https://www.dlsite.com/home/works/type/=/work_type_category/illust" class="illust" :class="{'is-active': currentWorkTypeCategoryEquals('illust')}">CG・イラスト</a>
            </li>
            <li class="headerNav-item">
              <a href="https://www.dlsite.com/home/ranking?date=30d" class="ranking">ランキング</a>
            </li>
            <li class="headerNav-item"><a href="https://www.dlsite.com/home/announce/list/day" class="announce">発売予告作品</a></li>
            <li class="headerNav-item"><a href="https://www.dlsite.com/home/new" class="calendar">発売カレンダー</a></li>

                                    
                            </ul>
        </div>
      </div>
      
    </div>
      </div>
</header>




            <script type="text/javascript">
var loginchecked = $.cookie('loginchecked');

if (loginchecked & 1) {
  // ユーザがログイン中の表示
  $("li#nav_login a").attr("title", "ログアウト").attr("href", "https://www.dlsite.com/home/logout/=/type/member");
  $("li#nav_login").attr("id", "nav_logout");
} else {
  // ユーザがログアウト中の表示
//  $("li#nav_logout a").attr("title", "ログイン").attr("href", "https://www.dlsite.com/home/login/=/type/member");
//  $("li#nav_logout").attr("id", "nav_login");
}
</script>

      
    </div>
    <!-- /header -->

        <!-- top_wrapper -->
    <div id="top_wrapper">
      <ul class="topicpath" itemscope="" itemtype="https://schema.org/BreadcrumbList">
  <li class="topicpath_item" itemprop="itemListElement" itemscope="" itemtype="https://schema.org/ListItem">
    <a itemtype="https://schema.org/Thing" itemprop="item" href="https://www.dlsite.com/home/">
    <span itemprop="name">
              同人
          </span>
    </a>
    <meta itemprop="position" content="1">
      </li>
                        
                        <li class="topicpath_item" itemprop="itemListElement" itemscope="" itemtype="https://schema.org/ListItem">
          <a itemtype="https://schema.org/Thing" itemprop="item" href="https://www.dlsite.com/home/worktype/list">
            <span itemprop="name">作品形式一覧</span>
          </a>
          <meta itemprop="position" content="2">
                  </li>
              <li class="topicpath_item" itemprop="itemListElement" itemscope="" itemtype="https://schema.org/ListItem">
          <a itemtype="https://schema.org/Thing" itemprop="item" href="https://www.dlsite.com/home/works/type/=/work_type_category/audio">
            <span itemprop="name">「ボイス・ASMR」の作品一覧</span>
          </a>
          <meta itemprop="position" content="3">
                  </li>
      
      </ul>
    </div>
    <!-- /top_wrapper -->
    
    <!-- wrapper -->
    <div id="wrapper" >

      
      <!-- main -->
      <div id="main">
        <!-- main_inner -->
        <div id="main_inner" >
          
            <div class="base_title_br clearfix"><h1><span class="original_name">検索結果</span></h1></div>

                    <div class="base_text_10">
  <p>
          DLsiteでは「ボイス・ASMR」の同人作品、3909件を販売中！      </p>
</div>
  
        <div class="top_center_banner male home" data-section_name="center_banner" data-top-center-banner-key="dlsiteworks_home_center2-audio">
  <div class="banner_container">
    <div class="top_banner_prev"><i></i></div>
    <div class="top_banner_next"><i></i></div>
    <div class="pagination pagenation_container"></div>
  </div>
</div>

<script>
if (/^.*(msie|trident\/).*$/i.test(window.navigator.userAgent)) {
  document.write("<script src='https:\/\/polyfill.io\/v2\/polyfill.min.js?features=IntersectionObserver'><\/script>");
}
</script>

<script src="/js/topbannerslider/center-banner.js?1638951819" async></script>
  
<div class="search_condition_box">

<table class="search_condition" cellspacing="0">
  <tr>
    <th>キーワード</th>
    <td>
      <form action="https://www.dlsite.com/home/fsr/=/language/jp/age_category%5B0%5D/general/work_category%5B0%5D/doujin/order%5B0%5D/trend/work_type_category%5B0%5D/audio/work_type_category_name%5B0%5D/%E3%83%9C%E3%82%A4%E3%82%B9%E3%83%BBASMR/options_and_or/and/options%5B0%5D/JPN/options%5B1%5D/NM/per_page/100/show_type/3/lang_options%5B0%5D/%E6%97%A5%E6%9C%AC%E8%AA%9E/lang_options%5B1%5D/%E8%A8%80%E8%AA%9E%E4%B8%8D%E8%A6%81" method="post" class="_research_form">
        <input type="text" id="research_text" name="keyword" value=""/><button value="search" type="submit" id="research_button">再検索</button>
      </form>
    </td>
    <td class="search_condition_2col" rowspan="3">
      <p class="to_search_details"><a href="https://www.dlsite.com/home/fs/=/language/jp/age_category%5B0%5D/general/work_category%5B0%5D/doujin/order%5B0%5D/trend/work_type_category%5B0%5D/audio/work_type_category_name%5B0%5D/%E3%83%9C%E3%82%A4%E3%82%B9%E3%83%BBASMR/options_and_or/and/options%5B0%5D/JPN/options%5B1%5D/NM/per_page/100/show_type/3/lang_options%5B0%5D/%E6%97%A5%E6%9C%AC%E8%AA%9E/lang_options%5B1%5D/%E8%A8%80%E8%AA%9E%E4%B8%8D%E8%A6%81">検索条件を変更する</a></p>
      <div data-vue-component="keyword-register" data-vue-async="true">

    <div v-if="!loggedin" class="save_search_condition_button" v-on:click="login()" v-show="!registered" v-cloak>
      検索条件の保存    </div>
    <div v-else class="save_search_condition_button" v-on:click="openPopup()" v-show="!registered" v-cloak>
      検索条件の保存    </div>
  <a href="https://www.dlsite.com/home/mypage/mygenre">
    <div class="save_search_condition_button btn_search_in" v-show="registered" v-cloak>検索条件の保存済み</div>
  </a>

    <div v-if="loggedin" class="popup_savesearch_add" :class="{'on':shownPopup}" v-on:click.stop>
    <p class="popup_close_button"><a href="#" v-on:click.prevent="hidePopup()">close</a></p>
    <table class="popup_my_genre search_condition" cellspacing="0">
      <tbody>
      <tr>
        <td colspan="2" class="box_left"><p>現在の検索キーワード・検索条件を保存していつでも呼び出せます。</p></td>
      </tr>
      <tr>
        <td><p>検索条件の保存名</p></td>
        <td><input name="mygenre_name" type="text" class="popup_my_genre_text" :class="{'input_error':errorMessage}" style="color: rgb(27, 27, 27);" v-model="title"></td>
      </tr>
      <tr v-show="errorMessage">
        <td colspan="2" class="pb0 pt0">
          <ul class="error_list">
            <li style="text-align: center;" v-text="errorMessage"></li>
          </ul>
        </td>
      </tr>
      <tr>
        <td class="popup_my_genre_text" colspan="2">
          <a href="#" class="savesearch_regist mt5" v-on:click.prevent="send('https://www.dlsite.com/home/fsr/=/language/jp/age_category%5B0%5D/general/work_category%5B0%5D/doujin/order%5B0%5D/trend/work_type_category%5B0%5D/audio/work_type_category_name%5B0%5D/%E3%83%9C%E3%82%A4%E3%82%B9%E3%83%BBASMR/options_and_or/and/options%5B0%5D/JPN/options%5B1%5D/NM/per_page/100/show_type/3/lang_options%5B0%5D/%E6%97%A5%E6%9C%AC%E8%AA%9E/lang_options%5B1%5D/%E8%A8%80%E8%AA%9E%E4%B8%8D%E8%A6%81')">検索条件の保存</a>
        </td>
      </tr>
      </tbody>
    </table>
  </div>
</div>
    </td>
  </tr>

    <tr>
    <th>検索条件</th>
        <td><ul class="condition_list">
        
  
  
  
  
  

  
  
  
  
  
      <li><dl><dt>年齢指定:</dt><dd>
          全年齢向け
                    </dd></dl></li>
  
      <li><dl><dt>対象作品:</dt><dd>
          同人作品
                    </dd></dl></li>
  
      <li><dl><dt>作品形式:</dt><dd>
                  ボイス・ASMR
                        
        </dd></dl></li>
  
  
  
      <li><dl><dt>対応言語:</dt><dd>
              日本語
        &nbsp;              言語不要
                  </dd></dl></li>
  
    
  
  
  
    </ul></td>
  </tr>
      </table>

</div>

<script type="text/javascript">
//<!CDATA[

jQuery(function($){
    var action = 'https://www.dlsite.com/home/fsr/=/language/jp/age_category%5B0%5D/general/work_category%5B0%5D/doujin/order%5B0%5D/trend/work_type_category%5B0%5D/audio/work_type_category_name%5B0%5D/%E3%83%9C%E3%82%A4%E3%82%B9%E3%83%BBASMR/options_and_or/and/options%5B0%5D/JPN/options%5B1%5D/NM/per_page/100/show_type/3/lang_options%5B0%5D/%E6%97%A5%E6%9C%AC%E8%AA%9E/lang_options%5B1%5D/%E8%A8%80%E8%AA%9E%E4%B8%8D%E8%A6%81/from/fsr.again';

    $('._research_form').submit(function(){
        var keyword = $('input[name=keyword]',this).val();
        location.href = action + (keyword ? '/keyword/'+ encodeURIComponent(keyword) : '');
        return false;
    });
});

//]]>
</script>






<div data-vue-component="campaign-tutorial-clear" data-api_url="https://www.dlsite.com/home/campaign/api/=/id/tutorials/mid/1" data-func="missionSearch" data-vue-async="true"></div>

<div class="sort_box border_b pb10">
  <div class="status_select">並び替え&nbsp;:
        <select name="order" id="query_order" class="_change_submit">
<option value="trend" selected="selected">人気順</option>
<option value="release_d" >発売日が新しい順</option>
<option value="release" >発売日が古い順</option>
<option value="dl_d" >販売数が多い順</option>
<option value="price" >価格が安い順</option>
<option value="price_d" >価格が高い順</option>
<option value="rate_d" >評価が高い順</option>
<option value="review_d" >レビューが多い順</option>
</select>
  </div>

  
      <div class="page_total">
    <strong>3909</strong>件中&nbsp;<strong>1～100</strong>件目
  </div>
  

  <div class="display_type_select">
    <span>表示形式&nbsp;：</span>
    <ul>
      <li class="display_normal"><a href="https://www.dlsite.com/home/works/type/=/language/jp/age_category%5B0%5D/general/work_category%5B0%5D/doujin/order%5B0%5D/trend/work_type_category%5B0%5D/audio/work_type_category_name%5B0%5D/%E3%83%9C%E3%82%A4%E3%82%B9%E3%83%BBASMR/options_and_or/and/options%5B0%5D/JPN/options%5B1%5D/NM/per_page/100/show_type/1/lang_options%5B0%5D/%E6%97%A5%E6%9C%AC%E8%AA%9E/lang_options%5B1%5D/%E8%A8%80%E8%AA%9E%E4%B8%8D%E8%A6%81/logged/1" data-value="1">通常表示</a></li>
      <li class="display_block on"><a href="https://www.dlsite.com/home/works/type/=/language/jp/age_category%5B0%5D/general/work_category%5B0%5D/doujin/order%5B0%5D/trend/work_type_category%5B0%5D/audio/work_type_category_name%5B0%5D/%E3%83%9C%E3%82%A4%E3%82%B9%E3%83%BBASMR/options_and_or/and/options%5B0%5D/JPN/options%5B1%5D/NM/per_page/100/show_type/3/lang_options%5B0%5D/%E6%97%A5%E6%9C%AC%E8%AA%9E/lang_options%5B1%5D/%E8%A8%80%E8%AA%9E%E4%B8%8D%E8%A6%81/logged/1" data-value="3">画像のみ</a></li>
    </ul>
  </div>

  <div class="display_num_select">
    <ul>
      <li class=""><a href="https://www.dlsite.com/home/works/type/=/language/jp/age_category%5B0%5D/general/work_category%5B0%5D/doujin/order%5B0%5D/trend/work_type_category%5B0%5D/audio/work_type_category_name%5B0%5D/%E3%83%9C%E3%82%A4%E3%82%B9%E3%83%BBASMR/options_and_or/and/options%5B0%5D/JPN/options%5B1%5D/NM/per_page/30/show_type/3/lang_options%5B0%5D/%E6%97%A5%E6%9C%AC%E8%AA%9E/lang_options%5B1%5D/%E8%A8%80%E8%AA%9E%E4%B8%8D%E8%A6%81/page/1" data-value="30">30</a></li>
      <li class=""><a href="https://www.dlsite.com/home/works/type/=/language/jp/age_category%5B0%5D/general/work_category%5B0%5D/doujin/order%5B0%5D/trend/work_type_category%5B0%5D/audio/work_type_category_name%5B0%5D/%E3%83%9C%E3%82%A4%E3%82%B9%E3%83%BBASMR/options_and_or/and/options%5B0%5D/JPN/options%5B1%5D/NM/per_page/50/show_type/3/lang_options%5B0%5D/%E6%97%A5%E6%9C%AC%E8%AA%9E/lang_options%5B1%5D/%E8%A8%80%E8%AA%9E%E4%B8%8D%E8%A6%81/page/1" data-value="50">50</a></li>
      <li class="on"><a href="https://www.dlsite.com/home/works/type/=/language/jp/age_category%5B0%5D/general/work_category%5B0%5D/doujin/order%5B0%5D/trend/work_type_category%5B0%5D/audio/work_type_category_name%5B0%5D/%E3%83%9C%E3%82%A4%E3%82%B9%E3%83%BBASMR/options_and_or/and/options%5B0%5D/JPN/options%5B1%5D/NM/per_page/100/show_type/3/lang_options%5B0%5D/%E6%97%A5%E6%9C%AC%E8%AA%9E/lang_options%5B1%5D/%E8%A8%80%E8%AA%9E%E4%B8%8D%E8%A6%81/page/1" data-value="100">100</a></li>
    </ul>
    <span>表示件数&nbsp;：</span>
  </div>
</div>

<div id="search_result_list" class="_search_result_list">
      
      <ul id="search_result_img_box" class="n_worklist">
            
                        
        
            <li class="search_result_img_box_inner type_exclusive_01 ">
<dl class="work_img_main"><dt class="search_img work_thumb" id="_link_RJ387206"><a href="https://www.dlsite.com/home/work/=/product_id/RJ387206.html" class="work_thumb_inner" data-vue-component="thumb-img-popup"  ><img class="lazy" alt="【アズールレーン】指揮官を癒やし隊！・綾波とゆっくり過ごす約1日間【ASMR】 [アトリエメール]" src="//img.dlsite.jp/resize/images2/work/doujin/RJ388000/RJ387206_img_main_240x240.jpg" ref="popup_img" @mouseenter="showPopupImg"><div v-cloak class="work_img_popover " :class="{ flip: is_flip }"><img src="data:image/gif;base64,R0lGODlhAQABAGAAACH5BAEKAP8ALAAAAAABAAEAAAgEAP8FBAA7" :src="is_show ? '//img.dlsite.jp/modpub/images2/work/doujin/RJ388000/RJ387206_img_main.jpg' : 'data:image/gif;base64,R0lGODlhAQABAGAAACH5BAEKAP8ALAAAAAABAAEAAAgEAP8FBAA7'" alt="【アズールレーン】指揮官を癒やし隊！・綾波とゆっくり過ごす約1日間【ASMR】 [アトリエメール]"></div></a></dt><dd class="work_category_free_sample"><div class="work_category type_SOU type_free_sample"><a href="https://www.dlsite.com/home/fsr/=/work_type/SOU">ボイス・ASMR</a></div><span data-vue-component="SampleViewMiniButton"data-product_id="RJ387206"data-view_samples="[{'thumb':'\/\/img.dlsite.jp\/modpub\/images2\/work\/doujin\/RJ388000\/RJ387206_img_main.jpg','width':551,'height':420},{'thumb':'\/\/img.dlsite.jp\/modpub\/images2\/work\/doujin\/RJ388000\/RJ387206_img_smp1.jpg','width':1638,'height':1250},{'thumb':'\/\/img.dlsite.jp\/modpub\/images2\/work\/doujin\/RJ388000\/RJ387206_img_smp2.jpg','width':1638,'height':1250},{'thumb':'\/\/img.dlsite.jp\/modpub\/images2\/work\/doujin\/RJ388000\/RJ387206_img_smp3.jpg','width':1638,'height':1250}]"data-worktype="SOU"></span></dd><dd class="work_name"><div class="icon_wrap"></div><div class="multiline_truncate"><a href="https://www.dlsite.com/home/work/=/product_id/RJ387206.html" title="【アズールレーン】指揮官を癒やし隊！・綾波とゆっくり過ごす約1日間【ASMR】">【アズールレーン】指揮官を癒やし隊！・綾波とゆっくり過ごす約1日間【ASMR】</a></div></dd><dd class="maker_name"><a href="https://www.dlsite.com/home/circle/profile/=/maker_id/RG64308.html">アトリエメール</a><span class="separator">/</span><span class="author"><a href="https://www.dlsite.com/home/fsr/=/keyword_creater/%22%E5%A4%A7%E5%9C%B0%E8%91%89%22" class="">大地葉</a></span></dd><dd class="work_price_wrap"><span class="work_price discount">1,320<i>円</i></span><span class="strike">1,650<i>円</i></span><span class="separator">/</span><span class="work_point">120pt</span></dd><dd class="work_genre"></dd><input type="hidden" class="__product_attributes" name="__product_attributes" id="_RJ387206" value="RG64308,male,SOU,SND,JPN,TRI,DLP,REV,497,056,496,051,008,503" disabled="disabled"><dd class="work_dl">販売数:&nbsp;<span class="_dl_count_RJ387206">7,590</span></dd><dd class="work_rating"><div class="star_rating star_50 mini">(1,312)</div><div class="work_review"><div title="レビューあり"><a href="https://www.dlsite.com/home/work/reviewlist/=/product_id/RJ387206.html">(25)</a></div></div></dd><dd class="work_deals work_labels">
<span class="icon_lead_01 type_sale">20%OFF</span></dd></dl><div data-vue-component="product-item" data-product_id="RJ387206" data-layout="image"><ul class="work_operation_btn table-ul"><li><p class="work_cart"><a href="https://www.dlsite.com/home/cart/=/product_id/RJ387206.html" class="btn_cart _btn_cart " id="_btn_cart_RJ387206">カートに追加</a></p></li><li><p class="work_favorite"><a href="https://www.dlsite.com/home/mypage/wishlist/=/product_id/RJ387206.html" class="btn_favorite" id="_btn_favorite_RJ387206">お気に入りに追加</a></p></li></ul></div>
</li>
                
                        
        
            <li class="search_result_img_box_inner type_exclusive_01 ">
<dl class="work_img_main"><dt class="search_img work_thumb" id="_link_RJ385913"><a href="https://www.dlsite.com/home/work/=/product_id/RJ385913.html" class="work_thumb_inner" data-vue-component="thumb-img-popup"  ><img class="lazy" alt="絶対に寝かしつける!からかい上手な後輩彼女のあまあま安眠耳かき [いちのや]" src="//img.dlsite.jp/resize/images2/work/doujin/RJ386000/RJ385913_img_main_240x240.jpg" ref="popup_img" @mouseenter="showPopupImg"><div v-cloak class="work_img_popover " :class="{ flip: is_flip }"><img src="data:image/gif;base64,R0lGODlhAQABAGAAACH5BAEKAP8ALAAAAAABAAEAAAgEAP8FBAA7" :src="is_show ? '//img.dlsite.jp/modpub/images2/work/doujin/RJ386000/RJ385913_img_main.jpg' : 'data:image/gif;base64,R0lGODlhAQABAGAAACH5BAEKAP8ALAAAAAABAAEAAAgEAP8FBAA7'" alt="絶対に寝かしつける!からかい上手な後輩彼女のあまあま安眠耳かき [いちのや]"></div></a></dt><dd class="work_category_free_sample"><div class="work_category type_SOU type_free_sample"><a href="https://www.dlsite.com/home/fsr/=/work_type/SOU">ボイス・ASMR</a></div><span data-vue-component="SampleViewMiniButton"data-product_id="RJ385913"data-view_samples="[{'thumb':'\/\/img.dlsite.jp\/modpub\/images2\/work\/doujin\/RJ386000\/RJ385913_img_main.jpg','width':560,'height':420},{'thumb':'\/\/img.dlsite.jp\/modpub\/images2\/work\/doujin\/RJ386000\/RJ385913_img_smp1.jpg','width':1680,'height':1260},{'thumb':'\/\/img.dlsite.jp\/modpub\/images2\/work\/doujin\/RJ386000\/RJ385913_img_smp2.jpg','width':1680,'height':1260},{'thumb':'\/\/img.dlsite.jp\/modpub\/images2\/work\/doujin\/RJ386000\/RJ385913_img_smp3.jpg','width':1680,'height':1260}]"data-worktype="SOU"></span></dd><dd class="work_name"><div class="icon_wrap"></div><div class="multiline_truncate"><a href="https://www.dlsite.com/home/work/=/product_id/RJ385913.html" title="絶対に寝かしつける!からかい上手な後輩彼女のあまあま安眠耳かき">絶対に寝かしつける!からかい上手な後輩彼女のあまあま安眠耳かき</a></div></dd><dd class="maker_name"><a href="https://www.dlsite.com/home/circle/profile/=/maker_id/RG46007.html">いちのや</a><span class="separator">/</span><span class="author"><a href="https://www.dlsite.com/home/fsr/=/keyword_creater/%22%E4%B8%80%E4%B9%8B%E7%80%AC%E3%82%8A%E3%81%A8%22" class="">一之瀬りと</a></span></dd><dd class="work_price_wrap"><span class="work_price discount">1,056<i>円</i></span><span class="strike">1,320<i>円</i></span><span class="separator">/</span><span class="work_point">96pt</span></dd><dd class="work_genre"></dd><input type="hidden" class="__product_attributes" name="__product_attributes" id="_RJ385913" value="RG46007,male,SOU,SND,JPN,DLP,TRI,REV,056,497,014,503,285,496,004,008" disabled="disabled"><dd class="work_dl">販売数:&nbsp;<span class="_dl_count_RJ385913">4,055</span></dd><dd class="work_rating"><div class="star_rating star_50 mini">(856)</div><div class="work_review"><div title="レビューあり"><a href="https://www.dlsite.com/home/work/reviewlist/=/product_id/RJ385913.html">(8)</a></div></div></dd><dd class="work_deals work_labels">
<span class="icon_lead_01 type_sale">20%OFF</span></dd></dl><div data-vue-component="product-item" data-product_id="RJ385913" data-layout="image"><ul class="work_operation_btn table-ul"><li><p class="work_cart"><a href="https://www.dlsite.com/home/cart/=/product_id/RJ385913.html" class="btn_cart _btn_cart " id="_btn_cart_RJ385913">カートに追加</a></p></li><li><p class="work_favorite"><a href="https://www.dlsite.com/home/mypage/wishlist/=/product_id/RJ385913.html" class="btn_favorite" id="_btn_favorite_RJ385913">お気に入りに追加</a></p></li></ul></div>
</li>
          </ul>
  </div>
    <table cellspacing="0" class="global_pagination">
    <tbody>
      <tr>
        <td class="page_no">
                    <ul>
                                                    <li><strong>1</strong></li>
                                                      <li><a href="https://www.dlsite.com/home/works/type/=/language/jp/age_category%5B0%5D/general/work_category%5B0%5D/doujin/order%5B0%5D/trend/work_type_category%5B0%5D/audio/work_type_category_name%5B0%5D/%E3%83%9C%E3%82%A4%E3%82%B9%E3%83%BBASMR/options_and_or/and/options%5B0%5D/JPN/options%5B1%5D/NM/per_page/100/show_type/3/lang_options%5B0%5D/%E6%97%A5%E6%9C%AC%E8%AA%9E/lang_options%5B1%5D/%E8%A8%80%E8%AA%9E%E4%B8%8D%E8%A6%81/page/2" data-value="2">2</a></li>
                                                      <li><a href="https://www.dlsite.com/home/works/type/=/language/jp/age_category%5B0%5D/general/work_category%5B0%5D/doujin/order%5B0%5D/trend/work_type_category%5B0%5D/audio/work_type_category_name%5B0%5D/%E3%83%9C%E3%82%A4%E3%82%B9%E3%83%BBASMR/options_and_or/and/options%5B0%5D/JPN/options%5B1%5D/NM/per_page/100/show_type/3/lang_options%5B0%5D/%E6%97%A5%E6%9C%AC%E8%AA%9E/lang_options%5B1%5D/%E8%A8%80%E8%AA%9E%E4%B8%8D%E8%A6%81/page/3" data-value="3">3</a></li>
                                                      <li><a href="https://www.dlsite.com/home/works/type/=/language/jp/age_category%5B0%5D/general/work_category%5B0%5D/doujin/order%5B0%5D/trend/work_type_category%5B0%5D/audio/work_type_category_name%5B0%5D/%E3%83%9C%E3%82%A4%E3%82%B9%E3%83%BBASMR/options_and_or/and/options%5B0%5D/JPN/options%5B1%5D/NM/per_page/100/show_type/3/lang_options%5B0%5D/%E6%97%A5%E6%9C%AC%E8%AA%9E/lang_options%5B1%5D/%E8%A8%80%E8%AA%9E%E4%B8%8D%E8%A6%81/page/4" data-value="4">4</a></li>
                                                      <li><a href="https://www.dlsite.com/home/works/type/=/language/jp/age_category%5B0%5D/general/work_category%5B0%5D/doujin/order%5B0%5D/trend/work_type_category%5B0%5D/audio/work_type_category_name%5B0%5D/%E3%83%9C%E3%82%A4%E3%82%B9%E3%83%BBASMR/options_and_or/and/options%5B0%5D/JPN/options%5B1%5D/NM/per_page/100/show_type/3/lang_options%5B0%5D/%E6%97%A5%E6%9C%AC%E8%AA%9E/lang_options%5B1%5D/%E8%A8%80%E8%AA%9E%E4%B8%8D%E8%A6%81/page/5" data-value="5">5</a></li>
                                                  <li><a href="https://www.dlsite.com/home/works/type/=/language/jp/age_category%5B0%5D/general/work_category%5B0%5D/doujin/order%5B0%5D/trend/work_type_category%5B0%5D/audio/work_type_category_name%5B0%5D/%E3%83%9C%E3%82%A4%E3%82%B9%E3%83%BBASMR/options_and_or/and/options%5B0%5D/JPN/options%5B1%5D/NM/per_page/100/show_type/3/lang_options%5B0%5D/%E6%97%A5%E6%9C%AC%E8%AA%9E/lang_options%5B1%5D/%E8%A8%80%E8%AA%9E%E4%B8%8D%E8%A6%81/page/2" data-value="2">次へ</a></li>
            <li><a href="https://www.dlsite.com/home/works/type/=/language/jp/age_category%5B0%5D/general/work_category%5B0%5D/doujin/order%5B0%5D/trend/work_type_category%5B0%5D/audio/work_type_category_name%5B0%5D/%E3%83%9C%E3%82%A4%E3%82%B9%E3%83%BBASMR/options_and_or/and/options%5B0%5D/JPN/options%5B1%5D/NM/per_page/100/show_type/3/lang_options%5B0%5D/%E6%97%A5%E6%9C%AC%E8%AA%9E/lang_options%5B1%5D/%E8%A8%80%E8%AA%9E%E4%B8%8D%E8%A6%81/page/40" data-value="40">最後へ</a></li>
                      </ul>
                  </td>
      </tr>
    </tbody>
  </table>
  

<div class="result_contents">
    
          <div class="search_more">
      <p>作品が見つかりませんか？</p>
      <ul>
                                            <li>
                <span>
                  <a href="https://www.dlsite.com/maniax/fsr/=/language/jp/age_category%5B0%5D/general/age_category%5B1%5D/r15/age_category%5B2%5D/adult/work_category%5B0%5D/doujin/order%5B0%5D/trend/work_type_category%5B0%5D/audio/work_type_category_name%5B0%5D/%E3%83%9C%E3%82%A4%E3%82%B9%E3%83%BBASMR/options_and_or/and/options%5B0%5D/JPN/options%5B1%5D/NM/options_name%5B0%5D/%E6%97%A5%E6%9C%AC%E8%AA%9E%E4%BD%9C%E5%93%81/options_name%5B1%5D/%E8%A8%80%E8%AA%9E%E4%B8%8D%E5%95%8F%E4%BD%9C%E5%93%81/per_page/100/from/fsr.more">
                    成人男性向け作品を含めて検索する
                    <span v-if="count > 0" v-cloak class="search_more_number" data-vue-component="search_result" data-vue-async="true" data-url="https://www.dlsite.com/maniax/sapi/=/language/jp/age_category%5B0%5D/general/age_category%5B1%5D/r15/age_category%5B2%5D/adult/work_category%5B0%5D/doujin/order%5B0%5D/trend/work_type_category%5B0%5D/audio/work_type_category_name%5B0%5D/%E3%83%9C%E3%82%A4%E3%82%B9%E3%83%BBASMR/options_and_or/and/options%5B0%5D/JPN/options%5B1%5D/NM/options_name%5B0%5D/%E6%97%A5%E6%9C%AC%E8%AA%9E%E4%BD%9C%E5%93%81/options_name%5B1%5D/%E8%A8%80%E8%AA%9E%E4%B8%8D%E5%95%8F%E4%BD%9C%E5%93%81/per_page/100/format/json/?cdn_cache=1">
                      (<span class="number">{{ count | number_format }}</span>件)
                    </span>
                  </a>
                </span>
              </li>
                              <li>
                  <span>
                    <a href="https://www.dlsite.com/app/fsr/=/language/jp/age_category%5B0%5D/general/work_category/app/order%5B0%5D/trend/work_type_category%5B0%5D/audio/work_type_category_name%5B0%5D/%E3%83%9C%E3%82%A4%E3%82%B9%E3%83%BBASMR/options_and_or/and/options%5B0%5D/JPN/options%5B1%5D/NM/per_page/100/show_type/3/lang_options%5B0%5D/%E6%97%A5%E6%9C%AC%E8%AA%9E/lang_options%5B1%5D/%E8%A8%80%E8%AA%9E%E4%B8%8D%E8%A6%81/from/fsr.more">
                      アプリ作品を検索する
                      <span v-if="count > 0" v-cloak class="search_more_number" data-vue-component="search_result" data-vue-async="true" data-url="https://www.dlsite.com/app/sapi/=/language/jp/age_category%5B0%5D/general/work_category/app/order%5B0%5D/trend/work_type_category%5B0%5D/audio/work_type_category_name%5B0%5D/%E3%83%9C%E3%82%A4%E3%82%B9%E3%83%BBASMR/options_and_or/and/options%5B0%5D/JPN/options%5B1%5D/NM/per_page/100/show_type/3/lang_options%5B0%5D/%E6%97%A5%E6%9C%AC%E8%AA%9E/lang_options%5B1%5D/%E8%A8%80%E8%AA%9E%E4%B8%8D%E8%A6%81/format/json/?cdn_cache=1">
                        (<span class="number">{{ count | number_format }}</span>件)
                      </span>
                    </a>
                  </span>
                </li>
                            <li>
                <span>
                  <a href="https://www.dlsite.com/girls/fsr/=/language/jp/age_category%5B0%5D/general/age_category%5B1%5D/r15/age_category%5B2%5D/adult/work_category%5B0%5D/doujin/order%5B0%5D/trend/work_type_category%5B0%5D/audio/work_type_category_name%5B0%5D/%E3%83%9C%E3%82%A4%E3%82%B9%E3%83%BBASMR/options_and_or/and/options%5B0%5D/JPN/options%5B1%5D/NM/options_name%5B0%5D/%E6%97%A5%E6%9C%AC%E8%AA%9E%E4%BD%9C%E5%93%81/options_name%5B1%5D/%E8%A8%80%E8%AA%9E%E4%B8%8D%E5%95%8F%E4%BD%9C%E5%93%81/per_page/100/from/fsr.more">
                    女性向け TL/乙女向け作品を検索する
                    <span v-if="count > 0" v-cloak class="search_more_number" data-vue-component="search_result" data-vue-async="true" data-url="https://www.dlsite.com/girls/sapi/=/language/jp/age_category%5B0%5D/general/age_category%5B1%5D/r15/age_category%5B2%5D/adult/work_category%5B0%5D/doujin/order%5B0%5D/trend/work_type_category%5B0%5D/audio/work_type_category_name%5B0%5D/%E3%83%9C%E3%82%A4%E3%82%B9%E3%83%BBASMR/options_and_or/and/options%5B0%5D/JPN/options%5B1%5D/NM/options_name%5B0%5D/%E6%97%A5%E6%9C%AC%E8%AA%9E%E4%BD%9C%E5%93%81/options_name%5B1%5D/%E8%A8%80%E8%AA%9E%E4%B8%8D%E5%95%8F%E4%BD%9C%E5%93%81/per_page/100/format/json/?cdn_cache=1">
                      (<span class="number">{{ count | number_format }}</span>件)
                    </span>
                  </a>
                </span>
              </li>
              <li>
                <span>
                  <a href="https://www.dlsite.com/bl/fsr/=/language/jp/age_category%5B0%5D/general/age_category%5B1%5D/r15/age_category%5B2%5D/adult/work_category%5B0%5D/doujin/order%5B0%5D/trend/work_type_category%5B0%5D/audio/work_type_category_name%5B0%5D/%E3%83%9C%E3%82%A4%E3%82%B9%E3%83%BBASMR/options_and_or/and/options%5B0%5D/JPN/options%5B1%5D/NM/options_name%5B0%5D/%E6%97%A5%E6%9C%AC%E8%AA%9E%E4%BD%9C%E5%93%81/options_name%5B1%5D/%E8%A8%80%E8%AA%9E%E4%B8%8D%E5%95%8F%E4%BD%9C%E5%93%81/per_page/100/from/fsr.more">
                    女性向け BL作品を検索する
                    <span v-if="count > 0" v-cloak class="search_more_number" data-vue-component="search_result" data-vue-async="true" data-url="https://www.dlsite.com/bl/sapi/=/language/jp/age_category%5B0%5D/general/age_category%5B1%5D/r15/age_category%5B2%5D/adult/work_category%5B0%5D/doujin/order%5B0%5D/trend/work_type_category%5B0%5D/audio/work_type_category_name%5B0%5D/%E3%83%9C%E3%82%A4%E3%82%B9%E3%83%BBASMR/options_and_or/and/options%5B0%5D/JPN/options%5B1%5D/NM/options_name%5B0%5D/%E6%97%A5%E6%9C%AC%E8%AA%9E%E4%BD%9C%E5%93%81/options_name%5B1%5D/%E8%A8%80%E8%AA%9E%E4%B8%8D%E5%95%8F%E4%BD%9C%E5%93%81/per_page/100/format/json/?cdn_cache=1">
                      (<span class="number">{{ count | number_format }}</span>件)
                    </span>
                  </a>
                </span>
              </li>
              <li>
                <span>
                  <a href="https://www.dlsite.com/home/fsr/=/language/jp/age_category%5B0%5D/general/work_category%5B0%5D/doujin/work_category%5B1%5D/books/work_category%5B2%5D/pc/order%5B0%5D/trend/work_type_category%5B0%5D/audio/work_type_category_name%5B0%5D/%E3%83%9C%E3%82%A4%E3%82%B9%E3%83%BBASMR/options_and_or/and/options%5B0%5D/JPN/options%5B1%5D/NM/options_name%5B0%5D/%E6%97%A5%E6%9C%AC%E8%AA%9E%E4%BD%9C%E5%93%81/options_name%5B1%5D/%E8%A8%80%E8%AA%9E%E4%B8%8D%E5%95%8F%E4%BD%9C%E5%93%81/per_page/100/from/fsr.more">
                    売り場を広げて検索する
                    <span v-if="count > 0" v-cloak class="search_more_number" data-vue-component="search_result" data-vue-async="true" data-url="https://www.dlsite.com/home/sapi/=/language/jp/age_category%5B0%5D/general/work_category%5B0%5D/doujin/work_category%5B1%5D/books/work_category%5B2%5D/pc/order%5B0%5D/trend/work_type_category%5B0%5D/audio/work_type_category_name%5B0%5D/%E3%83%9C%E3%82%A4%E3%82%B9%E3%83%BBASMR/options_and_or/and/options%5B0%5D/JPN/options%5B1%5D/NM/options_name%5B0%5D/%E6%97%A5%E6%9C%AC%E8%AA%9E%E4%BD%9C%E5%93%81/options_name%5B1%5D/%E8%A8%80%E8%AA%9E%E4%B8%8D%E5%95%8F%E4%BD%9C%E5%93%81/per_page/100/format/json/?cdn_cache=1">
                      (<span class="number">{{ count | number_format }}</span>件)
                    </span>
                  </a>
                </span>
              </li>
                                              <li>
          <span class="to_search_details">
            <a href="https://www.dlsite.com/home/fs/=/language/jp/age_category%5B0%5D/general/work_category%5B0%5D/doujin/order%5B0%5D/trend/work_type_category%5B0%5D/audio/work_type_category_name%5B0%5D/%E3%83%9C%E3%82%A4%E3%82%B9%E3%83%BBASMR/options_and_or/and/options%5B0%5D/JPN/options%5B1%5D/NM/per_page/100/show_type/3/lang_options%5B0%5D/%E6%97%A5%E6%9C%AC%E8%AA%9E/lang_options%5B1%5D/%E8%A8%80%E8%AA%9E%E4%B8%8D%E8%A6%81">検索条件を変更する</a>
          </span>
        </li>
      </ul>
    </div>
  </div>

<script>

jQuery(function($){
    //change order
    $('._change_submit').change(function(){
        var target = $(this);
        location.href = 'https://www.dlsite.com/home/works/type/=/language/jp/age_category%5B0%5D/general/work_category%5B0%5D/doujin/work_type_category%5B0%5D/audio/work_type_category_name%5B0%5D/%E3%83%9C%E3%82%A4%E3%82%B9%E3%83%BBASMR/options_and_or/and/options%5B0%5D/JPN/options%5B1%5D/NM/per_page/100/show_type/3/lang_options%5B0%5D/%E6%97%A5%E6%9C%AC%E8%AA%9E/lang_options%5B1%5D/%E8%A8%80%E8%AA%9E%E4%B8%8D%E8%A6%81/without_order/1/page/1' + '/order/'+ target.val() + '';
        target.attr({disabled: 'disabled'});
    });

    // 2017 gw campaign
    $('li.option_tab_item').filter(function(){ return $(this).children('a[data-tab=audio]').length }).addClass('selected');

    $('.display_type_select a').click(function(e){
      var show_type = $(this).data('value');
      $.cookie(
        'search_show_type',
        show_type,
        { path: '/', domain: '.dlsite.com' }
      );
    });
});

</script>


          <div id="_works_type_ranking" data-section_name="_works_type_ranking" class="contents_bottom">
  <div class="contents_bottom_inner border_t">
    <div class="contents_bottom_headline">同人 ボイス・ASMR 人気ランキング (7日間)</div>
        <div class="recommend_list type_bottom type_ranking  _top_total_ranking">
      <div id="_works_type_ranking_swiper_container" class="swiper-container swiper-container-horizontal" style="overflow: hidden;">
        <ul class="swiper-wrapper">
                  <li class="swiper-slide">
      <div class="recommend_work_item">

                  <div class="rank_number  type_rank01">
            <span>1</span>
          </div>
        
        <div>
          <a class="work_thumb" href="https://www.dlsite.com/home/work/=/product_id/RJ387519.html"  data-vue-component="thumb-img-popup">
            <img src="//img.dlsite.jp/resize/images2/work/doujin/RJ388000/RJ387519_img_main_240x240.jpg" alt="【催眠風音声・ヘッドマッサージ・耳かき】おしごとねいろ ～エステティシャン編～【CV.茅野愛衣】 [kotoneiro]" class="target_type"
              ref="popup_img"
               @mouseenter="showPopupImgInSwiper('//img.dlsite.jp/modpub/images2/work/doujin/RJ388000/RJ387519_img_main.jpg')" @mouseleave="hiddenPopupImgInSwiper()"             >
          </a>
        </div>

                  <div class="work_category type_SOU">
            <a href="https://www.dlsite.com/home/fsr/=/work_type/SOU">ボイス・ASMR</a>
          </div>
        
        <dl class="details">
          <dt class="work_name">
            <a href="https://www.dlsite.com/home/work/=/product_id/RJ387519.html" title="【催眠風音声・ヘッドマッサージ・耳かき】おしごとねいろ ～エステティシャン編～【CV.茅野愛衣】">【催眠風音声・ヘッドマッサージ・耳かき】おしごとねいろ ～エステティシャン編～【CV.茅野愛衣】</a>
          </dt>
          <dd class="maker_name">
            <a href="https://www.dlsite.com/home/circle/profile/=/maker_id/RG53557.html">kotoneiro</a>
                          <span class="separator">/</span>
              <span class="author"><a href="https://www.dlsite.com/home/fsr/=/keyword_creater/%22%E8%8C%85%E9%87%8E%E6%84%9B%E8%A1%A3%22" class="">茅野愛衣</a></span>
                      </dd>
          <dd class="work_price_wrap">
                        <span class="work_price">1,980<i>円</i></span>                      </dd>
                      <dd class="work_price_wrap">
              <span class="work_sales">販売数:&nbsp;<span class="_dl_count_RJ387519">3,936</span></span>
            </dd>
                    <dd class="work_label">
                                                            <span class="icon_option type_exclusive">専売</span>
                                                                </dd>
                    <template data-vue-component="product-item" data-product_id="RJ387519" data-layout="works_type"></template>
        </dl>
      </div>
    </li>
                  <li class="swiper-slide">
      <div class="recommend_work_item">

                  <div class="rank_number  type_rank02">
            <span>2</span>
          </div>
        
        <div>
          <a class="work_thumb" href="https://www.dlsite.com/home/work/=/product_id/RJ387206.html"  data-vue-component="thumb-img-popup">
            <img src="//img.dlsite.jp/resize/images2/work/doujin/RJ388000/RJ387206_img_main_240x240.jpg" alt="【アズールレーン】指揮官を癒やし隊！・綾波とゆっくり過ごす約1日間【ASMR】 [アトリエメール]" class="target_type"
              ref="popup_img"
               @mouseenter="showPopupImgInSwiper('//img.dlsite.jp/modpub/images2/work/doujin/RJ388000/RJ387206_img_main.jpg')" @mouseleave="hiddenPopupImgInSwiper()"             >
          </a>
        </div>

                  <div class="work_category type_SOU">
            <a href="https://www.dlsite.com/home/fsr/=/work_type/SOU">ボイス・ASMR</a>
          </div>
        
        <dl class="details">
          <dt class="work_name">
            <a href="https://www.dlsite.com/home/work/=/product_id/RJ387206.html" title="【アズールレーン】指揮官を癒やし隊！・綾波とゆっくり過ごす約1日間【ASMR】">【アズールレーン】指揮官を癒やし隊！・綾波とゆっくり過ごす約1日間【ASMR】</a>
          </dt>
          <dd class="maker_name">
            <a href="https://www.dlsite.com/home/circle/profile/=/maker_id/RG64308.html">アトリエメール</a>
                          <span class="separator">/</span>
              <span class="author"><a href="https://www.dlsite.com/home/fsr/=/keyword_creater/%22%E5%A4%A7%E5%9C%B0%E8%91%89%22" class="">大地葉</a></span>
                      </dd>
          <dd class="work_price_wrap">
                        <span class="work_price">1,320<i>円</i></span>            <span class="strike">1,650<i>円</i></span>                      </dd>
                      <dd class="work_price_wrap">
              <span class="work_sales">販売数:&nbsp;<span class="_dl_count_RJ387206">7,590</span></span>
            </dd>
                    <dd class="work_label">
                                                            <span class="icon_option type_exclusive">専売</span>
                                                                    <span class="icon_lead_01 type_sale">20%OFF</span>
                      </dd>
                    <template data-vue-component="product-item" data-product_id="RJ387206" data-layout="works_type"></template>
        </dl>
      </div>
    </li>
                  <li class="swiper-slide">
      <div class="recommend_work_item">

                  <div class="rank_number  type_rank03">
            <span>3</span>
          </div>
        
        <div>
          <a class="work_thumb" href="https://www.dlsite.com/home/work/=/product_id/RJ385913.html"  data-vue-component="thumb-img-popup">
            <img src="//img.dlsite.jp/resize/images2/work/doujin/RJ386000/RJ385913_img_main_240x240.jpg" alt="絶対に寝かしつける!からかい上手な後輩彼女のあまあま安眠耳かき [いちのや]" class="target_type"
              ref="popup_img"
               @mouseenter="showPopupImgInSwiper('//img.dlsite.jp/modpub/images2/work/doujin/RJ386000/RJ385913_img_main.jpg')" @mouseleave="hiddenPopupImgInSwiper()"             >
          </a>
        </div>

                  <div class="work_category type_SOU">
            <a href="https://www.dlsite.com/home/fsr/=/work_type/SOU">ボイス・ASMR</a>
          </div>
        
        <dl class="details">
          <dt class="work_name">
            <a href="https://www.dlsite.com/home/work/=/product_id/RJ385913.html" title="絶対に寝かしつける!からかい上手な後輩彼女のあまあま安眠耳かき">絶対に寝かしつける!からかい上手な後輩彼女のあまあま安眠耳かき</a>
          </dt>
          <dd class="maker_name">
            <a href="https://www.dlsite.com/home/circle/profile/=/maker_id/RG46007.html">いちのや</a>
                          <span class="separator">/</span>
              <span class="author"><a href="https://www.dlsite.com/home/fsr/=/keyword_creater/%22%E4%B8%80%E4%B9%8B%E7%80%AC%E3%82%8A%E3%81%A8%22" class="">一之瀬りと</a></span>
                      </dd>
          <dd class="work_price_wrap">
                        <span class="work_price">1,056<i>円</i></span>            <span class="strike">1,320<i>円</i></span>                      </dd>
                      <dd class="work_price_wrap">
              <span class="work_sales">販売数:&nbsp;<span class="_dl_count_RJ385913">4,055</span></span>
            </dd>
                    <dd class="work_label">
                                                            <span class="icon_option type_exclusive">専売</span>
                                                                    <span class="icon_lead_01 type_sale">20%OFF</span>
                      </dd>
                    <template data-vue-component="product-item" data-product_id="RJ385913" data-layout="works_type"></template>
        </dl>
      </div>
    </li>
                  <li class="swiper-slide">
      <div class="recommend_work_item">

                  <div class="rank_number ">
            <span>4</span>
          </div>
        
        <div>
          <a class="work_thumb" href="https://www.dlsite.com/home/work/=/product_id/RJ384573.html"  data-vue-component="thumb-img-popup">
            <img src="//img.dlsite.jp/resize/images2/work/doujin/RJ385000/RJ384573_img_main_240x240.jpg" alt="【耳かき・添い寝・看病・母性】癒やし満点彼女の甘やかし看病ASMR 〜天然おっとり保育士優衣にたっぷりお世話される2日間〜(CV.安野希世乃) [あんくりあさうんど]" class="target_type"
              ref="popup_img"
               @mouseenter="showPopupImgInSwiper('//img.dlsite.jp/modpub/images2/work/doujin/RJ385000/RJ384573_img_main.jpg')" @mouseleave="hiddenPopupImgInSwiper()"             >
          </a>
        </div>

                  <div class="work_category type_SOU">
            <a href="https://www.dlsite.com/home/fsr/=/work_type/SOU">ボイス・ASMR</a>
          </div>
        
        <dl class="details">
          <dt class="work_name">
            <a href="https://www.dlsite.com/home/work/=/product_id/RJ384573.html" title="【耳かき・添い寝・看病・母性】癒やし満点彼女の甘やかし看病ASMR 〜天然おっとり保育士優衣にたっぷりお世話される2日間〜(CV.安野希世乃)">【耳かき・添い寝・看病・母性】癒やし満点彼女の甘やかし看病ASMR 〜天然おっとり保育士優衣にたっぷりお世話される2日間〜(CV.安野希世乃)</a>
          </dt>
          <dd class="maker_name">
            <a href="https://www.dlsite.com/home/circle/profile/=/maker_id/RG54990.html">あんくりあさうんど</a>
                          <span class="separator">/</span>
              <span class="author"><a href="https://www.dlsite.com/home/fsr/=/keyword_creater/%22%E5%AE%89%E9%87%8E%E5%B8%8C%E4%B8%96%E4%B9%83%22" class="">安野希世乃</a></span>
                      </dd>
          <dd class="work_price_wrap">
                        <span class="work_price">1,584<i>円</i></span>            <span class="strike">1,980<i>円</i></span>                      </dd>
                      <dd class="work_price_wrap">
              <span class="work_sales">販売数:&nbsp;<span class="_dl_count_RJ384573">540</span></span>
            </dd>
                    <dd class="work_label">
                                                            <span class="icon_option type_exclusive">専売</span>
                                                                    <span class="icon_lead_01 type_sale">20%OFF</span>
                      </dd>
                    <template data-vue-component="product-item" data-product_id="RJ384573" data-layout="works_type"></template>
        </dl>
      </div>
    </li>
                  <li class="swiper-slide">
      <div class="recommend_work_item">

                  <div class="rank_number ">
            <span>5</span>
          </div>
        
        <div>
          <a class="work_thumb" href="https://www.dlsite.com/home/work/=/product_id/RJ385371.html"  data-vue-component="thumb-img-popup">
            <img src="//img.dlsite.jp/resize/images2/work/doujin/RJ386000/RJ385371_img_main_240x240.jpg" alt="【耳かき&amp;添い寝】圧倒的添い寝CD 〜素直でキュートな姪っ子にいっぱい癒されちゃう〜【CV:鬼頭明里】 [じゅじゅっとウェルダン]" class="target_type"
              ref="popup_img"
               @mouseenter="showPopupImgInSwiper('//img.dlsite.jp/modpub/images2/work/doujin/RJ386000/RJ385371_img_main.jpg')" @mouseleave="hiddenPopupImgInSwiper()"             >
          </a>
        </div>

                  <div class="work_category type_SOU">
            <a href="https://www.dlsite.com/home/fsr/=/work_type/SOU">ボイス・ASMR</a>
          </div>
        
        <dl class="details">
          <dt class="work_name">
            <a href="https://www.dlsite.com/home/work/=/product_id/RJ385371.html" title="【耳かき&amp;添い寝】圧倒的添い寝CD 〜素直でキュートな姪っ子にいっぱい癒されちゃう〜【CV:鬼頭明里】">【耳かき&amp;添い寝】圧倒的添い寝CD 〜素直でキュートな姪っ子にいっぱい癒されちゃう〜【CV:鬼頭明里】</a>
          </dt>
          <dd class="maker_name">
            <a href="https://www.dlsite.com/home/circle/profile/=/maker_id/RG50239.html">じゅじゅっとウェルダン</a>
                          <span class="separator">/</span>
              <span class="author"><a href="https://www.dlsite.com/home/fsr/=/keyword_creater/%22%E9%AC%BC%E9%A0%AD%E6%98%8E%E9%87%8C%22" class="">鬼頭明里</a></span>
                      </dd>
          <dd class="work_price_wrap">
                        <span class="work_price">990<i>円</i></span>            <span class="strike">1,320<i>円</i></span>                      </dd>
                      <dd class="work_price_wrap">
              <span class="work_sales">販売数:&nbsp;<span class="_dl_count_RJ385371">642</span></span>
            </dd>
                    <dd class="work_label">
                                                            <span class="icon_option type_exclusive">専売</span>
                                                                    <span class="icon_lead_01 type_sale">25%OFF</span>
                      </dd>
                    <template data-vue-component="product-item" data-product_id="RJ385371" data-layout="works_type"></template>
        </dl>
      </div>
    </li>
                  <li class="swiper-slide">
      <div class="recommend_work_item">

                  <div class="rank_number ">
            <span>6</span>
          </div>
        
        <div>
          <a class="work_thumb" href="https://www.dlsite.com/home/work/=/product_id/RJ386896.html"  data-vue-component="thumb-img-popup">
            <img src="//img.dlsite.jp/resize/images2/work/doujin/RJ387000/RJ386896_img_main_240x240.jpg" alt="【ASMR7時間30分/CV:こりす】ある森の小さな癒し処【耳かき・シャンプー・綿棒(ジェル・濡れ・乾き)】 [癒しの森の家]" class="target_type"
              ref="popup_img"
               @mouseenter="showPopupImgInSwiper('//img.dlsite.jp/modpub/images2/work/doujin/RJ387000/RJ386896_img_main.jpg')" @mouseleave="hiddenPopupImgInSwiper()"             >
          </a>
        </div>

                  <div class="work_category type_SOU">
            <a href="https://www.dlsite.com/home/fsr/=/work_type/SOU">ボイス・ASMR</a>
          </div>
        
        <dl class="details">
          <dt class="work_name">
            <a href="https://www.dlsite.com/home/work/=/product_id/RJ386896.html" title="【ASMR7時間30分/CV:こりす】ある森の小さな癒し処【耳かき・シャンプー・綿棒(ジェル・濡れ・乾き)】">【ASMR7時間30分/CV:こりす】ある森の小さな癒し処【耳かき・シャンプー・綿棒(ジェル・濡れ・乾き)】</a>
          </dt>
          <dd class="maker_name">
            <a href="https://www.dlsite.com/home/circle/profile/=/maker_id/RG47823.html">癒しの森の家</a>
                          <span class="separator">/</span>
              <span class="author"><a href="https://www.dlsite.com/home/fsr/=/keyword_creater/%22%E3%81%93%E3%82%8A%E3%81%99%22" class="">こりす</a></span>
                      </dd>
          <dd class="work_price_wrap">
                        <span class="work_price">1,155<i>円</i></span>            <span class="strike">1,650<i>円</i></span>                      </dd>
                      <dd class="work_price_wrap">
              <span class="work_sales">販売数:&nbsp;<span class="_dl_count_RJ386896">533</span></span>
            </dd>
                    <dd class="work_label">
                                                            <span class="icon_option type_exclusive">専売</span>
                                                                    <span class="icon_lead_01 type_sale">30%OFF</span>
                      </dd>
                    <template data-vue-component="product-item" data-product_id="RJ386896" data-layout="works_type"></template>
        </dl>
      </div>
    </li>
                  <li class="swiper-slide">
      <div class="recommend_work_item">

                  <div class="rank_number ">
            <span>7</span>
          </div>
        
        <div>
          <a class="work_thumb" href="https://www.dlsite.com/home/work/=/product_id/RJ387825.html"  data-vue-component="thumb-img-popup">
            <img src="//img.dlsite.jp/resize/images2/work/doujin/RJ388000/RJ387825_img_main_240x240.jpg" alt="【髪拭き・耳かき・添い寝】お姉さん系幼馴染と過ごすまったり雨宿りタイム【CV.大西亜玖璃】 [ほのぼの癒しのあまあま生活研究所]" class="target_type"
              ref="popup_img"
               @mouseenter="showPopupImgInSwiper('//img.dlsite.jp/modpub/images2/work/doujin/RJ388000/RJ387825_img_main.jpg')" @mouseleave="hiddenPopupImgInSwiper()"             >
          </a>
        </div>

                  <div class="work_category type_SOU">
            <a href="https://www.dlsite.com/home/fsr/=/work_type/SOU">ボイス・ASMR</a>
          </div>
        
        <dl class="details">
          <dt class="work_name">
            <a href="https://www.dlsite.com/home/work/=/product_id/RJ387825.html" title="【髪拭き・耳かき・添い寝】お姉さん系幼馴染と過ごすまったり雨宿りタイム【CV.大西亜玖璃】">【髪拭き・耳かき・添い寝】お姉さん系幼馴染と過ごすまったり雨宿りタイム【CV.大西亜玖璃】</a>
          </dt>
          <dd class="maker_name">
            <a href="https://www.dlsite.com/home/circle/profile/=/maker_id/RG64606.html">ほのぼの癒しのあまあま生活研究所</a>
                          <span class="separator">/</span>
              <span class="author"><a href="https://www.dlsite.com/home/fsr/=/keyword_creater/%22%E5%A4%A7%E8%A5%BF%E4%BA%9C%E7%8E%96%E7%92%83%22" class="">大西亜玖璃</a></span>
                      </dd>
          <dd class="work_price_wrap">
                        <span class="work_price">1,408<i>円</i></span>            <span class="strike">1,760<i>円</i></span>                      </dd>
                      <dd class="work_price_wrap">
              <span class="work_sales">販売数:&nbsp;<span class="_dl_count_RJ387825">376</span></span>
            </dd>
                    <dd class="work_label">
                                                            <span class="icon_option type_exclusive">専売</span>
                                                                    <span class="icon_lead_01 type_sale">20%OFF</span>
                      </dd>
                    <template data-vue-component="product-item" data-product_id="RJ387825" data-layout="works_type"></template>
        </dl>
      </div>
    </li>
                  <li class="swiper-slide">
      <div class="recommend_work_item">

                  <div class="rank_number ">
            <span>8</span>
          </div>
        
        <div>
          <a class="work_thumb" href="https://www.dlsite.com/home/work/=/product_id/RJ385780.html"  data-vue-component="thumb-img-popup">
            <img src="//img.dlsite.jp/resize/images2/work/doujin/RJ386000/RJ385780_img_main_240x240.jpg" alt="妹の友達のアイドルメスガキの大きなおともだち耳ほじりで俺は敗ける。 [でぶり]" class="target_type"
              ref="popup_img"
               @mouseenter="showPopupImgInSwiper('//img.dlsite.jp/modpub/images2/work/doujin/RJ386000/RJ385780_img_main.jpg')" @mouseleave="hiddenPopupImgInSwiper()"             >
          </a>
        </div>

                  <div class="work_category type_SOU">
            <a href="https://www.dlsite.com/home/fsr/=/work_type/SOU">ボイス・ASMR</a>
          </div>
        
        <dl class="details">
          <dt class="work_name">
            <a href="https://www.dlsite.com/home/work/=/product_id/RJ385780.html" title="妹の友達のアイドルメスガキの大きなおともだち耳ほじりで俺は敗ける。">妹の友達のアイドルメスガキの大きなおともだち耳ほじりで俺は敗ける。</a>
          </dt>
          <dd class="maker_name">
            <a href="https://www.dlsite.com/home/circle/profile/=/maker_id/RG49381.html">でぶり</a>
                          <span class="separator">/</span>
              <span class="author"><a href="https://www.dlsite.com/home/fsr/=/keyword_creater/%22%E7%8A%AC%E5%A1%9A%E3%81%84%E3%81%A1%E3%81%94%22" class="">犬塚いちご</a></span>
                      </dd>
          <dd class="work_price_wrap">
                        <span class="work_price">770<i>円</i></span>                      </dd>
                      <dd class="work_price_wrap">
              <span class="work_sales">販売数:&nbsp;<span class="_dl_count_RJ385780">887</span></span>
            </dd>
                    <dd class="work_label">
                                                            <span class="icon_option type_exclusive">専売</span>
                                                                </dd>
                    <template data-vue-component="product-item" data-product_id="RJ385780" data-layout="works_type"></template>
        </dl>
      </div>
    </li>
                  <li class="swiper-slide">
      <div class="recommend_work_item">

                  <div class="rank_number ">
            <span>9</span>
          </div>
        
        <div>
          <a class="work_thumb" href="https://www.dlsite.com/home/work/=/product_id/RJ387207.html"  data-vue-component="thumb-img-popup">
            <img src="//img.dlsite.jp/resize/images2/work/doujin/RJ388000/RJ387207_img_main_240x240.jpg" alt="【碧蓝航线】治愈指挥官小分队！・和绫波悠闲度过的1天【ASMR】 [アトリエメール]" class="target_type"
              ref="popup_img"
               @mouseenter="showPopupImgInSwiper('//img.dlsite.jp/modpub/images2/work/doujin/RJ388000/RJ387207_img_main.jpg')" @mouseleave="hiddenPopupImgInSwiper()"             >
          </a>
        </div>

                  <div class="work_category type_SOU">
            <a href="https://www.dlsite.com/home/fsr/=/work_type/SOU">ボイス・ASMR</a>
          </div>
        
        <dl class="details">
          <dt class="work_name">
            <a href="https://www.dlsite.com/home/work/=/product_id/RJ387207.html" title="【碧蓝航线】治愈指挥官小分队！・和绫波悠闲度过的1天【ASMR】">【碧蓝航线】治愈指挥官小分队！・和绫波悠闲度过的1天【ASMR】</a>
          </dt>
          <dd class="maker_name">
            <a href="https://www.dlsite.com/home/circle/profile/=/maker_id/RG64308.html">アトリエメール</a>
                          <span class="separator">/</span>
              <span class="author"><a href="https://www.dlsite.com/home/fsr/=/keyword_creater/%22%E5%A4%A7%E5%9C%B0%E8%91%89%22" class="">大地葉</a></span>
                      </dd>
          <dd class="work_price_wrap">
                        <span class="work_price">1,320<i>円</i></span>            <span class="strike">1,650<i>円</i></span>                      </dd>
                      <dd class="work_price_wrap">
              <span class="work_sales">販売数:&nbsp;<span class="_dl_count_RJ387207">1,278</span></span>
            </dd>
                    <dd class="work_label">
                                                            <span class="icon_option type_exclusive">専売</span>
                                                                    <span class="icon_lead_01 type_sale">20%OFF</span>
                      </dd>
                    <template data-vue-component="product-item" data-product_id="RJ387207" data-layout="works_type"></template>
        </dl>
      </div>
    </li>
                  <li class="swiper-slide">
      <div class="recommend_work_item">

                  <div class="rank_number ">
            <span>10</span>
          </div>
        
        <div>
          <a class="work_thumb" href="https://www.dlsite.com/home/work/=/product_id/RJ381666.html"  data-vue-component="thumb-img-popup">
            <img src="//img.dlsite.jp/resize/images2/work/doujin/RJ382000/RJ381666_img_main_240x240.jpg" alt="【初めて耳かき・耳裏、フェイスマッサージ】桜木学園癒やし部～2年C組・桑崎もえ 風紀委員長の不純異性交遊は絶対許さないASMR～【プレミアムサウンド2022】 [RaRo]" class="target_type"
              ref="popup_img"
               @mouseenter="showPopupImgInSwiper('//img.dlsite.jp/modpub/images2/work/doujin/RJ382000/RJ381666_img_main.jpg')" @mouseleave="hiddenPopupImgInSwiper()"             >
          </a>
        </div>

                  <div class="work_category type_SOU">
            <a href="https://www.dlsite.com/home/fsr/=/work_type/SOU">ボイス・ASMR</a>
          </div>
        
        <dl class="details">
          <dt class="work_name">
            <a href="https://www.dlsite.com/home/work/=/product_id/RJ381666.html" title="【初めて耳かき・耳裏、フェイスマッサージ】桜木学園癒やし部～2年C組・桑崎もえ 風紀委員長の不純異性交遊は絶対許さないASMR～【プレミアムサウンド2022】">【初めて耳かき・耳裏、フェイスマッサージ】桜木学園癒やし部～2年C組・桑崎もえ 風紀委員長の不純異性交遊は絶対許さないASMR～【プレミアムサウンド2022】</a>
          </dt>
          <dd class="maker_name">
            <a href="https://www.dlsite.com/home/circle/profile/=/maker_id/RG49556.html">RaRo</a>
                          <span class="separator">/</span>
              <span class="author"><a href="https://www.dlsite.com/home/fsr/=/keyword_creater/%22%E7%9B%B8%E5%9D%82%E5%84%AA%E6%AD%8C%22" class="">相坂優歌</a></span>
                      </dd>
          <dd class="work_price_wrap">
                        <span class="work_price">1,001<i>円</i></span>            <span class="strike">1,430<i>円</i></span>                      </dd>
                      <dd class="work_price_wrap">
              <span class="work_sales">販売数:&nbsp;<span class="_dl_count_RJ381666">1,745</span></span>
            </dd>
                    <dd class="work_label">
                                                            <span class="icon_option type_exclusive">専売</span>
                                                                    <span class="icon_lead_01 type_sale">30%OFF</span>
                      </dd>
                    <template data-vue-component="product-item" data-product_id="RJ381666" data-layout="works_type"></template>
        </dl>
      </div>
    </li>
                  <li class="swiper-slide">
      <div class="recommend_work_item">

                  <div class="rank_number ">
            <span>11</span>
          </div>
        
        <div>
          <a class="work_thumb" href="https://www.dlsite.com/home/work/=/product_id/RJ385489.html"  data-vue-component="thumb-img-popup">
            <img src="//img.dlsite.jp/resize/images2/work/doujin/RJ386000/RJ385489_img_main_240x240.jpg" alt="【百合観察】ユメリリ 〜 幼なじみカップル観察日記 〜【CV:山根綺 &amp; 根本京里】 [SukeraSono]" class="target_type"
              ref="popup_img"
               @mouseenter="showPopupImgInSwiper('//img.dlsite.jp/modpub/images2/work/doujin/RJ386000/RJ385489_img_main.jpg')" @mouseleave="hiddenPopupImgInSwiper()"             >
          </a>
        </div>

                  <div class="work_category type_SOU">
            <a href="https://www.dlsite.com/home/fsr/=/work_type/SOU">ボイス・ASMR</a>
          </div>
        
        <dl class="details">
          <dt class="work_name">
            <a href="https://www.dlsite.com/home/work/=/product_id/RJ385489.html" title="【百合観察】ユメリリ 〜 幼なじみカップル観察日記 〜【CV:山根綺 &amp; 根本京里】">【百合観察】ユメリリ 〜 幼なじみカップル観察日記 〜【CV:山根綺 &amp; 根本京里】</a>
          </dt>
          <dd class="maker_name">
            <a href="https://www.dlsite.com/home/circle/profile/=/maker_id/RG56373.html">SukeraSono</a>
                          <span class="separator">/</span>
              <span class="author"><a href="https://www.dlsite.com/home/fsr/=/keyword_creater/%22%E5%B1%B1%E6%A0%B9%E7%B6%BA%22" class="">山根綺</a>&nbsp;<a href="https://www.dlsite.com/home/fsr/=/keyword_creater/%22%E6%A0%B9%E6%9C%AC%E4%BA%AC%E9%87%8C%22" class="">根本京里</a></span>
                      </dd>
          <dd class="work_price_wrap">
                        <span class="work_price">1,584<i>円</i></span>            <span class="strike">1,980<i>円</i></span>                      </dd>
                      <dd class="work_price_wrap">
              <span class="work_sales">販売数:&nbsp;<span class="_dl_count_RJ385489">215</span></span>
            </dd>
                    <dd class="work_label">
                                                            <span class="icon_option type_exclusive">専売</span>
                                                                    <span class="icon_lead_01 type_sale">20%OFF</span>
                      </dd>
                    <template data-vue-component="product-item" data-product_id="RJ385489" data-layout="works_type"></template>
        </dl>
      </div>
    </li>
                  <li class="swiper-slide">
      <div class="recommend_work_item">

                  <div class="rank_number ">
            <span>12</span>
          </div>
        
        <div>
          <a class="work_thumb" href="https://www.dlsite.com/home/work/=/product_id/RJ329940.html"  data-vue-component="thumb-img-popup">
            <img src="//img.dlsite.jp/resize/images2/work/doujin/RJ330000/RJ329940_img_main_240x240.jpg" alt="【寝落ちASMR13時間】99.99%ぐ～っすり寝かせちゃう癒しの安眠屋さん。(極上耳かき・マッサージ・赤ちゃん綿棒・囁き) [周防パトラ]" class="target_type"
              ref="popup_img"
               @mouseenter="showPopupImgInSwiper('//img.dlsite.jp/modpub/images2/work/doujin/RJ330000/RJ329940_img_main.jpg')" @mouseleave="hiddenPopupImgInSwiper()"             >
          </a>
        </div>

                  <div class="work_category type_SOU">
            <a href="https://www.dlsite.com/home/fsr/=/work_type/SOU">ボイス・ASMR</a>
          </div>
        
        <dl class="details">
          <dt class="work_name">
            <a href="https://www.dlsite.com/home/work/=/product_id/RJ329940.html" title="【寝落ちASMR13時間】99.99%ぐ～っすり寝かせちゃう癒しの安眠屋さん。(極上耳かき・マッサージ・赤ちゃん綿棒・囁き)">【寝落ちASMR13時間】99.99%ぐ～っすり寝かせちゃう癒しの安眠屋さん。(極上耳かき・マッサージ・赤ちゃん綿棒・囁き)</a>
          </dt>
          <dd class="maker_name">
            <a href="https://www.dlsite.com/home/circle/profile/=/maker_id/RG51620.html">周防パトラ</a>
                          <span class="separator">/</span>
              <span class="author"><a href="https://www.dlsite.com/home/fsr/=/keyword_creater/%22%E5%91%A8%E9%98%B2%E3%83%91%E3%83%88%E3%83%A9%22" class="">周防パトラ</a></span>
                      </dd>
          <dd class="work_price_wrap">
                        <span class="work_price">1,430<i>円</i></span>                      </dd>
                      <dd class="work_price_wrap">
              <span class="work_sales">販売数:&nbsp;<span class="_dl_count_RJ329940">45,954</span></span>
            </dd>
                    <dd class="work_label">
                                                            <span class="icon_option type_exclusive">専売</span>
                                                                </dd>
                    <template data-vue-component="product-item" data-product_id="RJ329940" data-layout="works_type"></template>
        </dl>
      </div>
    </li>
                  <li class="swiper-slide">
      <div class="recommend_work_item">

                  <div class="rank_number ">
            <span>13</span>
          </div>
        
        <div>
          <a class="work_thumb" href="https://www.dlsite.com/home/work/=/product_id/RJ377034.html"  data-vue-component="thumb-img-popup">
            <img src="//img.dlsite.jp/resize/images2/work/doujin/RJ378000/RJ377034_img_main_240x240.jpg" alt="【耳かき・耳マッサージ・睡眠導入】おしごとねいろ ～睡眠カウンセラー編～【CV.富田美憂】 [kotoneiro]" class="target_type"
              ref="popup_img"
               @mouseenter="showPopupImgInSwiper('//img.dlsite.jp/modpub/images2/work/doujin/RJ378000/RJ377034_img_main.jpg')" @mouseleave="hiddenPopupImgInSwiper()"             >
          </a>
        </div>

                  <div class="work_category type_SOU">
            <a href="https://www.dlsite.com/home/fsr/=/work_type/SOU">ボイス・ASMR</a>
          </div>
        
        <dl class="details">
          <dt class="work_name">
            <a href="https://www.dlsite.com/home/work/=/product_id/RJ377034.html" title="【耳かき・耳マッサージ・睡眠導入】おしごとねいろ ～睡眠カウンセラー編～【CV.富田美憂】">【耳かき・耳マッサージ・睡眠導入】おしごとねいろ ～睡眠カウンセラー編～【CV.富田美憂】</a>
          </dt>
          <dd class="maker_name">
            <a href="https://www.dlsite.com/home/circle/profile/=/maker_id/RG53557.html">kotoneiro</a>
                          <span class="separator">/</span>
              <span class="author"><a href="https://www.dlsite.com/home/fsr/=/keyword_creater/%22%E5%AF%8C%E7%94%B0%E7%BE%8E%E6%86%82%22" class="">富田美憂</a></span>
                      </dd>
          <dd class="work_price_wrap">
                        <span class="work_price">1,980<i>円</i></span>                      </dd>
                      <dd class="work_price_wrap">
              <span class="work_sales">販売数:&nbsp;<span class="_dl_count_RJ377034">2,335</span></span>
            </dd>
                    <dd class="work_label">
                                                            <span class="icon_option type_exclusive">専売</span>
                                                                </dd>
                    <template data-vue-component="product-item" data-product_id="RJ377034" data-layout="works_type"></template>
        </dl>
      </div>
    </li>
                  <li class="swiper-slide">
      <div class="recommend_work_item">

                  <div class="rank_number ">
            <span>14</span>
          </div>
        
        <div>
          <a class="work_thumb" href="https://www.dlsite.com/home/work/=/product_id/RJ370398.html"  data-vue-component="thumb-img-popup">
            <img src="//img.dlsite.jp/resize/images2/work/doujin/RJ371000/RJ370398_img_main_240x240.jpg" alt="34時間で君の耳を幸せに出来るか?コンプリートパック [RaRo]" class="target_type"
              ref="popup_img"
               @mouseenter="showPopupImgInSwiper('//img.dlsite.jp/modpub/images2/work/doujin/RJ371000/RJ370398_img_main.jpg')" @mouseleave="hiddenPopupImgInSwiper()"             >
          </a>
        </div>

                  <div class="work_category type_SOU">
            <a href="https://www.dlsite.com/home/fsr/=/work_type/SOU">ボイス・ASMR</a>
          </div>
        
        <dl class="details">
          <dt class="work_name">
            <a href="https://www.dlsite.com/home/work/=/product_id/RJ370398.html" title="34時間で君の耳を幸せに出来るか?コンプリートパック">34時間で君の耳を幸せに出来るか?コンプリートパック</a>
          </dt>
          <dd class="maker_name">
            <a href="https://www.dlsite.com/home/circle/profile/=/maker_id/RG49556.html">RaRo</a>
                          <span class="separator">/</span>
              <span class="author"><a href="https://www.dlsite.com/home/fsr/=/keyword_creater/%22%E4%B8%89%E4%B8%8A%E6%9E%9D%E7%B9%94%22" class="">三上枝織</a>&nbsp;<a href="https://www.dlsite.com/home/fsr/=/keyword_creater/%22%E4%B8%89%E6%A3%AE%E3%81%99%E3%81%9A%E3%81%93%22" class="">三森すずこ</a>&nbsp;<a href="https://www.dlsite.com/home/fsr/=/keyword_creater/%22%E5%92%8C%E6%B0%A3%E3%81%82%E3%81%9A%E6%9C%AA%22" class="">和氣あず未</a>&nbsp;<a href="https://www.dlsite.com/home/fsr/=/keyword_creater/%22%E4%BC%8A%E8%97%A4%E3%81%8B%E3%81%AA%E6%81%B5%22" class="">伊藤かな恵</a>&nbsp;<a href="https://www.dlsite.com/home/fsr/=/keyword_creater/%22%E8%8A%B9%E6%BE%A4%E5%84%AA%22" class="">芹澤優</a>&nbsp;<a href="https://www.dlsite.com/home/fsr/=/keyword_creater/%22%E5%8F%A4%E8%B3%80%E8%91%B5%22" class="">古賀葵</a>&nbsp;<a href="https://www.dlsite.com/home/fsr/=/keyword_creater/%22%E9%80%9F%E6%B0%B4%E5%A5%A8%22" class="">速水奨</a></span>
                      </dd>
          <dd class="work_price_wrap">
                        <span class="work_price">3,520<i>円</i></span>            <span class="strike">8,800<i>円</i></span>                      </dd>
                      <dd class="work_price_wrap">
              <span class="work_sales">販売数:&nbsp;<span class="_dl_count_RJ370398">621</span></span>
            </dd>
                    <dd class="work_label">
                                                            <span class="icon_option type_exclusive">専売</span>
                                                                    <span class="icon_lead_01 type_sale">60%OFF</span>
                      </dd>
                    <template data-vue-component="product-item" data-product_id="RJ370398" data-layout="works_type"></template>
        </dl>
      </div>
    </li>
                  <li class="swiper-slide">
      <div class="recommend_work_item">

                  <div class="rank_number ">
            <span>15</span>
          </div>
        
        <div>
          <a class="work_thumb" href="https://www.dlsite.com/home/work/=/product_id/RJ385241.html"  data-vue-component="thumb-img-popup">
            <img src="//img.dlsite.jp/resize/images2/work/doujin/RJ386000/RJ385241_img_main_240x240.jpg" alt="【クール禁止耳かき・ブラッシング】桜木学園癒やし部～1年C組・五十嵐千歌 生意気クール娘の献身癒やし耳かき編～【CV.戸田めぐみ】 [RaRo]" class="target_type"
              ref="popup_img"
               @mouseenter="showPopupImgInSwiper('//img.dlsite.jp/modpub/images2/work/doujin/RJ386000/RJ385241_img_main.jpg')" @mouseleave="hiddenPopupImgInSwiper()"             >
          </a>
        </div>

                  <div class="work_category type_SOU">
            <a href="https://www.dlsite.com/home/fsr/=/work_type/SOU">ボイス・ASMR</a>
          </div>
        
        <dl class="details">
          <dt class="work_name">
            <a href="https://www.dlsite.com/home/work/=/product_id/RJ385241.html" title="【クール禁止耳かき・ブラッシング】桜木学園癒やし部～1年C組・五十嵐千歌 生意気クール娘の献身癒やし耳かき編～【CV.戸田めぐみ】">【クール禁止耳かき・ブラッシング】桜木学園癒やし部～1年C組・五十嵐千歌 生意気クール娘の献身癒やし耳かき編～【CV.戸田めぐみ】</a>
          </dt>
          <dd class="maker_name">
            <a href="https://www.dlsite.com/home/circle/profile/=/maker_id/RG49556.html">RaRo</a>
                          <span class="separator">/</span>
              <span class="author"><a href="https://www.dlsite.com/home/fsr/=/keyword_creater/%22%E6%88%B8%E7%94%B0%E3%82%81%E3%81%90%E3%81%BF%22" class="">戸田めぐみ</a></span>
                      </dd>
          <dd class="work_price_wrap">
                        <span class="work_price">1,386<i>円</i></span>            <span class="strike">1,980<i>円</i></span>                      </dd>
                      <dd class="work_price_wrap">
              <span class="work_sales">販売数:&nbsp;<span class="_dl_count_RJ385241">197</span></span>
            </dd>
                    <dd class="work_label">
                                                            <span class="icon_option type_exclusive">専売</span>
                                                                    <span class="icon_lead_01 type_sale">30%OFF</span>
                      </dd>
                    <template data-vue-component="product-item" data-product_id="RJ385241" data-layout="works_type"></template>
        </dl>
      </div>
    </li>
                  <li class="swiper-slide">
      <div class="recommend_work_item">

                  <div class="rank_number ">
            <span>16</span>
          </div>
        
        <div>
          <a class="work_thumb" href="https://www.dlsite.com/home/work/=/product_id/RJ363741.html"  data-vue-component="thumb-img-popup">
            <img src="//img.dlsite.jp/resize/images2/work/doujin/RJ364000/RJ363741_img_main_240x240.jpg" alt="【耳かきNo.1】唯一無二の最強耳かき音9時間 [ちろ猫ハウス(耳かき専門店)]" class="target_type"
              ref="popup_img"
               @mouseenter="showPopupImgInSwiper('//img.dlsite.jp/modpub/images2/work/doujin/RJ364000/RJ363741_img_main.jpg')" @mouseleave="hiddenPopupImgInSwiper()"             >
          </a>
        </div>

                  <div class="work_category type_SOU">
            <a href="https://www.dlsite.com/home/fsr/=/work_type/SOU">ボイス・ASMR</a>
          </div>
        
        <dl class="details">
          <dt class="work_name">
            <a href="https://www.dlsite.com/home/work/=/product_id/RJ363741.html" title="【耳かきNo.1】唯一無二の最強耳かき音9時間">【耳かきNo.1】唯一無二の最強耳かき音9時間</a>
          </dt>
          <dd class="maker_name">
            <a href="https://www.dlsite.com/home/circle/profile/=/maker_id/RG37770.html">ちろ猫ハウス(耳かき専門店)</a>
                          <span class="separator">/</span>
              <span class="author"><a href="https://www.dlsite.com/home/fsr/=/keyword_creater/%22%E6%81%8B%E7%8C%AB%E3%81%A1%E3%82%8D%E3%82%8B%22" class="">恋猫ちろる</a></span>
                      </dd>
          <dd class="work_price_wrap">
                        <span class="work_price">1,980<i>円</i></span>                      </dd>
                      <dd class="work_price_wrap">
              <span class="work_sales">販売数:&nbsp;<span class="_dl_count_RJ363741">9,483</span></span>
            </dd>
                    <dd class="work_label">
                                                            <span class="icon_option type_exclusive">専売</span>
                                                                </dd>
                    <template data-vue-component="product-item" data-product_id="RJ363741" data-layout="works_type"></template>
        </dl>
      </div>
    </li>
                  <li class="swiper-slide">
      <div class="recommend_work_item">

                  <div class="rank_number ">
            <span>17</span>
          </div>
        
        <div>
          <a class="work_thumb" href="https://www.dlsite.com/home/work/=/product_id/RJ379246.html"  data-vue-component="thumb-img-popup">
            <img src="//img.dlsite.jp/resize/images2/work/doujin/RJ380000/RJ379246_img_main_240x240.jpg" alt="『アサルトリリィ Last Bullet』ASMR 梨璃さんと一緒～お姉様にはナイショだよっ!～ [電撃G&#039;s magazine]" class="target_type"
              ref="popup_img"
               @mouseenter="showPopupImgInSwiper('//img.dlsite.jp/modpub/images2/work/doujin/RJ380000/RJ379246_img_main.jpg')" @mouseleave="hiddenPopupImgInSwiper()"             >
          </a>
        </div>

                  <div class="work_category type_SOU">
            <a href="https://www.dlsite.com/home/fsr/=/work_type/SOU">ボイス・ASMR</a>
          </div>
        
        <dl class="details">
          <dt class="work_name">
            <a href="https://www.dlsite.com/home/work/=/product_id/RJ379246.html" title="『アサルトリリィ Last Bullet』ASMR 梨璃さんと一緒～お姉様にはナイショだよっ!～">『アサルトリリィ Last Bullet』ASMR 梨璃さんと一緒～お姉様にはナイショだよっ!～</a>
          </dt>
          <dd class="maker_name">
            <a href="https://www.dlsite.com/home/circle/profile/=/maker_id/RG48399.html">電撃G&#039;s magazine</a>
                          <span class="separator">/</span>
              <span class="author"><a href="https://www.dlsite.com/home/fsr/=/keyword_creater/%22%E8%B5%A4%E5%B0%BE%E3%81%B2%E3%81%8B%E3%82%8B%22" class="">赤尾ひかる</a></span>
                      </dd>
          <dd class="work_price_wrap">
                        <span class="work_price">1,584<i>円</i></span>            <span class="strike">1,980<i>円</i></span>                      </dd>
                      <dd class="work_price_wrap">
              <span class="work_sales">販売数:&nbsp;<span class="_dl_count_RJ379246">1,115</span></span>
            </dd>
                    <dd class="work_label">
                                                                                              <span class="icon_lead_01 type_sale">20%OFF</span>
                      </dd>
                    <template data-vue-component="product-item" data-product_id="RJ379246" data-layout="works_type"></template>
        </dl>
      </div>
    </li>
                  <li class="swiper-slide">
      <div class="recommend_work_item">

                  <div class="rank_number ">
            <span>18</span>
          </div>
        
        <div>
          <a class="work_thumb" href="https://www.dlsite.com/home/work/=/product_id/RJ381822.html"  data-vue-component="thumb-img-popup">
            <img src="//img.dlsite.jp/resize/images2/work/doujin/RJ382000/RJ381822_img_main_240x240.jpg" alt="おさんぽASMR ～逍遥馬道を馬にゆられてカポカポお散歩のその後は、にんじんあげて、マッサージ受けて、耳かきをしてのんびり添い寝～ [RaRo]" class="target_type"
              ref="popup_img"
               @mouseenter="showPopupImgInSwiper('//img.dlsite.jp/modpub/images2/work/doujin/RJ382000/RJ381822_img_main.jpg')" @mouseleave="hiddenPopupImgInSwiper()"             >
          </a>
        </div>

                  <div class="work_category type_SOU">
            <a href="https://www.dlsite.com/home/fsr/=/work_type/SOU">ボイス・ASMR</a>
          </div>
        
        <dl class="details">
          <dt class="work_name">
            <a href="https://www.dlsite.com/home/work/=/product_id/RJ381822.html" title="おさんぽASMR ～逍遥馬道を馬にゆられてカポカポお散歩のその後は、にんじんあげて、マッサージ受けて、耳かきをしてのんびり添い寝～">おさんぽASMR ～逍遥馬道を馬にゆられてカポカポお散歩のその後は、にんじんあげて、マッサージ受けて、耳かきをしてのんびり添い寝～</a>
          </dt>
          <dd class="maker_name">
            <a href="https://www.dlsite.com/home/circle/profile/=/maker_id/RG49556.html">RaRo</a>
                          <span class="separator">/</span>
              <span class="author"><a href="https://www.dlsite.com/home/fsr/=/keyword_creater/%22%E7%9C%9E%E7%94%B0%E6%9C%B1%E9%9F%B3%22" class="">眞田朱音</a></span>
                      </dd>
          <dd class="work_price_wrap">
                        <span class="work_price">1,386<i>円</i></span>            <span class="strike">1,980<i>円</i></span>                      </dd>
                      <dd class="work_price_wrap">
              <span class="work_sales">販売数:&nbsp;<span class="_dl_count_RJ381822">200</span></span>
            </dd>
                    <dd class="work_label">
                                                            <span class="icon_option type_exclusive">専売</span>
                                                                    <span class="icon_lead_01 type_sale">30%OFF</span>
                      </dd>
                    <template data-vue-component="product-item" data-product_id="RJ381822" data-layout="works_type"></template>
        </dl>
      </div>
    </li>
                  <li class="swiper-slide">
      <div class="recommend_work_item">

                  <div class="rank_number ">
            <span>19</span>
          </div>
        
        <div>
          <a class="work_thumb" href="https://www.dlsite.com/home/work/=/product_id/RJ387210.html"  data-vue-component="thumb-img-popup">
            <img src="//img.dlsite.jp/resize/images2/work/doujin/RJ388000/RJ387210_img_main_240x240.jpg" alt="[Azur Lane ASMR] Commander Pampering Team! A Relaxing Day with Ayanami [アトリエメール]" class="target_type"
              ref="popup_img"
               @mouseenter="showPopupImgInSwiper('//img.dlsite.jp/modpub/images2/work/doujin/RJ388000/RJ387210_img_main.jpg')" @mouseleave="hiddenPopupImgInSwiper()"             >
          </a>
        </div>

                  <div class="work_category type_SOU">
            <a href="https://www.dlsite.com/home/fsr/=/work_type/SOU">ボイス・ASMR</a>
          </div>
        
        <dl class="details">
          <dt class="work_name">
            <a href="https://www.dlsite.com/home/work/=/product_id/RJ387210.html" title="[Azur Lane ASMR] Commander Pampering Team! A Relaxing Day with Ayanami">[Azur Lane ASMR] Commander Pampering Team! A Relaxing Day with Ayanami</a>
          </dt>
          <dd class="maker_name">
            <a href="https://www.dlsite.com/home/circle/profile/=/maker_id/RG64308.html">アトリエメール</a>
                          <span class="separator">/</span>
              <span class="author"><a href="https://www.dlsite.com/home/fsr/=/keyword_creater/%22%E5%A4%A7%E5%9C%B0%E8%91%89%22" class="">大地葉</a></span>
                      </dd>
          <dd class="work_price_wrap">
                        <span class="work_price">1,320<i>円</i></span>            <span class="strike">1,650<i>円</i></span>                      </dd>
                      <dd class="work_price_wrap">
              <span class="work_sales">販売数:&nbsp;<span class="_dl_count_RJ387210">375</span></span>
            </dd>
                    <dd class="work_label">
                                                            <span class="icon_option type_exclusive">専売</span>
                                                                    <span class="icon_lead_01 type_sale">20%OFF</span>
                      </dd>
                    <template data-vue-component="product-item" data-product_id="RJ387210" data-layout="works_type"></template>
        </dl>
      </div>
    </li>
                  <li class="swiper-slide">
      <div class="recommend_work_item">

                  <div class="rank_number ">
            <span>20</span>
          </div>
        
        <div>
          <a class="work_thumb" href="https://www.dlsite.com/home/work/=/product_id/RJ299717.html"  data-vue-component="thumb-img-popup">
            <img src="//img.dlsite.jp/resize/images2/work/doujin/RJ300000/RJ299717_img_main_240x240.jpg" alt="【寝落ちASMR】悪魔娘が最高に癒すのでものすごく眠れる(耳かき・囁き・マッサージ・泡オイル) [周防パトラ]" class="target_type"
              ref="popup_img"
               @mouseenter="showPopupImgInSwiper('//img.dlsite.jp/modpub/images2/work/doujin/RJ300000/RJ299717_img_main.jpg')" @mouseleave="hiddenPopupImgInSwiper()"             >
          </a>
        </div>

                  <div class="work_category type_SOU">
            <a href="https://www.dlsite.com/home/fsr/=/work_type/SOU">ボイス・ASMR</a>
          </div>
        
        <dl class="details">
          <dt class="work_name">
            <a href="https://www.dlsite.com/home/work/=/product_id/RJ299717.html" title="【寝落ちASMR】悪魔娘が最高に癒すのでものすごく眠れる(耳かき・囁き・マッサージ・泡オイル)">【寝落ちASMR】悪魔娘が最高に癒すのでものすごく眠れる(耳かき・囁き・マッサージ・泡オイル)</a>
          </dt>
          <dd class="maker_name">
            <a href="https://www.dlsite.com/home/circle/profile/=/maker_id/RG51620.html">周防パトラ</a>
                          <span class="separator">/</span>
              <span class="author"><a href="https://www.dlsite.com/home/fsr/=/keyword_creater/%22%E5%91%A8%E9%98%B2%E3%83%91%E3%83%88%E3%83%A9%22" class="">周防パトラ</a></span>
                      </dd>
          <dd class="work_price_wrap">
                        <span class="work_price">1,430<i>円</i></span>                      </dd>
                      <dd class="work_price_wrap">
              <span class="work_sales">販売数:&nbsp;<span class="_dl_count_RJ299717">49,583</span></span>
            </dd>
                    <dd class="work_label">
                                                                                          </dd>
                    <template data-vue-component="product-item" data-product_id="RJ299717" data-layout="works_type"></template>
        </dl>
      </div>
    </li>
  </ul>
      </div>
      <div class="swiper-button-next"></div>
      <div class="swiper-button-prev"></div>
    </div>
  </div>
</div>


  <script>
    $(function () {
      var element = document.getElementById('_works_type_ranking_swiper_container');
      new Swiper(element, {
        lazy: {
          loadPrevNext: true,
        },
        watchSlidesVisibility: true,
        slidesPerView: 3,
        slidesPerGroup: 3,
        spaceBetween: 10,
        breakpoints: {
          1250: {
            slidesPerView: 4,
            slidesPerGroup: 4,
          },
          1450: {
            slidesPerView: 5,
            slidesPerGroup: 5,
          },
          1650: {
            slidesPerView: 6,
            slidesPerGroup: 6,
          },
          1850: {
            slidesPerView: 7,
            slidesPerGroup: 7,
          },
          max: {
            slidesPerView: 8,
            slidesPerGroup: 8,
          }
        },
        navigation: {
          nextEl: element.parentNode.querySelector('.swiper-button-next'),
          prevEl: element.parentNode.querySelector('.swiper-button-prev'),
        },
        nextButton: element.parentNode.querySelector('.swiper-button-next'),
        prevButton: element.parentNode.querySelector('.swiper-button-prev'),
      });
    })
  </script>

          
          <div id="_works_type_recommend" data-section_name="_works_type_recommend" class="contents_bottom">
  <div class="contents_bottom_inner">
    <div class="contents_bottom_headline">おすすめ作品</div>
        <div class="recommend_list type_bottom type_ranking  _top_total_ranking">
      <div data-vue-component="WorksTypeRecommend" data-vue-async="true"
           data-parent_container_id="_works_type_recommend"
           data-endpoint="https://www.dlsite.com/home/works/type/recommend/ajax/=/work_type/audio/swiper_container_id/_works_type_recommend_swiper_container"
           id="_works_type_recommend_swiper_container" class="swiper-container swiper-container-horizontal"
           style="overflow: hidden;">

        <lazy-component @show="loadRecommend()"></lazy-component>

        <div class="loading_box" style="height: 70px;">
          <div class="loading"></div>
        </div>

      </div>
      <div class="swiper-button-next"></div>
      <div class="swiper-button-prev"></div>
    </div>
  </div>
</div>
            <div class="contents_bottom">
    <div class="contents_bottom_inner">
      <p class="contents_bottom_headline">他の作品形式で探す</p>
      <div class="filter_genre">
        <ul>
                      <li>
              <dl>
                <dt>ゲーム</dt>
                <div class="col_2">
                                                        <dd><a href="https://www.dlsite.com/home/works/type/=/work_type_category/game">すべて</a></dd>
                                                                                                <dd><a href="https://www.dlsite.com/home/works/type/=/work_type/ACN">アクション</a></dd>
                                          <dd><a href="https://www.dlsite.com/home/works/type/=/work_type/QIZ">クイズ</a></dd>
                                          <dd><a href="https://www.dlsite.com/home/works/type/=/work_type/ADV">アドベンチャー</a></dd>
                                          <dd><a href="https://www.dlsite.com/home/works/type/=/work_type/RPG">ロールプレイング</a></dd>
                                          <dd><a href="https://www.dlsite.com/home/works/type/=/work_type/TBL">テーブル</a></dd>
                                          <dd><a href="https://www.dlsite.com/home/works/type/=/work_type/DNV">デジタルノベル</a></dd>
                                          <dd><a href="https://www.dlsite.com/home/works/type/=/work_type/SLN">シミュレーション</a></dd>
                                          <dd><a href="https://www.dlsite.com/home/works/type/=/work_type/TYP">タイピング</a></dd>
                                          <dd><a href="https://www.dlsite.com/home/works/type/=/work_type/STG">シューティング</a></dd>
                                          <dd><a href="https://www.dlsite.com/home/works/type/=/work_type/PZL">パズル</a></dd>
                                          <dd><a href="https://www.dlsite.com/home/works/type/=/work_type/ETC">その他ゲーム</a></dd>
                                                      </div>
              </dl>
            </li>
                      <li>
              <dl>
                <dt>マンガ</dt>
                <div>
                                                        <dd><a href="https://www.dlsite.com/home/works/type/=/work_type_category/comic">すべて</a></dd>
                                                                      </div>
              </dl>
            </li>
                      <li>
              <dl>
                <dt>CG・イラスト</dt>
                <div>
                                                        <dd><a href="https://www.dlsite.com/home/works/type/=/work_type_category/illust">すべて</a></dd>
                                                                      </div>
              </dl>
            </li>
                      <li>
              <dl>
                <dt>ノベル</dt>
                <div>
                                                        <dd><a href="https://www.dlsite.com/home/works/type/=/work_type_category/novel">すべて</a></dd>
                                                                      </div>
              </dl>
            </li>
                      <li>
              <dl>
                <dt>動画</dt>
                <div>
                                                        <dd><a href="https://www.dlsite.com/home/works/type/=/work_type_category/movie">すべて</a></dd>
                                                                      </div>
              </dl>
            </li>
                      <li>
              <dl>
                <dt>音楽</dt>
                <div>
                                                        <dd><a href="https://www.dlsite.com/home/works/type/=/work_type_category/music">すべて</a></dd>
                                                                      </div>
              </dl>
            </li>
                      <li>
              <dl>
                <dt>ツール/アクセサリ</dt>
                <div>
                                                        <dd><a href="https://www.dlsite.com/home/works/type/=/work_type_category/tool">すべて</a></dd>
                                                                                                <dd><a href="https://www.dlsite.com/home/works/type/=/work_type/TOL">ツール/アクセサリ</a></dd>
                                          <dd><a href="https://www.dlsite.com/home/works/type/=/work_type/IMT">画像素材</a></dd>
                                          <dd><a href="https://www.dlsite.com/home/works/type/=/work_type/AMT">音素材</a></dd>
                                                      </div>
              </dl>
            </li>
                      <li>
              <dl>
                <dt>その他</dt>
                <div>
                                                        <dd><a href="https://www.dlsite.com/home/works/type/=/work_type_category/etc">すべて</a></dd>
                                                                      </div>
              </dl>
            </li>
                  </ul>
      </div>
    </div>
  </div>

              <div class="contents_bottom">
    <div class="contents_bottom_inner">
      <p class="contents_bottom_headline">ボイス・ASMRの人気ジャンルで絞り込む</p>
      <div class="recommend_item">
        <ul>
                      <li><a href="https://www.dlsite.com/home/works/type/=/work_type_category/audio/genre/497">ASMR</a></li>
                      <li><a href="https://www.dlsite.com/home/works/type/=/work_type_category/audio/genre/056">癒し</a></li>
                      <li><a href="https://www.dlsite.com/home/works/type/=/work_type_category/audio/genre/496">バイノーラル/ダミヘ</a></li>
                      <li><a href="https://www.dlsite.com/home/works/type/=/work_type_category/audio/genre/442">耳かき</a></li>
                      <li><a href="https://www.dlsite.com/home/works/type/=/work_type_category/audio/genre/503">ささやき</a></li>
                      <li><a href="https://www.dlsite.com/home/works/type/=/work_type_category/audio/genre/051">萌え</a></li>
                      <li><a href="https://www.dlsite.com/home/works/type/=/work_type_category/audio/genre/053">健全</a></li>
                      <li><a href="https://www.dlsite.com/home/works/type/=/work_type_category/audio/genre/004">ラブラブ/あまあま</a></li>
                  </ul>
      </div>
    </div>
  </div>

          
  
  <div class="contents_bottom">
    <div class="contents_bottom_inner">
      <p class="contents_bottom_headline">
                  DLsite（ディーエルサイト）について              </p>
      <div class="contents_bottom site_info">
        <p>
                      DLsiteは同人誌・漫画（マンガ）・コミック・美少女ゲームなどを取り扱っている二次元コンテンツダウンロードサイトです。                  </p>
        <p>パソコン・スマートフォンどちらからでも楽しめるコンテンツが盛りだくさん！</p>

                                    <p>「<strong>ボイス・ASMR</strong>」を始めとする様々な作品を取り揃えております！</p>
                
        <p>新着の作品も毎日更新中！</p>
      </div>
    </div>
  </div>

          
          
        </div>
        <!-- /main_inner -->
      </div>
      <!-- /main -->

    </div>
    <!-- /wrapper -->

    <!-- left -->
    <div id="left" >
      
      
      
      
      
      
      
      
      
      <div class="left_module arrow type_search" id="lm_search" data-section_name="left_search">
  <h3>同人作品を探す</h3>
  <div class="left_module_content">

    <div class="list_content">
      <ul class="list_content_text type-icon">
                <li class="list_content_text_item"><a class="free" href="https://www.dlsite.com/home/fsr/=/options%5B0%5D/JPN/options%5B1%5D/NM/work_category/doujin/is_free/1/exclude_work_type/vcm/order/trend/per_page/100">無料作品<span class="mark_label">pick up!</span></a></li>
        <li class="list_content_text_item"><a class="sale" href="https://www.dlsite.com/home/works/discount">割引中の作品<span class="mark_label">おすすめ</span></a></li>
        <li class="list_content_text_item"><a class="only" href="https://www.dlsite.com/home/works/exclusive">専売作品<span class="mark_label">pick up!</span></a></li>
        <li class="list_content_text_item"><a class="review" href="https://www.dlsite.com/home/new/review">新着レビュー</a></li>
              </ul>
    </div>

          <div id="slide_worktype" class="list_head"><h4><span class="slide_title">作品形式</span></h4></div>
      <div class="list_content">
        <ul class="list_content_text">
          <li class="list_content_text_item"><a href="https://www.dlsite.com/home/fsr/=/work_type_category/game/work_category/doujin/from/left_pain.work_type">ゲーム</a>
            <ul class="list_content_indent">
              <li class="list_text_indent"><a href="https://www.dlsite.com/home/fsr/=/work_type/RPG/work_category/doujin/from/left_pain.work_type">ロールプレイング</a></li>
              <li class="list_text_indent"><a href="https://www.dlsite.com/home/fsr/=/work_type/ACN/work_category/doujin/from/left_pain.work_type">アクション</a></li>
              <li class="list_text_indent"><a href="https://www.dlsite.com/home/fsr/=/work_type/SLN/work_category/doujin/from/left_pain.work_type">シミュレーション</a></li>
              <li class="list_text_indent"><a href="https://www.dlsite.com/home/fsr/=/work_type/ADV/work_category/doujin/from/left_pain.work_type">アドベンチャー</a></li>
              <li class="list_text_indent"><a href="https://www.dlsite.com/home/fsr/=/work_type/STG/work_category/doujin/from/left_pain.work_type">シューティング</a></li>
              <li class="list_text_indent"><a href="https://www.dlsite.com/home/fsr/=/work_type%5B0%5D/DNV/work_type%5B1%5D/QIZ/work_type%5B2%5D/ETC/work_type%5B3%5D/TBL/work_type%5B4%5D/TYP/work_type%5B5%5D/PZL/work_category/doujin/from/left_pain.work_type">その他ゲーム</a></li>
            </ul>
          </li>

          <li class="list_content_text_item"><a href="https://www.dlsite.com/home/fsr/=/work_type_category/comic/work_category/doujin/from/left_pain.work_type">マンガ</a></li>
          <li class="list_content_text_item"><a href="https://www.dlsite.com/home/fsr/=/work_type/ICG/work_category/doujin/from/left_pain.work_type">CG・イラスト</a></li>
          <li class="list_content_text_item"><a href="https://www.dlsite.com/home/fsr/=/work_type/MOV/work_category/doujin/from/left_pain.work_type">動画作品</a></li>
          <li class="list_content_text_item"><a href="https://www.dlsite.com/home/fsr/=/work_type/SOU/work_category/doujin/from/left_pain.work_type">ボイス・ASMR</a></li>
          <li class="list_content_text_item"><a href="https://www.dlsite.com/home/fsr/=/work_type/MUS/work_category/doujin/from/left_pain.work_type">音楽作品</a></li>
          <li class="list_content_text_item"><a href="https://www.dlsite.com/home/fsr/=/work_type_category/tool/work_category/doujin/from/left_pain.work_type">ツール/アクセサリ</a></li>
        </ul>
        <p class="link_list_item"><a href="https://www.dlsite.com/home/worktype/list">他の作品形式で探す</a></p>
      </div>
    
    <div id="slide_keyword" class="list_head"><h4><span class="slide_title">ジャンル</span></h4></div>
    <div class="list_content">
      <ul class="list_content_text clearfix">
                  <li class="list_content_text_item">
                          <a href="https://www.dlsite.com/home/works/genre/=/genre/442">耳かき <span>(1677)</span></a>
                      </li>
                  <li class="list_content_text_item">
                          <a href="https://www.dlsite.com/home/works/genre/=/genre/056">癒し <span>(5603)</span></a>
                      </li>
                  <li class="list_content_text_item">
                          <a href="https://www.dlsite.com/home/works/genre/=/genre/496">バイノーラル/ダミヘ <span>(1716)</span></a>
                      </li>
                  <li class="list_content_text_item">
                          <a href="https://www.dlsite.com/home/works/genre/=/genre/013">ほのぼの <span>(6771)</span></a>
                      </li>
                  <li class="list_content_text_item">
                          <a href="https://www.dlsite.com/home/works/genre/=/genre/004">ラブラブ/あまあま <span>(2746)</span></a>
                      </li>
                  <li class="list_content_text_item">
                          <a href="https://www.dlsite.com/home/works/genre/=/genre/158">百合 <span>(1621)</span></a>
                      </li>
                  <li class="list_content_text_item">
                          <a href="https://www.dlsite.com/home/works/genre/=/genre/245">性転換(TS) <span>(247)</span></a>
                      </li>
                  <li class="list_content_text_item">
                          <a href="https://www.dlsite.com/home/works/genre/=/genre/497">ASMR <span>(2115)</span></a>
                      </li>
                  <li class="list_content_text_item">
                          <a href="https://www.dlsite.com/home/works/genre/=/genre/016">ファンタジー <span>(4516)</span></a>
                      </li>
                  <li class="list_content_text_item">
                          <a href="https://www.dlsite.com/home/works/genre/=/genre/078">メイド <span>(825)</span></a>
                      </li>
                  <li class="list_content_text_item">
                          <a href="https://www.dlsite.com/home/works/genre/=/genre/240">魔法少女 <span>(851)</span></a>
                      </li>
              </ul>
      <div class="link_list_item"><a href="https://www.dlsite.com/home/genre/list">他のジャンルで探す</a></div>
    </div>

        <div id="slide_event" class="list_head"><h4><span class="slide_title">イベント</span></h4></div>
    <div class="list_content">
      <ul class="list_content_text">
        <li class="list_content_text_item"><a href="https://www.dlsite.com/home/fsr/=/ana_flg/all/options/C99/from/icon.work">コミックマーケット99</a></li>
      </ul>
    </div>
    
    <div class="list_content_border">
      <div class="list_head">
        <h4>その他の販売フロア</h4>
      </div>
      <ul class="link_list">
        <li class="link_list_item"><a href="https://www.dlsite.com/maniax/">男性成人向け作品へ</a></li>
        <li class="link_list_item"><a href="https://www.dlsite.com/girls/">女性向け作品へ</a></li>
      </ul>
    </div>

  </div>
</div>
      
      
      
      
      
      
      
      
      
      
      
      
      
      
      
      
      
      
      
      
      
      
      
      
      
      
      
      
      
      
      
      
      
    </div>
    <!-- /left -->

                  <!-- nijiyome_bnr -->
    <div id="footer_nijiyome_banner" class="bn_footer_1box" data-section_name="footer_banner">
      <div class="bcs_viewer" data-service="dlsite" data-site="common" data-frame="pc-long-banner-general" data-ext-action="random" data-ext-empty="footer_nijiyome_banner">
      </div>
    </div>
  
    
    
    
    <!-- footer -->
    <div id="footer">
      
      <div class="pagetop_block">
    <p class="pagetop"><a href="#header">このページの上部へ</a></p>
</div>
<div class="footer_floor_nav">
<ul class="floor_list">
<li class="floor_list_item"><a href="https://www.dlsite.com">総合トップ</a></li>
<li class="floor_list_item"><a href="https://www.dlsite.com/home/">同人</a></li>
<li class="floor_list_item"><a href="https://www.dlsite.com/comic/">コミック</a></li>
<li class="floor_list_item"><a href="https://www.dlsite.com/soft/">PCソフト</a></li>
<li class="floor_list_item"><a href="https://www.dlsite.com/app/">アプリ</a></li>
<li class="floor_list_item sp_switch"><a id="_touch_link" href="https://www.dlsite.com/home-touch/works/type/=/work_type_category/audio" data-platform="touch">スマホ版DLsiteへ</a></li></ul>
</div>

<div class="footer_section">
  <div class="footer_section_inner">
    <div class="link_list_wrap">
      <div class="link_list_box col_2">
                <div class="label">DLsiteについて</div>
                <ul class="link_list">
                    <li class="link_list_item"><a rel="noopener" href="https://www.eisys.co.jp/company/information" target="_blank">会社概要</a></li>
                              <li class="link_list_item"><a rel="noopener" href="https://www.eisys.co.jp/recruit" target="_blank">採用情報</a></li>
                    <li class="link_list_item"><a href="https://www.dlsite.com/home/user/regulations">利用規約</a></li>
          <li class="link_list_item"><a href="https://www.dlsite.com/home/guide/law">特定商取引法に基づく表示</a></li>
          <li class="link_list_item"><a href="https://www.dlsite.com/home/guide/settlement">資金決済法に基づく表記</a></li>
          <li class="link_list_item"><a href="https://www.dlsite.com/home/guide/privacy">個人情報の取扱いについて</a></li>
          <li class="link_list_item"><a href="https://www.dlsite.com/home/mosaic">コンプライアンスポリシー</a></li>
          <li class="link_list_item"><a href="https://www.dlsite.com/home/guide/copy">著作権</a></li>
                    <li class="link_list_item"><a href="https://www.dlsite.com/home/banners">リンクについて</a></li>
          <li class="link_list_item"><a href="https://www.dlsite.com/home/sitemap">サイトマップ</a></li>
                  </ul>
      </div>
      <div class="link_list_box">
        <div class="label">ヘルプ&amp;ガイド</div>
        <ul class="link_list">
                    <li class="link_list_item"><a href="https://www.dlsite.com/home/welcome">初めての方へ</a></li>
                    <li class="link_list_item"><a rel="noopener" href="https://www.dlsite.com/home/faq/=/type/user" target="_blank">よくある質問</a></li>
                    <li class="link_list_item"><a href="https://www.dlsite.com/home/opinion/contribution">作品リクエスト</a></li>
                  </ul>
      </div>
      <div class="link_list_box">
                <div class="label">DLsiteのサービス</div>
                <ul class="link_list">
                    <li class="link_list_item"><a href="https://www.dlsite.com/home/circle/invite">作品を販売したい方へ</a></li>
                              <li class="link_list_item"><a href="https://www.dlsite.com/home/sell/steam">公式翻訳サービス</a></li>
          <li class="link_list_item"><a href="https://www.dlsite.com/home/guide/affiliate">アフィリエイト</a></li>
                    <li class="link_list_item"><a href="https://www.dlsite.com/home/mypage/setting/mail">メールマガジン</a></li>
                    <li class="link_list_item"><a rel="noopener" href="https://play.dlsite.com/ja" target="_blank">Webビューア DLsite Play</a></li>
          <li class="link_list_item"><a href="https://www.dlsite.com/home/guide/dlnest">PCクライアント DLsite Nest</a></li>
                  </ul>
      </div>
    </div>
  </div>
  <div class="footer_section_inner payment">
    <div class="label">お支払い&amp;ポイント</div>
    <div class="footer_payment_box">
      <ul class="link_list">
        <li class="link_list_item"><a href="https://www.dlsite.com/home/guide/creditcard">お支払い方法</a></li>
                        <li class="link_list_item"><a href="https://www.dlsite.com/home/mypage/aboutpoint">ポイントについて</a></li>
                <li class="link_list_item"><a href="https://www.dlsite.com/home/guide/payment#link_point">ポイント購入方法</a></li>
      </ul>
      <a href="https://www.dlsite.com/home/guide/creditcard" class="footer_payment_img">
                <img src="/modpub/images/web/common/payment_wide.jpg" alt="多彩な決済方法">
              </a>
    </div>
  </div>
      <div class="footer_section_inner sns">
        <div class="label">Twitter公式アカウント</div>
    <ul class="footer_sns">
      <li class="footer_sns_item"><a rel="noopener" href="https://twitter.com/DLsite" target="_blank" class="twitter_g">全年齢</a></li>
      <li class="footer_sns_item"><a rel="noopener" href="https://twitter.com/DLsiteManiax" target="_blank" class="twitter_r18">R18</a></li>
      <li class="footer_sns_item"><a rel="noopener" href="https://twitter.com/GirlsManiax" target="_blank" class="twitter_bl">BL</a></li>
      <li class="footer_sns_item"><a rel="noopener" href="https://twitter.com/girlsmaniax_OTM" target="_blank" class="twitter_otm">乙女</a></li>
      <li class="footer_sns_item"><a rel="noopener" href="https://twitter.com/DLsiteC" target="_blank" class="twitter_comic">Comic</a></li>
      <li class="footer_sns_item"><a rel="noopener" href="https://twitter.com/DLsiteR18" target="_blank" class="twitter_comic_r18">Comic R18</a></li>
      <li class="footer_sns_item"><a rel="noopener" href="https://twitter.com/DLsiteEnglish" target="_blank" class="twitter_eng">English</a></li>
      <li class="footer_sns_item"><a rel="noopener" href="https://twitter.com/DLsite_info" target="_blank" class="twitter_info">広報</a></li>
    </ul>
    <div class="label">Youtube公式アカウント</div>
    <ul class="footer_sns">
      <li class="footer_sns_item"><a rel="noopener" href="https://www.youtube.com/channel/UCQEN3LsNnqottC2mXx3tAjA" target="_blank" class="youtube_info">広報</a></li>
    </ul>
        <!-- 日本語以外で共通で出すアイコン群 -->
      </div>
    <div class="footer_section_inner recruit">
    <div class="label">採用情報（グループ一括採用）</div>
    <ul class="job_list">
      <li class="job_list_item new">
        <a rel="noopener" href="https://hrmos.co/pages/vivion/jobs/0000272" target="_blank">
          <div class="job_icon recruit_01">
            <img src="/modpub/images/web/common/recruit/icon_recruit_02.png" loading="lazy">
          </div>
          <div class="job_info">
            <p class="job_name">webライター（アルバイト）</p>
          </div>
        </a>
      </li>
      <li class="job_list_item new">
        <a rel="noopener" href="https://hrmos.co/pages/vivion/jobs/0000181" target="_blank">
          <div class="job_icon recruit_02">
            <img src="/modpub/images/web/common/recruit/icon_recruit_07.png" loading="lazy">
          </div>
          <div class="job_info">
            <p class="job_name">動画制作クリエイター</p>
          </div>
        </a>
      </li>
      <li class="job_list_item new">
        <a rel="noopener" href="https://hrmos.co/pages/vivion/jobs/0000258" target="_blank">
          <div class="job_icon recruit_03">
            <img src="/modpub/images/web/common/recruit/icon_recruit_01.png" loading="lazy">
          </div>
          <div class="job_info">
            <p class="job_name">ゲームデバックスタッフ(アルバイト)</p>
          </div>
        </a>
      </li>
      <li class="job_list_item new">
        <a rel="noopener" href="https://hrmos.co/pages/vivion/jobs/0CLI003" target="_blank">
          <div class="job_icon recruit_04">
            <img src="/modpub/images/web/common/recruit/icon_recruit_02.png" loading="lazy">
          </div>
          <div class="job_info">
            <p class="job_name">デジタルコミックデータ登録(アルバイト)</p>
          </div>
        </a>
      </li>
    </ul>
    <a rel="noopener" href="https://www.eisys.co.jp/recruit" class="recruit_more" target="_blank">採用サイトへ</a>
  </div>
    </div>

<div id="copyright">
  <div class="container clearfix">
    <div id="system">推奨環境：最新版のMicrosoft Edge、Safari、Chrome、Firefox（JavaScript・Cookieを許可）</div>
        <p>&copy; 1996 DLsite</p>
      </div>
</div>

<div data-vue-component="thumb-img-popup-in-swiper"></div>
    </div>
    <!-- /footer -->

  </div>
  <!-- /container -->

    <div data-vue-component="custom-dialog" v-cloak></div>

    <!-- script_footer -->
  
<script type="application/ld+json">
{
  "@context": "https://schema.org",
  "@type": "WebSite",
  "name": "DLsite",
  "alternateName": "DLsite",
  "url": "https://www.dlsite.com/"
}
</script>


<div data-vue-component="cookie-policy" data-async="true"></div>

<script type="text/javascript" src="/vue/js/pc/vendor.js?cdn_cache=1&v=0.1.2&_=1643863895"></script>
<script type="text/javascript" src="/vue/js/pc/app.js?cdn_cache=1&v=0.1.2&_=1651458700"></script>


<script type="text/javascript">
(function(document, undefined) {
  "use strict";
  var $script = document.createElement('script');
  $script.src = document.location.protocol + '//banner.eisys-bcs.jp/js/bcs.min.js';
  document.body.appendChild($script);

  $(window).bind('guestbannerDelete', function(){
    var $guest_banner = $('.bcs_viewer').find('.__guest');
    if ($.cookie('loginchecked') >= 1 && $guest_banner.length > 0 ) {
      $guest_banner.remove();
    }
  });

  // フッターリンクQRコードのモーダル
  // 開き
  $('.qr-modal-open-link').bind('click', function(e) {
    e.preventDefault();
    var targetModal = $(this).attr('data-modal-open');
    $('[data-modal-type =' + targetModal + ']').addClass('is-show');
  })
  // 閉じ（ボタン）
  $('.global_modal_content_close').bind('click', function(e) {
    $('[data-modal-type]').removeClass('is-show');
  })
  // 閉じ（オーバーレイ）
  $('.global_modal_overlay.type_qr').bind('click', function(e) {
    $('[data-modal-type]').removeClass('is-show');
  })
  $('.global_modal_content.type_qr').bind('click', function(e) {
    e.stopPropagation();
  })
})(document);
</script>



<style>
div.measure_tag {
  height: 0 !important;
  width: 0 !important;
  line-height: 0 !important;
  font-size: 0 !important;
  margin-top: -10000px;
  margin-left: -10000px;
  float: left;
}
</style>

<div class="measure_tag"></div>
  <!-- /script_footer -->
    <script type="text/javascript">var contents = {"impression":[],"time":2.9087066650390625e-5};</script>
<script type="text/javascript">window.NREUM||(NREUM={});NREUM.info={"beacon":"bam.nr-data.net","licenseKey":"NRJS-26fb73475871c01d5bd","applicationID":"781370968","transactionName":"YwQBYkdWDRZQVxUKDlpOIENGQwwIHkMOEQpHPhdPRVI=","queueTime":0,"applicationTime":8182,"atts":"T0MCFA9MHhg=","errorBeacon":"bam.nr-data.net","agent":""}</script></body>

</html>`
