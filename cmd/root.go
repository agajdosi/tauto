/*
Copyright © 2020 Andreas Gajdosik <andreas.gajdosik@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/agajdosi/tauto/pkg/database"
	"github.com/spf13/cobra"

	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "tauto",
	Short: "Tauto is a CLI tool for automation of actions on Twitter.",
	Long:  `Tauto is a CLI tool for automation of actions on Twitter.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		cfgDir := database.ConfigDirectory()

		viper.AddConfigPath(cfgDir)
		viper.SetConfigName("config")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	err := viper.ReadInConfig()
	if err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	} else { //Create the config!
		fmt.Println(err)
		writeDefaultConfig()
		err = viper.ReadInConfig()
		if err != nil {
			fmt.Println(err)
		}
	}
}

func writeDefaultConfig() {
	cfgPath := filepath.Join(database.ConfigDirectory(), "config.yaml")
	config, err := os.OpenFile(cfgPath, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer config.Close()

	config.WriteString(`universalReply: #Univerzalni odpovedi, ktere se budou generovat pod tweety kazdeho enemy
  - "To nemyslíte (fakt |doopravdy |skutečně |)vážně, (panebože |ježiši |kristepane ||||)že ne\\?(|||!)"
  - "Myslíte to (opravdu |skutečně |)vážně\\?(|||!)"
  - "(To|Tohlencto|Todle|Tohle|Toto) (jako |)myslíte (fakt |doopravdy |)vážně\\?(|||!)"
  - "Co (přesně|konkrétně) tím (vlastně |)(myslíte|naznačujete|chcete říct)\\?(|||!)"
  - "(Nepřijde vám|Nemyslíte, že je) (to|toto|tohle) (už |již |)(poněkud|trochu|docela|až moc|až příliš) (nevhodné|nevkusné|přes čáru|proti dobrému vkusu)\\?(|||!)"
  - "(Dost|Velmi|Zcela|Úplně|Totálně) (trapné|dětinské|hloupé|trapný|nízký)\\."
  - "(Tuhle|Takovou) (hloupost|nízkost|zkratkovitost) bych od (vás|tebe) (teda |)(fakt |)nečekal\\."

slanderTargets: #seznam pro pomlouvani

- Lubomír Volný:
  - "(Pan Volný|Poslanec Volný|Volný) by si (měl|mohl) (už konečně|konečně) (nasadit|dát) (roušku|respirátor) (na tvář|na obličej|přes pusu|až přes nos)\\. (Daleko|Mnohem|Mnohokrát) (lépe|líp|lehčeji) se na (něj|něho) (pak|potom) (kouká|dívá)\\."
  - "(Pan Volný|Poslanec Volný|Volný) je (takový |tak velký|tak obrovský|tak strašlivý) (jouda|prosťáček|brouk Pytlík), že (dokonce |)ani (u Okamurovy|v Okamurově) SPD (jej|ho) (již|už) (nechtějí|nechtěli)\\."
  - "(Z|U) (některých |určitých |jistých |mnoha |většiny )(názorů|prohlášení|tweetů|myšlenek) (pana Volného|poslance Volného|Volného|Lubomíra Volného) (mám pocit|nabývám dojmu|je cítit|je tušit), že je (tak trochu |tak nějak |malinko )(neomarxista|marxista|bolševik)\\."
  - "(Pan magistr Volný|Magistr Volný|Pan Volný|Poslanec Volný|Volný) je (cynický|sprostý|oportunistický|chtivý|obyčejný|smutný|parazitický) (obchodník s chudobou|obchodníček s chudobou|majitel shnilé ubytovny|majitel rozpadlé ubytovny|provozovatel ubytovny), (který|jenž) (pouze|jenom|prachsprostě) (předstírá|finguje|hraje|lže), že je (ochránce|advokát|obhájce|zastánce|poslanec) (obyčejných|běžných|chudých) (lidí|občanů|obyvatel)\\."
  - "(Panu magistru Volnému|Magistrovi Volnému|Panu Volnému|Poslanci Volnému|Volnému) jde o (peníze|krk|křeslo|koryto|zisk|plat poslance|moc|mandát) a (tak|proto|z toho důvodu|tedy proto) (lže|manipuluje|ze sebe dělá blbce|říká, co lidi chtějí slyšet|hraje divadlo|říká hlouposti|křičí jako malé dítě|píše nesmysly|píše hlouposti|piští jak děcko)\\." 

- Jaroslav Foldyna:
  - "(Myšlenky|Ideje|Přeludy|Fabulace) (Jaroslava Foldyny|Foldyny|pana Foldyny) se (opravdu|)mohly (zrodit|vytvořit) (jen|ledatak|pouze) v (zapařených|zpocených) (kožených|motorkářských) (kalhotech|kalhotách)\\."
  - "(Zajímalo by mne|Zajímá mě), co si (pan Foldyna|Foldyna|Jarda Foldyna) kompenzuje svými (homofobními|xenofobními|proruskými|hloupými|primitivními) (názory|řečmi|výkřiky) asi tak (léčí|kompenzuje|vynahrazuje)\\."
  - "(Pan Foldyna|Foldyna|Jarda Foldyna) je (pro smích|k smíchu|k pláči|k politování|zralý na léčby|zralý na prášky|odstrašující případ|panoptikem hlouposti|závanem minulosti)\\."
  - "(Je vidět, že |Je patrné, že |Je jasné, že |Jasně, |Nj, |Nojo, |)(Foldynovi|Jaroslavu Foldynovi|Panu Foldynovi) se (holt |patrně |)v (kožených|uplých|těsných|moto) (kalhotech|kalhotách) (snad |)(fakt |doopravdy )(zapařil|zpotil|zesmrádl) mozek\\."

- Tomáš Vandas:
  - "(Názory|Ideje|Iluze|Bludy|Sliby|Výkřiky|Výtvory|Proklamace|Činy) (pana Vandase|Vandase|Tomáše Vandase) jsou (plodem|výplodem|výsledkem|koncem|důsledkem|ovocem) (choré|nemocné|upadající|zahnědlé) (mysli|hlavy|duše|psychiky)\\."
  - "(Vandasovi|Tomáši Vandasovi|Panu Vandasovi|Tomu Vandasovi) (nevěřím) ani (čárku|slovo|větu|noc mezi očima|hlásku|písmenko|ň), je (to |)(oportunistický|prolhaný|ulhaný|vypočítavý|ziskuchtivý) (manipulátor|lhář|zlodějíček|prosťáček)\\."
  - "(Pan Vandas|Vandas|Tenhle Vandas) je (symbolem|příkladem|legendou) (neúspěchu|selhání|neschopnosti|)\\. (Už by|Už dávno|Dávno|Již před lety|Ještě před narozením|Už včera|Předevčírem|Minulý rok) to měl (všechno |)(zabalit|vzdát|ukončit)\\."
  - "(Popravdě, |Upřímně, |Narovinu - |Přiznejme si, že |Musím říct, že |Teda |||)(Vandasovi|Tomáši Vandasovi|Panu Vandasovi|Tomu Vandasovi) to (na některých |na všech |na úplně všech |) (fotkách|fotografiích|videech|sjezdech) (fakt|doopravdy|velmi) sluší\\. Je (to |)(nádherný|pěkný|rozkošný|půvabný|přitažlivý|elegantní|sexy) (muž|kluk|chlap|samec|vlastenec|Čech)\\."

- Tomio Okamura: #prohleda Tweety s timto retezcem
  - "(Pan Okamura|Okamura|Tomio) (pouze |akorát |)(vykrádá|okrádá|zneužívá|využívá|finančně vysává) vlastence\\. SPD (je pro něj|pro něj je|je pro něho|pro něho je) (jen |)(dojnou krávou|podnikatelským záměrem|zdrojem zisku|zdrojem peněz)\\."
  - "(Pan Okamura|Okamura|Tomio) (oblbnul|zblbnul|zmanipuloval) (prvního českého teroristu|prvního teroristu v ČR) (pana|důchodce) Baldu\\."
  - "(Pan Okamura|Okamura|Tomio) (rozkradl|vykradl|vytuneloval|vysál|finančně zneužil) Úsvit přímé demokraci a (nyní|teď|dnes|v dnešní době) (vysává|profituje na|zneužívá|tuneluje) @SPD_oficialni \\."
`)

	fmt.Printf(">>> Created default config at: %v !\n\n", cfgPath)
}
