# rollpwd - dynamic password generator

This project generates a time (hour) dependent password for a given username and site combination. The use case is access to an embedded system which may have no or sporadic network connectivity and thus unable to synchronize the time with any standard time source. The intent is that the password generation is controlled and the generated password is shared at the time of usage with the user.

Practical applications include a field team visiting a customer site; the password generation is available through a website, accessible by the field team; the embedded device has this utility authenticate it.

## Challenges

The variation in the microprocessor clocks usually manifests as drifting time of day clocks. Due to such a drift, a time dependent password will fail under some border conditions. For a low to medium security requirement, the verification can be tolerant to some drift by generating passwords for 3 scenarios - the clock has drifted ahead, behind or not. The supplied password is accepted if it matches one of these 3. When the drift is < 1 hour then, there are only 2 unique passwords.

## Usage

The utility can be used to generate a password for a username for the specified host. The utility can also be used to authenticate the user on the host. 

        ../bin/rollpwd --help
        Usage of ../bin/rollpwd:
            --authenticate string   authenticate this password
            --verbose               high verbosity

## Examples

        $ date
        Wed Nov 27 20:36:22 EST 2019
        $ ../bin/rollpwd rs@toprllc.com
        2019/11/27 20:36:40 6c9f049218bd63f40610badeb396eb50

        $ date
        Wed Nov 27 20:37:32 EST 2019
        $ ../bin/rollpwd rs@toprllc.com --authenticate 6c9f049218bd63f40610badeb396eb50
        2019/11/27 20:37:55 Authenticated

        $ ../bin/rollpwd rs@toprllc.com --authenticate 7d9f049218bd63f40610badeb396eb50
        2019/11/27 20:39:50 Not authenticated
