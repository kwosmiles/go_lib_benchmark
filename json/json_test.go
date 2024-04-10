package jsontest

import (
	"runtime"
	"testing"

	"github.com/mailru/easyjson"
)

const (
	smallbytes  = `{"event":"published","room":{"sid":"RM_HJ9ZPdwpdX9K","name":"16ad04807d634f61a21e87e1fba794e2"},"participant":{"sid":"PA_ZX2GTaNeiS","identity":"DA_0198106858"},"track":{"sid":"TR_AMm8UpaDKmhBxz","source":2,"mime_type":"audio/opus","mid":"0","stream":"camera","version":{"unix_micro":1711753519746941}},"id":"EV_M8jdmHerSvjg","created_at":1711753519}`
	mediumbytes = `{"person":{"id":"d50887ca-a6ce-4e59-b89f-14f0b5d03b03","name":{"fullName":"LeonidBugaev","givenName":"Leonid","familyName":"Bugaev"},"email":"leonsbox@gmail.com","gender":"male","location":"SaintPetersburg,SaintPetersburg,RU","geo":{"city":"SaintPetersburg","state":"SaintPetersburg","country":"Russia","lat":59.9342802,"lng":30.3350986},"bio":"SeniorengineeratGranify.com","site":"http://flickfaver.com","avatar":"https://d1ts43dypk8bqh.cloudfront.net/v1/avatars/d50887ca-a6ce-4e59-b89f-14f0b5d03b03","employment":{"name":"www.latera.ru"},"facebook":{"handle":"leonid.bugaev"},"github":{"handle":"buger","id":14009,"avatar":"https://avatars.githubusercontent.com/u/14009?v=3","company":"Granify","blog":"http://leonsbox.com","followers":95,"following":10},"twitter":{"handle":"flickfaver","id":77004410,"bio":null,"followers":2,"site":"http://flickfaver.com","avatar":null},"linkedin":{"handle":"in/leonidbugaev"},"googleplus":{"handle":null},"angellist":{"handle":"leonid-bugaev","id":61541,"bio":"SeniorengineeratGranify.com","followers":41,"avatar":"https://d1qb2nb5cznatu.cloudfront.net/users/61541-medium_jpg?1405474390"},"klout":{"handle":null,"score":null},"foursquare":{"handle":null},"aboutme":{"handle":"leonid.bugaev","bio":null,"avatar":null},"gravatar":{"handle":"buger","urls":[],"avatar":"http://1.gravatar.com/avatar/f7c8edd577d13b8930d5522f281235","avatars":[{"url":"http://1.gravatar.com/avatar/f7c8edd577d13b8930d5522f28123510","type":"thumbnail"}]},"fuzzy":false},"company":null}`
	bigbytes    = `{"users":[{"id":-1,"username":"system","avatar_template":"/user_avatar/discourse.metabase.com/system/{size}/6_1.png"},{"id":89,"username":"zergot","avatar_template":"https://avatars.discourse.org/v2/letter/z/0ea827/{size}.png"},{"id":1,"username":"sameer","avatar_template":"https://avatars.discourse.org/v2/letter/s/bbce88/{size}.png"},{"id":84,"username":"HenryMirror","avatar_template":"https://avatars.discourse.org/v2/letter/h/ecd19e/{size}.png"},{"id":73,"username":"fimp","avatar_template":"https://avatars.discourse.org/v2/letter/f/ee59a6/{size}.png"},{"id":14,"username":"agilliland","avatar_template":"/user_avatar/discourse.metabase.com/agilliland/{size}/26_1.png"},{"id":87,"username":"amir","avatar_template":"https://avatars.discourse.org/v2/letter/a/c37758/{size}.png"},{"id":82,"username":"waseem","avatar_template":"https://avatars.discourse.org/v2/letter/w/9dc877/{size}.png"},{"id":78,"username":"tovenaar","avatar_template":"https://avatars.discourse.org/v2/letter/t/9de0a6/{size}.png"},{"id":74,"username":"Ben","avatar_template":"https://avatars.discourse.org/v2/letter/b/df788c/{size}.png"},{"id":71,"username":"MarkLaFay","avatar_template":"https://avatars.discourse.org/v2/letter/m/3bc359/{size}.png"},{"id":72,"username":"camsaul","avatar_template":"/user_avatar/discourse.metabase.com/camsaul/{size}/70_1.png"},{"id":53,"username":"mhjb","avatar_template":"/user_avatar/discourse.metabase.com/mhjb/{size}/54_1.png"},{"id":58,"username":"jbwiv","avatar_template":"https://avatars.discourse.org/v2/letter/j/6bbea6/{size}.png"},{"id":70,"username":"Maggs","avatar_template":"https://avatars.discourse.org/v2/letter/m/bbce88/{size}.png"},{"id":69,"username":"andrefaria","avatar_template":"/user_avatar/discourse.metabase.com/andrefaria/{size}/65_1.png"},{"id":60,"username":"bencarter78","avatar_template":"/user_avatar/discourse.metabase.com/bencarter78/{size}/59_1.png"},{"id":55,"username":"vikram","avatar_template":"https://avatars.discourse.org/v2/letter/v/e47774/{size}.png"},{"id":68,"username":"edchan77","avatar_template":"/user_avatar/discourse.metabase.com/edchan77/{size}/66_1.png"},{"id":9,"username":"karthikd","avatar_template":"https://avatars.discourse.org/v2/letter/k/cab0a1/{size}.png"},{"id":23,"username":"arthurz","avatar_template":"/user_avatar/discourse.metabase.com/arthurz/{size}/32_1.png"},{"id":3,"username":"tom","avatar_template":"/user_avatar/discourse.metabase.com/tom/{size}/21_1.png"},{"id":50,"username":"LeoNogueira","avatar_template":"/user_avatar/discourse.metabase.com/leonogueira/{size}/52_1.png"},{"id":66,"username":"ss06vi","avatar_template":"https://avatars.discourse.org/v2/letter/s/3ab097/{size}.png"},{"id":34,"username":"mattcollins","avatar_template":"/user_avatar/discourse.metabase.com/mattcollins/{size}/41_1.png"},{"id":51,"username":"krmmalik","avatar_template":"/user_avatar/discourse.metabase.com/krmmalik/{size}/53_1.png"},{"id":46,"username":"odysseas","avatar_template":"https://avatars.discourse.org/v2/letter/o/5f8ce5/{size}.png"},{"id":5,"username":"jonthewayne","avatar_template":"/user_avatar/discourse.metabase.com/jonthewayne/{size}/18_1.png"},{"id":11,"username":"anandiyer","avatar_template":"/user_avatar/discourse.metabase.com/anandiyer/{size}/23_1.png"},{"id":25,"username":"alnorth","avatar_template":"/user_avatar/discourse.metabase.com/alnorth/{size}/34_1.png"},{"id":52,"username":"j_at_svg","avatar_template":"https://avatars.discourse.org/v2/letter/j/96bed5/{size}.png"},{"id":42,"username":"styts","avatar_template":"/user_avatar/discourse.metabase.com/styts/{size}/47_1.png"}],"topics":{"can_create_topic":false,"more_topics_url":"/c/uncategorized/l/latest?page=1","draft":null,"draft_key":"new_topic","draft_sequence":null,"per_page":30,"topics":[{"id":8,"title":"WelcometoMetabase'sDiscussionForum","fancy_title":"WelcometoMetabase&rsquo;sDiscussionForum","slug":"welcome-to-metabases-discussion-forum","posts_count":1,"reply_count":0,"highest_post_number":1,"image_url":"/images/welcome/discourse-edit-post-animated.gif","created_at":"2015-10-17T00:14:49.526Z","last_posted_at":"2015-10-17T00:14:49.557Z","bumped":true,"bumped_at":"2015-10-21T02:32:22.486Z","unseen":false,"pinned":true,"unpinned":null,"excerpt":"WelcometoMetabase&#39;sdiscussionforum.Thisisaplacetogethelponinstallation,settingupaswellassharingtipsandtricks.","visible":true,"closed":false,"archived":false,"bookmarked":null,"liked":null,"views":197,"like_count":0,"has_summary":false,"archetype":"regular","last_poster_username":"system","category_id":1,"pinned_globally":true,"posters":[{"extras":"latestsingle","description":"OriginalPoster,MostRecentPoster","user_id":-1}]},{"id":169,"title":"FormattingDates","fancy_title":"FormattingDates","slug":"formatting-dates","posts_count":1,"reply_count":0,"highest_post_number":1,"image_url":null,"created_at":"2016-01-14T06:30:45.311Z","last_posted_at":"2016-01-14T06:30:45.397Z","bumped":true,"bumped_at":"2016-01-14T06:30:45.397Z","unseen":false,"pinned":false,"unpinned":null,"visible":true,"closed":false,"archived":false,"bookmarked":null,"liked":null,"views":11,"like_count":0,"has_summary":false,"archetype":"regular","last_poster_username":"zergot","category_id":1,"pinned_globally":false,"posters":[{"extras":"latestsingle","description":"OriginalPoster,MostRecentPoster","user_id":89}]},{"id":168,"title":"Settingforgoogleapikey","fancy_title":"Settingforgoogleapikey","slug":"setting-for-google-api-key","posts_count":2,"reply_count":0,"highest_post_number":2,"image_url":null,"created_at":"2016-01-13T17:14:31.799Z","last_posted_at":"2016-01-14T06:24:03.421Z","bumped":true,"bumped_at":"2016-01-14T06:24:03.421Z","unseen":false,"pinned":false,"unpinned":null,"visible":true,"closed":false,"archived":false,"bookmarked":null,"liked":null,"views":16,"like_count":0,"has_summary":false,"archetype":"regular","last_poster_username":"zergot","category_id":1,"pinned_globally":false,"posters":[{"extras":"latestsingle","description":"OriginalPoster,MostRecentPoster","user_id":89}]},{"id":167,"title":"Cannotseenon-UStimezonesontheadmin","fancy_title":"Cannotseenon-UStimezonesontheadmin","slug":"cannot-see-non-us-timezones-on-the-admin","posts_count":1,"reply_count":0,"highest_post_number":1,"image_url":null,"created_at":"2016-01-13T17:07:36.764Z","last_posted_at":"2016-01-13T17:07:36.831Z","bumped":true,"bumped_at":"2016-01-13T17:07:36.831Z","unseen":false,"pinned":false,"unpinned":null,"visible":true,"closed":false,"archived":false,"bookmarked":null,"liked":null,"views":11,"like_count":0,"has_summary":false,"archetype":"regular","last_poster_username":"zergot","category_id":1,"pinned_globally":false,"posters":[{"extras":"latestsingle","description":"OriginalPoster,MostRecentPoster","user_id":89}]},{"id":164,"title":"External(Metabaselevel)linkagesindataschema","fancy_title":"External(Metabaselevel)linkagesindataschema","slug":"external-metabase-level-linkages-in-data-schema","posts_count":4,"reply_count":1,"highest_post_number":4,"image_url":null,"created_at":"2016-01-11T13:51:02.286Z","last_posted_at":"2016-01-12T11:06:37.259Z","bumped":true,"bumped_at":"2016-01-12T11:06:37.259Z","unseen":false,"pinned":false,"unpinned":null,"visible":true,"closed":false,"archived":false,"bookmarked":null,"liked":null,"views":32,"like_count":0,"has_summary":false,"archetype":"regular","last_poster_username":"zergot","category_id":1,"pinned_globally":false,"posters":[{"extras":"latest","description":"OriginalPoster,MostRecentPoster","user_id":89},{"extras":null,"description":"FrequentPoster","user_id":1}]},{"id":155,"title":"Queryworkingon\"Questions\"butnotin\"Pulses\"","fancy_title":"Queryworkingon&ldquo;Questions&rdquo;butnotin&ldquo;Pulses&rdquo;","slug":"query-working-on-questions-but-not-in-pulses","posts_count":3,"reply_count":0,"highest_post_number":3,"image_url":null,"created_at":"2016-01-01T14:06:10.083Z","last_posted_at":"2016-01-08T22:37:51.772Z","bumped":true,"bumped_at":"2016-01-08T22:37:51.772Z","unseen":false,"pinned":false,"unpinned":null,"visible":true,"closed":false,"archived":false,"bookmarked":null,"liked":null,"views":72,"like_count":0,"has_summary":false,"archetype":"regular","last_poster_username":"agilliland","category_id":1,"pinned_globally":false,"posters":[{"extras":null,"description":"OriginalPoster","user_id":84},{"extras":null,"description":"FrequentPoster","user_id":73},{"extras":"latest","description":"MostRecentPoster","user_id":14}]},{"id":161,"title":"PulsespostedtoSlackdon'tshowquestionoutput","fancy_title":"PulsespostedtoSlackdon&rsquo;tshowquestionoutput","slug":"pulses-posted-to-slack-dont-show-question-output","posts_count":2,"reply_count":0,"highest_post_number":2,"image_url":"/uploads/default/original/1X/9d2806517bf3598b10be135b2c58923b47ba23e7.png","created_at":"2016-01-08T22:09:58.205Z","last_posted_at":"2016-01-08T22:28:44.685Z","bumped":true,"bumped_at":"2016-01-08T22:28:44.685Z","unseen":false,"pinned":false,"unpinned":null,"visible":true,"closed":false,"archived":false,"bookmarked":null,"liked":null,"views":34,"like_count":0,"has_summary":false,"archetype":"regular","last_poster_username":"sameer","category_id":1,"pinned_globally":false,"posters":[{"extras":null,"description":"OriginalPoster","user_id":87},{"extras":"latest","description":"MostRecentPoster","user_id":1}]},{"id":152,"title":"ShouldwebuildKafkaconnecterorKafkaplugin","fancy_title":"ShouldwebuildKafkaconnecterorKafkaplugin","slug":"should-we-build-kafka-connecter-or-kafka-plugin","posts_count":4,"reply_count":1,"highest_post_number":4,"image_url":null,"created_at":"2015-12-28T20:37:23.501Z","last_posted_at":"2015-12-31T18:16:45.477Z","bumped":true,"bumped_at":"2015-12-31T18:16:45.477Z","unseen":false,"pinned":false,"unpinned":null,"visible":true,"closed":false,"archived":false,"bookmarked":null,"liked":null,"views":84,"like_count":0,"has_summary":false,"archetype":"regular","last_poster_username":"sameer","category_id":1,"pinned_globally":false,"posters":[{"extras":null,"description":"OriginalPoster","user_id":82},{"extras":"latest","description":"MostRecentPoster,FrequentPoster","user_id":1}]},{"id":147,"title":"ChangeXandYongraph","fancy_title":"ChangeXandYongraph","slug":"change-x-and-y-on-graph","posts_count":1,"reply_count":0,"highest_post_number":1,"image_url":null,"created_at":"2015-12-21T17:52:46.581Z","last_posted_at":"2015-12-21T17:52:46.684Z","bumped":true,"bumped_at":"2015-12-21T18:19:13.003Z","unseen":false,"pinned":false,"unpinned":null,"visible":true,"closed":false,"archived":false,"bookmarked":null,"liked":null,"views":68,"like_count":0,"has_summary":false,"archetype":"regular","last_poster_username":"tovenaar","category_id":1,"pinned_globally":false,"posters":[{"extras":"latestsingle","description":"OriginalPoster,MostRecentPoster","user_id":78}]},{"id":142,"title":"Issuessendingmailviaoffice365relay","fancy_title":"Issuessendingmailviaoffice365relay","slug":"issues-sending-mail-via-office365-relay","posts_count":5,"reply_count":2,"highest_post_number":5,"image_url":null,"created_at":"2015-12-16T10:38:47.315Z","last_posted_at":"2015-12-21T09:26:27.167Z","bumped":true,"bumped_at":"2015-12-21T09:26:27.167Z","unseen":false,"pinned":false,"unpinned":null,"visible":true,"closed":false,"archived":false,"bookmarked":null,"liked":null,"views":122,"like_count":0,"has_summary":false,"archetype":"regular","last_poster_username":"Ben","category_id":1,"pinned_globally":false,"posters":[{"extras":"latest","description":"OriginalPoster,MostRecentPoster","user_id":74},{"extras":null,"description":"FrequentPoster","user_id":1}]},{"id":137,"title":"IseetriplicatesofmymongoDBcollections","fancy_title":"IseetriplicatesofmymongoDBcollections","slug":"i-see-triplicates-of-my-mongodb-collections","posts_count":3,"reply_count":0,"highest_post_number":3,"image_url":null,"created_at":"2015-12-14T13:33:03.426Z","last_posted_at":"2015-12-17T18:40:05.487Z","bumped":true,"bumped_at":"2015-12-17T18:40:05.487Z","unseen":false,"pinned":false,"unpinned":null,"visible":true,"closed":false,"archived":false,"bookmarked":null,"liked":null,"views":97,"like_count":0,"has_summary":false,"archetype":"regular","last_poster_username":"MarkLaFay","category_id":1,"pinned_globally":false,"posters":[{"extras":"latest","description":"OriginalPoster,MostRecentPoster","user_id":71},{"extras":null,"description":"FrequentPoster","user_id":14}]},{"id":140,"title":"GoogleAnalyticsplugin","fancy_title":"GoogleAnalyticsplugin","slug":"google-analytics-plugin","posts_count":1,"reply_count":0,"highest_post_number":1,"image_url":null,"created_at":"2015-12-15T13:00:55.644Z","last_posted_at":"2015-12-15T13:00:55.705Z","bumped":true,"bumped_at":"2015-12-15T13:00:55.705Z","unseen":false,"pinned":false,"unpinned":null,"visible":true,"closed":false,"archived":false,"bookmarked":null,"liked":null,"views":105,"like_count":0,"has_summary":false,"archetype":"regular","last_poster_username":"fimp","category_id":1,"pinned_globally":false,"posters":[{"extras":"latestsingle","description":"OriginalPoster,MostRecentPoster","user_id":73}]},{"id":138,"title":"With-mongo-connectionfailed:badconnectiondetails:","fancy_title":"With-mongo-connectionfailed:badconnectiondetails:","slug":"with-mongo-connection-failed-bad-connection-details","posts_count":1,"reply_count":0,"highest_post_number":1,"image_url":null,"created_at":"2015-12-14T17:28:11.041Z","last_posted_at":"2015-12-14T17:28:11.111Z","bumped":true,"bumped_at":"2015-12-14T17:28:11.111Z","unseen":false,"pinned":false,"unpinned":null,"visible":true,"closed":false,"archived":false,"bookmarked":null,"liked":null,"views":56,"like_count":0,"has_summary":false,"archetype":"regular","last_poster_username":"MarkLaFay","category_id":1,"pinned_globally":false,"posters":[{"extras":"latestsingle","description":"OriginalPoster,MostRecentPoster","user_id":71}]},{"id":133,"title":"\"Wecouldn'tunderstandyourquestion.\"whenIquerymongoDB","fancy_title":"&ldquo;Wecouldn&rsquo;tunderstandyourquestion.&rdquo;whenIquerymongoDB","slug":"we-couldnt-understand-your-question-when-i-query-mongodb","posts_count":3,"reply_count":0,"highest_post_number":3,"image_url":null,"created_at":"2015-12-11T17:38:30.576Z","last_posted_at":"2015-12-14T13:31:26.395Z","bumped":true,"bumped_at":"2015-12-14T13:31:26.395Z","unseen":false,"pinned":false,"unpinned":null,"visible":true,"closed":false,"archived":false,"bookmarked":null,"liked":null,"views":107,"like_count":0,"has_summary":false,"archetype":"regular","last_poster_username":"MarkLaFay","category_id":1,"pinned_globally":false,"posters":[{"extras":"latest","description":"OriginalPoster,MostRecentPoster","user_id":71},{"extras":null,"description":"FrequentPoster","user_id":72}]},{"id":129,"title":"Mybarchartsareallthin","fancy_title":"Mybarchartsareallthin","slug":"my-bar-charts-are-all-thin","posts_count":4,"reply_count":1,"highest_post_number":4,"image_url":"/uploads/default/original/1X/41bcf3b2a00dc7cfaff01cb3165d35d32a85bf1d.png","created_at":"2015-12-09T22:09:56.394Z","last_posted_at":"2015-12-11T19:00:45.289Z","bumped":true,"bumped_at":"2015-12-11T19:00:45.289Z","unseen":false,"pinned":false,"unpinned":null,"visible":true,"closed":false,"archived":false,"bookmarked":null,"liked":null,"views":116,"like_count":0,"has_summary":false,"archetype":"regular","last_poster_username":"mhjb","category_id":1,"pinned_globally":false,"posters":[{"extras":"latest","description":"OriginalPoster,MostRecentPoster","user_id":53},{"extras":null,"description":"FrequentPoster","user_id":1}]},{"id":106,"title":"WhatistheexpectedreturnorderofcolumnsforgraphingresultswhenusingrawSQL?","fancy_title":"WhatistheexpectedreturnorderofcolumnsforgraphingresultswhenusingrawSQL?","slug":"what-is-the-expected-return-order-of-columns-for-graphing-results-when-using-raw-sql","posts_count":3,"reply_count":0,"highest_post_number":3,"image_url":null,"created_at":"2015-11-24T19:07:14.561Z","last_posted_at":"2015-12-11T17:04:14.149Z","bumped":true,"bumped_at":"2015-12-11T17:04:14.149Z","unseen":false,"pinned":false,"unpinned":null,"visible":true,"closed":false,"archived":false,"bookmarked":null,"liked":null,"views":153,"like_count":0,"has_summary":false,"archetype":"regular","last_poster_username":"jbwiv","category_id":1,"pinned_globally":false,"posters":[{"extras":"latest","description":"OriginalPoster,MostRecentPoster","user_id":58},{"extras":null,"description":"FrequentPoster","user_id":14}]},{"id":131,"title":"Setsiteurlfromadminpanel","fancy_title":"Setsiteurlfromadminpanel","slug":"set-site-url-from-admin-panel","posts_count":2,"reply_count":0,"highest_post_number":2,"image_url":null,"created_at":"2015-12-10T06:22:46.042Z","last_posted_at":"2015-12-10T19:12:57.449Z","bumped":true,"bumped_at":"2015-12-10T19:12:57.449Z","unseen":false,"pinned":false,"unpinned":null,"visible":true,"closed":false,"archived":false,"bookmarked":null,"liked":null,"views":77,"like_count":0,"has_summary":false,"archetype":"regular","last_poster_username":"sameer","category_id":1,"pinned_globally":false,"posters":[{"extras":null,"description":"OriginalPoster","user_id":70},{"extras":"latest","description":"MostRecentPoster","user_id":1}]},{"id":127,"title":"Internationalization(i18n)","fancy_title":"Internationalization(i18n)","slug":"internationalization-i18n","posts_count":2,"reply_count":0,"highest_post_number":2,"image_url":null,"created_at":"2015-12-08T16:55:37.397Z","last_posted_at":"2015-12-09T16:49:55.816Z","bumped":true,"bumped_at":"2015-12-09T16:49:55.816Z","unseen":false,"pinned":false,"unpinned":null,"visible":true,"closed":false,"archived":false,"bookmarked":null,"liked":null,"views":85,"like_count":0,"has_summary":false,"archetype":"regular","last_poster_username":"agilliland","category_id":1,"pinned_globally":false,"posters":[{"extras":null,"description":"OriginalPoster","user_id":69},{"extras":"latest","description":"MostRecentPoster","user_id":14}]},{"id":109,"title":"ReturningrawdatawithnofiltersalwaysreturnsWecouldn'tunderstandyourquestion","fancy_title":"ReturningrawdatawithnofiltersalwaysreturnsWecouldn&rsquo;tunderstandyourquestion","slug":"returning-raw-data-with-no-filters-always-returns-we-couldnt-understand-your-question","posts_count":3,"reply_count":1,"highest_post_number":3,"image_url":null,"created_at":"2015-11-25T21:35:01.315Z","last_posted_at":"2015-12-09T10:26:12.255Z","bumped":true,"bumped_at":"2015-12-09T10:26:12.255Z","unseen":false,"pinned":false,"unpinned":null,"visible":true,"closed":false,"archived":false,"bookmarked":null,"liked":null,"views":133,"like_count":0,"has_summary":false,"archetype":"regular","last_poster_username":"bencarter78","category_id":1,"pinned_globally":false,"posters":[{"extras":"latest","description":"OriginalPoster,MostRecentPoster","user_id":60},{"extras":null,"description":"FrequentPoster","user_id":14}]},{"id":103,"title":"SupportforCassandra?","fancy_title":"SupportforCassandra?","slug":"support-for-cassandra","posts_count":5,"reply_count":1,"highest_post_number":5,"image_url":null,"created_at":"2015-11-20T06:45:31.741Z","last_posted_at":"2015-12-09T03:18:51.274Z","bumped":true,"bumped_at":"2015-12-09T03:18:51.274Z","unseen":false,"pinned":false,"unpinned":null,"visible":true,"closed":false,"archived":false,"bookmarked":null,"liked":null,"views":169,"like_count":0,"has_summary":false,"archetype":"regular","last_poster_username":"vikram","category_id":1,"pinned_globally":false,"posters":[{"extras":"latest","description":"OriginalPoster,MostRecentPoster","user_id":55},{"extras":null,"description":"FrequentPoster","user_id":1}]},{"id":128,"title":"MongoquerywithDatebreaks[solved:Mongo3.0required]","fancy_title":"MongoquerywithDatebreaks[solved:Mongo3.0required]","slug":"mongo-query-with-date-breaks-solved-mongo-3-0-required","posts_count":5,"reply_count":0,"highest_post_number":5,"image_url":null,"created_at":"2015-12-08T18:30:56.562Z","last_posted_at":"2015-12-08T21:03:02.421Z","bumped":true,"bumped_at":"2015-12-08T21:03:02.421Z","unseen":false,"pinned":false,"unpinned":null,"visible":true,"closed":false,"archived":false,"bookmarked":null,"liked":null,"views":102,"like_count":1,"has_summary":false,"archetype":"regular","last_poster_username":"edchan77","category_id":1,"pinned_globally":false,"posters":[{"extras":"latest","description":"OriginalPoster,MostRecentPoster","user_id":68},{"extras":null,"description":"FrequentPoster","user_id":1}]},{"id":23,"title":"CanthisconnecttoMSSQLServer?","fancy_title":"CanthisconnecttoMSSQLServer?","slug":"can-this-connect-to-ms-sql-server","posts_count":7,"reply_count":1,"highest_post_number":7,"image_url":null,"created_at":"2015-10-21T18:52:37.987Z","last_posted_at":"2015-12-07T17:41:51.609Z","bumped":true,"bumped_at":"2015-12-07T17:41:51.609Z","unseen":false,"pinned":false,"unpinned":null,"visible":true,"closed":false,"archived":false,"bookmarked":null,"liked":null,"views":367,"like_count":0,"has_summary":false,"archetype":"regular","last_poster_username":"sameer","category_id":1,"pinned_globally":false,"posters":[{"extras":null,"description":"OriginalPoster","user_id":9},{"extras":null,"description":"FrequentPoster","user_id":23},{"extras":null,"description":"FrequentPoster","user_id":3},{"extras":null,"description":"FrequentPoster","user_id":50},{"extras":"latest","description":"MostRecentPoster","user_id":1}]},{"id":121,"title":"Cannotrestartmetabaseindocker","fancy_title":"Cannotrestartmetabaseindocker","slug":"cannot-restart-metabase-in-docker","posts_count":5,"reply_count":1,"highest_post_number":5,"image_url":null,"created_at":"2015-12-04T21:28:58.137Z","last_posted_at":"2015-12-04T23:02:00.488Z","bumped":true,"bumped_at":"2015-12-04T23:02:00.488Z","unseen":false,"pinned":false,"unpinned":null,"visible":true,"closed":false,"archived":false,"bookmarked":null,"liked":null,"views":96,"like_count":0,"has_summary":false,"archetype":"regular","last_poster_username":"sameer","category_id":1,"pinned_globally":false,"posters":[{"extras":null,"description":"OriginalPoster","user_id":66},{"extras":"latest","description":"MostRecentPoster,FrequentPoster","user_id":1}]},{"id":85,"title":"EditMaxRowsCount","fancy_title":"EditMaxRowsCount","slug":"edit-max-rows-count","posts_count":4,"reply_count":2,"highest_post_number":4,"image_url":null,"created_at":"2015-11-11T23:46:52.917Z","last_posted_at":"2015-11-24T01:01:14.569Z","bumped":true,"bumped_at":"2015-11-24T01:01:14.569Z","unseen":false,"pinned":false,"unpinned":null,"visible":true,"closed":false,"archived":false,"bookmarked":null,"liked":null,"views":169,"like_count":1,"has_summary":false,"archetype":"regular","last_poster_username":"sameer","category_id":1,"pinned_globally":false,"posters":[{"extras":null,"description":"OriginalPoster","user_id":34},{"extras":"latest","description":"MostRecentPoster,FrequentPoster","user_id":1}]},{"id":96,"title":"Creatingchartsbyqueryingmorethanonetableatatime","fancy_title":"Creatingchartsbyqueryingmorethanonetableatatime","slug":"creating-charts-by-querying-more-than-one-table-at-a-time","posts_count":6,"reply_count":4,"highest_post_number":6,"image_url":null,"created_at":"2015-11-17T11:20:18.442Z","last_posted_at":"2015-11-21T02:12:25.995Z","bumped":true,"bumped_at":"2015-11-21T02:12:25.995Z","unseen":false,"pinned":false,"unpinned":null,"visible":true,"closed":false,"archived":false,"bookmarked":null,"liked":null,"views":217,"like_count":0,"has_summary":false,"archetype":"regular","last_poster_username":"sameer","category_id":1,"pinned_globally":false,"posters":[{"extras":null,"description":"OriginalPoster","user_id":51},{"extras":"latest","description":"MostRecentPoster,FrequentPoster","user_id":1}]},{"id":90,"title":"TryingtoaddRDSpostgresqlasthedatabasefailssilently","fancy_title":"TryingtoaddRDSpostgresqlasthedatabasefailssilently","slug":"trying-to-add-rds-postgresql-as-the-database-fails-silently","posts_count":4,"reply_count":2,"highest_post_number":4,"image_url":null,"created_at":"2015-11-14T23:45:02.967Z","last_posted_at":"2015-11-21T01:08:45.915Z","bumped":true,"bumped_at":"2015-11-21T01:08:45.915Z","unseen":false,"pinned":false,"unpinned":null,"visible":true,"closed":false,"archived":false,"bookmarked":null,"liked":null,"views":162,"like_count":0,"has_summary":false,"archetype":"regular","last_poster_username":"sameer","category_id":1,"pinned_globally":false,"posters":[{"extras":null,"description":"OriginalPoster","user_id":46},{"extras":"latest","description":"MostRecentPoster,FrequentPoster","user_id":1}]},{"id":17,"title":"DeploytoHerokuisn'tworking","fancy_title":"DeploytoHerokuisn&rsquo;tworking","slug":"deploy-to-heroku-isnt-working","posts_count":9,"reply_count":3,"highest_post_number":9,"image_url":null,"created_at":"2015-10-21T16:42:03.096Z","last_posted_at":"2015-11-20T18:34:14.044Z","bumped":true,"bumped_at":"2015-11-20T18:34:14.044Z","unseen":false,"pinned":false,"unpinned":null,"visible":true,"closed":false,"archived":false,"bookmarked":null,"liked":null,"views":332,"like_count":2,"has_summary":false,"archetype":"regular","last_poster_username":"agilliland","category_id":1,"pinned_globally":false,"posters":[{"extras":null,"description":"OriginalPoster","user_id":5},{"extras":null,"description":"FrequentPoster","user_id":3},{"extras":null,"description":"FrequentPoster","user_id":11},{"extras":null,"description":"FrequentPoster","user_id":25},{"extras":"latest","description":"MostRecentPoster","user_id":14}]},{"id":100,"title":"CanIuseDATEPART()inSQLqueries?","fancy_title":"CanIuseDATEPART()inSQLqueries?","slug":"can-i-use-datepart-in-sql-queries","posts_count":2,"reply_count":0,"highest_post_number":2,"image_url":null,"created_at":"2015-11-17T23:15:58.033Z","last_posted_at":"2015-11-18T00:19:48.763Z","bumped":true,"bumped_at":"2015-11-18T00:19:48.763Z","unseen":false,"pinned":false,"unpinned":null,"visible":true,"closed":false,"archived":false,"bookmarked":null,"liked":null,"views":112,"like_count":1,"has_summary":false,"archetype":"regular","last_poster_username":"sameer","category_id":1,"pinned_globally":false,"posters":[{"extras":null,"description":"OriginalPoster","user_id":53},{"extras":"latest","description":"MostRecentPoster","user_id":1}]},{"id":98,"title":"FeatureRequest:LDAPAuthentication","fancy_title":"FeatureRequest:LDAPAuthentication","slug":"feature-request-ldap-authentication","posts_count":1,"reply_count":0,"highest_post_number":1,"image_url":null,"created_at":"2015-11-17T17:22:44.484Z","last_posted_at":"2015-11-17T17:22:44.577Z","bumped":true,"bumped_at":"2015-11-17T17:22:44.577Z","unseen":false,"pinned":false,"unpinned":null,"visible":true,"closed":false,"archived":false,"bookmarked":null,"liked":null,"views":97,"like_count":0,"has_summary":false,"archetype":"regular","last_poster_username":"j_at_svg","category_id":1,"pinned_globally":false,"posters":[{"extras":"latestsingle","description":"OriginalPoster,MostRecentPoster","user_id":52}]},{"id":87,"title":"MigratingfrominternalH2toPostgres","fancy_title":"MigratingfrominternalH2toPostgres","slug":"migrating-from-internal-h2-to-postgres","posts_count":2,"reply_count":0,"highest_post_number":2,"image_url":null,"created_at":"2015-11-12T14:36:06.745Z","last_posted_at":"2015-11-12T18:05:10.796Z","bumped":true,"bumped_at":"2015-11-12T18:05:10.796Z","unseen":false,"pinned":false,"unpinned":null,"visible":true,"closed":false,"archived":false,"bookmarked":null,"liked":null,"views":111,"like_count":0,"has_summary":false,"archetype":"regular","last_poster_username":"sameer","category_id":1,"pinned_globally":false,"posters":[{"extras":null,"description":"OriginalPoster","user_id":42},{"extras":"latest","description":"MostRecentPoster","user_id":1}]}]},"person":{"person":{"id":"d50887ca-a6ce-4e59-b89f-14f0b5d03b03","name":{"fullName":"LeonidBugaev","givenName":"Leonid","familyName":"Bugaev"},"email":"leonsbox@gmail.com","gender":"male","location":"SaintPetersburg,SaintPetersburg,RU","geo":{"city":"SaintPetersburg","state":"SaintPetersburg","country":"Russia","lat":59.9342802,"lng":30.3350986},"bio":"SeniorengineeratGranify.com","site":"http://flickfaver.com","avatar":"https://d1ts43dypk8bqh.cloudfront.net/v1/avatars/d50887ca-a6ce-4e59-b89f-14f0b5d03b03","employment":{"name":"www.latera.ru"},"facebook":{"handle":"leonid.bugaev"},"github":{"handle":"buger","id":14009,"avatar":"https://avatars.githubusercontent.com/u/14009?v=3","company":"Granify","blog":"http://leonsbox.com","followers":95,"following":10},"twitter":{"handle":"flickfaver","id":77004410,"bio":null,"followers":2,"site":"http://flickfaver.com","avatar":null},"linkedin":{"handle":"in/leonidbugaev"},"googleplus":{"handle":null},"angellist":{"handle":"leonid-bugaev","id":61541,"bio":"SeniorengineeratGranify.com","followers":41,"avatar":"https://d1qb2nb5cznatu.cloudfront.net/users/61541-medium_jpg?1405474390"},"klout":{"handle":null,"score":null},"foursquare":{"handle":null},"aboutme":{"handle":"leonid.bugaev","bio":null,"avatar":null},"gravatar":{"handle":"buger","urls":[],"avatar":"http://1.gravatar.com/avatar/f7c8edd577d13b8930d5522f281235","avatars":[{"url":"http://1.gravatar.com/avatar/f7c8edd577d13b8930d5522f28123510","type":"thumbnail"}]},"fuzzy":false},"company":null},"record":{"event":"end","egress_info":{"egress_id":"EG_QmbJnbv8oZfN","room_id":"RM_7CxvNZvCeVaj","room_name":"943ec52f6812481a97224ae8da51de88","status":3,"started_at":1711753940605022425,"ended_at":1711753986735608003,"updated_at":1711753986735608003,"Request":{"Track":{"room_name":"943ec52f6812481a97224ae8da51de88","track_id":"TR_VCZLi","Output":{"File":{"filepath":"/recording/DA_0023920255-TR_VCZLi.mp4","Output":null}}}},"Result":{"File":{"filename":"/recording/DA_0023920255-TR_VCZLi.mp4","started_at":1711753940600335077,"duration":-1711753940600335077,"size":7920417,"location":"/recording/DA23920-TR_VCZLi.mp4"}},"file_results":[{"filename":"/recording/DA_0023920255-TR_VCZLi.mp4","started_at":1711753940600335077,"duration":-1711753940600335077,"size":7920417,"location":"/recording/DA_0023920255-TR_VCZLi.mp4"}]},"id":"EV_DkfNbvQwT2Jp","created_at":1711753986}}`
)

var (
	smallStruct  SmallStruct
	mediumStruct MediumStruct
	bigStruct    BigStruct
)

// benchmarkMarshal small struct
func benchmarkMarshalsmall(b *testing.B, m JsonLib) {
	b.ResetTimer()
	b.Run("Marshal-small", func(tb *testing.B) {
		for i := 0; i < tb.N; i++ {
			m.Marshal(smallStruct)
		}
	})
}

// benchmarkMarshal small struct
func benchmarkMarshalmedium(b *testing.B, m JsonLib) {
	b.ResetTimer()
	b.Run("Marshal-medium", func(tb *testing.B) {
		for i := 0; i < tb.N; i++ {
			m.Marshal(mediumStruct)
		}
	})
}

// benchmarkMarshal small struct
func benchmarkMarshalbig(b *testing.B, m JsonLib) {
	b.ResetTimer()
	b.Run("Marshal-big", func(tb *testing.B) {
		for i := 0; i < tb.N; i++ {
			m.Marshal(bigStruct)
		}
	})
}

// benchmarkUnmarshal
func benchmarkUnMarshalsmall(b *testing.B, m JsonLib) {
	runtime.GC()
	b.ResetTimer()

	b.Run("UnMarshal-small", func(tb *testing.B) {
		for i := 0; i < tb.N; i++ {
			m.UnMarshal([]byte(smallbytes), &smallStruct)
		}
	})
}

func benchmarkUnMarshalmedium(b *testing.B, m JsonLib) {
	runtime.GC()
	b.ResetTimer()

	b.Run("UnMarshal-medium", func(tb *testing.B) {
		for i := 0; i < tb.N; i++ {
			m.UnMarshal([]byte(mediumbytes), &mediumStruct)
		}
	})
}

func benchmarkUnMarshalbig(b *testing.B, m JsonLib) {
	runtime.GC()
	b.ResetTimer()

	b.Run("UnMarshal-big", func(tb *testing.B) {
		for i := 0; i < tb.N; i++ {
			m.UnMarshal([]byte(bigbytes), &bigStruct)
		}
	})
}

func benchmark(b *testing.B, m JsonLib) {
	b.ResetTimer()
	benchmarkUnMarshalsmall(b, m)
	benchmarkUnMarshalmedium(b, m)
	benchmarkUnMarshalbig(b, m)
	benchmarkMarshalsmall(b, m)
	benchmarkMarshalmedium(b, m)
	benchmarkMarshalbig(b, m)
}

func BenchmarkOrigin_json(b *testing.B) {
	benchmark(b, NewOriginLibrary())
}

func BenchmarkBytedance_sonic(b *testing.B) {
	benchmark(b, NewSonicLibrary())
}

func BenchmarkJson_iter(b *testing.B) {
	benchmark(b, NewJsoniterLibrary())
}

func BenchmarkPquerna_ffjson(b *testing.B) {
	benchmark(b, NewPquernaLibrary())
}

// func BenchmarkCosmWasm_tinyjson(b *testing.B) {
// 	benchmark(b, NewCosmWasmLibrary())
// }

//for easyjson

func BenchmarkMailru_easyjson(b *testing.B) {
	// benchmark(b, NewMailruLibrary())
	b.ResetTimer()
	benchmarkUnMarshal_M_e_small(b)
	benchmarkUnMarshal_M_e_medium(b)
	benchmarkUnMarshal_M_e_big(b)
	benchmarkMarshal_M_e_small(b)
	benchmarkMarshal_M_e_medium(b)
	benchmarkMarshal_M_e_big(b)
}

// benchmarkMarshal small struct
func benchmarkMarshal_M_e_small(b *testing.B) {
	b.ResetTimer()
	b.Run("Marshal-small", func(tb *testing.B) {
		for i := 0; i < tb.N; i++ {
			easyjson.Marshal(smallStruct)
		}
	})
}

// benchmarkMarshal small struct
func benchmarkMarshal_M_e_medium(b *testing.B) {
	b.ResetTimer()
	b.Run("Marshal-medium", func(tb *testing.B) {
		for i := 0; i < tb.N; i++ {
			easyjson.Marshal(mediumStruct)
		}
	})
}

// benchmarkMarshal small struct
func benchmarkMarshal_M_e_big(b *testing.B) {
	b.ResetTimer()
	b.Run("Marshal-big", func(tb *testing.B) {
		for i := 0; i < tb.N; i++ {
			easyjson.Marshal(bigStruct)
		}
	})
}

// benchmarkUnmarshal
func benchmarkUnMarshal_M_e_small(b *testing.B) {
	runtime.GC()
	b.ResetTimer()

	b.Run("UnMarshal-small", func(tb *testing.B) {
		for i := 0; i < tb.N; i++ {
			easyjson.Unmarshal([]byte(smallbytes), &smallStruct)
		}
	})
}

func benchmarkUnMarshal_M_e_medium(b *testing.B) {
	runtime.GC()
	b.ResetTimer()

	b.Run("UnMarshal-medium", func(tb *testing.B) {
		for i := 0; i < tb.N; i++ {
			easyjson.Unmarshal([]byte(mediumbytes), &mediumStruct)
		}
	})
}

func benchmarkUnMarshal_M_e_big(b *testing.B) {
	runtime.GC()
	b.ResetTimer()

	b.Run("UnMarshal-big", func(tb *testing.B) {
		for i := 0; i < tb.N; i++ {
			easyjson.Unmarshal([]byte(bigbytes), &bigStruct)
		}
	})
}