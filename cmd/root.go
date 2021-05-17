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
- Pavel Tykač: #prohleda Tweety s timto retezcem
  - "(Tenhle pan |Tenhle |Tento pan |Tento |Pan |)(P\\. Tykač|Tykač) je (opravdový|skutečný|neskutečný|neuvěřitelný|prolhaný) (devastátor|eko-terorista|ničitel (krajiny|naší země))\\." #a odpovi temito texty
  - "(Tenhle pan |Tenhle |Tento pan |Tento |Pan |)(P\\. Tykač|Tykač) (naprosto |)(cynicky|sobecky|bezohledně) (devastuje|ruinuje) (českou krajinu|naši zemi|naši budoucnost|budoucnost našich dětí)\\."
  - "(Tenhle pan |Tenhle |Tento pan |Tento |Pan |)(P\\. Tykač|Tykač) (smrdí|páchne|zavání) do oken (celé ČR|celé České republice|celému Česku)\\. Je to (hanba|odporné|nechutné|sobecké|tupé|sobec|cynik|zlý člověk)(\\.|\\.\\.\\.)"
  - "(Tenhle pan |Tenhle |Tento pan |Tento |Pan |)(P\\. Tykač|Tykač) by měl zkusit (čuchat|dýchat|dejchat|inhalovat) ten (jeho (jedovatej | toxickej |)smrad|(jedovatej | toxickej |)smrad jeho byznysu|(jedovatej |toxickej |)smrad z jeho elektráren)\\."
  - "(Tenhle pan |Tenhle |Tento pan |Tento |Pan |)(P\\. Tykač|Tykač) jede (svůj byznys |)(doslova |)přes mrtvoly\\. (Měl by se stydět|Měl by přestat|Kéž by přestal|Mohl by přestat|Měl by s tím skončit|(Páchne|Smrdí) nám (všem |)(přímo |)do oken)\\."
  - "(Tykačovi|Pavlu Tykačovi) (nevěřím ani slovo|nic nevěřím|nevěřte|nevěřte ani slovo)\\. Je to (zákeřný|neskutečný|rafinovaný) (manipulátor|cynik|sobec|lhář), (co|který|kterej) (lakuje věci|překrucuje věci) (nazeleno|narůžovo)\\."
- Daniel Křetínský: #prohleda Tweety s timto retezcem
  - "(Tenhle pan |Tenhle |Tento pan |Tento |Pan |)(Dan Křetínský|D\\. Křetínský|Křetínský) je (opravdový|skutečný|neskutečný|neuvěřitelný|prolhaný) (devastátor|eko-terorista|ničitel (krajiny|naší země))\\." #a odpovi temito texty
  - "(Tenhle pan |Tenhle |Tento pan |Tento |Pan |)(Dan Křetínský|D\\. Křetínský|Křetínský) (naprosto |)(cynicky|sobecky|bezohledně) (devastuje|ruinuje) (českou krajinu|naši zemi|naši budoucnost|budoucnost našich dětí)\\."
  - "(Tenhle pan |Tenhle |Tento pan |Tento |Pan |)(Dan Křetínský|D\\. Křetínský|Křetínský) (smrdí|páchne|zavání) do oken (celé ČR|celé České republice|celému Česku)\\. Je to (hanba|odporné|nechutné|sobecké|tupé|sobec|cynik|zlý člověk)(\\.|\\.\\.\\.)"
  - "(Tenhle pan |Tenhle |Tento pan |Tento |Pan |)(Dan Křetínský|D\\. Křetínský|Křetínský) by měl zkusit (čuchat|dýchat|dejchat|inhalovat) ten (jeho (jedovatej | toxickej |)smrad|(jedovatej | toxickej |)smrad jeho byznysu|(jedovatej |toxickej |)smrad z jeho elektráren)\\."
  - "(Tenhle pan |Tenhle |Tento pan |Tento |Pan |)(Dan Křetínský|D\\. Křetínský|Křetínský) jede (svůj byznys |)(doslova |)přes mrtvoly\\. (Měl by se stydět|Měl by přestat|Kéž by přestal|Mohl by přestat|Měl by s tím skončit|(Páchne|Smrdí) nám (všem |)(přímo |)do oken)\\."
  - "(Křetínskému|Křetínskýmu) (nevěřím ani slovo|nic nevěřím|nevěřte|nevěřte ani slovo)\\. Je to (zákeřný|neskutečný|rafinovaný) (manipulátor|cynik|sobec|lhář), (co|který|kterej) (lakuje věci|překrucuje věci) (nazeleno|narůžovo)\\."
`)

	fmt.Printf(">>> Created default config at: %v !\n\n", cfgPath)
}
