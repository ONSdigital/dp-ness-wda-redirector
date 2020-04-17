package main

var visualRedirects = map[string]string{
	"how-do-the-post-world-war-baby-boom-generations-compare":    "https://www.ons.gov.uk/peoplepopulationandcommunity/birthsdeathsandmarriages/ageing/articles/howdothepostworldwarbabyboomgenerationscompare/2018-03-06",
	"whats-in-the-basket-of-goods-70-years-of-shopping-history":  "https://www.ons.gov.uk/economy/inflationandpriceindices/articles/whatsinthebasketofgoods70yearsofshoppinghistory/2016-07-21",
	"what-affects-likelihood-of-smoking":                         "https://www.ons.gov.uk/peoplepopulationandcommunity/healthandsocialcare/drugusealcoholandsmoking/articles/likelihoodofsmokingfourtimeshigherinenglandsmostdeprivedareasthanleastdeprived/2018-03-14",
	"what-do-children-in-the-uk-spend-their-money-on":            "https://www.ons.gov.uk/peoplepopulationandcommunity/personalandhouseholdfinances/expenditure/articles/whatdochildrenintheukspendtheirmoneyon/2018-02-15",
	"men-enjoy-five-hours-more-leisure-time-per-week-than-women": "https://www.ons.gov.uk/peoplepopulationandcommunity/wellbeing/articles/menenjoyfivehoursmoreleisuretimeperweekthanwomen/2018-01-09",
	"uk-trade-partners": "https://www.ons.gov.uk/businessindustryandtrade/internationaltrade/articles/whodoestheuktradewith/2017-02-21",
	"young-people-spend-a-third-of-their-leisure-time-on-devices": "https://www.ons.gov.uk/peoplepopulationandcommunity/leisureandtourism/articles/youngpeoplespendathirdoftheirleisuretimeondevices/2017-12-19",
	"paddington-star-wars-and-the-rise-of-the-uk-film-industry":   "https://www.ons.gov.uk/economy/grossdomesticproductgdp/articles/paddingtonstarwarsandtheriseoftheukfilmindustry/2017-12-14",
	"what-is-my-life-expectancy-and-how-might-it-change":          "https://www.ons.gov.uk/peoplepopulationandcommunity/healthandsocialcare/healthandlifeexpectancies/articles/whatismylifeexpectancyandhowmightitchange/2017-12-01",
	"migration-since-the-brexit-vote-whats-changed-in-six-charts": "https://www.ons.gov.uk/peoplepopulationandcommunity/populationandmigration/internationalmigration/articles/migrationsincethebrexitvotewhatschangedinsixcharts/2017-11-30",
	"gender-pay-gap-changes-twenty-years":                         "https://www.ons.gov.uk/employmentandlabourmarket/peopleinwork/earningsandworkinghours/articles/londonhadthesmallestgenderpaygap20yearsagobutnowithasthelargest/2017-11-27",
	"interactive-how-well-does-my-job-pay":                        "https://www.ons.gov.uk/employmentandlabourmarket/peopleinwork/earningsandworkinghours/articles/howwelldoesmyjobpay/2015-01-15",
	"interactive-how-do-earnings-vary-across-the-country":         "https://www.ons.gov.uk/employmentandlabourmarket/peopleinwork/earningsandworkinghours/articles/howdoearningsvaryacrossgreatbritain/2015-01-15",
	"uk-perspectives-the-changing-population":                     "https://www.ons.gov.uk/peoplepopulationandcommunity/populationandmigration/migrationwithintheuk/articles/thechangingukpopulation/2015-01-15",
	"infographic-what-is-your-religion":                           "https://www.ons.gov.uk/peoplepopulationandcommunity/populationandmigration/migrationwithintheuk/articles/whatisyourreligion/2015-01-15",
	"uk-perspectives-a-recent-history-of-international-migration": "https://www.ons.gov.uk/peoplepopulationandcommunity/populationandmigration/internationalmigration/articles/internationalmigrationarecenthistory/2015-01-15",
	"uk-perspectives-housing-and-home-ownership-in-the-uk":        "https://www.ons.gov.uk/peoplepopulationandcommunity/populationandmigration/internationalmigration/articles/housingandhomeownershipintheuk/2015-01-22",
	"uk-perspectives-public-services-in-the-uk":                   "https://www.ons.gov.uk/economy/governmentpublicsectorandtaxes/publicspending/articles/spendingonpublicservicesintheuk/2015-02-10",
	"uk-perspectives-personal-and-household-finances-in-the-uk":   "https://www.ons.gov.uk/peoplepopulationandcommunity/personalandhouseholdfinances/incomeandwealth/articles/personalandhouseholdfinancesintheuk/2015-02-12",
	"how-we-travel-in-the-uk":                                     "https://www.ons.gov.uk/peoplepopulationandcommunity/leisureandtourism/articles/howwetravel/2015-02-19",
	"uk-perspectives-energy-and-emissions":                        "https://www.ons.gov.uk/economy/environmentalaccounts/articles/energyandemissionsintheuk/2015-02-19",
	"uk-perspectives-trends-in-the-uk-economy":                    "https://www.ons.gov.uk/economy/economicoutputandproductivity/productivitymeasures/articles/trendsintheukeconomy/2015-02-27",
	"uk-perspectives-an-overview-of-the-uk-labour-market":         "https://www.ons.gov.uk/employmentandlabourmarket/peopleinwork/articles/anoverviewoftheuklabourmarket/2015-02-27",
	"6-facts-about-pension-membership":                            "https://www.ons.gov.uk/peoplepopulationandcommunity/personalandhouseholdfinances/pensionssavingsandinvestments/articles/6factsaboutpensionmembership/2015-03-05",
	"how-does-your-family-size-compare":                           "https://www.ons.gov.uk/peoplepopulationandcommunity/birthsdeathsandmarriages/families/articles/howdoesyourfamilysizecompare/2015-03-13",
	"visualising-your-constituency":                               "https://www.ons.gov.uk/peoplepopulationandcommunity/housing/articles/visualisingyourconstituency/2015-03-26",
	"how-long-will-my-pension-need-to-last":                       "https://www.ons.gov.uk/peoplepopulationandcommunity/birthsdeathsandmarriages/lifeexpectancies/articles/howlongwillmypensionneedtolast/2015-03-27",
	"victory-in-europe-day-how-world-war-ii-changed-the-uk":       "https://www.ons.gov.uk/peoplepopulationandcommunity/birthsdeathsandmarriages/articles/victoryineuropedayhowworldwariichangedtheuk/2015-05-08",
	"travel-trends":        "https://www.ons.gov.uk/peoplepopulationandcommunity/populationandmigration/internationalmigration/articles/whovisitstheuktraveltrendsfrom20042014/2015-05-20",
	"2011-census-religion": "https://www.ons.gov.uk/peoplepopulationandcommunity/culturalidentity/religion/articles/howreligionhaschangedinenglandandwales/2015-06-04",
	"older-people-census":  "https://www.ons.gov.uk/peoplepopulationandcommunity/healthandsocialcare/healthcaresystem/articles/carehomepopulationstabilisesasunpaidcarerpopulationincreases/2015-06-08",
	"binge-drinking":       "https://www.ons.gov.uk/peoplepopulationandcommunity/healthandsocialcare/drugusealcoholandsmoking/articles/howmuchdopeoplebingedrinkingreatbritain/2015-06-12",
	"housing-census":       "https://www.ons.gov.uk/peoplepopulationandcommunity/housing/articles/homeownershipdownandrentingupforfirsttimeinacentury/2015-06-19",
	"a-decade-of-population-change-in-england-and-wales": "https://www.ons.gov.uk/peoplepopulationandcommunity/populationandmigration/populationestimates/articles/adecadeofpopulationchangeinenglandandwales/2015-06-25",
	"eu-budget":                    "https://www.ons.gov.uk/economy/governmentpublicsectorandtaxes/publicspending/articles/howdoestheukcontributetotheeubudget/2015-06-26",
	"ethnicity-2011-census":        "https://www.ons.gov.uk/peoplepopulationandcommunity/culturalidentity/ethnicity/articles/peopleidentifyingasotherwhitehasincreasedbyoveramillionsince2001/2015-06-26",
	"health-census":                "https://www.ons.gov.uk/peoplepopulationandcommunity/healthandsocialcare/healthandwellbeing/articles/8in10peoplehavegoodorverygoodhealthinenglandandwales/2015-07-02",
	"in-work-poverty":              "https://www.ons.gov.uk/employmentandlabourmarket/peopleinwork/employmentandemployeetypes/articles/doesgettingajobalwaysleadtopeopleleavingpoverty/2015-07-06",
	"productivity-puzzle":          "https://www.ons.gov.uk/employmentandlabourmarket/peopleinwork/labourproductivity/articles/whatistheproductivitypuzzle/2015-07-07",
	"language-census-2011":         "https://www.ons.gov.uk/peoplepopulationandcommunity/culturalidentity/language/articles/peoplewhocannotspeakenglishwellaremorelikelytobeinpoorhealth/2015-07-09",
	"managing-money":               "https://www.ons.gov.uk/peoplepopulationandcommunity/personalandhouseholdfinances/incomeandwealth/articles/whoisbestatmanagingmoney/2015-07-10",
	"disability-census":            "https://www.ons.gov.uk/peoplepopulationandcommunity/healthandsocialcare/disability/articles/nearlyoneinfivepeoplehadsomeformofdisabilityinenglandandwales/2015-07-13",
	"birthsanddeaths":              "https://www.ons.gov.uk/peoplepopulationandcommunity/birthsdeathsandmarriages/livebirths/articles/trendsinbirthsanddeathsoverthelastcentury/2015-07-15",
	"40-years-of-cancer":           "https://www.ons.gov.uk/peoplepopulationandcommunity/healthandsocialcare/conditionsanddiseases/articles/survivalfromcancerimprovingandmorepeoplebeingdiagnosed/2015-07-17",
	"families-census-2011":         "https://www.ons.gov.uk/peoplepopulationandcommunity/populationandmigration/populationestimates/articles/increaseinsinglepopulationoverthelastdecade/2015-07-24",
	"migration-census":             "https://www.ons.gov.uk/peoplepopulationandcommunity/populationandmigration/internationalmigration/articles/1in7peopleinenglandandwalesin2011werebornoutsideoftheuk/2015-07-30",
	"affordability-housing":        "https://www.ons.gov.uk/employmentandlabourmarket/peopleinwork/earningsandworkinghours/articles/houseprices24timesaveragesalaryinwestminster/2015-08-05",
	"social-housing-affordability": "https://www.ons.gov.uk/peoplepopulationandcommunity/housing/articles/socialhousingbecamelessaffordableoverpastdecade/2015-08-05",
	"work-travel-census":           "https://www.ons.gov.uk/employmentandlabourmarket/peopleinwork/earningsandworkinghours/articles/transitionfromamanufacturingtoserviceledlabourmarketoverpast170years/2015-08-06",
	"baby-names":                   "https://www.ons.gov.uk/peoplepopulationandcommunity/birthsdeathsandmarriages/livebirths/articles/10popcultureinfluencesonbabynamesgameofthronesmarvelfrozenandmore/2015-08-17",
	"top-10-baby-names-how-their-popularity-has-changed-over-the-last-decade": "https://www.ons.gov.uk/peoplepopulationandcommunity/birthsdeathsandmarriages/livebirths/articles/top10babynameshowtheirpopularityhaschangedoverthelastdecade/2015-08-17",
	"what-are-migration-levels-like-in-your-area":                             "https://www.ons.gov.uk/peoplepopulationandcommunity/populationandmigration/internationalmigration/articles/whataremigrationlevelslikeinyourarea/2015-08-28",
	"how-important-is-china-to-the-economy":                                   "https://www.ons.gov.uk/businessindustryandtrade/internationaltrade/articles/howimportantischinatotheukeconomy/2015-09-02",
	"deaths-involving-heroin-up-by-two-thirds-in-two-years":                   "https://www.ons.gov.uk/peoplepopulationandcommunity/birthsdeathsandmarriages/deaths/articles/deathsinvolvingheroinupbytwothirdsintwoyears/2015-09-03",
	"how-has-life-expectancy-changed-over-time":                               "https://www.ons.gov.uk/peoplepopulationandcommunity/birthsdeathsandmarriages/lifeexpectancies/articles/howhaslifeexpectancychangedovertime/2015-09-09",
	"the-history-of-strikes-in-britain":                                       "https://www.ons.gov.uk/employmentandlabourmarket/peopleinwork/employmentandemployeetypes/articles/thehistoryofstrikesintheuk/2015-09-21",
	"nine-things-you-might-not-know-about-older-people-in-the-uk":             "https://www.ons.gov.uk/peoplepopulationandcommunity/birthsdeathsandmarriages/ageing/articles/ninethingsyoumightnotknowaboutolderpeopleintheuk/2015-10-01",
	"peopleonthemove": "https://www.ons.gov.uk/peoplepopulationandcommunity/populationandmigration/migrationwithintheuk/articles/peopleonthemoveinenglandandwales/2015-10-08",
	"how-employee-numbers-have-changed-in-the-north-since-2009":                                                          "https://www.ons.gov.uk/employmentandlabourmarket/peopleinwork/employmentandemployeetypes/articles/howemployeenumbershavechangedinthenorthsince2009/2015-10-09",
	"how-many-jobs-are-paid-less-than-the-living-wage-in-your-area":                                                      "https://www.ons.gov.uk/employmentandlabourmarket/peopleinwork/earningsandworkinghours/articles/howmanyjobsarepaidlessthanthelivingwageinyourarea/2015-10-12",
	"more-than-a-quarter-of-children-who-spend-longer-on-social-networking-websites-report-mental-ill-health-symptoms-2": "https://www.ons.gov.uk/peoplepopulationandcommunity/wellbeing/articles/morechildrenusingsocialmediareportmentalillhealthsymptoms/2015-10-20",
	"most-affluent-man-now-outlives-the-average-woman-for-the-first-time":                                                "https://www.ons.gov.uk/peoplepopulationandcommunity/birthsdeathsandmarriages/lifeexpectancies/articles/mostaffluentmanoutlivestheaveragewomanforthefirsttime/2015-10-21",
	"how-big-will-the-uk-population-be-in-25-years-time":                                                                 "https://www.ons.gov.uk/peoplepopulationandcommunity/populationandmigration/populationprojections/articles/howbigwilltheukpopulationbein25yearstime/2015-10-29",
	"fewer-gb-companies-exporting-abroad":                                                                                "https://www.ons.gov.uk/businessindustryandtrade/business/businessservices/articles/fewergbcompaniesexportingabroad/2015-11-12",
	"how-long-will-you-live-in-good-health":                                                                              "https://www.ons.gov.uk/peoplepopulationandcommunity/birthsdeathsandmarriages/lifeexpectancies/articles/howlongwillyouliveingoodhealth/2015-11-20",
	"excesswintermortality":                                                                                              "https://www.ons.gov.uk/peoplepopulationandcommunity/birthsdeathsandmarriages/deaths/articles/highestnumberofexcesswinterdeathssince19992000/2015-11-25",
	"more-jobs-being-paid-close-to-the-minimum-wage":                                                                     "https://www.ons.gov.uk/economy/nationalaccounts/uksectoraccounts/articles/morejobsbeingpaidclosetotheminimumwage/2015-12-01",
	"households-spend-the-most-on-transport":                                                                             "https://www.ons.gov.uk/peoplepopulationandcommunity/personalandhouseholdfinances/incomeandwealth/articles/householdsspendthemostontransport/2015-12-08",
	"how-popular-is-your-birthday":                                                                                       "https://www.ons.gov.uk/peoplepopulationandcommunity/birthsdeathsandmarriages/livebirths/articles/howpopularisyourbirthday/2015-12-18",
	"what-are-your-chances-of-living-to-100":                                                                             "https://www.ons.gov.uk/peoplepopulationandcommunity/birthsdeathsandmarriages/lifeexpectancies/articles/whatareyourchancesoflivingto100/2016-01-14",
	"the-british-steel-industry-since-the-1970s":                                                                         "https://www.ons.gov.uk/economy/economicoutputandproductivity/output/articles/updatedthebritishsteelindustrysincethe1970s/2016-01-18",
	"fuel-prices-explained-a-breakdown-of-the-cost-of-petrol-and-diesel":                                                 "https://www.ons.gov.uk/economy/inflationandpriceindices/articles/fuelpricesexplainedabreakdownofthecostofpetrolanddiesel/2016-01-22",
	"how-do-survival-estimates-compare-for-common-cancers":                                                               "https://www.ons.gov.uk/peoplepopulationandcommunity/healthandsocialcare/conditionsanddiseases/articles/howdosurvivalestimatescompareforcommoncancers/2016-02-03",
	"how-does-getting-older-change-the-way-we-feel-about-our-lives":                                                      "https://www.ons.gov.uk/peoplepopulationandcommunity/wellbeing/articles/howdoesgettingolderchangethewaywefeelaboutourlives/2016-02-03",
	"most-people-use-e-cigarettes-to-help-them-quit-smoking-cigarettes":                                                  "https://www.ons.gov.uk/peoplepopulationandcommunity/healthandsocialcare/drugusealcoholandsmoking/articles/mostpeopleuseecigarettestohelpthemquitsmoking/2016-02-18",
	"living-with-parents":                                                                                                "https://www.ons.gov.uk/peoplepopulationandcommunity/birthsdeathsandmarriages/families/articles/whyaremoreyoungpeoplelivingwiththeirparents/2016-02-22",
	"household-income-and-inequality-where-do-you-fit-in":                                                                "https://www.ons.gov.uk/peoplepopulationandcommunity/personalandhouseholdfinances/incomeandwealth/articles/householdincomeandinequalitywheredoyoufitin/2016-02-23",
	"increase-in-cancer-patients-surviving-a-year-or-more-after-diagnosis-2":                                             "https://www.ons.gov.uk/peoplepopulationandcommunity/healthandsocialcare/conditionsanddiseases/articles/inequalityinoneyearcancersurvivalinenglandshrinks/2016-02-26",
	"teenage-pregnancies-perception-versus-reality":                                                                      "https://www.ons.gov.uk/peoplepopulationandcommunity/birthsdeathsandmarriages/conceptionandfertilityrates/articles/teenagepregnanciesperceptionversusreality/2016-03-09",
	"40-years-of-smoking-in-great-britain":                                                                               "https://www.ons.gov.uk/peoplepopulationandcommunity/healthandsocialcare/drugusealcoholandsmoking/articles/40yearsofsmokingingreatbritain/2016-03-09",
	"two-decades-of-sunday-trading":                                                                                      "https://www.ons.gov.uk/businessindustryandtrade/retailindustry/articles/twodecadesofsundaytrading/2016-03-09",
	"welfare-spending":                                                                                                   "https://www.ons.gov.uk/economy/governmentpublicsectorandtaxes/publicsectorfinance/articles/howisthewelfarebudgetspent/2016-03-16",
	"the-debt-and-deficit-of-the-uk-public-sector-explained":                                                             "https://www.ons.gov.uk/economy/governmentpublicsectorandtaxes/publicsectorfinance/articles/thedebtanddeficitoftheukpublicsectorexplained/2016-03-16",
	"million-pound-properties":                                                                                           "https://www.ons.gov.uk/peoplepopulationandcommunity/housing/articles/1millionpropertysalesincreasedover70foldsince1995/2016-03-24",
	"how-will-the-national-living-wage-affect-employees-and-businesses-in-the-uk":                                        "https://www.ons.gov.uk/employmentandlabourmarket/peopleinwork/earningsandworkinghours/articles/howwillthenationallivingwageaffectemployeesandbusinessesintheuk/2016-04-01",
	"dementiaalzheimers-and-flu-behind-biggest-annual-increase-in-deaths-since-the-1960s":                                "https://www.ons.gov.uk/peoplepopulationandcommunity/birthsdeathsandmarriages/deaths/articles/dementiaandrespiratorydiseasebehindbiggestannualdeathsincreasesincethe1960s/2016-04-07",
	"deaths-from-legal-highs":                                                                                            "https://www.ons.gov.uk/peoplepopulationandcommunity/birthsdeathsandmarriages/deaths/articles/deathsfromlegalhighs/2016-04-28",
	"how-do-childhood-circumstances-affect-your-chances-of-poverty-as-an-adult":                                          "https://www.ons.gov.uk/peoplepopulationandcommunity/educationandchildcare/articles/howdochildhoodcircumstancesaffectyourchancesofpovertyasanadult/2016-05-16",
	"uk-perspectives-2016-trends-in-the-uk-economy":                                                                      "https://www.ons.gov.uk/economy/grossdomesticproductgdp/articles/ukperspectives2016trendsintheukeconomy/2016-05-25",
	"uk-perspectives-2016-trade-with-the-eu-and-beyond":                                                                  "https://www.ons.gov.uk/businessindustryandtrade/internationaltrade/articles/ukperspectives2016tradewiththeeuandbeyond/2016-05-25",
	"uk-perspectives-2016-spending-on-public-services-in-the-uk":                                                         "https://www.ons.gov.uk/economy/governmentpublicsectorandtaxes/publicsectorfinance/articles/ukperspectives2016spendingonpublicservicesintheuk/2016-05-25",
	"uk-perspectives-2016-personal-and-household-finances-in-the-uk":                                                     "https://www.ons.gov.uk/peoplepopulationandcommunity/housing/articles/ukperspectives2016personalandhouseholdfinancesintheuk/2016-05-25",
	"uk-perspectives-2016-housing-and-home-ownership-in-the-uk":                                                          "https://www.ons.gov.uk/peoplepopulationandcommunity/housing/articles/ukperspectives2016housingandhomeownershipintheuk/2016-05-25",
	"uk-perspectives-2016-the-uk-contribution-to-the-eu-budget":                                                          "https://www.ons.gov.uk/economy/grossdomesticproductgdp/articles/ukperspectives2016theukcontributiontotheeubudget/2016-05-25",
	"uk-perspectives-2016-energy-and-emissions-in-the-uk":                                                                "https://www.ons.gov.uk/economy/environmentalaccounts/articles/ukperspectives2016energyandemissionsintheuk/2016-05-26",
	"uk-perspectives-2016-the-uk-in-an-european-context":                                                                 "https://www.ons.gov.uk/peoplepopulationandcommunity/populationandmigration/populationestimates/articles/ukperspectives2016theukinaeuropeancontext/2016-05-26",
	"uk-perspectives-2016-an-overview-of-the-uk-labour-market":                                                           "https://www.ons.gov.uk/employmentandlabourmarket/peopleinwork/employmentandemployeetypes/articles/ukperspectives2016anoverviewoftheuklabourmarket/2016-05-26",
	"uk-perspectives-2016-how-we-travel":                                                                                 "https://www.ons.gov.uk/peoplepopulationandcommunity/leisureandtourism/articles/ukperspectives2016howwetravel/2016-05-26",
	"uk-perspectives-2016-the-changing-uk-population":                                                                    "https://www.ons.gov.uk/peoplepopulationandcommunity/populationandmigration/populationestimates/articles/ukperspectives2016thechangingukpopulation/2016-05-26",
	"uk-perspectives-2016-international-migration-to-and-from-the-uk":                                                    "https://www.ons.gov.uk/peoplepopulationandcommunity/populationandmigration/populationestimates/articles/ukperspectives2016internationalmigrationtoandfromtheuk/2016-05-26",
	"does-our-sex-affect-what-we-die-from":                                                                               "https://www.ons.gov.uk/peoplepopulationandcommunity/birthsdeathsandmarriages/deaths/articles/doesoursexaffectwhatwediefrom/2016-07-13",
	"whats-in-the-basket-of-goods-80-years-of-shopping-history":                                                          "https://www.ons.gov.uk/economy/inflationandpriceindices/articles/whatsinthebasketofgoods70yearsofshoppinghistory/2016-07-21",
	"shopping-in-shops-that-have-no-shops":                                                                               "https://www.ons.gov.uk/businessindustryandtrade/retailindustry/articles/shoppinginshopsthathavenoshops/2016-07-29",
	"five-facts-about-strikes":                                                                                           "https://www.ons.gov.uk/employmentandlabourmarket/peopleinwork/workplacedisputesandworkingconditions/articles/fivefactsaboutstrikes/2016-08-12",
	"uk-energy-how-much-what-type-and-where-from":                                                                        "https://www.ons.gov.uk/economy/environmentalaccounts/articles/ukenergyhowmuchwhattypeandwherefrom/2016-08-15",
	"five-facts-about-housing":                                                                                           "https://www.ons.gov.uk/peoplepopulationandcommunity/housing/articles/fivefactsabouthousing/2016-08-17",
	"no-money-no-medals":                                                                                                 "https://www.ons.gov.uk/economy/grossdomesticproductgdp/articles/nomoneynomedals/2016-08-22",
	"whats-the-best-time-for-a-wedding":                                                                                  "https://www.ons.gov.uk/peoplepopulationandcommunity/birthsdeathsandmarriages/marriagecohabitationandcivilpartnerships/articles/whatsthebesttimeforawedding/2016-08-26",
	"baby-names-since-1904-how-has-yours-performed":                                                                      "https://www.ons.gov.uk/peoplepopulationandcommunity/birthsdeathsandmarriages/livebirths/articles/babynamessince1904howhasyoursperformed/2016-09-02",
	"the-popularity-of-the-name-muhammadmohammedmohammad":                                                                "https://www.ons.gov.uk/peoplepopulationandcommunity/birthsdeathsandmarriages/livebirths/articles/thepopularityofthenamemuhammadmohammedmohammad/2016-09-02",
	"how-has-the-student-population-changed":                                                                             "https://www.ons.gov.uk/peoplepopulationandcommunity/birthsdeathsandmarriages/livebirths/articles/howhasthestudentpopulationchanged/2016-09-20",
	"five-facts-about-cars":                                                                                              "https://www.ons.gov.uk/economy/environmentalaccounts/articles/fivefactsaboutcars/2016-09-22",
	"five-facts-about-the-uk-service-sector":                                                                             "https://www.ons.gov.uk/economy/economicoutputandproductivity/output/articles/fivefactsabouttheukservicesector/2016-09-29",
	"five-facts-about-older-people-at-work":                                                                              "https://www.ons.gov.uk/employmentandlabourmarket/peopleinwork/employmentandemployeetypes/articles/fivefactsaboutolderpeopleatwork/2016-10-01",
	"breadwinners-in-their-20s-how-are-they-doing-compared-with-previous-generations":                                    "https://www.ons.gov.uk/peoplepopulationandcommunity/personalandhouseholdfinances/incomeandwealth/articles/breadwinnersintheir20showaretheydoingcomparedwithpreviousgenerations/2016-10-11",
	"the-gender-pay-gap-what-is-it-and-what-affects-it":                                                                  "https://www.ons.gov.uk/employmentandlabourmarket/peopleinwork/earningsandworkinghours/articles/thegenderpaygapwhatisitandwhataffectsit/2016-10-26",
	"why-has-the-value-of-the-pound-been-falling-and-what-could-this-mean-for-people-in-the-uk":                          "https://www.ons.gov.uk/economy/inflationandpriceindices/articles/whyhasthevalueofthepoundbeenfallingandwhatcouldthismeanforpeopleintheuk/2016-10-28",
	"how-does-uk-healthcare-spending-compare-internationally":                                                            "https://www.ons.gov.uk/peoplepopulationandcommunity/healthandsocialcare/healthcaresystem/articles/howdoesukhealthcarespendingcompareinternationally/2016-11-01",
	"the-value-of-your-unpaid-work":                                                                                      "https://www.ons.gov.uk/employmentandlabourmarket/peopleinwork/earningsandworkinghours/articles/womenshouldertheresponsibilityofunpaidwork/2016-11-10",
	"what-is-gdp":                                                                                                        "https://www.ons.gov.uk/economy/grossdomesticproductgdp/articles/whatisgdp/2016-11-21",
	"the-challenges-of-measuring-gdp-in-the-digital-borderless-world":                                                    "https://www.ons.gov.uk/economy/grossdomesticproductgdp/articles/thechallengesofmeasuringgdpinthedigitalborderlessworld/2016-11-22",
	"gdp-and-special-events-in-history":                                                                                  "https://www.ons.gov.uk/economy/grossdomesticproductgdp/articles/gdpandspecialeventsinhistory/2016-11-25",
	"explore-50-years-of-international-migration":                                                                        "https://www.ons.gov.uk/peoplepopulationandcommunity/populationandmigration/internationalmigration/articles/explore50yearsofinternationalmigrationtoandfromtheuk/2016-12-01",
	"find-out-the-gender-pay-gap-for-your-job":                                                                           "https://www.ons.gov.uk/employmentandlabourmarket/peopleinwork/earningsandworkinghours/articles/findoutthegenderpaygapforyourjob/2016-12-09",
	"test-your-knowledge-on-the-gender-pay-gap":                                                                          "https://www.ons.gov.uk/employmentandlabourmarket/peopleinwork/earningsandworkinghours/articles/testyourknowledgeonthegenderpaygap/2016-12-09",
	"london-household-spending-outstrips-the-rest-of-the-uk":                                                             "https://www.ons.gov.uk/peoplepopulationandcommunity/personalandhouseholdfinances/expenditure/articles/londonhouseholdspendingoutstripstherestoftheuk/2017-02-16",
	"prospective-homeowners-struggling-to-get-onto-the-property-ladder":                                                  "https://www.ons.gov.uk/peoplepopulationandcommunity/housing/articles/prospectivehomeownersstrugglingtogetontopropertyladder/2017-02-24",
	"commonwealth-trade-in-focus-as-uk-prepares-for-brexit":                                                              "https://www.ons.gov.uk/businessindustryandtrade/internationaltrade/articles/commonwealthtradeinfocusasukpreparesforbrexit/2017-03-09",
	"hipsters-gin-and-the-basket-of-goods":                                                                               "https://www.ons.gov.uk/economy/inflationandpriceindices/articles/hipstersginandthebasketofgoods/2017-03-14",
	"billion-pound-loss-in-volunteering-effort-in-the-last-3-years":                                                      "https://www.ons.gov.uk/employmentandlabourmarket/peopleinwork/earningsandworkinghours/articles/billionpoundlossinvolunteeringeffort/2017-03-16",
	"gdp-and-me": "https://www.ons.gov.uk/economy/grossdomesticproductgdp/articles/gdpandme/2017-03-20",
	"trading-places-uk-goods-trade-with-eu-partners":                                         "https://www.ons.gov.uk/businessindustryandtrade/internationaltrade/articles/tradingplacesukgoodstradewitheupartners/2017-03-28",
	"are-we-training-enough-people-to-become-programmers":                                    "https://www.ons.gov.uk/businessindustryandtrade/itandinternetindustry/articles/arewetrainingenoughpeopletobecomeprogrammers/2017-06-19",
	"are-your-wages-keeping-up-with-inflation":                                               "https://www.ons.gov.uk/economy/inflationandpriceindices/articles/areyourwageskeepingupwithinflation/2017-06-20",
	"whats-changed-in-the-year-since-the-brexit-vote":                                        "https://www.ons.gov.uk/economy/grossdomesticproductgdp/articles/whatschangedsincethebrexitvote/2017-06-23",
	"how-inflation-changes-how-much-your-wages-are-worth":                                    "https://www.ons.gov.uk/economy/inflationandpriceindices/articles/howinflationchangeshowmuchyourwagesareworth/2017-06-26",
	"what-affects-an-areas-healthy-life-expectancy":                                          "https://www.ons.gov.uk/peoplepopulationandcommunity/healthandsocialcare/healthandlifeexpectancies/articles/whataffectsanareashealthylifeexpectancy/2017-06-28",
	"lesbian-gay-and-bisexual-people-say-they-experience-a-lower-quality-of-life":            "https://www.ons.gov.uk/peoplepopulationandcommunity/culturalidentity/sexuality/articles/lesbiangayandbisexualpeoplesaytheyexperiencealowerqualityoflife/2017-07-05",
	"unpaid-carers-provide-social-care-worth-57-billion":                                     "https://www.ons.gov.uk/peoplepopulationandcommunity/healthandsocialcare/healthandlifeexpectancies/articles/unpaidcarersprovidesocialcareworth57billion/2017-07-10",
	"the-changing-price-of-everyday-goods-and-services":                                      "https://www.ons.gov.uk/economy/inflationandpriceindices/articles/thechangingpriceofeverydaygoodsandservices/2017-07-11",
	"what-impact-could-lowering-the-uk-voting-age-to-16-have-on-the-shape-of-the-electorate": "https://www.ons.gov.uk/peoplepopulationandcommunity/elections/electoralregistration/articles/whatimpactcouldloweringtheukvotingageto16haveontheshapeoftheelectorate/2017-07-14",
	"where-are-industry-eyes-on-brexit":                                                      "https://www.ons.gov.uk/businessindustryandtrade/internationaltrade/articles/whereareindustryeyesonbrexit/2017-07-17",
	"marriage-and-divorce-on-the-rise-at-65-and-over":                                        "https://www.ons.gov.uk/peoplepopulationandcommunity/birthsdeathsandmarriages/marriagecohabitationandcivilpartnerships/articles/marriageanddivorceontheriseat65andover/2017-07-18",
	"migration-the-european-union-and-work-how-much-do-you-really-know":                      "https://www.ons.gov.uk/peoplepopulationandcommunity/populationandmigration/internationalmigration/articles/migrationtheeuropeanunionandworkhowmuchdoyoureallyknow/2017-07-19",
	"shrinkflation-and-the-changing-cost-of-chocolate":                                       "https://www.ons.gov.uk/economy/inflationandpriceindices/articles/shrinkflationandthechangingcostofchocolate/2017-07-24",
	"holidays-in-the-1990s-and-now":                                                          "https://www.ons.gov.uk/peoplepopulationandcommunity/leisureandtourism/articles/holidaysinthe1990sandnow/2017-08-07",
	"private-pensions-and-retired-households":                                                "https://www.ons.gov.uk/employmentandlabourmarket/peopleinwork/workplacepensions/articles/howareprivatepensionsaffectingtheincomeofretiredhouseholds/2017-08-08",
	"are-we-ready-to-switch-to-electric-cars":                                                "https://www.ons.gov.uk/economy/environmentalaccounts/articles/arewereadytoswitchtoelectriccars/2017-08-16",
	"migration-levels-what-do-you-know-about-your-area":                                      "https://www.ons.gov.uk/peoplepopulationandcommunity/populationandmigration/internationalmigration/articles/migrationlevelswhatdoyouknowaboutyourarea/2017-08-24",
	"pensioners-in-the-eu-and-uk":                                                            "https://www.ons.gov.uk/economy/investmentspensionsandtrusts/articles/pensionersintheeuanduk/2017-09-05",
	"people-greatly-overestimate-their-likelihood-of-being-robbed":                           "https://www.ons.gov.uk/peoplepopulationandcommunity/crimeandjustice/articles/peoplegreatlyoverestimatetheirlikelihoodofbeingrobbed/2017-09-07",
	"who-is-most-at-risk-of-suicide":                                                         "https://www.ons.gov.uk/peoplepopulationandcommunity/birthsdeathsandmarriages/deaths/articles/whoismostatriskofsuicide/2017-09-07",
	"causes-of-death-over-100-years":                                                         "https://www.ons.gov.uk/peoplepopulationandcommunity/birthsdeathsandmarriages/deaths/articles/causesofdeathover100years/2017-09-18",
	"more-mothers-with-young-children-working-full-time":                                     "https://www.ons.gov.uk/employmentandlabourmarket/peopleinwork/employmentandemployeetypes/articles/moremotherswithyoungchildrenworkingfulltime/2017-09-26",
	"people-who-were-abused-as-children-are-more-likely-to-be-abused-as-an-adult":            "https://www.ons.gov.uk/peoplepopulationandcommunity/crimeandjustice/articles/peoplewhowereabusedaschildrenaremorelikelytobeabusedasanadult/2017-09-27",
	"how-do-the-jobs-men-and-women-do-affect-the-gender-pay-gap":                             "https://www.ons.gov.uk/employmentandlabourmarket/peopleinwork/earningsandworkinghours/articles/howdothejobsmenandwomendoaffectthegenderpaygap/2017-10-06",
	"will-an-extension-increase-the-value-of-my-house":                                       "https://www.ons.gov.uk/peoplepopulationandcommunity/housing/articles/willanextensionincreasethevalueofmyhouse/2017-10-11",
	"house-prices-how-much-does-one-square-metre-cost-in-your-area":                          "https://www.ons.gov.uk/peoplepopulationandcommunity/housing/articles/housepriceshowmuchdoesonesquaremetrecostinyourarea/2017-10-11",
	"uk-drops-in-european-child-mortality-rankings":                                          "https://www.ons.gov.uk/peoplepopulationandcommunity/healthandsocialcare/childhealth/articles/ukdropsineuropeanchildmortalityrankings/2017-10-13",
	"60-years-of-change-bbc-today":                                                           "https://www.ons.gov.uk/peoplepopulationandcommunity/healthandsocialcare/healthandlifeexpectancies/articles/youdrawthecharts60yearsofchange/2017-10-24",
	"explore-the-gender-pay-gap-and-test-your-knowledge":                                     "https://www.ons.gov.uk/employmentandlabourmarket/peopleinwork/earningsandworkinghours/articles/explorethegenderpaygapandtestyourknowledge/2017-10-26",
	"the-uk-contribution-to-the-eu-budget":                                                   "https://www.ons.gov.uk/economy/governmentpublicsectorandtaxes/publicsectorfinance/articles/theukcontributiontotheeubudget/2017-10-31",
	"deprivation-by-leading-cause-of-death":                                                  "https://www.ons.gov.uk/peoplepopulationandcommunity/birthsdeathsandmarriages/deaths/articles/howdoesdeprivationvarybyleadingcauseofdeath/2017-11-01",
	"uk-interest-rate-rise-whats-changed-in-the-last-decade":                                 "https://www.ons.gov.uk/economy/grossdomesticproductgdp/articles/ukinterestraterisewhatschangedinthelastdecade/2017-11-02",
	"is-pay-higher-in-the-public-or-private-sector":                                          "https://www.ons.gov.uk/employmentandlabourmarket/peopleinwork/earningsandworkinghours/articles/ispayhigherinthepublicorprivatesector/2017-11-16",
}
