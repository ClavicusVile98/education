import pytest
from one_hot_encoder import fit_transform


#def test():
#    with pytest.raises(TypeError):
#        fit_transform()


def test_assert():
    exp = [
        ('banana', [0, 0, 1]),
        ('coconut', [0, 1, 0]),
        ('lemon', [1, 0, 0])
    ]
    assert fit_transform('banana', 'coconut', 'lemon') == exp


def test_assert_2():
    exp = [
        ('banana', [0, 1, 1]),
        ('coconut', [0, 1, 0]),
        ('lemon', [1, 0, 0])
    ]
    assert fit_transform('banana', 'coconut', 'lemon') == exp


@pytest.mark.parametrize('seasons, exp_seasons',[
    ('Autumn', 'Autumn[0, 0, 0]'),
    ('Winter', 'Winter[0, 0, 1]'),
    ('Spring', 'Spring[0, 1, 0]'),
    ('Summer', 'Summer[1, 0, 0]')
])
def test_param(seasons, exp_seasons):
    assert fit_transform() == exp_seasons


if __name__ == '__main__':
    pytest.main()