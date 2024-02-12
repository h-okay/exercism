// Enter your code below the lines of the families' information

// Secret knowledge of the Zhang family:
namespace zhang
{
    int bank_number_part(int secret_modifier)
    {
        int zhang_part{8'541};
        return (zhang_part * secret_modifier) % 10000;
    }
    namespace red
    {
        int code_fragment() { return 512; }
    }
    namespace blue
    {
        int code_fragment() { return 677; }
    }
}

// Secret knowledge of the Khan family:
namespace khan
{
    int bank_number_part(int secret_modifier)
    {
        int khan_part{4'142};
        return (khan_part * secret_modifier) % 10000;
    }
    namespace red
    {
        int code_fragment() { return 148; }
    }
    namespace blue
    {
        int code_fragment() { return 875; }
    }
}

// Secret knowledge of the Garcia family:
namespace garcia
{
    int bank_number_part(int secret_modifier)
    {
        int garcia_part{4'023};
        return (garcia_part * secret_modifier) % 10000;
    }
    namespace red
    {
        int code_fragment() { return 118; }
    }
    namespace blue
    {
        int code_fragment() { return 923; }
    }
}

// Enter your code below
namespace estate_executor
{
    int assemble_account_number(int secret_modifier)
    {
        int zhang_bank = zhang::bank_number_part(secret_modifier);
        int khan_bank = khan::bank_number_part(secret_modifier);
        int garcia_bank = garcia::bank_number_part(secret_modifier);
        return zhang_bank + khan_bank + garcia_bank;
    }

    int assemble_code() {
        int zhang_red_frag = zhang::red::code_fragment();
        int zhang_blue_frag = zhang::blue::code_fragment();
        int khan_red_frag = khan::red::code_fragment();
        int khan_blue_frag = khan::blue::code_fragment();
        int garcia_red_frag = garcia::red::code_fragment();
        int garcia_blue_frag = garcia::blue::code_fragment();
        int blue_frag = zhang_blue_frag + khan_blue_frag + garcia_blue_frag;
        int red_frag = zhang_red_frag + khan_red_frag + garcia_red_frag;
        return blue_frag * red_frag;
    }
}
