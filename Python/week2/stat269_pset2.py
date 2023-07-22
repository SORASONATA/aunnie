# Do NOT add any other import statements
import numpy as np
import pandas as pd

# Name:Worada Sajai
# StudentID:650510679

"""
************************IMPORTANT************************
For calculate_probs and calculate_cond_probs, do NOT 
modify the name of the functions. Do not add or remove 
parameters to them either. Moreover, make sure your 
return value is exactly as described in the PDF handout 
and in the provided function comments. You are free to
write helper functions if you so desire.
Do NOT rename this file.
************************IMPORTANT************************
"""

def calculate_probs(filename):
    """
    filename is a path to a csv file, e.g. "bats.csv".
    You must use the filename variable. Do NOT alter the 
    filename variable.

    You should return a numpy array of length 6, 
    call it probs. probs[i] should be P(G_i = 1) for 
    0 <= i <= 4. probs[5] should be P(T = 1).

    See the assignment handout for some advice on how 
    to use numpy to make your life easier in this 
    function.
    """
    data = np.genfromtxt(filename, delimiter=',')
    
    gene_probs = np.mean(data[:, :5], axis=0)  # Calculate P(Gi = 1) for each gene
    trait_prob = np.mean(data[:, 5])  # Calculate P(T = 1)
    
    return np.concatenate((gene_probs, [trait_prob]))

# Usage example:
probs = calculate_probs("bats.csv")
# print(probs)
#    / pass  # TODO: delete this line and implement the function


def calculate_cond_probs(filename):
    """
    filename is a path to a csv file, e.g. "bats.csv".
    Do NOT alter the filename variable.

    You should return a numpy array of length 5, call it 
    probs, where probs[i] is equal to P(T = 1 | G_i = 1). See 
    the assignment handout for some information on 
    numpy functionality that'll help you here.
    """
    data = np.genfromtxt(filename, delimiter=',')
    
    cond_probs = np.zeros(5)
    
    for i in range(5):
        subset = data[np.where(data[:, i] == 1)]  # Subset of bats where Gi = 1
        cond_probs[i] = np.mean(subset[:, 5])  # Calculate P(T = 1|Gi = 1)
    
    return cond_probs

# Usage example:
cond_probs = calculate_cond_probs("bats.csv")
# print(cond_probs)
    # pass  # TODO: delete this line and implement the function


def sandbox():
    """
    You can write anything in this function to help you
    calculate probabilities required to answer part (c).
    You should submit your answer to part (c) in your
    solution file.
    This function will get called by our provided
    main method. Feel free to do whatever you want here, 
    including leaving this function blank.
    """
    # data = pd.read_csv("bats.csv")
    # correlations = data.corr()
    # print(correlations.iloc[:5, 5])  # Print the correlation coefficients between genes and the trait

# Usage example:
# sandbox()
    pass


def main():
    """
    We've provided this for convenience, simply to call the 
    functions above. Feel free to modify this function however 
    you like. We won't grade anything in this function.
    """
    filename = 'bats.csv'

    print("****** Calling calculate_probs on 'bats.csv' *******")
    probs = calculate_probs(filename)
    print(f"\treturned {probs}")
    print("*********************************************\n")

    print("*** Beginning calculate_cond_probs on 'bats.csv' ***")
    cond_probs = calculate_cond_probs(filename)
    print(f"\treturned {cond_probs}")
    print("*********************************************\n")

    print("*** Beginning sandbox - convenience function ***")
    sandbox()
    print("*********************************************\n")


############################################################
# This if-condition is True if this file was executed directly.
# It's False if this file was executed indirectly, e.g. as part
# of an import statement.
if __name__ == "__main__":
    main()
