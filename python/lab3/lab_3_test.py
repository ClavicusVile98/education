import doctest

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


def encode(message: str) -> str:
    """
    Кодирует строку в соответсвие с таблицей азбуки Морзе
    >>> encode('SOS')
    '... --- ...'
    >>> encode('HIGHLOAD')
    '.... .. --. .... .-.. --- .- -..'
    >>> encode('HOW MANY TIMES I SLEPT?')
    '.... --- .-- -- .- -. -.-- - ..
    -- . ... .. ... .-.. . .--. - ..--..'
    >>> encode('Аннушка уже разлила масло...') #doctest: +ELLIPSIS
    '.-.-.- .-.-.-'
    """
    encoded_signs = [
        LETTER_TO_MORSE[letter] for letter in message
    ]

    return ' '.join(encoded_signs)


def decode(morse_message: str) -> str:
    """
    Декодирует строку из азбуки Морзе в английский
    >>> decode('.-.-.-_')
    '.'
    >>> decode('..... -....- --... -....- -.--. ...-- -..-. .---- -.--.-')
    '5-7-(3/1)'
    """
    decoded_letters = [
        MORSE_TO_LETTER[letter] for letter in morse_message.split()
    ]

    return ''.join(decoded_letters)


 # def test_assert():
    # actual = '.... .. --. .... .-.. --- .- -.. ..--..'
    # expected = '.... .. --. .... .-.. --- .- -..'
    # err_msg = f'{actual} != {expected}'
    # assert actual == expected, err_msg


if __name__ == '__main__':
    msg = 'MAI-PYTHON-2019'
    encoded_msg = '.... .. --. .... .-.. --- .- -.. ..--..'
    # print(decoded_msg)
    err_msg = f'{msg} != {encoded_msg}'

    assert msg == encode(encoded_msg), err_msg
    doctest.testmod()
