public class Resources {

    public void introDisplay(){
        System.out.println("""
                ______                                           ___            _\s
                |  _  \\                                         / _ \\          | |
                | | | |_   _ _ __   __ _  ___  ___  _ __  ___  / /_\\ \\_ __   __| |
                | | | | | | | '_ \\ / _` |/ _ \\/ _ \\| '_ \\/ __| |  _  | '_ \\ / _` |
                | |/ /| |_| | | | | (_| |  __/ (_) | | | \\__ \\ | | | | | | | (_| |
                |___/  \\__,_|_| |_|\\__, |\\___|\\___/|_| |_|___/ \\_| |_/_| |_|\\__,_|
                                    __/ |                                        \s
                                   |___/                                         \s
                   ___                   ______                                  \s
                  |_  |                  |  _  \\                                 \s
                    | | __ ___   ____ _  | | | |_ __ __ _  __ _  ___  _ __  ___  \s
                    | |/ _` \\ \\ / / _` | | | | | '__/ _` |/ _` |/ _ \\| '_ \\/ __| \s
                /\\__/ / (_| |\\ V / (_| | | |/ /| | | (_| | (_| | (_) | | | \\__ \\ \s
                \\____/ \\__,_| \\_/ \\__,_| |___/ |_|  \\__,_|\\__, |\\___/|_| |_|___/ \s
                                                           __/ |                 \s
                                                          |___/                  \s""");

    }

    public void clearScreen(){
        for(int i = 0; i < 20; i++){
            System.out.println();
        }
    }

    public void barbDisplay(){
        System.out.println("""
                      /___\\                                                \s
                     (|0 0|)                                                   \s
                   __/{\\U/}\\_ ___/vvv                                               \s
                  / \\  {~}   / _|_P|                                                \s
                  | /\\  ~   /_/   ||                                                \s
                  |_| (____)      ||                      \s
                  \\_]/______\\  /\\_||_/\\\s
                     _\\_||_/_ |] _||_ [|           \s
                     (_,_||_,_) \\/ [] \\/""");
    }

    public void dragonDisplay(){
        System.out.println("""
                                                             __----~~~~~~~~~~~------___
                                                  .  .   ~~//====......          __--~ ~~
                                  -.            \\_|//     |||\\\\  ~~~~~~::::... /~
                               ___-==_       _-~o~  \\/    |||  \\\\            _/~~-
                       __---~~~.==~||\\=_    -_--~/_-~|-   |\\\\   \\\\        _/~
                   _-~~     .=~    |  \\\\-_    '-~7  /-   /  ||    \\      /
                 .~       .~       |   \\\\ -_    /  /-   /   ||      \\   /
                /  ____  /         |     \\\\ ~-_/  /|- _/   .||       \\ /
                |~~    ~~|--~~~~--_ \\     ~==-/   | \\~--===~~        .\\
                         '         ~-|      /|    |-~\\~~       __--~~
                                     |-~~-_/ |    |   ~\\_   _-~            /\\
                                          /  \\     \\__   \\/~                \\__
                                      _--~ _/ | .-~~____--~-/                  ~~==.
                                     ((->/~   '.|||' -_|    ~~-/ ,              . _||
                                                -_     ~\\      ~~---l__i__i__i--~~_/
                                                _-~-__   ~)  \\--______________--~~
                                              //.-~~~-~_--~- |-------~~~~~~~~
                                                     //.-~~~--\\
                                                    \s
                                                    \s""");
    }

    public void introText(){
        System.out.println("""
                Deep within the heart of the Blackstone Mountains lies a cavern wreathed in legend. 
                Few dare to approach, for it is said that an ancient dragon slumbers within, guarding a hoard of untold riches. 
                The townsfolk whisper of adventurers who have entered, never to return. 
                Now, fate calls upon a new hero — will they uncover the truth behind the myth, claim the treasure, or fall to the beast’s fiery wrath?
                
                """);
    }

    public void encounterSetup(){
        System.out.println("""
                With each step, the air grows colder, and the scent of damp earth fills your lungs. 
                The mouth of the cave looms before you, jagged like the maw of a beast. 
                A faint, warm breeze drifts outward—carrying the unmistakable scent of sulfur. 
                The ground beneath your boots is littered with scorched bones, a silent warning of what lurks within.""");
    }

    public void dragonEnters(){
        System.out.println("""
                As you step deeper into the cavern, the dim torchlight flickers against the towering walls of stone. 
                The air is thick with the scent of sulfur, and the distant sound of something... breathing. 
                The ground trembles beneath your feet. Then, from the darkness, two immense eyes ignite like burning embers.
                
                A low, rumbling growl fills the chamber, shaking loose dust and pebbles from the ceiling. 
                A massive shape unfurls from the shadows—scales as black as night, claws sharp as razors, 
                and wings that stretch wide, casting you in darkness.""");
    }

    public void victoryText(){
        System.out.println("""
                The cavern trembles as the dragon lets out a final, deafening roar. 
                Its massive body collapses, sending a cloud of dust and embers into the air. 
                For a moment, all is silent—then, the glow of its fiery eyes fades, leaving only the flickering light of molten gold pooling beneath its fallen form.
                
                The beast is slain.
                
                The weight of the battle still lingers in your limbs, but you stand victorious. 
                Before you lies the dragon’s legendary hoard—glistening treasures, ancient artifacts, and relics of forgotten ages. 
                The world will speak of your triumph for generations to come.""");
    }

    public void defeatedDragon(){
        System.out.println("""
                  ,{
                   '_}
                ;._'.             .-'
                (             \\-./ ;.-;
                 ; _      |'--,| `  '< ___,
                   __)     \\`-.__.   /`.-'/
                 {;         `/o(o \\ | / ,'
                   ;  _  __.-'-'`-'  \\'`
                   ' (-,`  .-.     _ -.          _.-.
                      /     _))  _/    |      \s""");
    }
}
