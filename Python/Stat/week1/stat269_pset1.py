# Name:Worada Sajai
# StudentID:650510679

"""
*************** IMPORTANT ***************
Do NOT rename this file.
Do NOT rename the play_game function or alter its parameters in any way.

Hints:
- Notice the strict inequality in the looping condition
  (i.e. stop when S > 100, and S > 200)
- The function np.random.randint(low, high) is INCLUSIVE of low
  and EXCLUSIVE of high. Hence, we should have low=1 and high=101.
- Player 2 wins if and only if y > x, not when y >= x.
- Player 2 resumes adding from Player 1's sum.
  Player 2 does NOT start over at 0.
"""

# Do NOT add any import statements
# import numpy as np

# # Instructions:
# # Fill in the following function.
# # We will be testing your code by setting several random seeds.
# # DO NOT ALTER THIS FUNCTION OUTSIDE THE BEGIN/END CODE MARKERS.
# # You may add other helper functions if you wish.
# # Do not add or remove parameters to this function.
# def play_game(seed: int = 37, ntrials: int = 100000) -> float:
#     """
#     Plays a game described in the Pset 1 handout ntrials times 
#     with a predetermined seed.
#     :param seed: seed for the numpy random number generator.
#     :param ntrials: the number of trials to run.
#     :return: the probability as described in the written pset.
#     """
#     # np.random.seed(seed)
    

#     #################### BEGIN YOUR CODE ####################
#     # You MUST use the function np.random.randint from numpy
#     # to generate random numbers, NOT random.randint or
#     # any other package.
#     pass
#     ####################  END YOUR CODE  ####################

import numpy as np

def play_game(round):
    second_player_wins = 0
    for i in range(round):
        S = 0
        x = 0
        y = 0
        while S <= 100:
            number = np.random.randint(1, 101)
            S += number
            x = number
        
    # print(S)
        while S <= 200:
            number = np.random.randint(1, 101)
            S += number
            y = number
    # print(S)
        if y > x:
            second_player_wins += 1

    return second_player_wins / round

round = 100000
probability = play_game(round)
print(probability)