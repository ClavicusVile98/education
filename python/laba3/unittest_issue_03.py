import unittest

from typing import List, Tuple
from one_hot_encoder import fit_transform


class TestOneHotEncoder(unittest.TestCase):
    def test_(self):
        cities = ['Moscow', 'New York', 'Moscow', 'London']
        exp_transformed_cities = [
            ('Moscow', [0, 0, 1]),
            ('New York', [0, 1, 0]),
            ('Moscow', [0, 0, 1]),
            ('London', [1, 0, 0]),
        ]
        transformed_cities = fit_transform(cities)
        self.assertEqual(transformed_cities, exp_transformed_cities)

    def test_assert_NotIn(self):
        cities = ['Moscow', 'New York', 'Paris', 'London']
        exp_transformed_cities = [
            ('Moscow', [0, 0, 1]),
            ('New York', [0, 1, 0]),
            ('Ankara', [0, 0, 1]),
            ('London', [1, 0, 0]),
        ]
        transformed_cities = fit_transform(cities)
        self.assertNotIn(transformed_cities, exp_transformed_cities)

    def test_assert_exс(self):
        cities = ['Moscow', 'New York', 'Paris', 'London']
        with self.assertRaises(TypeError):
            fit_transform(cities, 2)

    def test_assertIsNotNone(self):
        exp_transformed_cities = [
        ('Moscow', [0, 0, 1]),
        ('New York', [0, 1, 0]),
        ('Moscow', [0, 0, 1]),
        ('London', [1, 0, 0]),
        ]
        self.assertIsNotNone(len(exp_transformed_cities))

if __name__ == '__main__':
    unittest.main()
