import pytest

"""Morse Code Translator"""

LETTER_TO_MORSE = {
    'A': '.-',
    'B': '-...',
    'C': '-.-.',
    'D': '-..',
    'E': '.',
    'F': '..-.',
    'G': '--.',
    'H': '....',
    'I': '..',
    'J': '.---',
    'K': '-.-',
    'L': '.-..',
    'M': '--',
    'N': '-.',
    'O': '---',
    'P': '.--.',
    'Q': '--.-',
    'R': '.-.',
    'S': '...',
    'T': '-',
    'U': '..-',
    'V': '...-',
    'W': '.--',
    'X': '-..-',
    'Y': '-.--',
    'Z': '--..',
    '1': '.----',
    '2': '..---',
    '3': '...--',
    '4': '....-',
    '5': '.....',
    '6': '-....',
    '7': '--...',
    '8': '---..',
    '9': '----.',
    '0': '-----',
    ', ': '--..--',
    '.': '.-.-.-',
    '?': '..--..',
    '/': '-..-.',
    '-': '-....-',
    '(': '-.--.',
    ')': '-.--.-',
    ' ': ' '
}

MORSE_TO_LETTER = {
    morse: letter
    for letter, morse in LETTER_TO_MORSE.items()
}


def decode(morse_message: str) -> str:
    """
    Декодирует строку из азбуки Морзе в английский
    """
    decoded_letters = [
        MORSE_TO_LETTER[letter] for letter in morse_message.split()
    ]
    return ''.join(decoded_letters)


@pytest.mark.parametrize('msg, exp',[
    ('-- .- .. -....- .--. -.-- - .... --- -. -....- ..--- ----- .---- ----.', 'MAI-PYTHON-2019'),
    ('.... .. --. .... .-.. --- .- -..', 'HIGHLOAD?'),
    ('.--. --- -.- . -- --- -.', 'POKEMON')
])
def test_decode(msg, exp):
    assert decode(msg) == exp


#if __name__ == '__main__':
# msg = '-- .- .. -....- .--. -.-- - .... --- -. -....- ..--- ----- .---- ----.'
# decoded_msg = decode(msg)
# print(decoded_msg)