// Dafny program main.dfy compiled into JavaScript
// Copyright by the contributors to the Dafny Project
// SPDX-License-Identifier: MIT

const BigNumber = require('bignumber.js');
BigNumber.config({ MODULO_MODE: BigNumber.EUCLID })
let _dafny = (function() {
  let $module = {};
  $module.areEqual = function(a, b) {
    if (typeof a === 'string' && b instanceof _dafny.Seq) {
      // Seq.equals(string) works as expected,
      // and the catch-all else block handles that direction.
      // But the opposite direction doesn't work; handle it here.
      return b.equals(a);
    } else if (typeof a === 'number' && BigNumber.isBigNumber(b)) {
      // This conditional would be correct even without the `typeof a` part,
      // but in most cases it's probably faster to short-circuit on a `typeof`
      // than to call `isBigNumber`. (But it remains to properly test this.)
      return b.isEqualTo(a);
    } else if (typeof a !== 'object' || a === null || b === null) {
      return a === b;
    } else if (BigNumber.isBigNumber(a)) {
      return a.isEqualTo(b);
    } else if (a._tname !== undefined || (Array.isArray(a) && a.constructor.name == "Array")) {
      return a === b;  // pointer equality
    } else {
      return a.equals(b);  // value-type equality
    }
  }
  $module.toString = function(a) {
    if (a === null) {
      return "null";
    } else if (typeof a === "number") {
      return a.toFixed();
    } else if (BigNumber.isBigNumber(a)) {
      return a.toFixed();
    } else if (a._tname !== undefined) {
      return a._tname;
    } else {
      return a.toString();
    }
  }
  $module.escapeCharacter = function(cp) {
    let s = String.fromCodePoint(cp.value)
    switch (s) {
      case '\n': return "\\n";
      case '\r': return "\\r";
      case '\t': return "\\t";
      case '\0': return "\\0";
      case '\'': return "\\'";
      case '\"': return "\\\"";
      case '\\': return "\\\\";
      default: return s;
    };
  }
  $module.NewObject = function() {
    return { _tname: "object" };
  }
  $module.InstanceOfTrait = function(obj, trait) {
    return obj._parentTraits !== undefined && obj._parentTraits().includes(trait);
  }
  $module.Rtd_bool = class {
    static get Default() { return false; }
  }
  $module.Rtd_char = class {
    static get Default() { return 'D'; }  // See CharType.DefaultValue in Dafny source code
  }
  $module.Rtd_codepoint = class {
    static get Default() { return new _dafny.CodePoint('D'.codePointAt(0)); }
  }
  $module.Rtd_int = class {
    static get Default() { return BigNumber(0); }
  }
  $module.Rtd_number = class {
    static get Default() { return 0; }
  }
  $module.Rtd_ref = class {
    static get Default() { return null; }
  }
  $module.Rtd_array = class {
    static get Default() { return []; }
  }
  $module.ZERO = new BigNumber(0);
  $module.ONE = new BigNumber(1);
  $module.NUMBER_LIMIT = new BigNumber(0x20).multipliedBy(0x1000000000000);  // 2^53
  $module.Tuple = class Tuple extends Array {
    constructor(...elems) {
      super(...elems);
    }
    toString() {
      return "(" + arrayElementsToString(this) + ")";
    }
    equals(other) {
      if (this === other) {
        return true;
      }
      for (let i = 0; i < this.length; i++) {
        if (!_dafny.areEqual(this[i], other[i])) {
          return false;
        }
      }
      return true;
    }
    static Default(...values) {
      return Tuple.of(...values);
    }
    static Rtd(...rtdArgs) {
      return {
        Default: Tuple.from(rtdArgs, rtd => rtd.Default)
      };
    }
  }
  $module.Set = class Set extends Array {
    constructor() {
      super();
    }
    static get Default() {
      return Set.Empty;
    }
    toString() {
      return "{" + arrayElementsToString(this) + "}";
    }
    static get Empty() {
      if (this._empty === undefined) {
        this._empty = new Set();
      }
      return this._empty;
    }
    static fromElements(...elmts) {
      let s = new Set();
      for (let k of elmts) {
        s.add(k);
      }
      return s;
    }
    contains(k) {
      for (let i = 0; i < this.length; i++) {
        if (_dafny.areEqual(this[i], k)) {
          return true;
        }
      }
      return false;
    }
    add(k) {  // mutates the Set; use only during construction
      if (!this.contains(k)) {
        this.push(k);
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.length !== other.length) {
        return false;
      }
      for (let e of this) {
        if (!other.contains(e)) {
          return false;
        }
      }
      return true;
    }
    get Elements() {
      return this;
    }
    Union(that) {
      if (this.length === 0) {
        return that;
      } else if (that.length === 0) {
        return this;
      } else {
        let s = Set.of(...this);
        for (let k of that) {
          s.add(k);
        }
        return s;
      }
    }
    Intersect(that) {
      if (this.length === 0) {
        return this;
      } else if (that.length === 0) {
        return that;
      } else {
        let s = new Set();
        for (let k of this) {
          if (that.contains(k)) {
            s.push(k);
          }
        }
        return s;
      }
    }
    Difference(that) {
      if (this.length == 0 || that.length == 0) {
        return this;
      } else {
        let s = new Set();
        for (let k of this) {
          if (!that.contains(k)) {
            s.push(k);
          }
        }
        return s;
      }
    }
    IsDisjointFrom(that) {
      for (let k of this) {
        if (that.contains(k)) {
          return false;
        }
      }
      return true;
    }
    IsSubsetOf(that) {
      if (that.length < this.length) {
        return false;
      }
      for (let k of this) {
        if (!that.contains(k)) {
          return false;
        }
      }
      return true;
    }
    IsProperSubsetOf(that) {
      if (that.length <= this.length) {
        return false;
      }
      for (let k of this) {
        if (!that.contains(k)) {
          return false;
        }
      }
      return true;
    }
    get AllSubsets() {
      return this.AllSubsets_();
    }
    *AllSubsets_() {
      // Start by putting all set elements into a list, but don't include null
      let elmts = Array.of(...this);
      let n = elmts.length;
      let which = new Array(n);
      which.fill(false);
      let a = [];
      while (true) {
        yield Set.of(...a);
        // "add 1" to "which", as if doing a carry chain.  For every digit changed, change the membership of the corresponding element in "a".
        let i = 0;
        for (; i < n && which[i]; i++) {
          which[i] = false;
          // remove elmts[i] from a
          for (let j = 0; j < a.length; j++) {
            if (_dafny.areEqual(a[j], elmts[i])) {
              // move the last element of a into slot j
              a[j] = a[-1];
              a.pop();
              break;
            }
          }
        }
        if (i === n) {
          // we have cycled through all the subsets
          break;
        }
        which[i] = true;
        a.push(elmts[i]);
      }
    }
  }
  $module.MultiSet = class MultiSet extends Array {
    constructor() {
      super();
    }
    static get Default() {
      return MultiSet.Empty;
    }
    toString() {
      let s = "multiset{";
      let sep = "";
      for (let e of this) {
        let [k, n] = e;
        let ks = _dafny.toString(k);
        while (!n.isZero()) {
          n = n.minus(1);
          s += sep + ks;
          sep = ", ";
        }
      }
      s += "}";
      return s;
    }
    static get Empty() {
      if (this._empty === undefined) {
        this._empty = new MultiSet();
      }
      return this._empty;
    }
    static fromElements(...elmts) {
      let s = new MultiSet();
      for (let e of elmts) {
        s.add(e, _dafny.ONE);
      }
      return s;
    }
    static FromArray(arr) {
      let s = new MultiSet();
      for (let e of arr) {
        s.add(e, _dafny.ONE);
      }
      return s;
    }
    cardinality() {
      let c = _dafny.ZERO;
      for (let e of this) {
        let [k, n] = e;
        c = c.plus(n);
      }
      return c;
    }
    clone() {
      let s = new MultiSet();
      for (let e of this) {
        let [k, n] = e;
        s.push([k, n]);  // make sure to create a new array [k, n] here
      }
      return s;
    }
    findIndex(k) {
      for (let i = 0; i < this.length; i++) {
        if (_dafny.areEqual(this[i][0], k)) {
          return i;
        }
      }
      return this.length;
    }
    get(k) {
      let i = this.findIndex(k);
      if (i === this.length) {
        return _dafny.ZERO;
      } else {
        return this[i][1];
      }
    }
    contains(k) {
      return !this.get(k).isZero();
    }
    add(k, n) {
      let i = this.findIndex(k);
      if (i === this.length) {
        this.push([k, n]);
      } else {
        let m = this[i][1];
        this[i] = [k, m.plus(n)];
      }
    }
    update(k, n) {
      let i = this.findIndex(k);
      if (i < this.length && this[i][1].isEqualTo(n)) {
        return this;
      } else if (i === this.length && n.isZero()) {
        return this;
      } else if (i === this.length) {
        let m = this.slice();
        m.push([k, n]);
        return m;
      } else {
        let m = this.slice();
        m[i] = [k, n];
        return m;
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      }
      for (let e of this) {
        let [k, n] = e;
        let m = other.get(k);
        if (!n.isEqualTo(m)) {
          return false;
        }
      }
      return this.cardinality().isEqualTo(other.cardinality());
    }
    get Elements() {
      return this.Elements_();
    }
    *Elements_() {
      for (let i = 0; i < this.length; i++) {
        let [k, n] = this[i];
        while (!n.isZero()) {
          yield k;
          n = n.minus(1);
        }
      }
    }
    get UniqueElements() {
      return this.UniqueElements_();
    }
    *UniqueElements_() {
      for (let e of this) {
        let [k, n] = e;
        if (!n.isZero()) {
          yield k;
        }
      }
    }
    Union(that) {
      if (this.length === 0) {
        return that;
      } else if (that.length === 0) {
        return this;
      } else {
        let s = this.clone();
        for (let e of that) {
          let [k, n] = e;
          s.add(k, n);
        }
        return s;
      }
    }
    Intersect(that) {
      if (this.length === 0) {
        return this;
      } else if (that.length === 0) {
        return that;
      } else {
        let s = new MultiSet();
        for (let e of this) {
          let [k, n] = e;
          let m = that.get(k);
          if (!m.isZero()) {
            s.push([k, m.isLessThan(n) ? m : n]);
          }
        }
        return s;
      }
    }
    Difference(that) {
      if (this.length === 0 || that.length === 0) {
        return this;
      } else {
        let s = new MultiSet();
        for (let e of this) {
          let [k, n] = e;
          let d = n.minus(that.get(k));
          if (d.isGreaterThan(0)) {
            s.push([k, d]);
          }
        }
        return s;
      }
    }
    IsDisjointFrom(that) {
      let intersection = this.Intersect(that);
      return intersection.cardinality().isZero();
    }
    IsSubsetOf(that) {
      for (let e of this) {
        let [k, n] = e;
        let m = that.get(k);
        if (!n.isLessThanOrEqualTo(m)) {
          return false;
        }
      }
      return true;
    }
    IsProperSubsetOf(that) {
      return this.IsSubsetOf(that) && this.cardinality().isLessThan(that.cardinality());
    }
  }
  $module.CodePoint = class CodePoint {
    constructor(value) {
      this.value = value
    }
    equals(other) {
      if (this === other) {
        return true;
      }
      return this.value === other.value
    }
    isLessThan(other) {
      return this.value < other.value
    }
    isLessThanOrEqual(other) {
      return this.value <= other.value
    }
    toString() {
      return "'" + $module.escapeCharacter(this) + "'";
    }
    static isCodePoint(i) {
      return (
        (_dafny.ZERO.isLessThanOrEqualTo(i) && i.isLessThan(new BigNumber(0xD800))) ||
        (new BigNumber(0xE000).isLessThanOrEqualTo(i) && i.isLessThan(new BigNumber(0x11_0000))))
    }
  }
  $module.Seq = class Seq extends Array {
    constructor(...elems) {
      super(...elems);
    }
    static get Default() {
      return Seq.of();
    }
    static Create(n, init) {
      return Seq.from({length: n}, (_, i) => init(new BigNumber(i)));
    }
    static UnicodeFromString(s) {
      return new Seq(...([...s].map(c => new _dafny.CodePoint(c.codePointAt(0)))))
    }
    toString() {
      return "[" + arrayElementsToString(this) + "]";
    }
    toVerbatimString(asLiteral) {
      if (asLiteral) {
        return '"' + this.map(c => _dafny.escapeCharacter(c)).join("") + '"';
      } else {
        return this.map(c => String.fromCodePoint(c.value)).join("");
      }
    }
    static update(s, i, v) {
      if (typeof s === "string") {
        let p = s.slice(0, i);
        let q = s.slice(i.toNumber() + 1);
        return p.concat(v, q);
      } else {
        let t = s.slice();
        t[i] = v;
        return t;
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.length !== other.length) {
        return false;
      }
      for (let i = 0; i < this.length; i++) {
        if (!_dafny.areEqual(this[i], other[i])) {
          return false;
        }
      }
      return true;
    }
    static contains(s, k) {
      if (typeof s === "string") {
        return s.includes(k);
      } else {
        for (let x of s) {
          if (_dafny.areEqual(x, k)) {
            return true;
          }
        }
        return false;
      }
    }
    get Elements() {
      return this;
    }
    get UniqueElements() {
      return _dafny.Set.fromElements(...this);
    }
    static Concat(a, b) {
      if (typeof a === "string" || typeof b === "string") {
        // string concatenation, so make sure both operands are strings before concatenating
        if (typeof a !== "string") {
          // a must be a Seq
          a = a.join("");
        }
        if (typeof b !== "string") {
          // b must be a Seq
          b = b.join("");
        }
        return a + b;
      } else {
        // ordinary concatenation
        let r = Seq.of(...a);
        r.push(...b);
        return r;
      }
    }
    static JoinIfPossible(x) {
      try { return x.join(""); } catch(_error) { return x; }
    }
    static IsPrefixOf(a, b) {
      if (b.length < a.length) {
        return false;
      }
      for (let i = 0; i < a.length; i++) {
        if (!_dafny.areEqual(a[i], b[i])) {
          return false;
        }
      }
      return true;
    }
    static IsProperPrefixOf(a, b) {
      if (b.length <= a.length) {
        return false;
      }
      for (let i = 0; i < a.length; i++) {
        if (!_dafny.areEqual(a[i], b[i])) {
          return false;
        }
      }
      return true;
    }
  }
  $module.Map = class Map extends Array {
    constructor() {
      super();
    }
    static get Default() {
      return Map.of();
    }
    toString() {
      return "map[" + this.map(maplet => _dafny.toString(maplet[0]) + " := " + _dafny.toString(maplet[1])).join(", ") + "]";
    }
    static get Empty() {
      if (this._empty === undefined) {
        this._empty = new Map();
      }
      return this._empty;
    }
    findIndex(k) {
      for (let i = 0; i < this.length; i++) {
        if (_dafny.areEqual(this[i][0], k)) {
          return i;
        }
      }
      return this.length;
    }
    get(k) {
      let i = this.findIndex(k);
      if (i === this.length) {
        return undefined;
      } else {
        return this[i][1];
      }
    }
    contains(k) {
      return this.findIndex(k) < this.length;
    }
    update(k, v) {
      let m = this.slice();
      m.updateUnsafe(k, v);
      return m;
    }
    // Similar to update, but make the modification in-place.
    // Meant to be used in the map constructor.
    updateUnsafe(k, v) {
      let m = this;
      let i = m.findIndex(k);
      m[i] = [k, v];
      return m;
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.length !== other.length) {
        return false;
      }
      for (let e of this) {
        let [k, v] = e;
        let w = other.get(k);
        if (w === undefined || !_dafny.areEqual(v, w)) {
          return false;
        }
      }
      return true;
    }
    get Keys() {
      let s = new _dafny.Set();
      for (let e of this) {
        let [k, v] = e;
        s.push(k);
      }
      return s;
    }
    get Values() {
      let s = new _dafny.Set();
      for (let e of this) {
        let [k, v] = e;
        s.add(v);
      }
      return s;
    }
    get Items() {
      let s = new _dafny.Set();
      for (let e of this) {
        let [k, v] = e;
        s.push(_dafny.Tuple.of(k, v));
      }
      return s;
    }
    Merge(that) {
      let m = that.slice();
      for (let e of this) {
        let [k, v] = e;
        let i = m.findIndex(k);
        if (i == m.length) {
          m[i] = [k, v];
        }
      }
      return m;
    }
    Subtract(keys) {
      if (this.length === 0 || keys.length === 0) {
        return this;
      }
      let m = new Map();
      for (let e of this) {
        let [k, v] = e;
        if (!keys.contains(k)) {
          m[m.length] = e;
        }
      }
      return m;
    }
  }
  $module.newArray = function(initValue, ...dims) {
    return { dims: dims, elmts: buildArray(initValue, ...dims) };
  }
  $module.BigOrdinal = class BigOrdinal {
    static get Default() {
      return _dafny.ZERO;
    }
    static IsLimit(ord) {
      return ord.isZero();
    }
    static IsSucc(ord) {
      return ord.isGreaterThan(0);
    }
    static Offset(ord) {
      return ord;
    }
    static IsNat(ord) {
      return true;  // at run time, every ORDINAL is a natural number
    }
  }
  $module.BigRational = class BigRational {
    static get ZERO() {
      if (this._zero === undefined) {
        this._zero = new BigRational(_dafny.ZERO);
      }
      return this._zero;
    }
    constructor (n, d) {
      // requires d === undefined || 1 <= d
      this.num = n;
      this.den = d === undefined ? _dafny.ONE : d;
      // invariant 1 <= den || (num == 0 && den == 0)
    }
    static get Default() {
      return _dafny.BigRational.ZERO;
    }
    // We need to deal with the special case `num == 0 && den == 0`, because
    // that's what C#'s default struct constructor will produce for BigRational. :(
    // To deal with it, we ignore `den` when `num` is 0.
    toString() {
      if (this.num.isZero() || this.den.isEqualTo(1)) {
        return this.num.toFixed() + ".0";
      }
      let answer = this.dividesAPowerOf10(this.den);
      if (answer !== undefined) {
        let n = this.num.multipliedBy(answer[0]);
        let log10 = answer[1];
        let sign, digits;
        if (this.num.isLessThan(0)) {
          sign = "-"; digits = n.negated().toFixed();
        } else {
          sign = ""; digits = n.toFixed();
        }
        if (log10 < digits.length) {
          let digitCount = digits.length - log10;
          return sign + digits.slice(0, digitCount) + "." + digits.slice(digitCount);
        } else {
          return sign + "0." + "0".repeat(log10 - digits.length) + digits;
        }
      } else {
        return "(" + this.num.toFixed() + ".0 / " + this.den.toFixed() + ".0)";
      }
    }
    isPowerOf10(x) {
      if (x.isZero()) {
        return undefined;
      }
      let log10 = 0;
      while (true) {  // invariant: x != 0 && x * 10^log10 == old(x)
        if (x.isEqualTo(1)) {
          return log10;
        } else if (x.mod(10).isZero()) {
          log10++;
          x = x.dividedToIntegerBy(10);
        } else {
          return undefined;
        }
      }
    }
    dividesAPowerOf10(i) {
      let factor = _dafny.ONE;
      let log10 = 0;
      if (i.isLessThanOrEqualTo(_dafny.ZERO)) {
        return undefined;
      }

      // invariant: 1 <= i && i * 10^log10 == factor * old(i)
      while (i.mod(10).isZero()) {
        i = i.dividedToIntegerBy(10);
       log10++;
      }

      while (i.mod(5).isZero()) {
        i = i.dividedToIntegerBy(5);
        factor = factor.multipliedBy(2);
        log10++;
      }
      while (i.mod(2).isZero()) {
        i = i.dividedToIntegerBy(2);
        factor = factor.multipliedBy(5);
        log10++;
      }

      if (i.isEqualTo(_dafny.ONE)) {
        return [factor, log10];
      } else {
        return undefined;
      }
    }
    toBigNumber() {
      if (this.num.isZero() || this.den.isEqualTo(1)) {
        return this.num;
      } else if (this.num.isGreaterThan(0)) {
        return this.num.dividedToIntegerBy(this.den);
      } else {
        return this.num.minus(this.den).plus(1).dividedToIntegerBy(this.den);
      }
    }
    isInteger() {
      return this.equals(new _dafny.BigRational(this.toBigNumber(), _dafny.ONE));
    }
    // Returns values such that aa/dd == a and bb/dd == b.
    normalize(b) {
      let a = this;
      let aa, bb, dd;
      if (a.num.isZero()) {
        aa = a.num;
        bb = b.num;
        dd = b.den;
      } else if (b.num.isZero()) {
        aa = a.num;
        dd = a.den;
        bb = b.num;
      } else {
        let gcd = BigNumberGcd(a.den, b.den);
        let xx = a.den.dividedToIntegerBy(gcd);
        let yy = b.den.dividedToIntegerBy(gcd);
        // We now have a == a.num / (xx * gcd) and b == b.num / (yy * gcd).
        aa = a.num.multipliedBy(yy);
        bb = b.num.multipliedBy(xx);
        dd = a.den.multipliedBy(yy);
      }
      return [aa, bb, dd];
    }
    compareTo(that) {
      // simple things first
      let asign = this.num.isZero() ? 0 : this.num.isLessThan(0) ? -1 : 1;
      let bsign = that.num.isZero() ? 0 : that.num.isLessThan(0) ? -1 : 1;
      if (asign < 0 && 0 <= bsign) {
        return -1;
      } else if (asign <= 0 && 0 < bsign) {
        return -1;
      } else if (bsign < 0 && 0 <= asign) {
        return 1;
      } else if (bsign <= 0 && 0 < asign) {
        return 1;
      }
      let [aa, bb, dd] = this.normalize(that);
      if (aa.isLessThan(bb)) {
        return -1;
      } else if (aa.isEqualTo(bb)){
        return 0;
      } else {
        return 1;
      }
    }
    equals(that) {
      return this.compareTo(that) === 0;
    }
    isLessThan(that) {
      return this.compareTo(that) < 0;
    }
    isAtMost(that) {
      return this.compareTo(that) <= 0;
    }
    plus(b) {
      let [aa, bb, dd] = this.normalize(b);
      return new BigRational(aa.plus(bb), dd);
    }
    minus(b) {
      let [aa, bb, dd] = this.normalize(b);
      return new BigRational(aa.minus(bb), dd);
    }
    negated() {
      return new BigRational(this.num.negated(), this.den);
    }
    multipliedBy(b) {
      return new BigRational(this.num.multipliedBy(b.num), this.den.multipliedBy(b.den));
    }
    dividedBy(b) {
      let a = this;
      // Compute the reciprocal of b
      let bReciprocal;
      if (b.num.isGreaterThan(0)) {
        bReciprocal = new BigRational(b.den, b.num);
      } else {
        // this is the case b.num < 0
        bReciprocal = new BigRational(b.den.negated(), b.num.negated());
      }
      return a.multipliedBy(bReciprocal);
    }
  }
  $module.EuclideanDivisionNumber = function(a, b) {
    if (0 <= a) {
      if (0 <= b) {
        // +a +b: a/b
        return Math.floor(a / b);
      } else {
        // +a -b: -(a/(-b))
        return -Math.floor(a / -b);
      }
    } else {
      if (0 <= b) {
        // -a +b: -((-a-1)/b) - 1
        return -Math.floor((-a-1) / b) - 1;
      } else {
        // -a -b: ((-a-1)/(-b)) + 1
        return Math.floor((-a-1) / -b) + 1;
      }
    }
  }
  $module.EuclideanDivision = function(a, b) {
    if (a.isGreaterThanOrEqualTo(0)) {
      if (b.isGreaterThanOrEqualTo(0)) {
        // +a +b: a/b
        return a.dividedToIntegerBy(b);
      } else {
        // +a -b: -(a/(-b))
        return a.dividedToIntegerBy(b.negated()).negated();
      }
    } else {
      if (b.isGreaterThanOrEqualTo(0)) {
        // -a +b: -((-a-1)/b) - 1
        return a.negated().minus(1).dividedToIntegerBy(b).negated().minus(1);
      } else {
        // -a -b: ((-a-1)/(-b)) + 1
        return a.negated().minus(1).dividedToIntegerBy(b.negated()).plus(1);
      }
    }
  }
  $module.EuclideanModuloNumber = function(a, b) {
    let bp = Math.abs(b);
    if (0 <= a) {
      // +a: a % bp
      return a % bp;
    } else {
      // c = ((-a) % bp)
      // -a: bp - c if c > 0
      // -a: 0 if c == 0
      let c = (-a) % bp;
      return c === 0 ? c : bp - c;
    }
  }
  $module.ShiftLeft = function(b, n) {
    return b.multipliedBy(new BigNumber(2).exponentiatedBy(n));
  }
  $module.ShiftRight = function(b, n) {
    return b.dividedToIntegerBy(new BigNumber(2).exponentiatedBy(n));
  }
  $module.RotateLeft = function(b, n, w) {  // truncate(b << n) | (b >> (w - n))
    let x = _dafny.ShiftLeft(b, n).mod(new BigNumber(2).exponentiatedBy(w));
    let y = _dafny.ShiftRight(b, w - n);
    return x.plus(y);
  }
  $module.RotateRight = function(b, n, w) {  // (b >> n) | truncate(b << (w - n))
    let x = _dafny.ShiftRight(b, n);
    let y = _dafny.ShiftLeft(b, w - n).mod(new BigNumber(2).exponentiatedBy(w));;
    return x.plus(y);
  }
  $module.BitwiseAnd = function(a, b) {
    let r = _dafny.ZERO;
    const m = _dafny.NUMBER_LIMIT;  // 2^53
    let h = _dafny.ONE;
    while (!a.isZero() && !b.isZero()) {
      let a0 = a.mod(m);
      let b0 = b.mod(m);
      r = r.plus(h.multipliedBy(a0 & b0));
      a = a.dividedToIntegerBy(m);
      b = b.dividedToIntegerBy(m);
      h = h.multipliedBy(m);
    }
    return r;
  }
  $module.BitwiseOr = function(a, b) {
    let r = _dafny.ZERO;
    const m = _dafny.NUMBER_LIMIT;  // 2^53
    let h = _dafny.ONE;
    while (!a.isZero() && !b.isZero()) {
      let a0 = a.mod(m);
      let b0 = b.mod(m);
      r = r.plus(h.multipliedBy(a0 | b0));
      a = a.dividedToIntegerBy(m);
      b = b.dividedToIntegerBy(m);
      h = h.multipliedBy(m);
    }
    r = r.plus(h.multipliedBy(a | b));
    return r;
  }
  $module.BitwiseXor = function(a, b) {
    let r = _dafny.ZERO;
    const m = _dafny.NUMBER_LIMIT;  // 2^53
    let h = _dafny.ONE;
    while (!a.isZero() && !b.isZero()) {
      let a0 = a.mod(m);
      let b0 = b.mod(m);
      r = r.plus(h.multipliedBy(a0 ^ b0));
      a = a.dividedToIntegerBy(m);
      b = b.dividedToIntegerBy(m);
      h = h.multipliedBy(m);
    }
    r = r.plus(h.multipliedBy(a | b));
    return r;
  }
  $module.BitwiseNot = function(a, bits) {
    let r = _dafny.ZERO;
    let h = _dafny.ONE;
    for (let i = 0; i < bits; i++) {
      let bit = a.mod(2);
      if (bit.isZero()) {
        r = r.plus(h);
      }
      a = a.dividedToIntegerBy(2);
      h = h.multipliedBy(2);
    }
    return r;
  }
  $module.Quantifier = function(vals, frall, pred) {
    for (let u of vals) {
      if (pred(u) !== frall) { return !frall; }
    }
    return frall;
  }
  $module.PlusChar = function(a, b) {
    return String.fromCharCode(a.charCodeAt(0) + b.charCodeAt(0));
  }
  $module.UnicodePlusChar = function(a, b) {
    return new _dafny.CodePoint(a.value + b.value);
  }
  $module.MinusChar = function(a, b) {
    return String.fromCharCode(a.charCodeAt(0) - b.charCodeAt(0));
  }
  $module.UnicodeMinusChar = function(a, b) {
    return new _dafny.CodePoint(a.value - b.value);
  }
  $module.AllBooleans = function*() {
    yield false;
    yield true;
  }
  $module.AllChars = function*() {
    for (let i = 0; i < 0x10000; i++) {
      yield String.fromCharCode(i);
    }
  }
  $module.AllUnicodeChars = function*() {
    for (let i = 0; i < 0xD800; i++) {
      yield new _dafny.CodePoint(i);
    }
    for (let i = 0xE0000; i < 0x110000; i++) {
      yield new _dafny.CodePoint(i);
    }
  }
  $module.AllIntegers = function*() {
    yield _dafny.ZERO;
    for (let j = _dafny.ONE;; j = j.plus(1)) {
      yield j;
      yield j.negated();
    }
  }
  $module.IntegerRange = function*(lo, hi) {
    if (lo === null) {
      while (true) {
        hi = hi.minus(1);
        yield hi;
      }
    } else if (hi === null) {
      while (true) {
        yield lo;
        lo = lo.plus(1);
      }
    } else {
      while (lo.isLessThan(hi)) {
        yield lo;
        lo = lo.plus(1);
      }
    }
  }
  $module.SingleValue = function*(v) {
    yield v;
  }
  $module.HaltException = class HaltException extends Error {
    constructor(message) {
      super(message)
    }
  }
  $module.HandleHaltExceptions = function(f) {
    try {
      f()
    } catch (e) {
      if (e instanceof _dafny.HaltException) {
        process.stdout.write("[Program halted] " + e.message + "\n")
        process.exitCode = 1
      } else {
        throw e
      }
    }
  }
  $module.FromMainArguments = function(args) {
    var a = [...args];
    a.splice(0, 2, args[0] + " " + args[1]);
    return a;
  }
  $module.UnicodeFromMainArguments = function(args) {
    return $module.FromMainArguments(args).map(_dafny.Seq.UnicodeFromString);
  }
  return $module;

  // What follows are routines private to the Dafny runtime
  function buildArray(initValue, ...dims) {
    if (dims.length === 0) {
      return initValue;
    } else {
      let a = Array(dims[0].toNumber());
      let b = Array.from(a, (x) => buildArray(initValue, ...dims.slice(1)));
      return b;
    }
  }
  function arrayElementsToString(a) {
    // like `a.join(", ")`, but calling _dafny.toString(x) on every element x instead of x.toString()
    let s = "";
    let sep = "";
    for (let x of a) {
      s += sep + _dafny.toString(x);
      sep = ", ";
    }
    return s;
  }
  function BigNumberGcd(a, b){  // gcd of two non-negative BigNumber's
    while (true) {
      if (a.isZero()) {
        return b;
      } else if (b.isZero()) {
        return a;
      }
      if (a.isLessThan(b)) {
        b = b.modulo(a);
      } else {
        a = a.modulo(b);
      }
    }
  }
})();
// Dafny program systemModulePopulator.dfy compiled into JavaScript
let _System = (function() {
  let $module = {};

  $module.nat = class nat {
    constructor () {
    }
    static get Default() {
      return _dafny.ZERO;
    }
    static _Is(__source) {
      let _0_x = (__source);
      return (_dafny.ZERO).isLessThanOrEqualTo(_0_x);
    }
  };

  return $module;
})(); // end of module _System
let _module = (function() {
  let $module = {};

  $module.__default = class __default {
    constructor () {
      this._tname = "_module._default";
    }
    _parentTraits() {
      return [];
    }
    static abs(x) {
      if ((x).isLessThan(_dafny.ZERO)) {
        return (new BigNumber(-1)).multipliedBy(x);
      } else {
        return x;
      }
    };
    static safeIndex(x, length) {
      if ((x).isLessThan(_dafny.ZERO)) {
        return _dafny.ZERO;
      } else if ((length).isLessThanOrEqualTo(x)) {
        return (x).mod(length);
      } else {
        return x;
      }
    };
    static safeDivisionInt(x1, x2) {
      if ((x2).isEqualTo(_dafny.ZERO)) {
        return x1;
      } else {
        return _dafny.EuclideanDivision(x1, x2);
      }
    };
    static safeModuloInt(x1, x2) {
      if ((x2).isEqualTo(_dafny.ZERO)) {
        return x1;
      } else {
        return (x1).mod(x2);
      }
    };
    static fm2(p0, p1, p2, globalState) {
      return ((_dafny.ZERO).minus(((true) ? (new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(false),new _dafny.CodePoint('n'.codePointAt(0)))).length)) : (new BigNumber(98))))).isLessThanOrEqualTo((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("bdayefsjg"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(543)), function (_0_i0) {
        return new _dafny.CodePoint('w'.codePointAt(0));
      }))).length)));
    };
    static fm3(globalState) {
      return _dafny.Seq.Create(_module.__default.abs(new BigNumber(988)), function (_1_i0) {
        return _dafny.Map.Empty.slice().updateUnsafe(true,!(true));
      });
    };
    static fm4(p0, p1, p2, p3, globalState) {
      return _module.D1.create_DC1(_dafny.Seq.Create(_module.__default.abs(new BigNumber(61)), function (_2_i0) {
  return _dafny.Map.Empty.slice().updateUnsafe(false,false);
}));
    };
    static fm5(p0, globalState) {
      return new _dafny.CodePoint('o'.codePointAt(0));
    };
    static fm6(p0, p1, p2, globalState) {
      return _dafny.Seq.of(false, !((_dafny.Map.Empty.slice().updateUnsafe(true,true)).equals(_dafny.Map.Empty.slice().updateUnsafe(true,true))), !(false));
    };
    static fm8(p0, globalState) {
      return _dafny.Seq.Create(_module.__default.abs(new BigNumber(32)), function (_3_i0) {
        return new _dafny.CodePoint('x'.codePointAt(0));
      });
    };
    static fm9(p0, globalState) {
      return ((_dafny.ZERO).minus(_module.__default.safeDivisionInt(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Set.fromElements(true)).length))).cardinality()), new BigNumber((function () {
        let _coll0 = new _dafny.Map();
        for (const _compr_0 of _dafny.IntegerRange(new BigNumber(468), new BigNumber(-404))) {
          let _4_v0 = _compr_0;
          if (((new BigNumber(468)).isLessThanOrEqualTo(_4_v0)) && ((_4_v0).isLessThan(new BigNumber(-404)))) {
            _coll0.push([(_4_v0).multipliedBy(new BigNumber(203)),new BigNumber(415)]);
          }
        }
        return _coll0;
      }()).length)))).plus(new BigNumber(665));
    };
    static fm10(p0, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.of(true), _dafny.Seq.Concat(_dafny.Seq.of(true), _dafny.Seq.of(false, true)));
    };
    static fm11(p0, p1, p2, globalState) {
      return ((_dafny.MultiSet.fromElements((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("kjyj")).length)), new BigNumber(-123))).Intersect(_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("pmifknqvs")).length)))).Difference(_dafny.MultiSet.FromArray(_dafny.Seq.Concat(_dafny.Seq.of(new BigNumber(781)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-166)), function (_5_i0) {
        return new BigNumber(406);
      }))));
    };
    static fm12(p0, p1, p2, p3, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(684)), function (_6_i0) {
        return (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("fxkcy")).length));
      }), _dafny.Seq.Create(_module.__default.abs(new BigNumber(173)), function (_7_i1) {
        return new BigNumber(-110);
      })), _dafny.Seq.of(new BigNumber(194)));
    };
    static fm15(p0, p1, globalState) {
      return ((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.of(_dafny.Seq.of(new BigNumber(57), new BigNumber(825)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(13)), function (_8_i0) {
        return new BigNumber(897);
      }), _dafny.Seq.of(new BigNumber(-853), new BigNumber((_dafny.Seq.UnicodeFromString("inrv")).length)))).length), new BigNumber(-354), new BigNumber(-817))).Difference(_dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber(4))))).Union((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(5))).length), (_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.fromElements(false)).cardinality())), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,false)).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("uvdonglnj")).length),false)).length))).Intersect(_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(577)), function (_9_i1) {
        return new BigNumber(527);
      })).length))));
    };
    static fm16(p0, p1, p2, globalState) {
      return ((_dafny.Set.fromElements(false, true, false)).Difference(_dafny.Set.fromElements(!(true)))).Intersect(((true) ? (_dafny.Set.fromElements(false, true, (_module.D4.create_DC9(new BigNumber(293), new _dafny.CodePoint('i'.codePointAt(0)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(726)), function (_10_i0) {
  return new BigNumber(767);
}), new BigNumber(37), false)).dtor_cf20, true)) : (_dafny.Set.fromElements(true, true))));
    };
    static fm17(p0, p1, p2, p3, globalState) {
      return ((_dafny.ZERO).minus((new BigNumber(190)).minus((_dafny.ZERO).minus(new BigNumber(-831))))).multipliedBy((_dafny.ZERO).minus(_module.__default.safeModuloInt(new BigNumber((_dafny.Seq.of(false, true)).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!(false),new _dafny.CodePoint('r'.codePointAt(0)))).length))));
    };
    static fm18(p0, p1, globalState) {
      return _module.D1.create_DC2(new BigNumber(503), false, (new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(true, !(true))).length)),true)).length)).plus(new BigNumber(829)));
    };
    static fm21(p0, p1, p2, p3, globalState) {
      return (new BigNumber((_dafny.Seq.Concat(_dafny.Seq.of(!(!(true)), false, false), _dafny.Seq.of(true))).length)).isLessThan((_dafny.ZERO).minus(new BigNumber(-948)));
    };
    static fm22(p0, p1, p2, globalState) {
      return _module.D2.create_DC3((_dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.FromArray(_dafny.Seq.Create(_module.__default.abs(new BigNumber(884)), function (_11_i0) {
  return new BigNumber(925);
})),new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(78)), function (_12_i1) {
  return new BigNumber((_dafny.Seq.UnicodeFromString("fqwbp")).length);
})).length))).Merge(_dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.Create(_module.__default.abs(new BigNumber(769)), function (_13_i2) {
  return new BigNumber(837);
}))).cardinality()), new BigNumber(471))),new BigNumber(819))));
    };
    static fm23(p0, p1, p2, p3, globalState) {
      return ((new BigNumber(239)).minus(new BigNumber(156))).isLessThanOrEqualTo(new BigNumber((_dafny.Seq.Concat(_dafny.Seq.of(new BigNumber(394)), _dafny.Seq.of(new BigNumber(904)))).length));
    };
    static fm24(p0, globalState) {
      return (_dafny.ZERO).minus(((new BigNumber(938)).multipliedBy(new BigNumber((_dafny.Seq.of(false)).length))).multipliedBy((new BigNumber(634)).multipliedBy(new BigNumber(619))));
    };
    static fm25(p0, p1, globalState) {
      let _source0 = _module.D13.create_DC25((_dafny.ZERO).minus((_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,false)).length)))), new BigNumber(787), false);
      if (_source0.is_DC25) {
        let _14___mcc_h0 = (_source0).cf42;
        let _15___mcc_h1 = (_source0).cf43;
        let _16___mcc_h2 = (_source0).cf44;
        let _17_cf44 = _16___mcc_h2;
        let _18_cf43 = _15___mcc_h1;
        let _19_cf42 = _14___mcc_h0;
        return _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("li"), _dafny.Seq.UnicodeFromString("ggkk"));
      } else {
        let _20___mcc_h3 = (_source0).cf41;
        let _21_cf41 = _20___mcc_h3;
        return _dafny.Seq.UnicodeFromString("vnno");
      }
    };
    static fm26(p0, p1, p2, globalState) {
      return new _dafny.CodePoint('v'.codePointAt(0));
    };
    static fm27(p0, p1, p2, globalState) {
      if (_dafny.Seq.IsPrefixOf(_dafny.Seq.UnicodeFromString("hls"), _dafny.Seq.UnicodeFromString("nhuwclrv"))) {
        return _dafny.Seq.Create(_module.__default.abs(new BigNumber(472)), function (_22_i0) {
          return (_dafny.ZERO).minus(new BigNumber(-387));
        });
      } else {
        return _dafny.Seq.Concat(_dafny.Seq.of(new BigNumber(((_module.D14.create_DC26(_dafny.MultiSet.fromElements(false))).dtor_cf45).cardinality())), (_module.D4.create_DC9(new BigNumber(317), new _dafny.CodePoint('l'.codePointAt(0)), _dafny.Seq.of(new BigNumber(111), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new _dafny.CodePoint('v'.codePointAt(0)))).length)), new BigNumber(447), !(true))).dtor_cf18);
      }
    };
    static fm28(p0, globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Set.fromElements(true)).length),new BigNumber((_dafny.Seq.of(true, false)).length))).length),new BigNumber(729))).Merge(function () {
        let _coll1 = new _dafny.Map();
        for (const _compr_1 of _dafny.IntegerRange(new BigNumber(24), new BigNumber(239))) {
          let _23_v0 = _compr_1;
          if (((new BigNumber(24)).isLessThanOrEqualTo(_23_v0)) && ((_23_v0).isLessThan(new BigNumber(239)))) {
            _coll1.push([(_23_v0).multipliedBy(new BigNumber(758)),new BigNumber(627)]);
          }
        }
        return _coll1;
      }())).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(167),new BigNumber((_dafny.Seq.of(true)).length)));
    };
    static fm30(p0, p1, p2, globalState) {
      if (!(!(true))) {
        return _module.D2.create_DC5(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.of(!(true))).length), new BigNumber(-659))).length), !(true), true, true, !(false));
      } else {
        return _module.D2.create_DC5(new BigNumber(63), true, !(false), false, false);
      }
    };
    static fm31(p0, p1, p2, globalState) {
      return (_dafny.Set.fromElements(false, false)).Union((_dafny.Set.fromElements(false, true)).Union(_dafny.Set.fromElements(true)));
    };
    static fm32(p0, p1, p2, p3, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.of(false, !(!(!(true)))), _dafny.Seq.of(true, true)), _dafny.Seq.of(!(true)));
    };
    static fm33(p0, p1, p2, p3, globalState) {
      return (_dafny.MultiSet.fromElements(_dafny.Seq.of(true))).Intersect(_dafny.MultiSet.fromElements(_dafny.Seq.of(false, false)));
    };
    static fm34(p0, p1, p2, globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe((_module.D3.create_DC7(new BigNumber(397), new BigNumber((_dafny.Seq.of(!(true))).length))).dtor_cf14,false)).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Set.fromElements(false)).length),true))).Merge((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(637))).cardinality()), new BigNumber(462))).length),false)).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(449),!(true))));
    };
    static fm35(globalState) {
      if (!(false) || (true)) {
        return _dafny.Set.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(397)), function (_24_i0) {
          return new _dafny.CodePoint('b'.codePointAt(0));
        }));
      } else {
        return _dafny.Set.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-615)), function (_25_i1) {
          return new _dafny.CodePoint('o'.codePointAt(0));
        }), _dafny.Seq.UnicodeFromString("b"));
      }
    };
    static fm36(p0, globalState) {
      return _module.D7.create_DC14((new BigNumber(814)).minus(new BigNumber((_dafny.Seq.UnicodeFromString("sqkmgu")).length)));
    };
    static fm37(p0, p1, globalState) {
      return _dafny.Seq.of((_dafny.Map.Empty.slice().updateUnsafe(true,false)).Merge(_dafny.Map.Empty.slice().updateUnsafe(false,false)));
    };
    static fm38(p0, p1, p2, p3, globalState) {
      return (_dafny.Set.fromElements(new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!(false),(_dafny.ZERO).minus(new BigNumber(-689)))).length), new BigNumber(-952), new BigNumber(228))).length))).Union((_dafny.Set.fromElements(new BigNumber(621))).Difference(_dafny.Set.fromElements(new BigNumber(-276))));
    };
    static fm39(p0, globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe(!(false),true)).Merge(_dafny.Map.Empty.slice().updateUnsafe(false,true))).Merge((_dafny.Map.Empty.slice().updateUnsafe(false,!(true))).Merge(_dafny.Map.Empty.slice().updateUnsafe(false,!(true))));
    };
    static fm40(p0, p1, p2, p3, globalState) {
      return ((_dafny.MultiSet.fromElements(!(!(true)), false)).Intersect(_dafny.MultiSet.fromElements(true))).Intersect((_dafny.MultiSet.fromElements(false)).Intersect(_dafny.MultiSet.fromElements(false)));
    };
    static fm41(p0, p1, globalState) {
      let _source1 = _module.D1.create_DC2(new BigNumber((_dafny.Seq.of(true, true)).length), true, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,true)).length));
      if (_source1.is_DC2) {
        let _26___mcc_h0 = (_source1).cf2;
        let _27___mcc_h1 = (_source1).cf3;
        let _28___mcc_h2 = (_source1).cf4;
        let _29_cf4 = _28___mcc_h2;
        let _30_cf3 = _27___mcc_h1;
        let _31_cf2 = _26___mcc_h0;
        return _module.D8.create_DC16(_29_cf4, _30_cf3, _30_cf3, _31_cf2);
      } else {
        let _32___mcc_h3 = (_source1).cf1;
        let _33_cf1 = _32___mcc_h3;
        return _module.D8.create_DC16(new BigNumber((_dafny.Set.fromElements(false)).length), false, true, new BigNumber((_dafny.Set.fromElements(!(false))).length));
      }
    };
    static fm42(p0, p1, p2, globalState) {
      return function () {
        let _coll2 = new _dafny.Map();
        for (const _compr_2 of ((function () {
          let _coll3 = new _dafny.Set();
          for (const _compr_3 of (function () {
            let _coll4 = new _dafny.Map();
            for (const _compr_4 of (_dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("sl"), _dafny.Seq.UnicodeFromString("qwc"), _dafny.Seq.UnicodeFromString("prgm"))).Elements) {
              let _34_v1 = _compr_4;
              if ((_dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("sl"), _dafny.Seq.UnicodeFromString("qwc"), _dafny.Seq.UnicodeFromString("prgm"))).contains(_34_v1)) {
                _coll4.push([_34_v1,false]);
              }
            }
            return _coll4;
          }()).Keys.Elements) {
            let _35_v2 = _compr_3;
            if ((function () {
              let _coll5 = new _dafny.Map();
              for (const _compr_5 of (_dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("sl"), _dafny.Seq.UnicodeFromString("qwc"), _dafny.Seq.UnicodeFromString("prgm"))).Elements) {
                let _34_v1 = _compr_5;
                if ((_dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("sl"), _dafny.Seq.UnicodeFromString("qwc"), _dafny.Seq.UnicodeFromString("prgm"))).contains(_34_v1)) {
                  _coll5.push([_34_v1,false]);
                }
              }
              return _coll5;
            }()).contains(_35_v2)) {
              _coll3.add(_35_v2);
            }
          }
          return _coll3;
        }()).Difference(_dafny.Set.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(580)), function (_36_i0) {
          return new _dafny.CodePoint('k'.codePointAt(0));
        }), _dafny.Seq.UnicodeFromString("kbeim")))).Elements) {
          let _37_v0 = _compr_2;
          if (((function () {
            let _coll6 = new _dafny.Set();
            for (const _compr_6 of (function () {
              let _coll7 = new _dafny.Map();
              for (const _compr_7 of (_dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("sl"), _dafny.Seq.UnicodeFromString("qwc"), _dafny.Seq.UnicodeFromString("prgm"))).Elements) {
                let _34_v1 = _compr_7;
                if ((_dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("sl"), _dafny.Seq.UnicodeFromString("qwc"), _dafny.Seq.UnicodeFromString("prgm"))).contains(_34_v1)) {
                  _coll7.push([_34_v1,false]);
                }
              }
              return _coll7;
            }()).Keys.Elements) {
              let _38_v2 = _compr_6;
              if ((function () {
                let _coll8 = new _dafny.Map();
                for (const _compr_8 of (_dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("sl"), _dafny.Seq.UnicodeFromString("qwc"), _dafny.Seq.UnicodeFromString("prgm"))).Elements) {
                  let _34_v1 = _compr_8;
                  if ((_dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("sl"), _dafny.Seq.UnicodeFromString("qwc"), _dafny.Seq.UnicodeFromString("prgm"))).contains(_34_v1)) {
                    _coll8.push([_34_v1,false]);
                  }
                }
                return _coll8;
              }()).contains(_38_v2)) {
                _coll6.add(_38_v2);
              }
            }
            return _coll6;
          }()).Difference(_dafny.Set.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(580)), function (_36_i0) {
            return new _dafny.CodePoint('k'.codePointAt(0));
          }), _dafny.Seq.UnicodeFromString("kbeim")))).contains(_37_v0)) {
            _coll2.push([_37_v0,(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('w'.codePointAt(0)),true)).Merge(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('i'.codePointAt(0)),false))]);
          }
        }
        return _coll2;
      }();
    };
    static fm43(p0, p1, p2, globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements(false),_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("snxgcwim")).length), new BigNumber(430)))).Merge(_dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.FromArray(_dafny.Seq.of(false, false)),_dafny.MultiSet.fromElements(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("ejalkcx")).length))).cardinality()))))).Merge(_dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements(!(false)),_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.of(new _dafny.CodePoint('w'.codePointAt(0)), new _dafny.CodePoint('r'.codePointAt(0)), new _dafny.CodePoint('y'.codePointAt(0)), new _dafny.CodePoint('p'.codePointAt(0)), new _dafny.CodePoint('r'.codePointAt(0)))).length))));
    };
    static fm44(globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.of(true, false), _dafny.Seq.of(true));
    };
    static fm45(p0, p1, globalState) {
      return _module.D9.create_DC20(!(!(!_dafny.areEqual(_dafny.Seq.Create(_module.__default.abs(new BigNumber(429)), function (_39_i0) {
  return new _dafny.CodePoint('n'.codePointAt(0));
}), _dafny.Seq.UnicodeFromString("fc")))));
    };
    static m0(p0, globalState) {
      let r0 = false;
      let r1 = false;
      let r2 = _dafny.ZERO;
      let _40_v0;
      _40_v0 = true;
      let _41_v1;
      _41_v1 = new _dafny.CodePoint('r'.codePointAt(0));
      let _42_v2;
      _42_v2 = _dafny.Seq.of(_41_v1, _41_v1);
      let _43_v3;
      _43_v3 = _dafny.Seq.of(_42_v2);
      let _44_v4;
      let _nw0 = Array((new BigNumber(13)).toNumber()).fill(false);
      _44_v4 = _nw0;
      let _45_v5;
      _45_v5 = _dafny.Map.Empty.slice().updateUnsafe(_40_v0,_44_v4);
      let _46_v6;
      let _nw1 = Array((new BigNumber(13)).toNumber()).fill(_dafny.Map.Empty);
      _46_v6 = _nw1;
      let _47_v7;
      let _nw2 = new _module.C4();
      _nw2.__ctor((_40_v0) === (_40_v0), _dafny.MultiSet.FromArray(_dafny.Seq.of((_dafny.ZERO).minus(_module.__default.fm17(_43_v3, new BigNumber((_45_v5).length), p0, new BigNumber((_42_v2).length), globalState)))), _46_v6);
      _47_v7 = _nw2;
      _40_v0 = (_47_v7).f7;
      let _48_v8;
      _48_v8 = _dafny.Seq.of((_47_v7).f7, (_47_v7).f7, _40_v0);
      let _49_v9;
      _49_v9 = _dafny.Map.Empty.slice().updateUnsafe((_47_v7).f7,new BigNumber((_dafny.MultiSet.FromArray(_48_v8)).cardinality()));
      let _50_v10;
      let _nw3 = new _module.C6();
      _nw3.__ctor((_49_v9).Merge((_dafny.Map.Empty.slice().updateUnsafe(_40_v0,p0)).update((_47_v7).f7, p0)), p0);
      _50_v10 = _nw3;
      let _51_v11;
      let _nw4 = new _module.C5();
      _nw4.__ctor(new _dafny.CodePoint('m'.codePointAt(0)));
      _51_v11 = _nw4;
      let _52_v12;
      _52_v12 = _dafny.Map.Empty.slice().updateUnsafe(_51_v11,_50_v10.f5);
      _52_v12 = (_52_v12).update(_51_v11, new BigNumber(821));
      (_50_v10).f5 = _module.__default.safeDivisionInt(p0, _50_v10.f5);
      let _53_v13;
      let _nw5 = new _module.C2();
      _nw5.__ctor(_40_v0, (_dafny.ZERO).minus(p0), _46_v6);
      _53_v13 = _nw5;
      let _54_v14;
      let _nw6 = Array((new BigNumber(20)).toNumber());
      _nw6[(_dafny.ZERO).toNumber()] = _53_v13;
      _nw6[(_dafny.ONE).toNumber()] = _53_v13;
      _nw6[(new BigNumber(2)).toNumber()] = _53_v13;
      _nw6[(new BigNumber(3)).toNumber()] = _53_v13;
      _nw6[(new BigNumber(4)).toNumber()] = _53_v13;
      _nw6[(new BigNumber(5)).toNumber()] = _53_v13;
      _nw6[(new BigNumber(6)).toNumber()] = _53_v13;
      _nw6[(new BigNumber(7)).toNumber()] = _53_v13;
      _nw6[(new BigNumber(8)).toNumber()] = _53_v13;
      _nw6[(new BigNumber(9)).toNumber()] = _53_v13;
      _nw6[(new BigNumber(10)).toNumber()] = _53_v13;
      _nw6[(new BigNumber(11)).toNumber()] = _53_v13;
      _nw6[(new BigNumber(12)).toNumber()] = _53_v13;
      _nw6[(new BigNumber(13)).toNumber()] = _53_v13;
      _nw6[(new BigNumber(14)).toNumber()] = _53_v13;
      _nw6[(new BigNumber(15)).toNumber()] = _53_v13;
      _nw6[(new BigNumber(16)).toNumber()] = _53_v13;
      _nw6[(new BigNumber(17)).toNumber()] = _53_v13;
      _nw6[(new BigNumber(18)).toNumber()] = _53_v13;
      _nw6[(new BigNumber(19)).toNumber()] = _53_v13;
      _54_v14 = _nw6;
      let _index0 = _module.__default.safeIndex(new BigNumber(813), new BigNumber((_54_v14).length));
      (_54_v14)[_index0] = _53_v13;
      let _index1 = _module.__default.safeIndex(new BigNumber(813), new BigNumber((_54_v14).length));
      (_54_v14)[_index1] = _53_v13;
      r0 = false;
      r1 = ((_53_v13).f12).isLessThan(p0);
      let _55_v15;
      _55_v15 = _module.D6.create_DC11(_dafny.Seq.UnicodeFromString("tw"));
      let _pat_let_tv0 = _42_v2;
      let _pat_let_tv1 = _42_v2;
      r2 = new BigNumber((function (_source2) {
        if (_source2.is_DC12) {
          let _56___mcc_h0 = (_source2).cf23;
          let _57___mcc_h1 = (_source2).cf24;
          let _58_cf24 = _57___mcc_h1;
          let _59_cf23 = _56___mcc_h0;
          return _pat_let_tv0;
        } else {
          let _60___mcc_h2 = (_source2).cf22;
          let _61_cf22 = _60___mcc_h2;
          return _pat_let_tv1;
        }
      }(_55_v15)).length);
      return [r0, r1, r2];
    }
    static Main(__noArgsParameter) {
      let _62_v0;
      _62_v0 = false;
      let _63_v1;
      _63_v1 = _dafny.Seq.of(_62_v0);
      let _64_v2;
      _64_v2 = _63_v1;
      let _65_v3;
      _65_v3 = new BigNumber(124);
      let _66_v4;
      let _nw7 = Array((new BigNumber(11)).toNumber());
      _nw7[(_dafny.ZERO).toNumber()] = _63_v1;
      _nw7[(_dafny.ONE).toNumber()] = _63_v1;
      _nw7[(new BigNumber(2)).toNumber()] = (_64_v2);
      _nw7[(new BigNumber(3)).toNumber()] = _63_v1;
      _nw7[(new BigNumber(4)).toNumber()] = _dafny.Seq.update(_63_v1, _module.__default.safeIndex(_65_v3, new BigNumber((_63_v1).length)), _62_v0);
      _nw7[(new BigNumber(5)).toNumber()] = _63_v1;
      _nw7[(new BigNumber(6)).toNumber()] = _63_v1;
      _nw7[(new BigNumber(7)).toNumber()] = _63_v1;
      _nw7[(new BigNumber(8)).toNumber()] = _63_v1;
      _nw7[(new BigNumber(9)).toNumber()] = _dafny.Seq.of(_62_v0);
      _nw7[(new BigNumber(10)).toNumber()] = _63_v1;
      _66_v4 = _nw7;
      let _67_globalState;
      let _nw8 = new _module.GlobalState();
      _nw8.__ctor(_66_v4, false, new BigNumber(47));
      _67_globalState = _nw8;
      let _68_v5;
      _68_v5 = _dafny.Set.fromElements(_65_v3);
      if ((_68_v5).IsProperSubsetOf(_68_v5)) {
        let _69_v6;
        let _70_v7;
        let _71_v8;
        let _out0;
        let _out1;
        let _out2;
        let _outcollector0 = _module.__default.m0(_65_v3, _67_globalState);
        _out0 = _outcollector0[0];
        _out1 = _outcollector0[1];
        _out2 = _outcollector0[2];
        _69_v6 = _out0;
        _70_v7 = _out1;
        _71_v8 = _out2;
        _62_v0 = _62_v0;
        if (!(_70_v7)) {
          let _72_v9;
          _72_v9 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-674),false);
          let _73_v11;
          _73_v11 = _dafny.Seq.of(_65_v3);
          let _74_v12;
          _74_v12 = _72_v9;
          let _75_v13;
          _75_v13 = new _dafny.CodePoint('v'.codePointAt(0));
          let _76_v14;
          _76_v14 = _dafny.Seq.of(_75_v13);
          let _77_v17;
          let _nw9 = Array((new BigNumber(17)).toNumber());
          _nw9[(_dafny.ZERO).toNumber()] = _72_v9;
          _nw9[(_dafny.ONE).toNumber()] = _72_v9;
          _nw9[(new BigNumber(2)).toNumber()] = _72_v9;
          _nw9[(new BigNumber(3)).toNumber()] = _72_v9;
          _nw9[(new BigNumber(4)).toNumber()] = _72_v9;
          _nw9[(new BigNumber(5)).toNumber()] = function () {
            let _coll9 = new _dafny.Map();
            for (const _compr_9 of (_73_v11).Elements) {
              let _78_v10 = _compr_9;
              if (_dafny.Seq.contains(_73_v11, _78_v10)) {
                _coll9.push([(_78_v10).multipliedBy(new BigNumber(789)),_70_v7]);
              }
            }
            return _coll9;
          }();
          _nw9[(new BigNumber(6)).toNumber()] = _72_v9;
          _nw9[(new BigNumber(7)).toNumber()] = (_74_v12);
          _nw9[(new BigNumber(8)).toNumber()] = _module.__default.fm34(_62_v0, _75_v13, _73_v11, _67_globalState);
          _nw9[(new BigNumber(9)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm17(_dafny.Seq.of(_76_v14, _76_v14), _71_v8, _71_v8, _71_v8, _67_globalState),_69_v6);
          _nw9[(new BigNumber(10)).toNumber()] = _72_v9;
          _nw9[(new BigNumber(11)).toNumber()] = _72_v9;
          _nw9[(new BigNumber(12)).toNumber()] = function () {
            let _coll10 = new _dafny.Map();
            for (const _compr_10 of (_dafny.Map.Empty.slice().updateUnsafe(_71_v8,_71_v8)).Keys.Elements) {
              let _79_v15 = _compr_10;
              if ((_dafny.Map.Empty.slice().updateUnsafe(_71_v8,_71_v8)).contains(_79_v15)) {
                _coll10.push([(_79_v15).plus(_65_v3),_62_v0]);
              }
            }
            return _coll10;
          }();
          _nw9[(new BigNumber(13)).toNumber()] = function () {
            let _coll11 = new _dafny.Map();
            for (const _compr_11 of (_73_v11).Elements) {
              let _80_v16 = _compr_11;
              if (_dafny.Seq.contains(_73_v11, _80_v16)) {
                _coll11.push([(_80_v16).minus(_71_v8),_70_v7]);
              }
            }
            return _coll11;
          }();
          _nw9[(new BigNumber(14)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_65_v3,_70_v7);
          _nw9[(new BigNumber(15)).toNumber()] = _72_v9;
          _nw9[(new BigNumber(16)).toNumber()] = _72_v9;
          _77_v17 = _nw9;
          let _81_v18;
          let _nw10 = new _module.C2();
          _nw10.__ctor(_69_v6, (_dafny.ZERO).minus(_65_v3), _77_v17);
          _81_v18 = _nw10;
          let _82_v19;
          _82_v19 = _module.D6.create_DC11(_76_v14);
          let _83_v20;
          let _nw11 = Array((new BigNumber(7)).toNumber());
          _nw11[(_dafny.ZERO).toNumber()] = _82_v19;
          _nw11[(_dafny.ONE).toNumber()] = _82_v19;
          _nw11[(new BigNumber(2)).toNumber()] = _82_v19;
          _nw11[(new BigNumber(3)).toNumber()] = _82_v19;
          _nw11[(new BigNumber(4)).toNumber()] = _82_v19;
          _nw11[(new BigNumber(5)).toNumber()] = _82_v19;
          _nw11[(new BigNumber(6)).toNumber()] = _82_v19;
          _83_v20 = _nw11;
          let _rhs0 = (_81_v18).f12;
          let _rhs1 = _65_v3;
          let _rhs2 = ((_62_v0) ? (_83_v20) : (_83_v20));
          let _rhs3 = (_65_v3).plus(_65_v3);
          let _lhs0 = _67_globalState;
          let _lhs1 = _67_globalState;
          _lhs0.f2 = _rhs0;
          _lhs1.f2 = _rhs1;
          _83_v20 = _rhs2;
          _65_v3 = _rhs3;
          _69_v6 = true;
          let _84_v21;
          _84_v21 = _dafny.Map.Empty.slice().updateUnsafe((_81_v18).f11,(_81_v18).f12);
          let _85_v22;
          _85_v22 = _dafny.Map.Empty.slice().updateUnsafe(!(_69_v6),(((_84_v21).contains((_81_v18).f11)) ? ((_84_v21).get((_81_v18).f11)) : (_71_v8)));
          _71_v8 = (((_85_v22).contains(true)) ? ((_85_v22).get(true)) : (new BigNumber(325)));
          _85_v22 = (_85_v22).update(_dafny.Seq.IsProperPrefixOf(_dafny.Seq.Create(_module.__default.abs(new BigNumber(218)), ((_86_v13) => function (_87_i0) {
            return _86_v13;
          })(_75_v13)), _76_v14), _71_v8);
        } else {
          let _88_v23;
          let _nw12 = Array((new BigNumber(6)).toNumber());
          _88_v23 = _nw12;
          _88_v23 = _88_v23;
          let _89_v24;
          _89_v24 = _dafny.Seq.of(_65_v3, _71_v8, _71_v8);
          let _90_v25;
          _90_v25 = _dafny.Seq.UnicodeFromString("ro");
          let _91_v26;
          _91_v26 = new _dafny.CodePoint('j'.codePointAt(0));
          let _92_v27;
          let _nw13 = Array((new BigNumber(9)).toNumber());
          _nw13[(_dafny.ZERO).toNumber()] = ((_62_v0) ? (_70_v7) : (_62_v0));
          _nw13[(_dafny.ONE).toNumber()] = _69_v6;
          _nw13[(new BigNumber(2)).toNumber()] = _module.__default.fm21(new BigNumber((_90_v25).length), _91_v26, false, _70_v7, _67_globalState);
          _nw13[(new BigNumber(3)).toNumber()] = _70_v7;
          _nw13[(new BigNumber(4)).toNumber()] = false;
          _nw13[(new BigNumber(5)).toNumber()] = ((_69_v6) ? (_module.__default.fm21(_71_v8, _91_v26, _70_v7, !(_70_v7), _67_globalState)) : (_69_v6));
          _nw13[(new BigNumber(6)).toNumber()] = _62_v0;
          _nw13[(new BigNumber(7)).toNumber()] = !(((_70_v7) ? (!(_module.__default.fm2(new _dafny.CodePoint('f'.codePointAt(0)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(745)), ((_93_v7) => function (_94_i1) {
            return _dafny.Map.Empty.slice().updateUnsafe(_93_v7,_93_v7);
          })(_70_v7)), _70_v7, _67_globalState))) : (_70_v7)));
          _nw13[(new BigNumber(8)).toNumber()] = !(true) || (_70_v7);
          _92_v27 = _nw13;
          let _rhs4 = _89_v24;
          let _rhs5 = ((_69_v6) ? ((_69_v6) === (_70_v7)) : (_62_v0));
          let _rhs6 = _92_v27;
          let _rhs7 = !_dafny.areEqual(_89_v24, _89_v24);
          let _rhs8 = _65_v3;
          _89_v24 = _rhs4;
          _69_v6 = _rhs5;
          _92_v27 = _rhs6;
          _69_v6 = _rhs7;
          _71_v8 = _rhs8;
          let _95_v28;
          let _nw14 = Array((new BigNumber(13)).toNumber()).fill(_dafny.Map.Empty);
          _95_v28 = _nw14;
          let _96_v29;
          let _nw15 = new _module.C1();
          _nw15.__ctor(_71_v8, _95_v28);
          _96_v29 = _nw15;
          let _97_v30;
          let _nw16 = Array((new BigNumber(14)).toNumber());
          _nw16[(_dafny.ZERO).toNumber()] = _64_v2;
          _nw16[(_dafny.ONE).toNumber()] = _64_v2;
          _nw16[(new BigNumber(2)).toNumber()] = _64_v2;
          _nw16[(new BigNumber(3)).toNumber()] = _64_v2;
          _nw16[(new BigNumber(4)).toNumber()] = _module.__default.fm44(_67_globalState);
          _nw16[(new BigNumber(5)).toNumber()] = _module.__default.fm44(_67_globalState);
          _nw16[(new BigNumber(6)).toNumber()] = _64_v2;
          _nw16[(new BigNumber(7)).toNumber()] = _64_v2;
          _nw16[(new BigNumber(8)).toNumber()] = _64_v2;
          _nw16[(new BigNumber(9)).toNumber()] = _module.__default.fm44(_67_globalState);
          _nw16[(new BigNumber(10)).toNumber()] = _64_v2;
          _nw16[(new BigNumber(11)).toNumber()] = _64_v2;
          _nw16[(new BigNumber(12)).toNumber()] = _64_v2;
          _nw16[(new BigNumber(13)).toNumber()] = _64_v2;
          _97_v30 = _nw16;
          let _index2 = _module.__default.safeIndex(new BigNumber(154), new BigNumber((_97_v30).length));
          (_97_v30)[_index2] = _64_v2;
          let _index3 = _module.__default.safeIndex(new BigNumber(154), new BigNumber((_97_v30).length));
          let _rhs9 = _64_v2;
          let _rhs10 = new BigNumber((_dafny.Seq.Concat(_90_v25, _90_v25)).length);
          let _rhs11 = (_96_v29).f14;
          let _rhs12 = (_63_v1)[_module.__default.safeIndex((_96_v29).f14, new BigNumber((_63_v1).length))];
          let _rhs13 = _91_v26;
          let _lhs2 = _97_v30;
          let _lhs3 = _module.__default.safeIndex(new BigNumber(154), new BigNumber((_97_v30).length));
          let _lhs4 = _67_globalState;
          _lhs2[_lhs3] = _rhs9;
          _65_v3 = _rhs10;
          _lhs4.f2 = _rhs11;
          _70_v7 = _rhs12;
          _91_v26 = _rhs13;
          let _98_v31;
          _98_v31 = _dafny.Seq.of(_90_v25);
          let _99_v32;
          _99_v32 = _dafny.Set.fromElements(_92_v27, _92_v27);
          let _100_v33;
          _100_v33 = _dafny.Set.fromElements(_96_v29);
          _69_v6 = ((new BigNumber((_90_v25).length)).minus(_module.__default.fm17(_98_v31, (_89_v24)[_module.__default.safeIndex((_96_v29).f14, new BigNumber((_89_v24).length))], new BigNumber((_99_v32).length), (_96_v29).f14, _67_globalState))).isLessThanOrEqualTo(_module.__default.fm17(_dafny.Seq.of(_90_v25, _90_v25, _dafny.Seq.update(_90_v25, _module.__default.safeIndex(new BigNumber(653), new BigNumber((_90_v25).length)), new _dafny.CodePoint('b'.codePointAt(0)))), (_96_v29).f14, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_65_v3,new BigNumber((_100_v33).length))).length), _65_v3, _67_globalState));
        }
        _62_v0 = _70_v7;
        let _101_v34;
        let _102_v35;
        let _103_v36;
        let _out3;
        let _out4;
        let _out5;
        let _outcollector1 = _module.__default.m0((_dafny.ZERO).minus(_65_v3), _67_globalState);
        _out3 = _outcollector1[0];
        _out4 = _outcollector1[1];
        _out5 = _outcollector1[2];
        _101_v34 = _out3;
        _102_v35 = _out4;
        _103_v36 = _out5;
      } else {
        _62_v0 = _62_v0;
        let _104_v37;
        _104_v37 = _dafny.Map.Empty.slice().updateUnsafe(_62_v0,!(_62_v0) || (_62_v0));
        let _105_v38;
        _105_v38 = new _dafny.CodePoint('b'.codePointAt(0));
        _104_v37 = (_104_v37).update(_module.__default.fm21(_65_v3, _105_v38, _62_v0, false, _67_globalState), _62_v0);
        let _106_v39;
        _106_v39 = _dafny.Seq.UnicodeFromString("fgobb");
        _106_v39 = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(440)), ((_107_v38) => function (_108_i2) {
          return _107_v38;
        })(_105_v38)), _106_v39);
        _62_v0 = ((_62_v0) ? ((_68_v5).IsDisjointFrom(_68_v5)) : (_62_v0));
        if (_62_v0) {
          _62_v0 = (false) && (_62_v0);
          let _109_v40;
          _109_v40 = _dafny.Map.Empty.slice().updateUnsafe(_65_v3,_65_v3);
          let _110_v41;
          let _111_v42;
          let _112_v43;
          let _out6;
          let _out7;
          let _out8;
          let _outcollector2 = _module.__default.m0((((_109_v40).contains(_65_v3)) ? ((_109_v40).get(_65_v3)) : (_65_v3)), _67_globalState);
          _out6 = _outcollector2[0];
          _out7 = _outcollector2[1];
          _out8 = _outcollector2[2];
          _110_v41 = _out6;
          _111_v42 = _out7;
          _112_v43 = _out8;
          _110_v41 = _62_v0;
          let _113_v44;
          _113_v44 = _dafny.MultiSet.fromElements(_110_v41);
          let _114_v45;
          _114_v45 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_106_v39).length),(_113_v44).update(_110_v41, _module.__default.abs(_65_v3)));
          _114_v45 = (_114_v45).update(_65_v3, _113_v44);
          let _115_v46;
          _115_v46 = _dafny.MultiSet.fromElements(_112_v43, _65_v3, _112_v43, _65_v3);
          let _116_v47;
          let _nw17 = new _module.C3();
          _nw17.__ctor(_115_v46, _110_v41);
          _116_v47 = _nw17;
          let _117_v48;
          _117_v48 = _dafny.Seq.of(_116_v47);
          let _118_v49;
          _118_v49 = _dafny.Set.fromElements(_109_v40);
          let _119_v50;
          let _nw18 = Array((new BigNumber(6)).toNumber());
          _nw18[(_dafny.ZERO).toNumber()] = _116_v47;
          _nw18[(_dafny.ONE).toNumber()] = _116_v47;
          _nw18[(new BigNumber(2)).toNumber()] = (_117_v48)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.of(new BigNumber((_118_v49).length), _112_v43)).length), new BigNumber((_117_v48).length))];
          _nw18[(new BigNumber(3)).toNumber()] = _116_v47;
          _nw18[(new BigNumber(4)).toNumber()] = _116_v47;
          _nw18[(new BigNumber(5)).toNumber()] = _116_v47;
          _119_v50 = _nw18;
          _119_v50 = _119_v50;
        } else {
          let _120_v51;
          let _nw19 = Array((new BigNumber(7)).toNumber()).fill(_dafny.ZERO);
          _120_v51 = _nw19;
          let _121_v52;
          let _nw20 = Array((new BigNumber(10)).toNumber());
          _nw20[(_dafny.ZERO).toNumber()] = _62_v0;
          _nw20[(_dafny.ONE).toNumber()] = !(_62_v0);
          _nw20[(new BigNumber(2)).toNumber()] = _62_v0;
          _nw20[(new BigNumber(3)).toNumber()] = _module.__default.fm21(_65_v3, _105_v38, _62_v0, _62_v0, _67_globalState);
          _nw20[(new BigNumber(4)).toNumber()] = _62_v0;
          _nw20[(new BigNumber(5)).toNumber()] = _62_v0;
          _nw20[(new BigNumber(6)).toNumber()] = _module.__default.fm23(_106_v39, _65_v3, _65_v3, _62_v0, _67_globalState);
          _nw20[(new BigNumber(7)).toNumber()] = false;
          _nw20[(new BigNumber(8)).toNumber()] = _62_v0;
          _nw20[(new BigNumber(9)).toNumber()] = _62_v0;
          _121_v52 = _nw20;
          let _122_v53;
          _122_v53 = _dafny.Set.fromElements(_121_v52, _121_v52, _121_v52, _121_v52, _121_v52);
          let _index4 = _module.__default.safeIndex(new BigNumber(699), new BigNumber((_120_v51).length));
          (_120_v51)[_index4] = new BigNumber((_122_v53).length);
          let _123_v54;
          _123_v54 = _dafny.Seq.of(_module.__default.fm25(_62_v0, _65_v3, _67_globalState));
          let _124_v55;
          _124_v55 = _dafny.Map.Empty.slice().updateUnsafe(_62_v0,_65_v3);
          let _125_v56;
          _125_v56 = _dafny.Seq.of(_124_v55);
          let _126_v57;
          _126_v57 = _module.D4.create_DC9(new BigNumber((_125_v56).length), _105_v38, _dafny.Seq.of(_65_v3), new BigNumber((_dafny.Seq.UnicodeFromString("muudfjey")).length), _module.__default.fm23(_106_v39, _65_v3, new BigNumber((_dafny.MultiSet.fromElements(_module.__default.fm17(_123_v54, _65_v3, new BigNumber(-758), _65_v3, _67_globalState))).cardinality()), false, _67_globalState));
          let _index5 = _module.__default.safeIndex(new BigNumber(699), new BigNumber((_120_v51).length));
          (_120_v51)[_index5] = _module.__default.fm17(_123_v54, new BigNumber((_106_v39).length), (_126_v57).dtor_cf19, _65_v3, _67_globalState);
          let _index6 = _module.__default.safeIndex(new BigNumber(689), new BigNumber((_121_v52).length));
          (_121_v52)[_index6] = _62_v0;
          let _index7 = _module.__default.safeIndex(new BigNumber(689), new BigNumber((_121_v52).length));
          (_121_v52)[_index7] = _62_v0;
          _124_v55 = (_124_v55).update(_62_v0, (_120_v51)[_module.__default.safeIndex(new BigNumber(699), new BigNumber((_120_v51).length))]);
          let _127_v58;
          _127_v58 = _dafny.Map.Empty.slice().updateUnsafe(_105_v38,(new BigNumber((_106_v39).length)).plus(_65_v3));
          _127_v58 = (_127_v58).update(_105_v38, (_120_v51)[_module.__default.safeIndex(new BigNumber(699), new BigNumber((_120_v51).length))]);
          let _128_v59;
          let _init0 = ((_129_v3, _130_v0) => function (_131_i3) {
            return _dafny.Map.Empty.slice().updateUnsafe(_129_v3,_130_v0);
          })(_65_v3, _62_v0);
          let _nw21 = Array((_dafny.ONE).toNumber());
          for (let _i0_0 = 0; _i0_0 < new BigNumber(_nw21.length); _i0_0++) {
            _nw21[_i0_0] = _init0(new BigNumber(_i0_0));
          }
          _128_v59 = _nw21;
          let _132_v60;
          let _nw22 = new _module.C7();
          _nw22.__ctor(_128_v59);
          _132_v60 = _nw22;
          _132_v60 = _132_v60;
        }
      }
      let _hi0 = (_65_v3).plus(_65_v3);
      for (let _133_i4 = _65_v3; _133_i4.isLessThan(_hi0); _133_i4 = _133_i4.plus(_dafny.ONE)) {
        let _134_v61;
        _134_v61 = new _dafny.CodePoint('k'.codePointAt(0));
        let _135_v62;
        _135_v62 = _dafny.MultiSet.fromElements(_134_v61, _module.__default.fm5(_65_v3, _67_globalState));
        let _136_v63;
        _136_v63 = _module.D1.create_DC2(new BigNumber((_135_v62).cardinality()), _62_v0, new BigNumber(502));
        let _137_v64;
        _137_v64 = _dafny.Set.fromElements(false, false);
        let _138_v65;
        _138_v65 = _dafny.Seq.of(_137_v64, _137_v64);
        let _139_v66;
        let _nw23 = Array((_dafny.ONE).toNumber());
        _nw23[(_dafny.ZERO).toNumber()] = new BigNumber(((_138_v65)[_module.__default.safeIndex((_dafny.ZERO).minus((_dafny.ZERO).minus(_65_v3)), new BigNumber((_138_v65).length))]).length);
        _139_v66 = _nw23;
        let _140_v67;
        _140_v67 = _dafny.Seq.UnicodeFromString("qm");
        let _141_v68;
        _141_v68 = _dafny.Map.Empty.slice().updateUnsafe(_139_v66,new BigNumber((_140_v67).length));
        let _142_v69;
        let _nw24 = Array((new BigNumber(14)).toNumber());
        _nw24[(_dafny.ZERO).toNumber()] = _62_v0;
        _nw24[(_dafny.ONE).toNumber()] = ((_62_v0) ? (_62_v0) : (_62_v0));
        _nw24[(new BigNumber(2)).toNumber()] = _62_v0;
        _nw24[(new BigNumber(3)).toNumber()] = (_136_v63).dtor_cf3;
        _nw24[(new BigNumber(4)).toNumber()] = _62_v0;
        _nw24[(new BigNumber(5)).toNumber()] = _62_v0;
        _nw24[(new BigNumber(6)).toNumber()] = _62_v0;
        _nw24[(new BigNumber(7)).toNumber()] = _62_v0;
        _nw24[(new BigNumber(8)).toNumber()] = _62_v0;
        _nw24[(new BigNumber(9)).toNumber()] = _62_v0;
        _nw24[(new BigNumber(10)).toNumber()] = false;
        _nw24[(new BigNumber(11)).toNumber()] = !(_141_v68).contains(_139_v66);
        _nw24[(new BigNumber(12)).toNumber()] = true;
        _nw24[(new BigNumber(13)).toNumber()] = _62_v0;
        _142_v69 = _nw24;
        let _index8 = _module.__default.safeIndex(new BigNumber(422), new BigNumber((_142_v69).length));
        (_142_v69)[_index8] = _62_v0;
        let _index9 = _module.__default.safeIndex(new BigNumber(422), new BigNumber((_142_v69).length));
        (_142_v69)[_index9] = _62_v0;
        _62_v0 = !(_62_v0);
        let _143_v70;
        _143_v70 = _dafny.MultiSet.fromElements(new BigNumber((_140_v67).length));
        let _144_v71;
        let _nw25 = new _module.C3();
        _nw25.__ctor(_143_v70, true);
        _144_v71 = _nw25;
        let _145_v72;
        _145_v72 = _dafny.Map.Empty.slice().updateUnsafe(_62_v0,(_144_v71).f10);
        let _index10 = _module.__default.safeIndex(new BigNumber(422), new BigNumber((_142_v69).length));
        (_142_v69)[_index10] = (((_145_v72).contains(_62_v0)) ? ((_145_v72).get(_62_v0)) : ((_142_v69)[_module.__default.safeIndex(new BigNumber(422), new BigNumber((_142_v69).length))]));
      }
      _62_v0 = _62_v0;
      let _146_v73;
      let _nw26 = Array((new BigNumber(29)).toNumber()).fill(false);
      _146_v73 = _nw26;
      _146_v73 = _146_v73;
      let _147_v74;
      _147_v74 = _dafny.MultiSet.fromElements(_62_v0);
      let _148_i5;
      _148_i5 = _dafny.ZERO;
      L0: {
        while (!(((_62_v0) ? (_147_v74) : (_147_v74))).equals((_147_v74).Union(_147_v74))) {
          C0: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_148_i5)) {
              break L0;
            }
            _148_i5 = (_148_i5).plus(_dafny.ONE);
            let _149_v75;
            _149_v75 = _dafny.Seq.of(_68_v5);
            let _150_v76;
            _150_v76 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeModuloInt(new BigNumber((_149_v75).length), new BigNumber((_63_v1).length)),_dafny.Seq.of(_65_v3, _65_v3));
            _150_v76 = _150_v76;
            let _151_v77;
            _151_v77 = _dafny.Map.Empty.slice().updateUnsafe(_62_v0,_65_v3);
            let _152_v78;
            _152_v78 = _dafny.Seq.of(new BigNumber(742));
            let _153_v79;
            _153_v79 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_152_v78).length),_62_v0);
            let _154_v80;
            _154_v80 = _153_v79;
            let _155_v81;
            _155_v81 = _dafny.MultiSet.fromElements(_154_v80, _154_v80);
            let _156_v82;
            _156_v82 = _dafny.Map.Empty.slice().updateUnsafe(_152_v78,_155_v81);
            let _157_v83;
            let _nw27 = new _module.C6();
            _nw27.__ctor(_151_v77, new BigNumber(((_156_v82).Merge(_156_v82)).length));
            _157_v83 = _nw27;
            let _158_v84;
            _158_v84 = new _dafny.CodePoint('m'.codePointAt(0));
            _158_v84 = _158_v84;
            _63_v1 = _63_v1;
          }
        }
      }
      let _159_v85;
      _159_v85 = new _dafny.CodePoint('y'.codePointAt(0));
      let _160_v86;
      let _nw28 = new _module.C5();
      _nw28.__ctor(_159_v85);
      _160_v86 = _nw28;
      let _161_i6;
      _161_i6 = _dafny.ZERO;
      L1: {
        while (!(_62_v0)) {
          C1: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_161_i6)) {
              break L1;
            }
            _161_i6 = (_161_i6).plus(_dafny.ONE);
            let _index11 = _module.__default.safeIndex(new BigNumber(724), new BigNumber((_146_v73).length));
            (_146_v73)[_index11] = _62_v0;
            let _index12 = _module.__default.safeIndex(new BigNumber(724), new BigNumber((_146_v73).length));
            (_146_v73)[_index12] = _62_v0;
            (_67_globalState).f2 = _65_v3;
            let _162_v87;
            let _init1 = ((_163_v3, _164_v0) => function (_165_i7) {
              return _dafny.Map.Empty.slice().updateUnsafe(_163_v3,_164_v0);
            })(_65_v3, _62_v0);
            let _nw29 = Array((new BigNumber(11)).toNumber());
            for (let _i0_1 = 0; _i0_1 < new BigNumber(_nw29.length); _i0_1++) {
              _nw29[_i0_1] = _init1(new BigNumber(_i0_1));
            }
            _162_v87 = _nw29;
            let _166_v88;
            let _nw30 = new _module.C2();
            _nw30.__ctor(_62_v0, (_dafny.ZERO).minus((_dafny.ZERO).minus(_65_v3)), _162_v87);
            _166_v88 = _nw30;
            let _167_v89;
            let _init2 = ((_168_v3) => function (_169_i8) {
              return (_169_i8).multipliedBy(_168_v3);
            })(_65_v3);
            let _nw31 = Array((new BigNumber(10)).toNumber());
            for (let _i0_2 = 0; _i0_2 < new BigNumber(_nw31.length); _i0_2++) {
              _nw31[_i0_2] = _init2(new BigNumber(_i0_2));
            }
            _167_v89 = _nw31;
            let _170_v90;
            _170_v90 = _dafny.Map.Empty.slice().updateUnsafe(_65_v3,_167_v89);
            let _171_v91;
            _171_v91 = _module.D2.create_DC5(new BigNumber((_170_v90).length), (_146_v73)[_module.__default.safeIndex(new BigNumber(724), new BigNumber((_146_v73).length))], _62_v0, !(_62_v0), (_146_v73)[_module.__default.safeIndex(new BigNumber(724), new BigNumber((_146_v73).length))]);
            let _172_v92;
            _172_v92 = _dafny.Seq.of(_171_v91);
            let _173_v93;
            _173_v93 = _module.D9.create_DC19(_166_v88, new BigNumber((_dafny.MultiSet.fromElements(_172_v92)).cardinality()), _65_v3);
            let _174_v94;
            let _nw32 = Array((new BigNumber(5)).toNumber());
            _nw32[(_dafny.ZERO).toNumber()] = _166_v88;
            _nw32[(_dafny.ONE).toNumber()] = _166_v88;
            _nw32[(new BigNumber(2)).toNumber()] = (_173_v93).dtor_cf34;
            _nw32[(new BigNumber(3)).toNumber()] = _166_v88;
            _nw32[(new BigNumber(4)).toNumber()] = _166_v88;
            _174_v94 = _nw32;
            let _175_v95;
            _175_v95 = _dafny.Seq.of(_166_v88, _166_v88);
            let _index13 = _module.__default.safeIndex(new BigNumber(160), new BigNumber((_174_v94).length));
            (_174_v94)[_index13] = (_175_v95)[_module.__default.safeIndex(_65_v3, new BigNumber((_175_v95).length))];
            let _index14 = _module.__default.safeIndex(new BigNumber(160), new BigNumber((_174_v94).length));
            (_174_v94)[_index14] = (_173_v93).dtor_cf34;
            let _176_v96;
            _176_v96 = _dafny.Seq.UnicodeFromString("jfj");
            _176_v96 = _dafny.Seq.Concat(_176_v96, _dafny.Seq.Concat(_176_v96, _dafny.Seq.UnicodeFromString("wye")));
          }
        }
      }
      let _177_v97;
      _177_v97 = _dafny.Set.fromElements(_62_v0, _62_v0, _62_v0, false, _62_v0);
      let _178_v98;
      _178_v98 = _177_v97;
      let _179_v99;
      _179_v99 = _dafny.Map.Empty.slice().updateUnsafe(_62_v0,_178_v98);
      let _rhs14 = (_dafny.Map.Empty.slice().updateUnsafe(_62_v0,_178_v98)).Merge(_dafny.Map.Empty.slice().updateUnsafe(_62_v0,_178_v98));
      let _rhs15 = _65_v3;
      let _lhs5 = _67_globalState;
      _179_v99 = _rhs14;
      _lhs5.f2 = _rhs15;
      let _180_v100;
      _180_v100 = _dafny.Map.Empty.slice().updateUnsafe(_65_v3,_62_v0);
      let _181_v101;
      _181_v101 = _dafny.Seq.of(_159_v85);
      let _182_v102;
      _182_v102 = _dafny.MultiSet.fromElements(_65_v3, new BigNumber((_181_v101).length), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(652)), function (_183_i9) {
        return new _dafny.CodePoint('i'.codePointAt(0));
      })).length), _65_v3);
      let _184_v103;
      _184_v103 = _dafny.Seq.of(new BigNumber((_182_v102).cardinality()), new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_159_v85,_182_v102)).length))).length));
      let _185_v106;
      _185_v106 = _180_v100;
      let _186_v107;
      _186_v107 = _dafny.Map.Empty.slice().updateUnsafe(_62_v0,_65_v3);
      let _187_v108;
      let _nw33 = new _module.C6();
      _nw33.__ctor(_186_v107, _65_v3);
      _187_v108 = _nw33;
      let _188_v109;
      let _nw34 = new _module.C0();
      _nw34.__ctor(_module.__default.fm9(_65_v3, _67_globalState));
      _188_v109 = _nw34;
      let _189_v110;
      _189_v110 = _dafny.Seq.of(_188_v109, _188_v109);
      let _190_v111;
      _190_v111 = _dafny.Map.Empty.slice().updateUnsafe(_187_v108,(_dafny.ZERO).minus(new BigNumber((_189_v110).length)));
      let _191_v112;
      let _nw35 = Array((new BigNumber(29)).toNumber());
      _nw35[(_dafny.ZERO).toNumber()] = _180_v100;
      _nw35[(_dafny.ONE).toNumber()] = _180_v100;
      _nw35[(new BigNumber(2)).toNumber()] = _180_v100;
      _nw35[(new BigNumber(3)).toNumber()] = _180_v100;
      _nw35[(new BigNumber(4)).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe(_65_v3,_62_v0)).update(new BigNumber(147), _62_v0);
      _nw35[(new BigNumber(5)).toNumber()] = _180_v100;
      _nw35[(new BigNumber(6)).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe(_65_v3,_62_v0)).update(_65_v3, _62_v0);
      _nw35[(new BigNumber(7)).toNumber()] = _module.__default.fm34(_62_v0, (_160_v86).f6, _184_v103, _67_globalState);
      _nw35[(new BigNumber(8)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_181_v101).length),_62_v0);
      _nw35[(new BigNumber(9)).toNumber()] = _180_v100;
      _nw35[(new BigNumber(10)).toNumber()] = _180_v100;
      _nw35[(new BigNumber(11)).toNumber()] = function () {
        let _coll12 = new _dafny.Map();
        for (const _compr_12 of (_182_v102).Elements) {
          let _192_v104 = _compr_12;
          if ((_182_v102).contains(_192_v104)) {
            _coll12.push([(_192_v104).plus(_65_v3),_62_v0]);
          }
        }
        return _coll12;
      }();
      _nw35[(new BigNumber(12)).toNumber()] = _180_v100;
      _nw35[(new BigNumber(13)).toNumber()] = _180_v100;
      _nw35[(new BigNumber(14)).toNumber()] = (_180_v100).update(_65_v3, _62_v0);
      _nw35[(new BigNumber(15)).toNumber()] = _180_v100;
      _nw35[(new BigNumber(16)).toNumber()] = function () {
        let _coll13 = new _dafny.Map();
        for (const _compr_13 of _dafny.IntegerRange(new BigNumber(80), new BigNumber(355))) {
          let _193_v105 = _compr_13;
          if (((new BigNumber(80)).isLessThanOrEqualTo(_193_v105)) && ((_193_v105).isLessThan(new BigNumber(355)))) {
            _coll13.push([(_193_v105).plus(_65_v3),_62_v0]);
          }
        }
        return _coll13;
      }();
      _nw35[(new BigNumber(17)).toNumber()] = _module.__default.fm34(_62_v0, _159_v85, _184_v103, _67_globalState);
      _nw35[(new BigNumber(18)).toNumber()] = (_185_v106);
      _nw35[(new BigNumber(19)).toNumber()] = _180_v100;
      _nw35[(new BigNumber(20)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_65_v3,_62_v0);
      _nw35[(new BigNumber(21)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe((((_190_v111).contains(_187_v108)) ? ((_190_v111).get(_187_v108)) : (new BigNumber(684))),_62_v0);
      _nw35[(new BigNumber(22)).toNumber()] = _180_v100;
      _nw35[(new BigNumber(23)).toNumber()] = _180_v100;
      _nw35[(new BigNumber(24)).toNumber()] = _180_v100;
      _nw35[(new BigNumber(25)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.fromElements(_65_v3)).cardinality()),_62_v0);
      _nw35[(new BigNumber(26)).toNumber()] = _180_v100;
      _nw35[(new BigNumber(27)).toNumber()] = _180_v100;
      _nw35[(new BigNumber(28)).toNumber()] = _180_v100;
      _191_v112 = _nw35;
      let _194_v113;
      let _nw36 = new _module.C7();
      _nw36.__ctor(_191_v112);
      _194_v113 = _nw36;
      let _195_v114;
      let _nw37 = new _module.C7();
      _nw37.__ctor(_191_v112);
      _195_v114 = _nw37;
      let _196_v115;
      _196_v115 = _dafny.Seq.of(_195_v114, _195_v114);
      let _197_v116;
      let _nw38 = new _module.C2();
      _nw38.__ctor(_62_v0, (_184_v103)[_module.__default.safeIndex(new BigNumber((_196_v115).length), new BigNumber((_184_v103).length))], _191_v112);
      _197_v116 = _nw38;
      let _198_v117;
      let _199_v118;
      let _out9;
      let _out10;
      let _outcollector3 = (_194_v113).m1(_dafny.Seq.Create(_module.__default.abs(new BigNumber(119)), ((_200_v3) => function (_201_i10) {
        return _200_v3;
      })(_65_v3)), ((_188_v109).f13).isEqualTo((_197_v116).f12), _181_v101, _67_globalState);
      _out9 = _outcollector3[0];
      _out10 = _outcollector3[1];
      _198_v117 = _out9;
      _199_v118 = _out10;
      let _index15 = _module.__default.safeIndex(new BigNumber(902), new BigNumber((_146_v73).length));
      (_146_v73)[_index15] = (_dafny.Set.fromElements((_188_v109).f13)).IsSubsetOf(_68_v5);
      let _index16 = _module.__default.safeIndex(new BigNumber(902), new BigNumber((_146_v73).length));
      (_146_v73)[_index16] = _module.__default.fm21(_module.__default.safeDivisionInt(new BigNumber(71), _198_v117), _199_v118, !((_197_v116).f11) || (_62_v0), _62_v0, _67_globalState);
      _62_v0 = !(((_197_v116).f11) === ((_147_v74).IsSubsetOf(_147_v74)));
      let _202_v119;
      _202_v119 = _module.D8.create_DC15(_146_v73);
      let _203_v120;
      _203_v120 = _dafny.MultiSet.fromElements(_202_v119);
      let _204_v121;
      _204_v121 = _dafny.Seq.of(_182_v102, _182_v102);
      let _205_v122;
      _205_v122 = _dafny.Map.Empty.slice().updateUnsafe((_146_v73)[_module.__default.safeIndex(new BigNumber(902), new BigNumber((_146_v73).length))],_62_v0);
      let _206_v123;
      _206_v123 = _dafny.MultiSet.fromElements(_184_v103);
      let _207_v124;
      _207_v124 = _dafny.Map.Empty.slice().updateUnsafe(_181_v101,(_dafny.ZERO).minus(new BigNumber((_147_v74).cardinality())));
      let _208_v125;
      let _nw39 = Array((new BigNumber(9)).toNumber());
      _nw39[(_dafny.ZERO).toNumber()] = new BigNumber(((((_146_v73)[_module.__default.safeIndex(new BigNumber(902), new BigNumber((_146_v73).length))]) ? (_203_v120) : (_203_v120))).cardinality());
      _nw39[(_dafny.ONE).toNumber()] = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.update(_204_v121, _module.__default.safeIndex((_197_v116).f12, new BigNumber((_204_v121).length)), _dafny.MultiSet.fromElements(_198_v117, (_197_v116).f12)), _204_v121)).length);
      _nw39[(new BigNumber(2)).toNumber()] = ((_62_v0) ? ((_197_v116).f12) : (new BigNumber(((_182_v102).update(new BigNumber((_205_v122).length), _module.__default.abs(_198_v117))).cardinality())));
      _nw39[(new BigNumber(3)).toNumber()] = (((_206_v123).contains(_dafny.Seq.of(_65_v3))) ? ((_206_v123).get(_dafny.Seq.of(_65_v3))) : (_65_v3));
      _nw39[(new BigNumber(4)).toNumber()] = (((_207_v124).contains(_181_v101)) ? ((_207_v124).get(_181_v101)) : (_187_v108.f5));
      _nw39[(new BigNumber(5)).toNumber()] = (_65_v3).multipliedBy((_184_v103)[_module.__default.safeIndex(new BigNumber((_181_v101).length), new BigNumber((_184_v103).length))]);
      _nw39[(new BigNumber(6)).toNumber()] = (_188_v109).f13;
      _nw39[(new BigNumber(7)).toNumber()] = new BigNumber(689);
      _nw39[(new BigNumber(8)).toNumber()] = (_188_v109).f13;
      _208_v125 = _nw39;
      let _index17 = _module.__default.safeIndex(new BigNumber(288), new BigNumber((_208_v125).length));
      (_208_v125)[_index17] = (_188_v109).f13;
      let _209_v126;
      _209_v126 = _dafny.Map.Empty.slice().updateUnsafe(_208_v125,(_146_v73)[_module.__default.safeIndex(new BigNumber(902), new BigNumber((_146_v73).length))]);
      let _index18 = _module.__default.safeIndex(new BigNumber(288), new BigNumber((_208_v125).length));
      (_208_v125)[_index18] = new BigNumber((_209_v126).length);
      let _source3 = _module.D3.create_DC7(new BigNumber((_63_v1).length), _module.__default.safeDivisionInt(_187_v108.f5, (_208_v125)[_module.__default.safeIndex(new BigNumber(288), new BigNumber((_208_v125).length))]));
      if (_source3.is_DC7) {
        let _210___mcc_h0 = (_source3).cf13;
        let _211___mcc_h1 = (_source3).cf14;
        let _212_cf14 = _211___mcc_h1;
        let _213_cf13 = _210___mcc_h0;
        let _214_v127;
        let _215_v128;
        let _out11;
        let _out12;
        let _outcollector4 = (_194_v113).m1(_module.__default.fm12(!((_197_v116).f11), _187_v108.f5, _177_v97, _dafny.Map.Empty.slice().updateUnsafe((_197_v116).f12,new BigNumber((_68_v5).length)), _67_globalState), (_197_v116).f11, _181_v101, _67_globalState);
        _out11 = _outcollector4[0];
        _out12 = _outcollector4[1];
        _214_v127 = _out11;
        _215_v128 = _out12;
        let _index19 = _module.__default.safeIndex(new BigNumber(902), new BigNumber((_146_v73).length));
        (_146_v73)[_index19] = (_146_v73)[_module.__default.safeIndex(new BigNumber(902), new BigNumber((_146_v73).length))];
        let _216_v130;
        _216_v130 = _dafny.Map.Empty.slice().updateUnsafe(_212_cf14,_213_cf13);
        let _217_v131;
        _217_v131 = _dafny.Map.Empty.slice().updateUnsafe((_module.__default.fm24(function () {
          let _coll14 = new _dafny.Set();
          for (const _compr_14 of _dafny.IntegerRange(new BigNumber(676), new BigNumber(163))) {
            let _218_v129 = _compr_14;
            if (((new BigNumber(676)).isLessThanOrEqualTo(_218_v129)) && ((_218_v129).isLessThan(new BigNumber(163)))) {
              _coll14.add(_module.__default.safeModuloInt(_218_v129, new BigNumber((_177_v97).length)));
            }
          }
          return _coll14;
        }(), _67_globalState)).multipliedBy(new BigNumber((_216_v130).length)),(_188_v109).f13);
        _216_v130 = _217_v131;
        let _219_v132;
        let _nw40 = Array((new BigNumber(10)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
        _219_v132 = _nw40;
        let _220_v133;
        _220_v133 = _module.D9.create_DC18(_219_v132);
        let _index20 = _module.__default.safeIndex(new BigNumber(288), new BigNumber((_208_v125).length));
        let _rhs16 = _220_v133;
        let _rhs17 = _module.__default.safeModuloInt(new BigNumber((_181_v101).length), ((_197_v116).f12).plus(_212_cf14));
        let _rhs18 = _146_v73;
        let _rhs19 = ((new BigNumber(332)).plus(new BigNumber((_181_v101).length))).plus(_187_v108.f5);
        let _lhs6 = _208_v125;
        let _lhs7 = _module.__default.safeIndex(new BigNumber(288), new BigNumber((_208_v125).length));
        _220_v133 = _rhs16;
        _213_cf13 = _rhs17;
        _146_v73 = _rhs18;
        _lhs6[_lhs7] = _rhs19;
      } else {
        let _221___mcc_h2 = (_source3).cf12;
        let _222_cf12 = _221___mcc_h2;
        _182_v102 = (_182_v102).Intersect((_182_v102).Union(_182_v102));
        _68_v5 = (((_197_v116).f11) ? (_68_v5) : (_68_v5));
        _65_v3 = ((_62_v0) ? (_module.__default.safeModuloInt(new BigNumber((_dafny.Seq.of(_197_v116)).length), _187_v108.f5)) : ((_197_v116).f12));
        _62_v0 = true;
      }
      let _223_v134;
      _223_v134 = _module.D4.create_DC8(_188_v109);
      let _source4 = _223_v134;
      if (_source4.is_DC9) {
        let _224___mcc_h3 = (_source4).cf16;
        let _225___mcc_h4 = (_source4).cf17;
        let _226___mcc_h5 = (_source4).cf18;
        let _227___mcc_h6 = (_source4).cf19;
        let _228___mcc_h7 = (_source4).cf20;
        let _229_cf20 = _228___mcc_h7;
        let _230_cf19 = _227___mcc_h6;
        let _231_cf18 = _226___mcc_h5;
        let _232_cf17 = _225___mcc_h4;
        let _233_cf16 = _224___mcc_h3;
        _62_v0 = _229_cf20;
        _198_v117 = (_194_v113).fm0((_188_v109).f13, _67_globalState);
        let _234_v135;
        let _235_v136;
        let _out13;
        let _out14;
        let _outcollector5 = (_197_v116).m1(_dafny.Seq.of(new BigNumber((_63_v1).length)), _module.__default.fm23(_181_v101, new BigNumber((_182_v102).cardinality()), (_197_v116).f12, _229_cf20, _67_globalState), _module.__default.fm25(true, _187_v108.f5, _67_globalState), _67_globalState);
        _out13 = _outcollector5[0];
        _out14 = _outcollector5[1];
        _234_v135 = _out13;
        _235_v136 = _out14;
        if (_62_v0) {
          _186_v107 = (_186_v107).update((_197_v116).f11, _65_v3);
          let _236_v137;
          let _init3 = ((_237_cf18) => function (_238_i11) {
            return _237_cf18;
          })(_231_cf18);
          let _nw41 = Array((new BigNumber(16)).toNumber());
          for (let _i0_3 = 0; _i0_3 < new BigNumber(_nw41.length); _i0_3++) {
            _nw41[_i0_3] = _init3(new BigNumber(_i0_3));
          }
          _236_v137 = _nw41;
          let _index21 = _module.__default.safeIndex(new BigNumber(642), new BigNumber((_236_v137).length));
          (_236_v137)[_index21] = _184_v103;
          let _index22 = _module.__default.safeIndex(new BigNumber(642), new BigNumber((_236_v137).length));
          let _index23 = _module.__default.safeIndex(new BigNumber(902), new BigNumber((_146_v73).length));
          let _rhs20 = (((_187_v108).f4).Merge(_186_v107)).Merge((_187_v108).f4);
          let _rhs21 = _198_v117;
          let _rhs22 = (_188_v109).f13;
          let _rhs23 = _231_cf18;
          let _rhs24 = !(!((_dafny.MultiSet.FromArray(_63_v1)).update((_197_v116).f11, _module.__default.abs(new BigNumber((((_205_v122).update(false, (_146_v73)[_module.__default.safeIndex(new BigNumber(902), new BigNumber((_146_v73).length))])).update(_62_v0, (_197_v116).f11)).length)))).equals(_dafny.MultiSet.FromArray(_63_v1)));
          let _lhs8 = _67_globalState;
          let _lhs9 = _236_v137;
          let _lhs10 = _module.__default.safeIndex(new BigNumber(642), new BigNumber((_236_v137).length));
          let _lhs11 = _146_v73;
          let _lhs12 = _module.__default.safeIndex(new BigNumber(902), new BigNumber((_146_v73).length));
          _186_v107 = _rhs20;
          _lhs8.f2 = _rhs21;
          _65_v3 = _rhs22;
          _lhs9[_lhs10] = _rhs23;
          _lhs11[_lhs12] = _rhs24;
          let _239_v138;
          _239_v138 = _dafny.Map.Empty.slice().updateUnsafe(_194_v113,(_236_v137)[_module.__default.safeIndex(new BigNumber(642), new BigNumber((_236_v137).length))]);
          let _240_v139;
          _240_v139 = _dafny.Seq.of((_236_v137)[_module.__default.safeIndex(new BigNumber(642), new BigNumber((_236_v137).length))], (_236_v137)[_module.__default.safeIndex(new BigNumber(642), new BigNumber((_236_v137).length))]);
          _239_v138 = ((_dafny.Map.Empty.slice().updateUnsafe(_194_v113,_184_v103)).update(_194_v113, (_240_v139)[_module.__default.safeIndex((_188_v109).f13, new BigNumber((_240_v139).length))])).Merge(_239_v138);
          let _241_v140;
          _241_v140 = _module.D9.create_DC20((_197_v116).f11);
          _241_v140 = _module.__default.fm45((_147_v74).IsProperSubsetOf(_dafny.MultiSet.fromElements(_module.__default.fm21(_234_v135, new _dafny.CodePoint('f'.codePointAt(0)), (_197_v116).f11, !((_197_v116).f11), _67_globalState))), (_197_v116).f12, _67_globalState);
          let _242_v141;
          _242_v141 = _dafny.Map.Empty.slice().updateUnsafe(_198_v117,_181_v101);
          _208_v125 = (((_160_v86).fm7(new BigNumber(380), (_197_v116).f11, _242_v141, _67_globalState)) ? (_208_v125) : (_208_v125));
        } else {
          let _243_v142;
          _243_v142 = _dafny.Seq.of(_dafny.Seq.of(false, (((_180_v100).contains((_197_v116).f12)) ? ((_180_v100).get((_197_v116).f12)) : (_229_cf20))), (((_146_v73)[_module.__default.safeIndex(new BigNumber(902), new BigNumber((_146_v73).length))]) ? (_63_v1) : (_63_v1)), _63_v1);
          let _244_v143;
          let _nw42 = new _module.C2();
          _nw42.__ctor((_146_v73)[_module.__default.safeIndex(new BigNumber(902), new BigNumber((_146_v73).length))], _187_v108.f5, _191_v112);
          _244_v143 = _nw42;
          let _245_v144;
          _245_v144 = _dafny.Map.Empty.slice().updateUnsafe(_244_v143,_243_v142);
          _243_v142 = (((_245_v144).contains(_244_v143)) ? ((_245_v144).get(_244_v143)) : (_243_v142));
          _235_v136 = (_160_v86).f6;
          _189_v110 = _189_v110;
          _194_v113 = (((_182_v102).IsProperSubsetOf(_182_v102)) ? (_194_v113) : (_194_v113));
          let _246_v145;
          let _nw43 = Array((new BigNumber(12)).toNumber());
          _nw43[(_dafny.ZERO).toNumber()] = (_181_v101)[_module.__default.safeIndex((_188_v109).f13, new BigNumber((_181_v101).length))];
          _nw43[(_dafny.ONE).toNumber()] = _159_v85;
          _nw43[(new BigNumber(2)).toNumber()] = (_181_v101)[_module.__default.safeIndex(new BigNumber((_181_v101).length), new BigNumber((_181_v101).length))];
          _nw43[(new BigNumber(3)).toNumber()] = _159_v85;
          _nw43[(new BigNumber(4)).toNumber()] = (_160_v86).f6;
          _nw43[(new BigNumber(5)).toNumber()] = _199_v118;
          _nw43[(new BigNumber(6)).toNumber()] = ((!(true)) ? ((_module.D3.create_DC6(_199_v118)).dtor_cf12) : ((_160_v86).f6));
          _nw43[(new BigNumber(7)).toNumber()] = (_160_v86).f6;
          _nw43[(new BigNumber(8)).toNumber()] = (_160_v86).f6;
          _nw43[(new BigNumber(9)).toNumber()] = _232_cf17;
          _nw43[(new BigNumber(10)).toNumber()] = (_160_v86).f6;
          _nw43[(new BigNumber(11)).toNumber()] = (_160_v86).f6;
          _246_v145 = _nw43;
          let _index24 = _module.__default.safeIndex(new BigNumber(825), new BigNumber((_246_v145).length));
          (_246_v145)[_index24] = (_160_v86).f6;
          let _247_v146;
          _247_v146 = _dafny.Map.Empty.slice().updateUnsafe(_62_v0,_177_v97);
          let _248_v147;
          _248_v147 = (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_181_v101).length),new BigNumber(((((_247_v146).contains(_229_cf20)) ? ((_247_v146).get(_229_cf20)) : (_177_v97))).length))).update(_187_v108.f5, _65_v3);
          let _index25 = _module.__default.safeIndex(new BigNumber(131), new BigNumber((_246_v145).length));
          (_246_v145)[_index25] = ((_229_cf20) ? (_235_v136) : ((_160_v86).f6));
          let _index26 = _module.__default.safeIndex(new BigNumber(825), new BigNumber((_246_v145).length));
          let _index27 = _module.__default.safeIndex(new BigNumber(131), new BigNumber((_246_v145).length));
          let _rhs25 = (_197_v116).f11;
          let _rhs26 = _199_v118;
          let _rhs27 = _248_v147;
          let _rhs28 = (_181_v101)[_module.__default.safeIndex(_module.__default.safeDivisionInt(new BigNumber(-851), new BigNumber((_dafny.Set.fromElements(_dafny.Seq.UnicodeFromString("lbovk"), _dafny.Seq.UnicodeFromString("kydcrfrb"))).length)), new BigNumber((_181_v101).length))];
          let _lhs13 = _246_v145;
          let _lhs14 = _module.__default.safeIndex(new BigNumber(825), new BigNumber((_246_v145).length));
          let _lhs15 = _246_v145;
          let _lhs16 = _module.__default.safeIndex(new BigNumber(131), new BigNumber((_246_v145).length));
          _62_v0 = _rhs25;
          _lhs13[_lhs14] = _rhs26;
          _248_v147 = _rhs27;
          _lhs15[_lhs16] = _rhs28;
        }
      } else {
        let _249___mcc_h8 = (_source4).cf15;
        let _250_cf15 = _249___mcc_h8;
        if (((_250_cf15).f13).isEqualTo((new BigNumber(-485)).minus((_188_v109).f13))) {
          _204_v121 = _204_v121;
          (_67_globalState).f2 = (_dafny.ZERO).minus(_module.__default.safeDivisionInt(_module.__default.fm9(new BigNumber(836), _67_globalState), (_208_v125)[_module.__default.safeIndex(new BigNumber(288), new BigNumber((_208_v125).length))]));
          (_187_v108).f5 = (((_186_v107).contains(((true) ? ((_146_v73)[_module.__default.safeIndex(new BigNumber(902), new BigNumber((_146_v73).length))]) : (_module.__default.fm2(new _dafny.CodePoint('r'.codePointAt(0)), _dafny.Seq.of(_205_v122), false, _67_globalState))))) ? ((_186_v107).get(((true) ? ((_146_v73)[_module.__default.safeIndex(new BigNumber(902), new BigNumber((_146_v73).length))]) : (_module.__default.fm2(new _dafny.CodePoint('r'.codePointAt(0)), _dafny.Seq.of(_205_v122), false, _67_globalState))))) : ((_198_v117).plus((_250_cf15).f13)));
          (_67_globalState).f2 = _198_v117;
          _65_v3 = _module.__default.fm9(((_250_cf15).f13).multipliedBy((_197_v116).f12), _67_globalState);
        } else {
          let _index28 = _module.__default.safeIndex(new BigNumber(902), new BigNumber((_146_v73).length));
          (_146_v73)[_index28] = true;
          _68_v5 = (_68_v5).Difference(_68_v5);
          let _251_v148;
          let _nw44 = Array((new BigNumber(7)).toNumber());
          _251_v148 = _nw44;
          _251_v148 = _251_v148;
          let _252_v149;
          _252_v149 = _dafny.Seq.of(_197_v116, _197_v116);
          let _253_v150;
          _253_v150 = _module.D13.create_DC24(_197_v116);
          let _254_v151;
          let _nw45 = Array((new BigNumber(23)).toNumber());
          _nw45[(_dafny.ZERO).toNumber()] = _197_v116;
          _nw45[(_dafny.ONE).toNumber()] = _197_v116;
          _nw45[(new BigNumber(2)).toNumber()] = _197_v116;
          _nw45[(new BigNumber(3)).toNumber()] = (_252_v149)[_module.__default.safeIndex((_197_v116).f12, new BigNumber((_252_v149).length))];
          _nw45[(new BigNumber(4)).toNumber()] = _197_v116;
          _nw45[(new BigNumber(5)).toNumber()] = _197_v116;
          _nw45[(new BigNumber(6)).toNumber()] = _197_v116;
          _nw45[(new BigNumber(7)).toNumber()] = _197_v116;
          _nw45[(new BigNumber(8)).toNumber()] = (((_146_v73)[_module.__default.safeIndex(new BigNumber(902), new BigNumber((_146_v73).length))]) ? (_197_v116) : (_197_v116));
          _nw45[(new BigNumber(9)).toNumber()] = _197_v116;
          _nw45[(new BigNumber(10)).toNumber()] = _197_v116;
          _nw45[(new BigNumber(11)).toNumber()] = _197_v116;
          _nw45[(new BigNumber(12)).toNumber()] = _197_v116;
          _nw45[(new BigNumber(13)).toNumber()] = _197_v116;
          _nw45[(new BigNumber(14)).toNumber()] = _197_v116;
          _nw45[(new BigNumber(15)).toNumber()] = (((_197_v116).f11) ? (_197_v116) : ((_253_v150).dtor_cf41));
          _nw45[(new BigNumber(16)).toNumber()] = _197_v116;
          _nw45[(new BigNumber(17)).toNumber()] = (_252_v149)[_module.__default.safeIndex((_208_v125)[_module.__default.safeIndex(new BigNumber(288), new BigNumber((_208_v125).length))], new BigNumber((_252_v149).length))];
          _nw45[(new BigNumber(18)).toNumber()] = _197_v116;
          _nw45[(new BigNumber(19)).toNumber()] = _197_v116;
          _nw45[(new BigNumber(20)).toNumber()] = _197_v116;
          _nw45[(new BigNumber(21)).toNumber()] = _197_v116;
          _nw45[(new BigNumber(22)).toNumber()] = _197_v116;
          _254_v151 = _nw45;
          let _index29 = _module.__default.safeIndex(new BigNumber(385), new BigNumber((_254_v151).length));
          (_254_v151)[_index29] = _197_v116;
          let _index30 = _module.__default.safeIndex(new BigNumber(385), new BigNumber((_254_v151).length));
          (_254_v151)[_index30] = _197_v116;
          let _255_v152;
          let _nw46 = new _module.C7();
          _nw46.__ctor(_191_v112);
          _255_v152 = _nw46;
        }
        let _256_v153;
        let _nw47 = new _module.C4();
        _nw47.__ctor((_146_v73)[_module.__default.safeIndex(new BigNumber(902), new BigNumber((_146_v73).length))], _182_v102, _191_v112);
        _256_v153 = _nw47;
        let _257_v154;
        _257_v154 = _dafny.Map.Empty.slice().updateUnsafe(_256_v153,_62_v0);
        _181_v101 = _dafny.Seq.Concat(_dafny.Seq.update(_dafny.Seq.Concat(_181_v101, _181_v101), _module.__default.safeIndex(new BigNumber((_257_v154).length), new BigNumber((_dafny.Seq.Concat(_181_v101, _181_v101)).length)), (_160_v86).f6), _181_v101);
        let _258_v155;
        _258_v155 = _dafny.Seq.of(_181_v101);
        let _259_v156;
        _259_v156 = _dafny.Map.Empty.slice().updateUnsafe(_198_v117,_module.__default.fm17(_258_v155, _module.__default.fm9((_188_v109).f13, _67_globalState), (_197_v116).f12, (_197_v116).f12, _67_globalState));
        let _index31 = _module.__default.safeIndex(new BigNumber(288), new BigNumber((_208_v125).length));
        (_208_v125)[_index31] = (((_259_v156).contains((_188_v109).f13)) ? ((_259_v156).get((_188_v109).f13)) : (new BigNumber(-30)));
        let _260_v157;
        let _261_v158;
        let _out15;
        let _out16;
        let _outcollector6 = (_187_v108).m2((_197_v116).f11, _67_globalState);
        _out15 = _outcollector6[0];
        _out16 = _outcollector6[1];
        _260_v157 = _out15;
        _261_v158 = _out16;
      }
      process.stdout.write(_dafny.toString(_62_v0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_63_v1, _dafny.Seq.of(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_64_v2), _dafny.Seq.of(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_65_v3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_66_v4)[_dafny.ZERO], _dafny.Seq.of(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_66_v4)[_dafny.ONE], _dafny.Seq.of(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_66_v4)[new BigNumber(2)], _dafny.Seq.of(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_66_v4)[new BigNumber(3)], _dafny.Seq.of(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_66_v4)[new BigNumber(4)], _dafny.Seq.of(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_66_v4)[new BigNumber(5)], _dafny.Seq.of(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_66_v4)[new BigNumber(6)], _dafny.Seq.of(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_66_v4)[new BigNumber(7)], _dafny.Seq.of(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_66_v4)[new BigNumber(8)], _dafny.Seq.of(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_66_v4)[new BigNumber(9)], _dafny.Seq.of(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_66_v4)[new BigNumber(10)], _dafny.Seq.of(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(((_67_globalState).f0)[_dafny.ZERO], _dafny.Seq.of(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(((_67_globalState).f0)[_dafny.ONE], _dafny.Seq.of(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(((_67_globalState).f0)[new BigNumber(2)], _dafny.Seq.of(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(((_67_globalState).f0)[new BigNumber(3)], _dafny.Seq.of(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(((_67_globalState).f0)[new BigNumber(4)], _dafny.Seq.of(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(((_67_globalState).f0)[new BigNumber(5)], _dafny.Seq.of(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(((_67_globalState).f0)[new BigNumber(6)], _dafny.Seq.of(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(((_67_globalState).f0)[new BigNumber(7)], _dafny.Seq.of(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(((_67_globalState).f0)[new BigNumber(8)], _dafny.Seq.of(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(((_67_globalState).f0)[new BigNumber(9)], _dafny.Seq.of(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(((_67_globalState).f0)[new BigNumber(10)], _dafny.Seq.of(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_67_globalState).f1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_67_globalState.f2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_68_v5).equals(_dafny.Set.fromElements())));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_146_v73)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_146_v73)[new BigNumber(28)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_147_v74).equals(_dafny.MultiSet.fromElements(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_148_i5));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_159_v85));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_160_v86).f6));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_161_i6));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_177_v97).equals(_dafny.Set.fromElements(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_178_v98)).equals(_dafny.Set.fromElements(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_179_v99).equals(_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.Set.fromElements(false)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_180_v100).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(124),false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_181_v101, _dafny.Seq.of(new _dafny.CodePoint('y'.codePointAt(0)), new _dafny.CodePoint('y'.codePointAt(0)), new _dafny.CodePoint('y'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_182_v102).equals(_dafny.MultiSet.fromElements(new BigNumber(124), new BigNumber(124), _dafny.ONE, new BigNumber(652)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_184_v103, _dafny.Seq.of(new BigNumber(4), _dafny.ONE))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_185_v106)).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(124),false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_186_v107).equals(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(124)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_187_v108).f4).equals(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(124)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_187_v108.f5));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_188_v109).f13));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_189_v110).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_190_v111).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_191_v112)[_dafny.ZERO]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(124),false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_191_v112)[_dafny.ONE]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(124),false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_191_v112)[new BigNumber(2)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(124),false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_191_v112)[new BigNumber(3)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(124),false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_191_v112)[new BigNumber(4)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(124),false).updateUnsafe(new BigNumber(147),false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_191_v112)[new BigNumber(5)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(124),false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_191_v112)[new BigNumber(6)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(124),false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_191_v112)[new BigNumber(7)]).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,true).updateUnsafe(new BigNumber(2),false).updateUnsafe(new BigNumber(449),false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_191_v112)[new BigNumber(8)]).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_191_v112)[new BigNumber(9)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(124),false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_191_v112)[new BigNumber(10)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(124),false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_191_v112)[new BigNumber(11)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(248),false).updateUnsafe(new BigNumber(125),false).updateUnsafe(new BigNumber(776),false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_191_v112)[new BigNumber(12)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(124),false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_191_v112)[new BigNumber(13)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(124),false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_191_v112)[new BigNumber(14)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(124),false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_191_v112)[new BigNumber(15)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(124),false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_191_v112)[new BigNumber(16)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(204),false).updateUnsafe(new BigNumber(205),false).updateUnsafe(new BigNumber(206),false).updateUnsafe(new BigNumber(207),false).updateUnsafe(new BigNumber(208),false).updateUnsafe(new BigNumber(209),false).updateUnsafe(new BigNumber(210),false).updateUnsafe(new BigNumber(211),false).updateUnsafe(new BigNumber(212),false).updateUnsafe(new BigNumber(213),false).updateUnsafe(new BigNumber(214),false).updateUnsafe(new BigNumber(215),false).updateUnsafe(new BigNumber(216),false).updateUnsafe(new BigNumber(217),false).updateUnsafe(new BigNumber(218),false).updateUnsafe(new BigNumber(219),false).updateUnsafe(new BigNumber(220),false).updateUnsafe(new BigNumber(221),false).updateUnsafe(new BigNumber(222),false).updateUnsafe(new BigNumber(223),false).updateUnsafe(new BigNumber(224),false).updateUnsafe(new BigNumber(225),false).updateUnsafe(new BigNumber(226),false).updateUnsafe(new BigNumber(227),false).updateUnsafe(new BigNumber(228),false).updateUnsafe(new BigNumber(229),false).updateUnsafe(new BigNumber(230),false).updateUnsafe(new BigNumber(231),false).updateUnsafe(new BigNumber(232),false).updateUnsafe(new BigNumber(233),false).updateUnsafe(new BigNumber(234),false).updateUnsafe(new BigNumber(235),false).updateUnsafe(new BigNumber(236),false).updateUnsafe(new BigNumber(237),false).updateUnsafe(new BigNumber(238),false).updateUnsafe(new BigNumber(239),false).updateUnsafe(new BigNumber(240),false).updateUnsafe(new BigNumber(241),false).updateUnsafe(new BigNumber(242),false).updateUnsafe(new BigNumber(243),false).updateUnsafe(new BigNumber(244),false).updateUnsafe(new BigNumber(245),false).updateUnsafe(new BigNumber(246),false).updateUnsafe(new BigNumber(247),false).updateUnsafe(new BigNumber(248),false).updateUnsafe(new BigNumber(249),false).updateUnsafe(new BigNumber(250),false).updateUnsafe(new BigNumber(251),false).updateUnsafe(new BigNumber(252),false).updateUnsafe(new BigNumber(253),false).updateUnsafe(new BigNumber(254),false).updateUnsafe(new BigNumber(255),false).updateUnsafe(new BigNumber(256),false).updateUnsafe(new BigNumber(257),false).updateUnsafe(new BigNumber(258),false).updateUnsafe(new BigNumber(259),false).updateUnsafe(new BigNumber(260),false).updateUnsafe(new BigNumber(261),false).updateUnsafe(new BigNumber(262),false).updateUnsafe(new BigNumber(263),false).updateUnsafe(new BigNumber(264),false).updateUnsafe(new BigNumber(265),false).updateUnsafe(new BigNumber(266),false).updateUnsafe(new BigNumber(267),false).updateUnsafe(new BigNumber(268),false).updateUnsafe(new BigNumber(269),false).updateUnsafe(new BigNumber(270),false).updateUnsafe(new BigNumber(271),false).updateUnsafe(new BigNumber(272),false).updateUnsafe(new BigNumber(273),false).updateUnsafe(new BigNumber(274),false).updateUnsafe(new BigNumber(275),false).updateUnsafe(new BigNumber(276),false).updateUnsafe(new BigNumber(277),false).updateUnsafe(new BigNumber(278),false).updateUnsafe(new BigNumber(279),false).updateUnsafe(new BigNumber(280),false).updateUnsafe(new BigNumber(281),false).updateUnsafe(new BigNumber(282),false).updateUnsafe(new BigNumber(283),false).updateUnsafe(new BigNumber(284),false).updateUnsafe(new BigNumber(285),false).updateUnsafe(new BigNumber(286),false).updateUnsafe(new BigNumber(287),false).updateUnsafe(new BigNumber(288),false).updateUnsafe(new BigNumber(289),false).updateUnsafe(new BigNumber(290),false).updateUnsafe(new BigNumber(291),false).updateUnsafe(new BigNumber(292),false).updateUnsafe(new BigNumber(293),false).updateUnsafe(new BigNumber(294),false).updateUnsafe(new BigNumber(295),false).updateUnsafe(new BigNumber(296),false).updateUnsafe(new BigNumber(297),false).updateUnsafe(new BigNumber(298),false).updateUnsafe(new BigNumber(299),false).updateUnsafe(new BigNumber(300),false).updateUnsafe(new BigNumber(301),false).updateUnsafe(new BigNumber(302),false).updateUnsafe(new BigNumber(303),false).updateUnsafe(new BigNumber(304),false).updateUnsafe(new BigNumber(305),false).updateUnsafe(new BigNumber(306),false).updateUnsafe(new BigNumber(307),false).updateUnsafe(new BigNumber(308),false).updateUnsafe(new BigNumber(309),false).updateUnsafe(new BigNumber(310),false).updateUnsafe(new BigNumber(311),false).updateUnsafe(new BigNumber(312),false).updateUnsafe(new BigNumber(313),false).updateUnsafe(new BigNumber(314),false).updateUnsafe(new BigNumber(315),false).updateUnsafe(new BigNumber(316),false).updateUnsafe(new BigNumber(317),false).updateUnsafe(new BigNumber(318),false).updateUnsafe(new BigNumber(319),false).updateUnsafe(new BigNumber(320),false).updateUnsafe(new BigNumber(321),false).updateUnsafe(new BigNumber(322),false).updateUnsafe(new BigNumber(323),false).updateUnsafe(new BigNumber(324),false).updateUnsafe(new BigNumber(325),false).updateUnsafe(new BigNumber(326),false).updateUnsafe(new BigNumber(327),false).updateUnsafe(new BigNumber(328),false).updateUnsafe(new BigNumber(329),false).updateUnsafe(new BigNumber(330),false).updateUnsafe(new BigNumber(331),false).updateUnsafe(new BigNumber(332),false).updateUnsafe(new BigNumber(333),false).updateUnsafe(new BigNumber(334),false).updateUnsafe(new BigNumber(335),false).updateUnsafe(new BigNumber(336),false).updateUnsafe(new BigNumber(337),false).updateUnsafe(new BigNumber(338),false).updateUnsafe(new BigNumber(339),false).updateUnsafe(new BigNumber(340),false).updateUnsafe(new BigNumber(341),false).updateUnsafe(new BigNumber(342),false).updateUnsafe(new BigNumber(343),false).updateUnsafe(new BigNumber(344),false).updateUnsafe(new BigNumber(345),false).updateUnsafe(new BigNumber(346),false).updateUnsafe(new BigNumber(347),false).updateUnsafe(new BigNumber(348),false).updateUnsafe(new BigNumber(349),false).updateUnsafe(new BigNumber(350),false).updateUnsafe(new BigNumber(351),false).updateUnsafe(new BigNumber(352),false).updateUnsafe(new BigNumber(353),false).updateUnsafe(new BigNumber(354),false).updateUnsafe(new BigNumber(355),false).updateUnsafe(new BigNumber(356),false).updateUnsafe(new BigNumber(357),false).updateUnsafe(new BigNumber(358),false).updateUnsafe(new BigNumber(359),false).updateUnsafe(new BigNumber(360),false).updateUnsafe(new BigNumber(361),false).updateUnsafe(new BigNumber(362),false).updateUnsafe(new BigNumber(363),false).updateUnsafe(new BigNumber(364),false).updateUnsafe(new BigNumber(365),false).updateUnsafe(new BigNumber(366),false).updateUnsafe(new BigNumber(367),false).updateUnsafe(new BigNumber(368),false).updateUnsafe(new BigNumber(369),false).updateUnsafe(new BigNumber(370),false).updateUnsafe(new BigNumber(371),false).updateUnsafe(new BigNumber(372),false).updateUnsafe(new BigNumber(373),false).updateUnsafe(new BigNumber(374),false).updateUnsafe(new BigNumber(375),false).updateUnsafe(new BigNumber(376),false).updateUnsafe(new BigNumber(377),false).updateUnsafe(new BigNumber(378),false).updateUnsafe(new BigNumber(379),false).updateUnsafe(new BigNumber(380),false).updateUnsafe(new BigNumber(381),false).updateUnsafe(new BigNumber(382),false).updateUnsafe(new BigNumber(383),false).updateUnsafe(new BigNumber(384),false).updateUnsafe(new BigNumber(385),false).updateUnsafe(new BigNumber(386),false).updateUnsafe(new BigNumber(387),false).updateUnsafe(new BigNumber(388),false).updateUnsafe(new BigNumber(389),false).updateUnsafe(new BigNumber(390),false).updateUnsafe(new BigNumber(391),false).updateUnsafe(new BigNumber(392),false).updateUnsafe(new BigNumber(393),false).updateUnsafe(new BigNumber(394),false).updateUnsafe(new BigNumber(395),false).updateUnsafe(new BigNumber(396),false).updateUnsafe(new BigNumber(397),false).updateUnsafe(new BigNumber(398),false).updateUnsafe(new BigNumber(399),false).updateUnsafe(new BigNumber(400),false).updateUnsafe(new BigNumber(401),false).updateUnsafe(new BigNumber(402),false).updateUnsafe(new BigNumber(403),false).updateUnsafe(new BigNumber(404),false).updateUnsafe(new BigNumber(405),false).updateUnsafe(new BigNumber(406),false).updateUnsafe(new BigNumber(407),false).updateUnsafe(new BigNumber(408),false).updateUnsafe(new BigNumber(409),false).updateUnsafe(new BigNumber(410),false).updateUnsafe(new BigNumber(411),false).updateUnsafe(new BigNumber(412),false).updateUnsafe(new BigNumber(413),false).updateUnsafe(new BigNumber(414),false).updateUnsafe(new BigNumber(415),false).updateUnsafe(new BigNumber(416),false).updateUnsafe(new BigNumber(417),false).updateUnsafe(new BigNumber(418),false).updateUnsafe(new BigNumber(419),false).updateUnsafe(new BigNumber(420),false).updateUnsafe(new BigNumber(421),false).updateUnsafe(new BigNumber(422),false).updateUnsafe(new BigNumber(423),false).updateUnsafe(new BigNumber(424),false).updateUnsafe(new BigNumber(425),false).updateUnsafe(new BigNumber(426),false).updateUnsafe(new BigNumber(427),false).updateUnsafe(new BigNumber(428),false).updateUnsafe(new BigNumber(429),false).updateUnsafe(new BigNumber(430),false).updateUnsafe(new BigNumber(431),false).updateUnsafe(new BigNumber(432),false).updateUnsafe(new BigNumber(433),false).updateUnsafe(new BigNumber(434),false).updateUnsafe(new BigNumber(435),false).updateUnsafe(new BigNumber(436),false).updateUnsafe(new BigNumber(437),false).updateUnsafe(new BigNumber(438),false).updateUnsafe(new BigNumber(439),false).updateUnsafe(new BigNumber(440),false).updateUnsafe(new BigNumber(441),false).updateUnsafe(new BigNumber(442),false).updateUnsafe(new BigNumber(443),false).updateUnsafe(new BigNumber(444),false).updateUnsafe(new BigNumber(445),false).updateUnsafe(new BigNumber(446),false).updateUnsafe(new BigNumber(447),false).updateUnsafe(new BigNumber(448),false).updateUnsafe(new BigNumber(449),false).updateUnsafe(new BigNumber(450),false).updateUnsafe(new BigNumber(451),false).updateUnsafe(new BigNumber(452),false).updateUnsafe(new BigNumber(453),false).updateUnsafe(new BigNumber(454),false).updateUnsafe(new BigNumber(455),false).updateUnsafe(new BigNumber(456),false).updateUnsafe(new BigNumber(457),false).updateUnsafe(new BigNumber(458),false).updateUnsafe(new BigNumber(459),false).updateUnsafe(new BigNumber(460),false).updateUnsafe(new BigNumber(461),false).updateUnsafe(new BigNumber(462),false).updateUnsafe(new BigNumber(463),false).updateUnsafe(new BigNumber(464),false).updateUnsafe(new BigNumber(465),false).updateUnsafe(new BigNumber(466),false).updateUnsafe(new BigNumber(467),false).updateUnsafe(new BigNumber(468),false).updateUnsafe(new BigNumber(469),false).updateUnsafe(new BigNumber(470),false).updateUnsafe(new BigNumber(471),false).updateUnsafe(new BigNumber(472),false).updateUnsafe(new BigNumber(473),false).updateUnsafe(new BigNumber(474),false).updateUnsafe(new BigNumber(475),false).updateUnsafe(new BigNumber(476),false).updateUnsafe(new BigNumber(477),false).updateUnsafe(new BigNumber(478),false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_191_v112)[new BigNumber(17)]).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,true).updateUnsafe(new BigNumber(2),false).updateUnsafe(new BigNumber(449),false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_191_v112)[new BigNumber(18)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(124),false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_191_v112)[new BigNumber(19)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(124),false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_191_v112)[new BigNumber(20)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(124),false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_191_v112)[new BigNumber(21)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-2),false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_191_v112)[new BigNumber(22)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(124),false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_191_v112)[new BigNumber(23)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(124),false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_191_v112)[new BigNumber(24)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(124),false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_191_v112)[new BigNumber(25)]).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_191_v112)[new BigNumber(26)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(124),false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_191_v112)[new BigNumber(27)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(124),false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_191_v112)[new BigNumber(28)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(124),false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_196_v115).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_197_v116).f11));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_197_v116).f12));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_198_v117));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_199_v118));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_202_v119).dtor_cf27)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_202_v119).dtor_cf27)[new BigNumber(28)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_203_v120).cardinality())));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_204_v121, _dafny.Seq.of(_dafny.MultiSet.fromElements(new BigNumber(124), new BigNumber(124), _dafny.ONE, new BigNumber(652)), _dafny.MultiSet.fromElements(new BigNumber(124), new BigNumber(124), _dafny.ONE, new BigNumber(652))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_205_v122).equals(_dafny.Map.Empty.slice().updateUnsafe(true,true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_206_v123).equals(_dafny.MultiSet.fromElements(_dafny.Seq.of(new BigNumber(4), _dafny.ONE)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_207_v124).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(new _dafny.CodePoint('y'.codePointAt(0))),new BigNumber(-1)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_208_v125)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_208_v125)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_208_v125)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_208_v125)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_208_v125)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_208_v125)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_208_v125)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_208_v125)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_208_v125)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_209_v126).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_223_v134).dtor_cf15).f13));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      return;
    }
  };

  $module.D0 = class D0 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC0(cf0) {
      let $dt = new D0(0);
      $dt.cf0 = cf0;
      return $dt;
    }
    get is_DC0() { return this.$tag === 0; }
    get dtor_cf0() { return this.cf0; }
    toString() {
      if (this.$tag === 0) {
        return "D0.DC0" + "(" + _dafny.toString(this.cf0) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf0, other.cf0);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _dafny.Seq.of();
    }
    static Rtd() {
      return class {
        static get Default() {
          return D0.Default();
        }
      };
    }
  }

  $module.D1 = class D1 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC2(cf2, cf3, cf4) {
      let $dt = new D1(0);
      $dt.cf2 = cf2;
      $dt.cf3 = cf3;
      $dt.cf4 = cf4;
      return $dt;
    }
    static create_DC1(cf1) {
      let $dt = new D1(1);
      $dt.cf1 = cf1;
      return $dt;
    }
    get is_DC2() { return this.$tag === 0; }
    get is_DC1() { return this.$tag === 1; }
    get dtor_cf2() { return this.cf2; }
    get dtor_cf3() { return this.cf3; }
    get dtor_cf4() { return this.cf4; }
    get dtor_cf1() { return this.cf1; }
    toString() {
      if (this.$tag === 0) {
        return "D1.DC2" + "(" + _dafny.toString(this.cf2) + ", " + _dafny.toString(this.cf3) + ", " + _dafny.toString(this.cf4) + ")";
      } else if (this.$tag === 1) {
        return "D1.DC1" + "(" + _dafny.toString(this.cf1) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf2, other.cf2) && this.cf3 === other.cf3 && _dafny.areEqual(this.cf4, other.cf4);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf1, other.cf1);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D1.create_DC2(_dafny.ZERO, false, _dafny.ZERO);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D1.Default();
        }
      };
    }
  }

  $module.D2 = class D2 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC4(cf6) {
      let $dt = new D2(0);
      $dt.cf6 = cf6;
      return $dt;
    }
    static create_DC5(cf7, cf8, cf9, cf10, cf11) {
      let $dt = new D2(1);
      $dt.cf7 = cf7;
      $dt.cf8 = cf8;
      $dt.cf9 = cf9;
      $dt.cf10 = cf10;
      $dt.cf11 = cf11;
      return $dt;
    }
    static create_DC3(cf5) {
      let $dt = new D2(2);
      $dt.cf5 = cf5;
      return $dt;
    }
    get is_DC4() { return this.$tag === 0; }
    get is_DC5() { return this.$tag === 1; }
    get is_DC3() { return this.$tag === 2; }
    get dtor_cf6() { return this.cf6; }
    get dtor_cf7() { return this.cf7; }
    get dtor_cf8() { return this.cf8; }
    get dtor_cf9() { return this.cf9; }
    get dtor_cf10() { return this.cf10; }
    get dtor_cf11() { return this.cf11; }
    get dtor_cf5() { return this.cf5; }
    toString() {
      if (this.$tag === 0) {
        return "D2.DC4" + "(" + _dafny.toString(this.cf6) + ")";
      } else if (this.$tag === 1) {
        return "D2.DC5" + "(" + _dafny.toString(this.cf7) + ", " + _dafny.toString(this.cf8) + ", " + _dafny.toString(this.cf9) + ", " + _dafny.toString(this.cf10) + ", " + _dafny.toString(this.cf11) + ")";
      } else if (this.$tag === 2) {
        return "D2.DC3" + "(" + _dafny.toString(this.cf5) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf6, other.cf6);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf7, other.cf7) && this.cf8 === other.cf8 && this.cf9 === other.cf9 && this.cf10 === other.cf10 && this.cf11 === other.cf11;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf5, other.cf5);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D2.create_DC4(_dafny.ZERO);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D2.Default();
        }
      };
    }
  }

  $module.D3 = class D3 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC7(cf13, cf14) {
      let $dt = new D3(0);
      $dt.cf13 = cf13;
      $dt.cf14 = cf14;
      return $dt;
    }
    static create_DC6(cf12) {
      let $dt = new D3(1);
      $dt.cf12 = cf12;
      return $dt;
    }
    get is_DC7() { return this.$tag === 0; }
    get is_DC6() { return this.$tag === 1; }
    get dtor_cf13() { return this.cf13; }
    get dtor_cf14() { return this.cf14; }
    get dtor_cf12() { return this.cf12; }
    toString() {
      if (this.$tag === 0) {
        return "D3.DC7" + "(" + _dafny.toString(this.cf13) + ", " + _dafny.toString(this.cf14) + ")";
      } else if (this.$tag === 1) {
        return "D3.DC6" + "(" + _dafny.toString(this.cf12) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf13, other.cf13) && _dafny.areEqual(this.cf14, other.cf14);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf12, other.cf12);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D3.create_DC7(_dafny.ZERO, _dafny.ZERO);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D3.Default();
        }
      };
    }
  }

  $module.D4 = class D4 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC9(cf16, cf17, cf18, cf19, cf20) {
      let $dt = new D4(0);
      $dt.cf16 = cf16;
      $dt.cf17 = cf17;
      $dt.cf18 = cf18;
      $dt.cf19 = cf19;
      $dt.cf20 = cf20;
      return $dt;
    }
    static create_DC8(cf15) {
      let $dt = new D4(1);
      $dt.cf15 = cf15;
      return $dt;
    }
    get is_DC9() { return this.$tag === 0; }
    get is_DC8() { return this.$tag === 1; }
    get dtor_cf16() { return this.cf16; }
    get dtor_cf17() { return this.cf17; }
    get dtor_cf18() { return this.cf18; }
    get dtor_cf19() { return this.cf19; }
    get dtor_cf20() { return this.cf20; }
    get dtor_cf15() { return this.cf15; }
    toString() {
      if (this.$tag === 0) {
        return "D4.DC9" + "(" + _dafny.toString(this.cf16) + ", " + _dafny.toString(this.cf17) + ", " + _dafny.toString(this.cf18) + ", " + _dafny.toString(this.cf19) + ", " + _dafny.toString(this.cf20) + ")";
      } else if (this.$tag === 1) {
        return "D4.DC8" + "(" + _dafny.toString(this.cf15) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf16, other.cf16) && _dafny.areEqual(this.cf17, other.cf17) && _dafny.areEqual(this.cf18, other.cf18) && _dafny.areEqual(this.cf19, other.cf19) && this.cf20 === other.cf20;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf15 === other.cf15;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D4.create_DC9(_dafny.ZERO, new _dafny.CodePoint('D'.codePointAt(0)), _dafny.Seq.of(), _dafny.ZERO, false);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D4.Default();
        }
      };
    }
  }

  $module.D5 = class D5 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC10(cf21) {
      let $dt = new D5(0);
      $dt.cf21 = cf21;
      return $dt;
    }
    get is_DC10() { return this.$tag === 0; }
    get dtor_cf21() { return this.cf21; }
    toString() {
      if (this.$tag === 0) {
        return "D5.DC10" + "(" + _dafny.toString(this.cf21) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf21, other.cf21);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _dafny.Set.Empty;
    }
    static Rtd() {
      return class {
        static get Default() {
          return D5.Default();
        }
      };
    }
  }

  $module.D6 = class D6 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC12(cf23, cf24) {
      let $dt = new D6(0);
      $dt.cf23 = cf23;
      $dt.cf24 = cf24;
      return $dt;
    }
    static create_DC11(cf22) {
      let $dt = new D6(1);
      $dt.cf22 = cf22;
      return $dt;
    }
    get is_DC12() { return this.$tag === 0; }
    get is_DC11() { return this.$tag === 1; }
    get dtor_cf23() { return this.cf23; }
    get dtor_cf24() { return this.cf24; }
    get dtor_cf22() { return this.cf22; }
    toString() {
      if (this.$tag === 0) {
        return "D6.DC12" + "(" + _dafny.toString(this.cf23) + ", " + _dafny.toString(this.cf24) + ")";
      } else if (this.$tag === 1) {
        return "D6.DC11" + "(" + this.cf22.toVerbatimString(true) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf23 === other.cf23 && _dafny.areEqual(this.cf24, other.cf24);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf22, other.cf22);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D6.create_DC12(false, new _dafny.CodePoint('D'.codePointAt(0)));
    }
    static Rtd() {
      return class {
        static get Default() {
          return D6.Default();
        }
      };
    }
  }

  $module.D7 = class D7 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC14(cf26) {
      let $dt = new D7(0);
      $dt.cf26 = cf26;
      return $dt;
    }
    static create_DC13(cf25) {
      let $dt = new D7(1);
      $dt.cf25 = cf25;
      return $dt;
    }
    get is_DC14() { return this.$tag === 0; }
    get is_DC13() { return this.$tag === 1; }
    get dtor_cf26() { return this.cf26; }
    get dtor_cf25() { return this.cf25; }
    toString() {
      if (this.$tag === 0) {
        return "D7.DC14" + "(" + _dafny.toString(this.cf26) + ")";
      } else if (this.$tag === 1) {
        return "D7.DC13" + "(" + _dafny.toString(this.cf25) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf26, other.cf26);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf25, other.cf25);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D7.create_DC14(_dafny.ZERO);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D7.Default();
        }
      };
    }
  }

  $module.D8 = class D8 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC16(cf28, cf29, cf30, cf31) {
      let $dt = new D8(0);
      $dt.cf28 = cf28;
      $dt.cf29 = cf29;
      $dt.cf30 = cf30;
      $dt.cf31 = cf31;
      return $dt;
    }
    static create_DC15(cf27) {
      let $dt = new D8(1);
      $dt.cf27 = cf27;
      return $dt;
    }
    static create_DC17(cf32) {
      let $dt = new D8(2);
      $dt.cf32 = cf32;
      return $dt;
    }
    get is_DC16() { return this.$tag === 0; }
    get is_DC15() { return this.$tag === 1; }
    get is_DC17() { return this.$tag === 2; }
    get dtor_cf28() { return this.cf28; }
    get dtor_cf29() { return this.cf29; }
    get dtor_cf30() { return this.cf30; }
    get dtor_cf31() { return this.cf31; }
    get dtor_cf27() { return this.cf27; }
    get dtor_cf32() { return this.cf32; }
    toString() {
      if (this.$tag === 0) {
        return "D8.DC16" + "(" + _dafny.toString(this.cf28) + ", " + _dafny.toString(this.cf29) + ", " + _dafny.toString(this.cf30) + ", " + _dafny.toString(this.cf31) + ")";
      } else if (this.$tag === 1) {
        return "D8.DC15" + "(" + _dafny.toString(this.cf27) + ")";
      } else if (this.$tag === 2) {
        return "D8.DC17" + "(" + _dafny.toString(this.cf32) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf28, other.cf28) && this.cf29 === other.cf29 && this.cf30 === other.cf30 && _dafny.areEqual(this.cf31, other.cf31);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf27 === other.cf27;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf32, other.cf32);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D8.create_DC16(_dafny.ZERO, false, false, _dafny.ZERO);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D8.Default();
        }
      };
    }
  }

  $module.D9 = class D9 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC19(cf34, cf35, cf36) {
      let $dt = new D9(0);
      $dt.cf34 = cf34;
      $dt.cf35 = cf35;
      $dt.cf36 = cf36;
      return $dt;
    }
    static create_DC20(cf37) {
      let $dt = new D9(1);
      $dt.cf37 = cf37;
      return $dt;
    }
    static create_DC18(cf33) {
      let $dt = new D9(2);
      $dt.cf33 = cf33;
      return $dt;
    }
    get is_DC19() { return this.$tag === 0; }
    get is_DC20() { return this.$tag === 1; }
    get is_DC18() { return this.$tag === 2; }
    get dtor_cf34() { return this.cf34; }
    get dtor_cf35() { return this.cf35; }
    get dtor_cf36() { return this.cf36; }
    get dtor_cf37() { return this.cf37; }
    get dtor_cf33() { return this.cf33; }
    toString() {
      if (this.$tag === 0) {
        return "D9.DC19" + "(" + _dafny.toString(this.cf34) + ", " + _dafny.toString(this.cf35) + ", " + _dafny.toString(this.cf36) + ")";
      } else if (this.$tag === 1) {
        return "D9.DC20" + "(" + _dafny.toString(this.cf37) + ")";
      } else if (this.$tag === 2) {
        return "D9.DC18" + "(" + _dafny.toString(this.cf33) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf34 === other.cf34 && _dafny.areEqual(this.cf35, other.cf35) && _dafny.areEqual(this.cf36, other.cf36);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf37 === other.cf37;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf33 === other.cf33;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D9.create_DC19(null, _dafny.ZERO, _dafny.ZERO);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D9.Default();
        }
      };
    }
  }

  $module.D10 = class D10 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC21(cf38) {
      let $dt = new D10(0);
      $dt.cf38 = cf38;
      return $dt;
    }
    get is_DC21() { return this.$tag === 0; }
    get dtor_cf38() { return this.cf38; }
    toString() {
      if (this.$tag === 0) {
        return "D10.DC21" + "(" + _dafny.toString(this.cf38) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf38 === other.cf38;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return [];
    }
    static Rtd() {
      return class {
        static get Default() {
          return D10.Default();
        }
      };
    }
  }

  $module.D11 = class D11 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC22(cf39) {
      let $dt = new D11(0);
      $dt.cf39 = cf39;
      return $dt;
    }
    get is_DC22() { return this.$tag === 0; }
    get dtor_cf39() { return this.cf39; }
    toString() {
      if (this.$tag === 0) {
        return "D11.DC22" + "(" + _dafny.toString(this.cf39) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf39, other.cf39);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _dafny.Map.Empty;
    }
    static Rtd() {
      return class {
        static get Default() {
          return D11.Default();
        }
      };
    }
  }

  $module.D12 = class D12 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC23(cf40) {
      let $dt = new D12(0);
      $dt.cf40 = cf40;
      return $dt;
    }
    get is_DC23() { return this.$tag === 0; }
    get dtor_cf40() { return this.cf40; }
    toString() {
      if (this.$tag === 0) {
        return "D12.DC23" + "(" + _dafny.toString(this.cf40) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf40, other.cf40);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _dafny.Map.Empty;
    }
    static Rtd() {
      return class {
        static get Default() {
          return D12.Default();
        }
      };
    }
  }

  $module.D13 = class D13 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC25(cf42, cf43, cf44) {
      let $dt = new D13(0);
      $dt.cf42 = cf42;
      $dt.cf43 = cf43;
      $dt.cf44 = cf44;
      return $dt;
    }
    static create_DC24(cf41) {
      let $dt = new D13(1);
      $dt.cf41 = cf41;
      return $dt;
    }
    get is_DC25() { return this.$tag === 0; }
    get is_DC24() { return this.$tag === 1; }
    get dtor_cf42() { return this.cf42; }
    get dtor_cf43() { return this.cf43; }
    get dtor_cf44() { return this.cf44; }
    get dtor_cf41() { return this.cf41; }
    toString() {
      if (this.$tag === 0) {
        return "D13.DC25" + "(" + _dafny.toString(this.cf42) + ", " + _dafny.toString(this.cf43) + ", " + _dafny.toString(this.cf44) + ")";
      } else if (this.$tag === 1) {
        return "D13.DC24" + "(" + _dafny.toString(this.cf41) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf42, other.cf42) && _dafny.areEqual(this.cf43, other.cf43) && this.cf44 === other.cf44;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf41 === other.cf41;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D13.create_DC25(_dafny.ZERO, _dafny.ZERO, false);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D13.Default();
        }
      };
    }
  }

  $module.D14 = class D14 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC27(cf46, cf47, cf48, cf49, cf50) {
      let $dt = new D14(0);
      $dt.cf46 = cf46;
      $dt.cf47 = cf47;
      $dt.cf48 = cf48;
      $dt.cf49 = cf49;
      $dt.cf50 = cf50;
      return $dt;
    }
    static create_DC28(cf51, cf52, cf53, cf54) {
      let $dt = new D14(1);
      $dt.cf51 = cf51;
      $dt.cf52 = cf52;
      $dt.cf53 = cf53;
      $dt.cf54 = cf54;
      return $dt;
    }
    static create_DC26(cf45) {
      let $dt = new D14(2);
      $dt.cf45 = cf45;
      return $dt;
    }
    static create_DC29(cf55) {
      let $dt = new D14(3);
      $dt.cf55 = cf55;
      return $dt;
    }
    get is_DC27() { return this.$tag === 0; }
    get is_DC28() { return this.$tag === 1; }
    get is_DC26() { return this.$tag === 2; }
    get is_DC29() { return this.$tag === 3; }
    get dtor_cf46() { return this.cf46; }
    get dtor_cf47() { return this.cf47; }
    get dtor_cf48() { return this.cf48; }
    get dtor_cf49() { return this.cf49; }
    get dtor_cf50() { return this.cf50; }
    get dtor_cf51() { return this.cf51; }
    get dtor_cf52() { return this.cf52; }
    get dtor_cf53() { return this.cf53; }
    get dtor_cf54() { return this.cf54; }
    get dtor_cf45() { return this.cf45; }
    get dtor_cf55() { return this.cf55; }
    toString() {
      if (this.$tag === 0) {
        return "D14.DC27" + "(" + _dafny.toString(this.cf46) + ", " + _dafny.toString(this.cf47) + ", " + _dafny.toString(this.cf48) + ", " + _dafny.toString(this.cf49) + ", " + _dafny.toString(this.cf50) + ")";
      } else if (this.$tag === 1) {
        return "D14.DC28" + "(" + _dafny.toString(this.cf51) + ", " + _dafny.toString(this.cf52) + ", " + _dafny.toString(this.cf53) + ", " + _dafny.toString(this.cf54) + ")";
      } else if (this.$tag === 2) {
        return "D14.DC26" + "(" + _dafny.toString(this.cf45) + ")";
      } else if (this.$tag === 3) {
        return "D14.DC29" + "(" + _dafny.toString(this.cf55) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf46, other.cf46) && this.cf47 === other.cf47 && this.cf48 === other.cf48 && this.cf49 === other.cf49 && _dafny.areEqual(this.cf50, other.cf50);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf51 === other.cf51 && this.cf52 === other.cf52 && this.cf53 === other.cf53 && _dafny.areEqual(this.cf54, other.cf54);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf45, other.cf45);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf55, other.cf55);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D14.create_DC27(_dafny.ZERO, false, false, false, _dafny.ZERO);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D14.Default();
        }
      };
    }
  }

  $module.T0 = class T0 {
  };

  $module.T1 = class T1 {
  };

  $module.GlobalState = class GlobalState {
    constructor () {
      this._tname = "_module.GlobalState";
      this.f2 = _dafny.ZERO;
      this._f0 = [];
      this._f1 = false;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f0, f1, f2) {
      let _this = this;
      (_this)._f0 = f0;
      (_this)._f1 = f1;
      (_this).f2 = f2;
      return;
    }
    get f0() {
      let _this = this;
      return _this._f0;
    };
    get f1() {
      let _this = this;
      return _this._f1;
    };
  };

  $module.C0 = class C0 {
    constructor () {
      this._tname = "_module.C0";
      this._f13 = _dafny.ZERO;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f13) {
      let _this = this;
      (_this)._f13 = f13;
      return;
    }
    get f13() {
      let _this = this;
      return _this._f13;
    };
  };

  $module.C1 = class C1 {
    constructor () {
      this._tname = "_module.C1";
      this._f3 = [];
      this._f14 = _dafny.ZERO;
    }
    _parentTraits() {
      return [_module.T0];
    }
    get f3() {
      let _this = this;
      return _this._f3;
    };
    __ctor(f14, f3) {
      let _this = this;
      (_this)._f14 = f14;
      (_this)._f3 = f3;
      return;
    }
    fm29(p0, p1, p2, globalState) {
      let _this = this;
      return new _dafny.CodePoint('y'.codePointAt(0));
    };
    get f14() {
      let _this = this;
      return _this._f14;
    };
  };

  $module.C2 = class C2 {
    constructor () {
      this._tname = "_module.C2";
      this._f3 = [];
      this._f11 = false;
      this._f12 = _dafny.ZERO;
    }
    _parentTraits() {
      return [_module.T0, _module.T1];
    }
    get f3() {
      let _this = this;
      return _this._f3;
    };
    __ctor(f11, f12, f3) {
      let _this = this;
      (_this)._f11 = f11;
      (_this)._f12 = f12;
      (_this)._f3 = f3;
      return;
    }
    m1(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = new _dafny.CodePoint('D'.codePointAt(0));
      let _262_v0;
      _262_v0 = true;
      _262_v0 = (_this).f11;
      let _263_v1;
      _263_v1 = _dafny.Map.Empty.slice().updateUnsafe((_this).f11,_module.__default.fm23(p2, (_this).f12, (_this).f12, true, globalState));
      let _264_v2;
      _264_v2 = _dafny.Seq.of(_263_v1, _263_v1);
      let _265_v3;
      _265_v3 = _dafny.Map.Empty.slice().updateUnsafe((p0)[_module.__default.safeIndex(_module.__default.fm24(_dafny.Set.fromElements((_this).f12, (_this).f12, new BigNumber((p2).length)), globalState), new BigNumber((p0).length))],_264_v2);
      let _266_v4;
      _266_v4 = _module.D1.create_DC1(_dafny.Seq.Concat(_264_v2, (((_265_v3).contains((_this).f12)) ? ((_265_v3).get((_this).f12)) : (_dafny.Seq.of(_263_v1)))));
      let _source5 = _266_v4;
      if (_source5.is_DC2) {
        let _267___mcc_h0 = (_source5).cf2;
        let _268___mcc_h1 = (_source5).cf3;
        let _269___mcc_h2 = (_source5).cf4;
        let _270_cf4 = _269___mcc_h2;
        let _271_cf3 = _268___mcc_h1;
        let _272_cf2 = _267___mcc_h0;
        let _273_v5;
        let _nw48 = Array((new BigNumber(22)).toNumber()).fill(false);
        _273_v5 = _nw48;
        let _index32 = _module.__default.safeIndex(new BigNumber(309), new BigNumber((_273_v5).length));
        (_273_v5)[_index32] = _262_v0;
        let _274_v6;
        _274_v6 = _dafny.Set.fromElements(_272_cf2, new BigNumber((p2).length), new BigNumber(-72));
        let _index33 = _module.__default.safeIndex(new BigNumber(309), new BigNumber((_273_v5).length));
        (_273_v5)[_index33] = (_274_v6).IsSubsetOf((_274_v6).Difference(_274_v6));
        _274_v6 = _274_v6;
        if (_271_cf3) {
          let _275_v7;
          _275_v7 = _dafny.Seq.of(true, _262_v0);
          let _276_v8;
          let _nw49 = new _module.C0();
          _nw49.__ctor((_270_cf4).minus(new BigNumber((_275_v7).length)));
          _276_v8 = _nw49;
          let _277_v9;
          _277_v9 = _dafny.Seq.UnicodeFromString("hdgtom");
          _277_v9 = ((!(false)) ? (_277_v9) : (_277_v9));
          let _278_v10;
          _278_v10 = _dafny.Set.fromElements((_this).f11, p1);
          let _279_v11;
          let _nw50 = Array((new BigNumber(14)).toNumber());
          _nw50[(_dafny.ZERO).toNumber()] = _dafny.Set.fromElements(_262_v0);
          _nw50[(_dafny.ONE).toNumber()] = (_278_v10).Union(_278_v10);
          _nw50[(new BigNumber(2)).toNumber()] = _278_v10;
          _nw50[(new BigNumber(3)).toNumber()] = _dafny.Set.fromElements((_273_v5)[_module.__default.safeIndex(new BigNumber(309), new BigNumber((_273_v5).length))], true);
          _nw50[(new BigNumber(4)).toNumber()] = _278_v10;
          _nw50[(new BigNumber(5)).toNumber()] = _278_v10;
          _nw50[(new BigNumber(6)).toNumber()] = _278_v10;
          _nw50[(new BigNumber(7)).toNumber()] = _278_v10;
          _nw50[(new BigNumber(8)).toNumber()] = ((true) ? (_278_v10) : (_278_v10));
          _nw50[(new BigNumber(9)).toNumber()] = (_dafny.Set.fromElements(_262_v0)).Difference(_278_v10);
          _nw50[(new BigNumber(10)).toNumber()] = _278_v10;
          _nw50[(new BigNumber(11)).toNumber()] = _278_v10;
          _nw50[(new BigNumber(12)).toNumber()] = _278_v10;
          _nw50[(new BigNumber(13)).toNumber()] = _278_v10;
          _279_v11 = _nw50;
          _279_v11 = _279_v11;
          let _280_v12;
          _280_v12 = _dafny.MultiSet.fromElements((_this).f12, new BigNumber(634), (_276_v8).f13);
          let _281_v13;
          _281_v13 = _dafny.Map.Empty.slice().updateUnsafe(_280_v12,_270_cf4);
          let _282_v14;
          _282_v14 = _module.D2.create_DC3(_281_v13);
          let _283_v15;
          _283_v15 = _dafny.Map.Empty.slice().updateUnsafe(_282_v14,p1);
          let _284_v16;
          _284_v16 = _dafny.Map.Empty.slice().updateUnsafe(_272_cf2,(new BigNumber((_283_v15).length)).isLessThanOrEqualTo(new BigNumber((_274_v6).length)));
          _284_v16 = (_284_v16).update((_276_v8).f13, (_273_v5)[_module.__default.safeIndex(new BigNumber(309), new BigNumber((_273_v5).length))]);
          let _285_v17;
          let _nw51 = Array((new BigNumber(26)).toNumber()).fill(_dafny.ZERO);
          _285_v17 = _nw51;
          let _286_v18;
          _286_v18 = _dafny.Seq.of(_274_v6);
          let _287_v19;
          _287_v19 = _dafny.Map.Empty.slice().updateUnsafe(_285_v17,new BigNumber((_dafny.Seq.Concat(_286_v18, _dafny.Seq.of(_274_v6, _274_v6, _274_v6))).length));
          _287_v19 = (_287_v19).update(_285_v17, (_this).f12);
        } else {
          _262_v0 = !(_271_cf3);
          let _288_v20;
          let _289_v21;
          let _290_v22;
          let _out17;
          let _out18;
          let _out19;
          let _outcollector7 = _module.__default.m0(_270_cf4, globalState);
          _out17 = _outcollector7[0];
          _out18 = _outcollector7[1];
          _out19 = _outcollector7[2];
          _288_v20 = _out17;
          _289_v21 = _out18;
          _290_v22 = _out19;
          let _291_v23;
          _291_v23 = _dafny.Seq.of(_271_cf3, p1);
          let _292_v25;
          _292_v25 = _dafny.Seq.of(function () {
            let _coll15 = new _dafny.Set();
            for (const _compr_15 of (p0).Elements) {
              let _293_v24 = _compr_15;
              if (_dafny.Seq.contains(p0, _293_v24)) {
                _coll15.add((_293_v24).multipliedBy(new BigNumber(495)));
              }
            }
            return _coll15;
          }());
          let _294_v26;
          _294_v26 = _dafny.Map.Empty.slice().updateUnsafe((new BigNumber((_dafny.Seq.of((_291_v23)[_module.__default.safeIndex(new BigNumber(896), new BigNumber((_291_v23).length))], _271_cf3)).length)).multipliedBy(_module.__default.fm24((_292_v25)[_module.__default.safeIndex((p0)[_module.__default.safeIndex(_290_v22, new BigNumber((p0).length))], new BigNumber((_292_v25).length))], globalState)),_262_v0);
          _294_v26 = (_294_v26).update((_this).f12, !(!(_288_v20)) || (_262_v0));
          let _295_v27;
          let _nw52 = Array((_dafny.ONE).toNumber());
          _nw52[(_dafny.ZERO).toNumber()] = (_this).f12;
          _295_v27 = _nw52;
          let _296_v28;
          _296_v28 = _dafny.Map.Empty.slice().updateUnsafe(_262_v0,_274_v6);
          let _297_v29;
          _297_v29 = _dafny.MultiSet.fromElements(!(false), _271_cf3);
          let _298_v30;
          _298_v30 = _dafny.Map.Empty.slice().updateUnsafe((_this).f12,(((_297_v29).contains(_288_v20)) ? ((_297_v29).get(_288_v20)) : (new BigNumber(293))));
          let _index34 = _module.__default.safeIndex(new BigNumber(976), new BigNumber((_295_v27).length));
          (_295_v27)[_index34] = _module.__default.safeModuloInt(new BigNumber(((((_296_v28).contains((((_263_v1).contains(p1)) ? ((_263_v1).get(p1)) : (false)))) ? ((_296_v28).get((((_263_v1).contains(p1)) ? ((_263_v1).get(p1)) : (false)))) : ((_292_v25)[_module.__default.safeIndex(new BigNumber((_298_v30).length), new BigNumber((_292_v25).length))]))).length), (_dafny.ZERO).minus(_module.__default.fm24(_274_v6, globalState)));
          let _index35 = _module.__default.safeIndex(new BigNumber(976), new BigNumber((_295_v27).length));
          (_295_v27)[_index35] = (_290_v22).minus(_272_cf2);
          let _299_v31;
          let _nw53 = new _module.C0();
          _nw53.__ctor((_dafny.ZERO).minus(new BigNumber(-817)));
          _299_v31 = _nw53;
        }
        (globalState).f2 = (_this).f12;
      } else {
        let _300___mcc_h3 = (_source5).cf1;
        let _301_cf1 = _300___mcc_h3;
        let _302_v32;
        let _init4 = function (_303_i0) {
          return (_this).f11;
        };
        let _nw54 = Array((new BigNumber(10)).toNumber());
        for (let _i0_4 = 0; _i0_4 < new BigNumber(_nw54.length); _i0_4++) {
          _nw54[_i0_4] = _init4(new BigNumber(_i0_4));
        }
        _302_v32 = _nw54;
        let _index36 = _module.__default.safeIndex(new BigNumber(478), new BigNumber((_302_v32).length));
        (_302_v32)[_index36] = _262_v0;
        let _304_v33;
        _304_v33 = _module.D2.create_DC5((_this).f12, (_this).f11, (_this).f11, _262_v0, _262_v0);
        let _index37 = _module.__default.safeIndex(new BigNumber(478), new BigNumber((_302_v32).length));
        (_302_v32)[_index37] = (_304_v33).dtor_cf11;
        let _305_v34;
        _305_v34 = _dafny.MultiSet.fromElements(_302_v32);
        let _index38 = _module.__default.safeIndex(new BigNumber(478), new BigNumber((_302_v32).length));
        (_302_v32)[_index38] = ((_305_v34).Difference(_305_v34)).IsSubsetOf(_305_v34);
        let _306_v35;
        _306_v35 = _dafny.Set.fromElements((_this).f11, _262_v0, (_302_v32)[_module.__default.safeIndex(new BigNumber(478), new BigNumber((_302_v32).length))]);
        let _307_v36;
        _307_v36 = _dafny.Map.Empty.slice().updateUnsafe(_306_v35,_306_v35);
        if ((_306_v35).IsDisjointFrom((((_307_v36).contains(_dafny.Set.fromElements(_262_v0, _262_v0))) ? ((_307_v36).get(_dafny.Set.fromElements(_262_v0, _262_v0))) : (_dafny.Set.fromElements(_262_v0))))) {
          let _308_v37;
          _308_v37 = _dafny.Map.Empty.slice().updateUnsafe((_302_v32)[_module.__default.safeIndex(new BigNumber(478), new BigNumber((_302_v32).length))],p0);
          let _309_v38;
          _309_v38 = _dafny.Seq.of((_302_v32)[_module.__default.safeIndex(new BigNumber(478), new BigNumber((_302_v32).length))], (_this).f11, (_302_v32)[_module.__default.safeIndex(new BigNumber(478), new BigNumber((_302_v32).length))], false);
          let _310_v39;
          _310_v39 = _dafny.Map.Empty.slice().updateUnsafe((_302_v32)[_module.__default.safeIndex(new BigNumber(478), new BigNumber((_302_v32).length))],(_this).f12);
          let _311_v40;
          let _nw55 = Array((new BigNumber(6)).toNumber());
          _nw55[(_dafny.ZERO).toNumber()] = p0;
          _nw55[(_dafny.ONE).toNumber()] = _dafny.Seq.Concat((((_308_v37).contains(!(_module.__default.fm23(p2, (_this).f12, new BigNumber(784), (_302_v32)[_module.__default.safeIndex(new BigNumber(478), new BigNumber((_302_v32).length))], globalState)))) ? ((_308_v37).get(!(_module.__default.fm23(p2, (_this).f12, new BigNumber(784), (_302_v32)[_module.__default.safeIndex(new BigNumber(478), new BigNumber((_302_v32).length))], globalState)))) : (p0)), _dafny.Seq.update(p0, _module.__default.safeIndex(new BigNumber((_309_v38).length), new BigNumber((p0).length)), new BigNumber((_310_v39).length)));
          _nw55[(new BigNumber(2)).toNumber()] = _dafny.Seq.of((_this).f12, (_this).f12);
          _nw55[(new BigNumber(3)).toNumber()] = p0;
          _nw55[(new BigNumber(4)).toNumber()] = p0;
          _nw55[(new BigNumber(5)).toNumber()] = p0;
          _311_v40 = _nw55;
          let _312_v41;
          _312_v41 = _dafny.Map.Empty.slice().updateUnsafe((_this).f12,_311_v40);
          let _313_v42;
          _313_v42 = _dafny.Seq.of((((_312_v41).contains((_this).f12)) ? ((_312_v41).get((_this).f12)) : (_311_v40)));
          _311_v40 = (_313_v42)[_module.__default.safeIndex((_this).f12, new BigNumber((_313_v42).length))];
          let _314_v43;
          _314_v43 = new _dafny.CodePoint('l'.codePointAt(0));
          r1 = _314_v43;
          let _index39 = _module.__default.safeIndex(new BigNumber(478), new BigNumber((_302_v32).length));
          (_302_v32)[_index39] = (!(_dafny.Seq.contains(_dafny.Seq.update(_dafny.Seq.of(p2), _module.__default.safeIndex((_this).f12, new BigNumber((_dafny.Seq.of(p2)).length)), p2), p2))) && (!(_module.__default.fm23(p2, (_this).f12, (_this).f12, (_302_v32)[_module.__default.safeIndex(new BigNumber(478), new BigNumber((_302_v32).length))], globalState)));
          let _315_v44;
          let _nw56 = Array((new BigNumber(3)).toNumber());
          _315_v44 = _nw56;
          let _316_v45;
          let _nw57 = new _module.C0();
          _nw57.__ctor((_this).f12);
          _316_v45 = _nw57;
          let _index40 = _module.__default.safeIndex(new BigNumber(607), new BigNumber((_315_v44).length));
          (_315_v44)[_index40] = _316_v45;
          let _index41 = _module.__default.safeIndex(new BigNumber(607), new BigNumber((_315_v44).length));
          (_315_v44)[_index41] = _316_v45;
          let _index42 = _module.__default.safeIndex(new BigNumber(567), new BigNumber((_311_v40).length));
          (_311_v40)[_index42] = p0;
          let _317_v46;
          _317_v46 = _dafny.Seq.of(_309_v38);
          let _318_v47;
          _318_v47 = _dafny.MultiSet.fromElements(_309_v38, (_317_v46)[_module.__default.safeIndex((_316_v45).f13, new BigNumber((_317_v46).length))]);
          let _319_v48;
          _319_v48 = _dafny.MultiSet.fromElements(_262_v0);
          let _320_v49;
          _320_v49 = _dafny.Map.Empty.slice().updateUnsafe(_318_v47,_dafny.Seq.Concat(_dafny.Seq.update(p0, _module.__default.safeIndex((p0)[_module.__default.safeIndex((((_319_v48).contains((_302_v32)[_module.__default.safeIndex(new BigNumber(478), new BigNumber((_302_v32).length))])) ? ((_319_v48).get((_302_v32)[_module.__default.safeIndex(new BigNumber(478), new BigNumber((_302_v32).length))])) : (new BigNumber((p2).length))), new BigNumber((p0).length))], new BigNumber((p0).length)), (_this).f12), p0));
          let _index43 = _module.__default.safeIndex(new BigNumber(567), new BigNumber((_311_v40).length));
          (_311_v40)[_index43] = (((_320_v49).contains(_318_v47)) ? ((_320_v49).get(_318_v47)) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(651)), function (_321_i1) {
            return new BigNumber(-644);
          })));
        } else {
          let _322_v50;
          _322_v50 = _dafny.Set.fromElements((_this).f12);
          let _323_v51;
          _323_v51 = _dafny.Seq.of(true, p1);
          let _324_v52;
          _324_v52 = _dafny.MultiSet.fromElements(p1, p1);
          let _325_v53;
          _325_v53 = _dafny.Map.Empty.slice().updateUnsafe((_this).f11,(_this).f12);
          let _326_v54;
          let _nw58 = Array((new BigNumber(24)).toNumber());
          _nw58[(_dafny.ZERO).toNumber()] = new BigNumber(413);
          _nw58[(_dafny.ONE).toNumber()] = (_this).f12;
          _nw58[(new BigNumber(2)).toNumber()] = (_this).f12;
          _nw58[(new BigNumber(3)).toNumber()] = (_this).f12;
          _nw58[(new BigNumber(4)).toNumber()] = (_this).f12;
          _nw58[(new BigNumber(5)).toNumber()] = (_this).f12;
          _nw58[(new BigNumber(6)).toNumber()] = (_dafny.ZERO).minus((_this).f12);
          _nw58[(new BigNumber(7)).toNumber()] = _module.__default.fm24(_322_v50, globalState);
          _nw58[(new BigNumber(8)).toNumber()] = (_dafny.ZERO).minus(new BigNumber((_323_v51).length));
          _nw58[(new BigNumber(9)).toNumber()] = new BigNumber((p2).length);
          _nw58[(new BigNumber(10)).toNumber()] = (_dafny.ZERO).minus((_this).f12);
          _nw58[(new BigNumber(11)).toNumber()] = (_dafny.ZERO).minus((_this).f12);
          _nw58[(new BigNumber(12)).toNumber()] = (_this).f12;
          _nw58[(new BigNumber(13)).toNumber()] = (_this).f12;
          _nw58[(new BigNumber(14)).toNumber()] = new BigNumber(825);
          _nw58[(new BigNumber(15)).toNumber()] = (_this).f12;
          _nw58[(new BigNumber(16)).toNumber()] = (_dafny.ZERO).minus(new BigNumber((p2).length));
          _nw58[(new BigNumber(17)).toNumber()] = (_this).f12;
          _nw58[(new BigNumber(18)).toNumber()] = (_this).f12;
          _nw58[(new BigNumber(19)).toNumber()] = (((_324_v52).contains((_this).f11)) ? ((_324_v52).get((_this).f11)) : (new BigNumber((_325_v53).length)));
          _nw58[(new BigNumber(20)).toNumber()] = (_this).f12;
          _nw58[(new BigNumber(21)).toNumber()] = new BigNumber((_324_v52).cardinality());
          _nw58[(new BigNumber(22)).toNumber()] = (_this).f12;
          _nw58[(new BigNumber(23)).toNumber()] = new BigNumber(241);
          _326_v54 = _nw58;
          let _327_v55;
          _327_v55 = _dafny.Seq.of(_326_v54, _326_v54);
          let _328_v56;
          let _nw59 = new _module.C0();
          _nw59.__ctor((_this).f12);
          _328_v56 = _nw59;
          let _rhs29 = _dafny.Seq.of(_326_v54, _326_v54, _326_v54);
          let _rhs30 = _328_v56;
          _327_v55 = _rhs29;
          _328_v56 = _rhs30;
          let _329_v57;
          _329_v57 = _dafny.Seq.of(_322_v50, _dafny.Set.fromElements((_this).f12), _dafny.Set.fromElements(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(23)), ((_330_p2, _331_v0) => function (_332_i2) {
            return new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_330_p2,_331_v0)).length);
          })(p2, _262_v0))).length)));
          let _333_v58;
          _333_v58 = _dafny.MultiSet.fromElements(new BigNumber(848), _module.__default.fm24((_329_v57)[_module.__default.safeIndex((_this).f12, new BigNumber((_329_v57).length))], globalState), new BigNumber((_dafny.Set.fromElements(_328_v56, _328_v56, _328_v56)).length));
          let _334_v59;
          _334_v59 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_263_v1).length),(_this).f12);
          (globalState).f2 = (((_333_v58).contains(new BigNumber(725))) ? ((_333_v58).get(new BigNumber(725))) : (new BigNumber((((_module.__default.fm23(p2, (_328_v56).f13, new BigNumber((_334_v59).length), (_302_v32)[_module.__default.safeIndex(new BigNumber(478), new BigNumber((_302_v32).length))], globalState)) ? (_dafny.Seq.Create(_module.__default.abs(new BigNumber(300)), function (_335_i3) {
            return new _dafny.CodePoint('y'.codePointAt(0));
          })) : (p2))).length)));
          let _336_v60;
          _336_v60 = _dafny.Seq.of(_323_v51);
          _262_v0 = _dafny.areEqual(_336_v60, _336_v60);
          (globalState).f2 = (((!(_262_v0)) ? (_304_v33) : (_304_v33))).dtor_cf7;
          r0 = ((_this).f12).minus((_328_v56).f13);
        }
        let _337_v61;
        let _nw60 = new _module.C0();
        _nw60.__ctor((new BigNumber(733)).multipliedBy((_this).f12));
        _337_v61 = _nw60;
      }
      let _338_v62;
      _338_v62 = _dafny.Map.Empty.slice().updateUnsafe(p0,p2);
      _338_v62 = (_338_v62).update(p0, p2);
      let _339_v63;
      let _init5 = function (_340_i5) {
        return (_340_i5).minus(new BigNumber(-892));
      };
      let _nw61 = Array((_dafny.ONE).toNumber());
      for (let _i0_5 = 0; _i0_5 < new BigNumber(_nw61.length); _i0_5++) {
        _nw61[_i0_5] = _init5(new BigNumber(_i0_5));
      }
      _339_v63 = _nw61;
      for (const _guard_loop_0 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_339_v63).length))) {
        let _341_i4 = _guard_loop_0;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_341_i4)) && ((_341_i4).isLessThan(new BigNumber((_339_v63).length))))) {
          (_339_v63)[(_341_i4)] = _module.__default.safeModuloInt(_341_i4, (_this).f12);
        }
      }
      r0 = (_this).f12;
      let _342_v64;
      _342_v64 = _dafny.Set.fromElements(new _dafny.CodePoint('y'.codePointAt(0)));
      let _hi1 = new BigNumber((_342_v64).length);
      for (let _343_i6 = (_dafny.ZERO).minus((_this).f12); _343_i6.isLessThan(_hi1); _343_i6 = _343_i6.plus(_dafny.ONE)) {
        let _344_v65;
        _344_v65 = _module.D3.create_DC7((_343_i6).plus(_343_i6), (_this).f12);
        let _source6 = _344_v65;
        if (_source6.is_DC7) {
          let _345___mcc_h4 = (_source6).cf13;
          let _346___mcc_h5 = (_source6).cf14;
          let _347_cf14 = _346___mcc_h5;
          let _348_cf13 = _345___mcc_h4;
          let _349_v66;
          _349_v66 = _dafny.Map.Empty.slice().updateUnsafe(_348_cf13,!((_this).f11));
          (globalState).f2 = (_348_cf13).plus(new BigNumber((_349_v66).length));
          _262_v0 = !(_dafny.Seq.IsPrefixOf(p2, p2));
          let _350_v67;
          let _nw62 = new _module.C0();
          _nw62.__ctor(_347_cf14);
          _350_v67 = _nw62;
          _262_v0 = true;
        } else {
          let _351___mcc_h6 = (_source6).cf12;
          let _352_cf12 = _351___mcc_h6;
          _262_v0 = (new BigNumber((p2).length)).isLessThan((_343_i6).minus(_343_i6));
          let _353_v68;
          _353_v68 = _dafny.Seq.UnicodeFromString("chue");
          let _354_v69;
          let _nw63 = Array((new BigNumber(25)).toNumber()).fill(false);
          _354_v69 = _nw63;
          let _355_v70;
          _355_v70 = _dafny.Set.fromElements(_354_v69, _354_v69);
          let _356_v71;
          _356_v71 = _dafny.MultiSet.fromElements(((true) ? (_343_i6) : (new BigNumber((_355_v70).length))));
          let _rhs31 = p2;
          let _rhs32 = (_dafny.ZERO).minus((((true) || ((_this).f11)) ? (_343_i6) : ((_this).f12)));
          let _rhs33 = (_356_v71).Difference((_dafny.MultiSet.FromArray(p0)).Intersect(_356_v71));
          _353_v68 = _rhs31;
          r0 = _rhs32;
          _356_v71 = _rhs33;
          _354_v69 = _354_v69;
          _266_v4 = _266_v4;
        }
        let _357_v72;
        let _nw64 = new _module.C0();
        _nw64.__ctor((_this).f12);
        _357_v72 = _nw64;
        (globalState).f2 = new BigNumber(99);
        let _358_v73;
        _358_v73 = _dafny.MultiSet.fromElements(p1, (_this).f11);
        let _359_v74;
        _359_v74 = _dafny.Map.Empty.slice().updateUnsafe(_358_v73,!(p1));
        let _360_v75;
        _360_v75 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.Map.Empty.slice().updateUnsafe(_358_v73,p1)).Merge(_359_v74),_module.__default.fm25(p1, (_357_v72).f13, globalState));
        _360_v75 = (_360_v75).update((_359_v74).Merge(_dafny.Map.Empty.slice().updateUnsafe(_358_v73,p1)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(147)), function (_361_i7) {
          return new _dafny.CodePoint('j'.codePointAt(0));
        }));
      }
      r0 = (_dafny.ZERO).minus(_module.__default.safeDivisionInt((_this).f12, (_module.D2.create_DC4((_this).f12)).dtor_cf6));
      let _362_v76;
      _362_v76 = new _dafny.CodePoint('q'.codePointAt(0));
      r1 = _362_v76;
      return [r0, r1];
    }
    m6(globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = _dafny.ZERO;
      let r2 = false;
      r2 = !((_this).f12).isEqualTo(((_this).f12).multipliedBy((_this).f12));
      let _363_v0;
      _363_v0 = _dafny.Map.Empty.slice().updateUnsafe((_this).f11,true);
      let _364_v1;
      _364_v1 = _dafny.Map.Empty.slice().updateUnsafe(!(true),_363_v0);
      let _365_v2;
      _365_v2 = _dafny.Set.fromElements(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(760)), function (_366_i1) {
        return new _dafny.CodePoint('s'.codePointAt(0));
      })).length), (_this).f12, new BigNumber((_364_v1).length));
      let _hi2 = _module.__default.fm24(_365_v2, globalState);
      for (let _367_i0 = (_this).f12; _367_i0.isLessThan(_hi2); _367_i0 = _367_i0.plus(_dafny.ONE)) {
        r2 = (((_this).f11) ? ((_this).f11) : ((_this).f11));
        (globalState).f2 = _367_i0;
        let _368_v3;
        _368_v3 = _dafny.Seq.UnicodeFromString("y");
        if (!(_module.__default.fm23(_368_v3, _module.__default.safeDivisionInt(_367_i0, new BigNumber(-836)), _367_i0, (_this).f11, globalState))) {
          _368_v3 = _dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.Concat(_368_v3, _368_v3), _368_v3), _module.__default.safeIndex((_this).f12, new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Concat(_368_v3, _368_v3), _368_v3)).length)), new _dafny.CodePoint('f'.codePointAt(0)));
          let _369_v4;
          let _nw65 = Array((new BigNumber(16)).toNumber()).fill(false);
          _369_v4 = _nw65;
          let _370_v5;
          _370_v5 = _module.D2.create_DC5(_367_i0, (_this).f11, (_this).f11, (_this).f11, (_this).f11);
          let _index44 = _module.__default.safeIndex(new BigNumber(561), new BigNumber((_369_v4).length));
          (_369_v4)[_index44] = ((_370_v5).dtor_cf9) || ((_this).f11);
          let _371_v6;
          let _init6 = ((_372_v5) => function (_373_i2) {
            return _372_v5;
          })(_370_v5);
          let _nw66 = Array((new BigNumber(22)).toNumber());
          for (let _i0_6 = 0; _i0_6 < new BigNumber(_nw66.length); _i0_6++) {
            _nw66[_i0_6] = _init6(new BigNumber(_i0_6));
          }
          _371_v6 = _nw66;
          let _374_v7;
          _374_v7 = _dafny.Seq.of(!((_this).f11), (_this).f11, true, (_this).f11, !((_this).f11));
          let _375_v8;
          _375_v8 = _dafny.MultiSet.fromElements((_this).f11);
          let _index45 = _module.__default.safeIndex(new BigNumber(561), new BigNumber((_369_v4).length));
          let _rhs34 = !((_this).f12).isEqualTo(new BigNumber((_374_v7).length));
          let _rhs35 = !((_375_v8).update(false, _module.__default.abs(_367_i0))).equals((_375_v8).Difference(_375_v8));
          let _rhs36 = (_this).f11;
          let _rhs37 = _371_v6;
          let _lhs17 = _369_v4;
          let _lhs18 = _module.__default.safeIndex(new BigNumber(561), new BigNumber((_369_v4).length));
          r2 = _rhs34;
          _lhs17[_lhs18] = _rhs35;
          r2 = _rhs36;
          _371_v6 = _rhs37;
          let _376_v9;
          _376_v9 = new _dafny.CodePoint('j'.codePointAt(0));
          let _377_v10;
          _377_v10 = _module.D1.create_DC2((_this).f12, (_369_v4)[_module.__default.safeIndex(new BigNumber(561), new BigNumber((_369_v4).length))], new BigNumber(-552));
          let _378_v11;
          let _nw67 = Array((new BigNumber(4)).toNumber());
          _nw67[(_dafny.ZERO).toNumber()] = _374_v7;
          _nw67[(_dafny.ONE).toNumber()] = _dafny.Seq.Concat(_374_v7, _374_v7);
          _nw67[(new BigNumber(2)).toNumber()] = _374_v7;
          _nw67[(new BigNumber(3)).toNumber()] = (((_377_v10).dtor_cf3) ? (_374_v7) : (_dafny.Seq.of((_369_v4)[_module.__default.safeIndex(new BigNumber(561), new BigNumber((_369_v4).length))])));
          _378_v11 = _nw67;
          let _index46 = _module.__default.safeIndex(new BigNumber(807), new BigNumber((_378_v11).length));
          (_378_v11)[_index46] = _374_v7;
          let _379_v12;
          _379_v12 = _dafny.Map.Empty.slice().updateUnsafe((_this).f11,_376_v9);
          let _380_v13;
          _380_v13 = _dafny.Set.fromElements((((_379_v12).contains((_369_v4)[_module.__default.safeIndex(new BigNumber(561), new BigNumber((_369_v4).length))])) ? ((_379_v12).get((_369_v4)[_module.__default.safeIndex(new BigNumber(561), new BigNumber((_369_v4).length))])) : (new _dafny.CodePoint('w'.codePointAt(0)))), _376_v9);
          let _381_v14;
          let _nw68 = Array((new BigNumber(7)).toNumber());
          _nw68[(_dafny.ZERO).toNumber()] = new BigNumber((_368_v3).length);
          _nw68[(_dafny.ONE).toNumber()] = new BigNumber((_380_v13).length);
          _nw68[(new BigNumber(2)).toNumber()] = (_this).f12;
          _nw68[(new BigNumber(3)).toNumber()] = _367_i0;
          _nw68[(new BigNumber(4)).toNumber()] = (_this).f12;
          _nw68[(new BigNumber(5)).toNumber()] = _367_i0;
          _nw68[(new BigNumber(6)).toNumber()] = (_this).f12;
          _381_v14 = _nw68;
          let _382_v15;
          _382_v15 = _dafny.Seq.of(_381_v14, _381_v14, _381_v14);
          let _383_v16;
          let _nw69 = Array((new BigNumber(3)).toNumber());
          _nw69[(_dafny.ZERO).toNumber()] = (_382_v15)[_module.__default.safeIndex(_367_i0, new BigNumber((_382_v15).length))];
          _nw69[(_dafny.ONE).toNumber()] = _381_v14;
          _nw69[(new BigNumber(2)).toNumber()] = _381_v14;
          _383_v16 = _nw69;
          let _index47 = _module.__default.safeIndex(new BigNumber(113), new BigNumber((_383_v16).length));
          (_383_v16)[_index47] = _381_v14;
          let _384_v17;
          _384_v17 = _dafny.Set.fromElements((_369_v4)[_module.__default.safeIndex(new BigNumber(561), new BigNumber((_369_v4).length))], (_369_v4)[_module.__default.safeIndex(new BigNumber(561), new BigNumber((_369_v4).length))]);
          let _index48 = _module.__default.safeIndex(new BigNumber(807), new BigNumber((_378_v11).length));
          let _index49 = _module.__default.safeIndex(new BigNumber(113), new BigNumber((_383_v16).length));
          let _rhs38 = _module.__default.fm26(_module.__default.safeDivisionInt(_367_i0, _367_i0), new BigNumber(((_dafny.Set.fromElements((_this).f11)).Union(_384_v17)).length), (_369_v4)[_module.__default.safeIndex(new BigNumber(561), new BigNumber((_369_v4).length))], globalState);
          let _rhs39 = _dafny.Seq.Concat(_374_v7, _374_v7);
          let _rhs40 = (_dafny.ZERO).minus((((new BigNumber((_368_v3).length)).isLessThanOrEqualTo(_367_i0)) ? (new BigNumber((_dafny.Seq.Concat(_374_v7, _374_v7)).length)) : ((_this).f12)));
          let _rhs41 = _381_v14;
          let _rhs42 = _dafny.Seq.update(_374_v7, _module.__default.safeIndex((_this).f12, new BigNumber((_374_v7).length)), ((true) ? ((_this).f11) : ((_369_v4)[_module.__default.safeIndex(new BigNumber(561), new BigNumber((_369_v4).length))])));
          let _lhs19 = _378_v11;
          let _lhs20 = _module.__default.safeIndex(new BigNumber(807), new BigNumber((_378_v11).length));
          let _lhs21 = globalState;
          let _lhs22 = _383_v16;
          let _lhs23 = _module.__default.safeIndex(new BigNumber(113), new BigNumber((_383_v16).length));
          _376_v9 = _rhs38;
          _lhs19[_lhs20] = _rhs39;
          _lhs21.f2 = _rhs40;
          _lhs22[_lhs23] = _rhs41;
          _374_v7 = _rhs42;
          (globalState).f2 = _367_i0;
          let _index50 = _module.__default.safeIndex(new BigNumber(561), new BigNumber((_369_v4).length));
          (_369_v4)[_index50] = (_this).f11;
        } else {
          let _385_v18;
          let _init7 = function (_386_i3) {
            return (_386_i3).minus(new BigNumber((_dafny.Seq.UnicodeFromString("kisi")).length));
          };
          let _nw70 = Array((new BigNumber(7)).toNumber());
          for (let _i0_7 = 0; _i0_7 < new BigNumber(_nw70.length); _i0_7++) {
            _nw70[_i0_7] = _init7(new BigNumber(_i0_7));
          }
          _385_v18 = _nw70;
          let _387_v19;
          _387_v19 = _dafny.MultiSet.fromElements(_385_v18);
          (globalState).f2 = (((_387_v19).contains(_385_v18)) ? ((_387_v19).get(_385_v18)) : (new BigNumber((_dafny.Seq.UnicodeFromString("styupjoxw")).length)));
          let _388_v20;
          _388_v20 = _dafny.Seq.of((_this).f11);
          let _389_v21;
          let _nw71 = new _module.C0();
          _nw71.__ctor(_367_i0);
          _389_v21 = _nw71;
          let _390_v22;
          _390_v22 = _module.D4.create_DC8(_389_v21);
          let _391_v23;
          _391_v23 = _dafny.Map.Empty.slice().updateUnsafe((_389_v21).f13,_389_v21);
          let _392_v24;
          let _nw72 = Array((new BigNumber(27)).toNumber());
          _nw72[(_dafny.ZERO).toNumber()] = _389_v21;
          _nw72[(_dafny.ONE).toNumber()] = _389_v21;
          _nw72[(new BigNumber(2)).toNumber()] = _389_v21;
          _nw72[(new BigNumber(3)).toNumber()] = _389_v21;
          _nw72[(new BigNumber(4)).toNumber()] = _389_v21;
          _nw72[(new BigNumber(5)).toNumber()] = _389_v21;
          _nw72[(new BigNumber(6)).toNumber()] = _389_v21;
          _nw72[(new BigNumber(7)).toNumber()] = _389_v21;
          _nw72[(new BigNumber(8)).toNumber()] = _389_v21;
          _nw72[(new BigNumber(9)).toNumber()] = _389_v21;
          _nw72[(new BigNumber(10)).toNumber()] = (_390_v22).dtor_cf15;
          _nw72[(new BigNumber(11)).toNumber()] = _389_v21;
          _nw72[(new BigNumber(12)).toNumber()] = (((_391_v23).contains((_389_v21).f13)) ? ((_391_v23).get((_389_v21).f13)) : (_389_v21));
          _nw72[(new BigNumber(13)).toNumber()] = _389_v21;
          _nw72[(new BigNumber(14)).toNumber()] = _389_v21;
          _nw72[(new BigNumber(15)).toNumber()] = _389_v21;
          _nw72[(new BigNumber(16)).toNumber()] = _389_v21;
          _nw72[(new BigNumber(17)).toNumber()] = _389_v21;
          _nw72[(new BigNumber(18)).toNumber()] = _389_v21;
          _nw72[(new BigNumber(19)).toNumber()] = _389_v21;
          _nw72[(new BigNumber(20)).toNumber()] = _389_v21;
          _nw72[(new BigNumber(21)).toNumber()] = _389_v21;
          _nw72[(new BigNumber(22)).toNumber()] = _389_v21;
          _nw72[(new BigNumber(23)).toNumber()] = (((_391_v23).contains((_389_v21).f13)) ? ((_391_v23).get((_389_v21).f13)) : (_389_v21));
          _nw72[(new BigNumber(24)).toNumber()] = _389_v21;
          _nw72[(new BigNumber(25)).toNumber()] = _389_v21;
          _nw72[(new BigNumber(26)).toNumber()] = _389_v21;
          _392_v24 = _nw72;
          let _393_v25;
          _393_v25 = _dafny.Map.Empty.slice().updateUnsafe(_392_v24,(_389_v21).f13);
          let _394_v26;
          _394_v26 = _dafny.Set.fromElements((_this).f11, (new BigNumber((_388_v20).length)).isEqualTo(new BigNumber((_393_v25).length)), (_this).f11);
          let _395_v27;
          _395_v27 = _394_v26;
          _394_v26 = (_394_v26).Intersect((_395_v27));
          let _396_v28;
          _396_v28 = _dafny.Seq.of(_367_i0);
          (globalState).f2 = ((!(_dafny.Seq.IsProperPrefixOf(_396_v28, _dafny.Seq.Create(_module.__default.abs(new BigNumber(297)), function (_397_i4) {
            return new BigNumber(885);
          })))) ? ((_this).f12) : ((_this).f12));
          let _398_v29;
          _398_v29 = _dafny.MultiSet.fromElements((_this).f11);
          (globalState).f2 = (((_398_v29).contains(true)) ? ((_398_v29).get(true)) : (_367_i0));
          r2 = (_this).f11;
        }
        let _399_v30;
        let _init8 = function (_400_i5) {
          return (_400_i5).multipliedBy((_this).f12);
        };
        let _nw73 = Array((new BigNumber(12)).toNumber());
        for (let _i0_8 = 0; _i0_8 < new BigNumber(_nw73.length); _i0_8++) {
          _nw73[_i0_8] = _init8(new BigNumber(_i0_8));
        }
        _399_v30 = _nw73;
        _399_v30 = _399_v30;
      }
      let _401_v31;
      let _nw74 = Array((new BigNumber(5)).toNumber()).fill(_dafny.ZERO);
      _401_v31 = _nw74;
      let _402_v32;
      let _nw75 = Array((new BigNumber(22)).toNumber());
      _nw75[(_dafny.ZERO).toNumber()] = _401_v31;
      _nw75[(_dafny.ONE).toNumber()] = _401_v31;
      _nw75[(new BigNumber(2)).toNumber()] = _401_v31;
      _nw75[(new BigNumber(3)).toNumber()] = _401_v31;
      _nw75[(new BigNumber(4)).toNumber()] = _401_v31;
      _nw75[(new BigNumber(5)).toNumber()] = _401_v31;
      _nw75[(new BigNumber(6)).toNumber()] = _401_v31;
      _nw75[(new BigNumber(7)).toNumber()] = _401_v31;
      _nw75[(new BigNumber(8)).toNumber()] = _401_v31;
      _nw75[(new BigNumber(9)).toNumber()] = _401_v31;
      _nw75[(new BigNumber(10)).toNumber()] = _401_v31;
      _nw75[(new BigNumber(11)).toNumber()] = _401_v31;
      _nw75[(new BigNumber(12)).toNumber()] = _401_v31;
      _nw75[(new BigNumber(13)).toNumber()] = _401_v31;
      _nw75[(new BigNumber(14)).toNumber()] = _401_v31;
      _nw75[(new BigNumber(15)).toNumber()] = _401_v31;
      _nw75[(new BigNumber(16)).toNumber()] = _401_v31;
      _nw75[(new BigNumber(17)).toNumber()] = _401_v31;
      _nw75[(new BigNumber(18)).toNumber()] = _401_v31;
      _nw75[(new BigNumber(19)).toNumber()] = _401_v31;
      _nw75[(new BigNumber(20)).toNumber()] = _401_v31;
      _nw75[(new BigNumber(21)).toNumber()] = _401_v31;
      _402_v32 = _nw75;
      let _index51 = _module.__default.safeIndex(new BigNumber(923), new BigNumber((_402_v32).length));
      (_402_v32)[_index51] = _401_v31;
      let _403_v33;
      let _init9 = function (_404_i6) {
        return (_module.D6.create_DC11((_module.D6.create_DC11(_dafny.Seq.UnicodeFromString("fwqyvo"))).dtor_cf22)).dtor_cf22;
      };
      let _nw76 = Array((new BigNumber(24)).toNumber());
      for (let _i0_9 = 0; _i0_9 < new BigNumber(_nw76.length); _i0_9++) {
        _nw76[_i0_9] = _init9(new BigNumber(_i0_9));
      }
      _403_v33 = _nw76;
      let _405_v34;
      _405_v34 = _dafny.Seq.UnicodeFromString("g");
      let _406_v35;
      _406_v35 = new _dafny.CodePoint('x'.codePointAt(0));
      let _407_v36;
      _407_v36 = _module.D6.create_DC11(_dafny.Seq.update(_405_v34, _module.__default.safeIndex(new BigNumber(261), new BigNumber((_405_v34).length)), _406_v35));
      let _index52 = _module.__default.safeIndex(new BigNumber(772), new BigNumber((_403_v33).length));
      (_403_v33)[_index52] = (_407_v36).dtor_cf22;
      let _index53 = _module.__default.safeIndex(new BigNumber(923), new BigNumber((_402_v32).length));
      let _index54 = _module.__default.safeIndex(new BigNumber(772), new BigNumber((_403_v33).length));
      let _rhs43 = (((_this).f11) ? (_401_v31) : (_401_v31));
      let _rhs44 = _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(48)), ((_408_v35) => function (_409_i7) {
        return _408_v35;
      })(_406_v35)), _405_v34), _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(569)), ((_410_v35) => function (_411_i8) {
        return _410_v35;
      })(_406_v35)), _405_v34));
      let _rhs45 = (_this).f12;
      let _rhs46 = (_this).f11;
      let _lhs24 = _402_v32;
      let _lhs25 = _module.__default.safeIndex(new BigNumber(923), new BigNumber((_402_v32).length));
      let _lhs26 = _403_v33;
      let _lhs27 = _module.__default.safeIndex(new BigNumber(772), new BigNumber((_403_v33).length));
      _lhs24[_lhs25] = _rhs43;
      _lhs26[_lhs27] = _rhs44;
      r1 = _rhs45;
      r2 = _rhs46;
      let _412_v37;
      let _413_v38;
      let _414_v39;
      let _out20;
      let _out21;
      let _out22;
      let _outcollector8 = _module.__default.m0((_this).f12, globalState);
      _out20 = _outcollector8[0];
      _out21 = _outcollector8[1];
      _out22 = _outcollector8[2];
      _412_v37 = _out20;
      _413_v38 = _out21;
      _414_v39 = _out22;
      let _415_v40;
      _415_v40 = _dafny.Seq.of((_this).f12);
      let _hi3 = new BigNumber((_415_v40).length);
      for (let _416_i9 = _414_v39; _416_i9.isLessThan(_hi3); _416_i9 = _416_i9.plus(_dafny.ONE)) {
        let _417_v41;
        _417_v41 = _dafny.Seq.of(_module.__default.fm27(_415_v40, !(_412_v37), _412_v37, globalState));
        r0 = new BigNumber((_417_v41).length);
        let _arr0 = (_402_v32)[_module.__default.safeIndex(new BigNumber(923), new BigNumber((_402_v32).length))];
        let _index55 = _module.__default.safeIndex(new BigNumber(816), new BigNumber(((_402_v32)[_module.__default.safeIndex(new BigNumber(923), new BigNumber((_402_v32).length))]).length));
        _arr0[_index55] = new BigNumber(483);
        let _418_v42;
        _418_v42 = _dafny.Seq.of((_363_v0).update(_413_v38, false));
        let _419_v43;
        _419_v43 = _module.D1.create_DC1(_418_v42);
        let _arr1 = (_402_v32)[_module.__default.safeIndex(new BigNumber(923), new BigNumber((_402_v32).length))];
        let _index56 = _module.__default.safeIndex(new BigNumber(816), new BigNumber(((_402_v32)[_module.__default.safeIndex(new BigNumber(923), new BigNumber((_402_v32).length))]).length));
        let _rhs47 = new _dafny.CodePoint('r'.codePointAt(0));
        let _rhs48 = new BigNumber(707);
        let _rhs49 = _419_v43;
        let _lhs28 = (_402_v32)[_module.__default.safeIndex(new BigNumber(923), new BigNumber((_402_v32).length))];
        let _lhs29 = _module.__default.safeIndex(new BigNumber(816), new BigNumber(((_402_v32)[_module.__default.safeIndex(new BigNumber(923), new BigNumber((_402_v32).length))]).length));
        _406_v35 = _rhs47;
        _lhs28[_lhs29] = _rhs48;
        _419_v43 = _rhs49;
        _414_v39 = _module.__default.safeModuloInt(((_402_v32)[_module.__default.safeIndex(new BigNumber(923), new BigNumber((_402_v32).length))])[_module.__default.safeIndex(new BigNumber(816), new BigNumber(((_402_v32)[_module.__default.safeIndex(new BigNumber(923), new BigNumber((_402_v32).length))]).length))], (((_402_v32)[_module.__default.safeIndex(new BigNumber(923), new BigNumber((_402_v32).length))])[_module.__default.safeIndex(new BigNumber(816), new BigNumber(((_402_v32)[_module.__default.safeIndex(new BigNumber(923), new BigNumber((_402_v32).length))]).length))]).multipliedBy(new BigNumber((_415_v40).length)));
        let _420_v44;
        _420_v44 = _dafny.Set.fromElements(_401_v31);
        let _421_v45;
        _421_v45 = _module.D1.create_DC2(((_402_v32)[_module.__default.safeIndex(new BigNumber(923), new BigNumber((_402_v32).length))])[_module.__default.safeIndex(new BigNumber(816), new BigNumber(((_402_v32)[_module.__default.safeIndex(new BigNumber(923), new BigNumber((_402_v32).length))]).length))], _412_v37, ((_402_v32)[_module.__default.safeIndex(new BigNumber(923), new BigNumber((_402_v32).length))])[_module.__default.safeIndex(new BigNumber(816), new BigNumber(((_402_v32)[_module.__default.safeIndex(new BigNumber(923), new BigNumber((_402_v32).length))]).length))]);
        let _arr2 = (_402_v32)[_module.__default.safeIndex(new BigNumber(923), new BigNumber((_402_v32).length))];
        let _index57 = _module.__default.safeIndex(new BigNumber(816), new BigNumber(((_402_v32)[_module.__default.safeIndex(new BigNumber(923), new BigNumber((_402_v32).length))]).length));
        _arr2[_index57] = (new BigNumber((_420_v44).length)).plus((_421_v45).dtor_cf4);
      }
      let _pat_let_tv2 = _414_v39;
      let _pat_let_tv3 = _414_v39;
      let _pat_let_tv4 = _414_v39;
      let _pat_let_tv5 = _414_v39;
      let _pat_let_tv6 = _414_v39;
      let _pat_let_tv7 = _414_v39;
      let _pat_let_tv8 = _414_v39;
      let _pat_let_tv9 = _414_v39;
      let _pat_let_tv10 = _414_v39;
      let _source7 = function (_source8) {
        if (_source8.is_DC2) {
          let _422___mcc_h7 = (_source8).cf2;
          let _423___mcc_h8 = (_source8).cf3;
          let _424___mcc_h9 = (_source8).cf4;
          let _425_cf4 = _424___mcc_h9;
          let _426_cf3 = _423___mcc_h8;
          let _427_cf2 = _422___mcc_h7;
          return _module.D2.create_DC3(function () {
  let _coll16 = new _dafny.Map();
  for (const _compr_16 of (_dafny.Set.fromElements(_dafny.MultiSet.fromElements(_pat_let_tv2), _dafny.MultiSet.fromElements(_pat_let_tv3, _pat_let_tv4, (_this).f12, _pat_let_tv5))).Elements) {
    let _428_v46 = _compr_16;
    if ((_dafny.Set.fromElements(_dafny.MultiSet.fromElements(_pat_let_tv6), _dafny.MultiSet.fromElements(_pat_let_tv7, _pat_let_tv8, (_this).f12, _pat_let_tv9))).contains(_428_v46)) {
      _coll16.push([_428_v46,_427_cf2]);
    }
  }
  return _coll16;
}());
        } else {
          let _429___mcc_h10 = (_source8).cf1;
          let _430_cf1 = _429___mcc_h10;
          return _module.D2.create_DC3(_dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements(new BigNumber(-146)),_pat_let_tv10));
        }
      }(_module.D1.create_DC2(_414_v39, (_this).f11, (_this).f12));
      if (_source7.is_DC4) {
        let _431___mcc_h0 = (_source7).cf6;
        let _432_cf6 = _431___mcc_h0;
        _412_v37 = (_412_v37) === (false);
        let _433_v47;
        let _nw77 = new _module.C0();
        _nw77.__ctor(_432_cf6);
        _433_v47 = _nw77;
        let _index58 = _module.__default.safeIndex(new BigNumber(311), new BigNumber((_401_v31).length));
        (_401_v31)[_index58] = new BigNumber(809);
        let _index59 = _module.__default.safeIndex(new BigNumber(311), new BigNumber((_401_v31).length));
        (_401_v31)[_index59] = (_414_v39).multipliedBy((_432_cf6).multipliedBy(_432_cf6));
        _432_cf6 = (_dafny.ZERO).minus((_433_v47).f13);
      } else if (_source7.is_DC5) {
        let _434___mcc_h1 = (_source7).cf7;
        let _435___mcc_h2 = (_source7).cf8;
        let _436___mcc_h3 = (_source7).cf9;
        let _437___mcc_h4 = (_source7).cf10;
        let _438___mcc_h5 = (_source7).cf11;
        let _439_cf11 = _438___mcc_h5;
        let _440_cf10 = _437___mcc_h4;
        let _441_cf9 = _436___mcc_h3;
        let _442_cf8 = _435___mcc_h2;
        let _443_cf7 = _434___mcc_h1;
        let _444_v48;
        _444_v48 = _dafny.Seq.of(_405_v34);
        let _445_v49;
        _445_v49 = _dafny.Map.Empty.slice().updateUnsafe(false,_406_v35);
        let _446_v50;
        let _nw78 = Array((new BigNumber(12)).toNumber());
        _nw78[(_dafny.ZERO).toNumber()] = _441_cf9;
        _nw78[(_dafny.ONE).toNumber()] = !(true);
        _nw78[(new BigNumber(2)).toNumber()] = _412_v37;
        _nw78[(new BigNumber(3)).toNumber()] = _440_cf10;
        _nw78[(new BigNumber(4)).toNumber()] = _413_v38;
        _nw78[(new BigNumber(5)).toNumber()] = _412_v37;
        _nw78[(new BigNumber(6)).toNumber()] = !(_439_cf11);
        _nw78[(new BigNumber(7)).toNumber()] = _412_v37;
        _nw78[(new BigNumber(8)).toNumber()] = _442_cf8;
        _nw78[(new BigNumber(9)).toNumber()] = _440_cf10;
        _nw78[(new BigNumber(10)).toNumber()] = !_dafny.areEqual(_444_v48, _444_v48);
        _nw78[(new BigNumber(11)).toNumber()] = _dafny.Seq.contains(_dafny.Seq.UnicodeFromString("odq"), (((_445_v49).contains((((_363_v0).contains(_439_cf11)) ? ((_363_v0).get(_439_cf11)) : (_440_cf10)))) ? ((_445_v49).get((((_363_v0).contains(_439_cf11)) ? ((_363_v0).get(_439_cf11)) : (_440_cf10)))) : (_406_v35)));
        _446_v50 = _nw78;
        let _index60 = _module.__default.safeIndex(new BigNumber(510), new BigNumber((_446_v50).length));
        (_446_v50)[_index60] = _439_cf11;
        let _447_v52;
        _447_v52 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(((_403_v33)[_module.__default.safeIndex(new BigNumber(772), new BigNumber((_403_v33).length))]).length),_413_v38);
        let _index61 = _module.__default.safeIndex(new BigNumber(510), new BigNumber((_446_v50).length));
        let _rhs50 = _443_cf7;
        let _rhs51 = (_dafny.ZERO).minus(new BigNumber((((function () {
          let _coll17 = new _dafny.Map();
          for (const _compr_17 of _dafny.IntegerRange(new BigNumber(865), new BigNumber(185))) {
            let _448_v51 = _compr_17;
            if (((new BigNumber(865)).isLessThanOrEqualTo(_448_v51)) && ((_448_v51).isLessThan(new BigNumber(185)))) {
              _coll17.push([_module.__default.safeDivisionInt(_448_v51, new BigNumber(235)),true]);
            }
          }
          return _coll17;
        }()).Merge(((_447_v52).update((_this).f12, _439_cf11)).update(new BigNumber((_415_v40).length), _module.__default.fm23((_403_v33)[_module.__default.safeIndex(new BigNumber(772), new BigNumber((_403_v33).length))], (_this).f12, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(145)), ((_449_v35) => function (_450_i10) {
          return _449_v35;
        })(_406_v35))).length), _439_cf11, globalState)))).Merge((_447_v52).update((_this).f12, _439_cf11))).length));
        let _rhs52 = _442_cf8;
        let _rhs53 = true;
        let _lhs30 = _446_v50;
        let _lhs31 = _module.__default.safeIndex(new BigNumber(510), new BigNumber((_446_v50).length));
        r1 = _rhs50;
        r1 = _rhs51;
        _lhs30[_lhs31] = _rhs52;
        _441_cf9 = _rhs53;
        let _451_v53;
        _451_v53 = _module.D3.create_DC6(_406_v35);
        _406_v35 = (_451_v53).dtor_cf12;
        r2 = (_446_v50)[_module.__default.safeIndex(new BigNumber(510), new BigNumber((_446_v50).length))];
        let _452_v54;
        _452_v54 = _module.D2.create_DC5(_443_cf7, _439_cf11, (_this).f11, _439_cf11, _442_cf8);
        let _pat_let_tv11 = _439_cf11;
        let _pat_let_tv12 = _446_v50;
        let _pat_let_tv13 = _446_v50;
        _442_cf8 = (function (_pat_let0_0) {
          return function (_453_dt__update__tmp_h0) {
            return function (_pat_let1_0) {
              return function (_454_dt__update_hcf9_h0) {
                return function (_pat_let2_0) {
                  return function (_455_dt__update_hcf8_h0) {
                    return function (_pat_let3_0) {
                      return function (_456_dt__update_hcf11_h0) {
                        return _module.D2.create_DC5((_453_dt__update__tmp_h0).dtor_cf7, _455_dt__update_hcf8_h0, _454_dt__update_hcf9_h0, (_453_dt__update__tmp_h0).dtor_cf10, _456_dt__update_hcf11_h0);
                      }(_pat_let3_0);
                    }((_this).f11);
                  }(_pat_let2_0);
                }((_pat_let_tv13)[_module.__default.safeIndex(new BigNumber(510), new BigNumber((_pat_let_tv12).length))]);
              }(_pat_let1_0);
            }(!(!(_pat_let_tv11)));
          }(_pat_let0_0);
        }(_452_v54)).dtor_cf11;
      } else {
        let _457___mcc_h6 = (_source7).cf5;
        let _458_cf5 = _457___mcc_h6;
        r1 = _414_v39;
        let _459_v55;
        let _nw79 = new _module.C0();
        _nw79.__ctor((_this).f12);
        _459_v55 = _nw79;
        let _460_v56;
        _460_v56 = _dafny.Map.Empty.slice().updateUnsafe((_this).f12,(_459_v55).f13);
        _413_v38 = ((_460_v56).contains(_414_v39)) && ((_this).f11);
        r1 = (_this).f12;
      }
      r0 = _414_v39;
      let _461_v57;
      _461_v57 = _dafny.MultiSet.fromElements(true);
      let _462_v58;
      _462_v58 = _dafny.Seq.of((_this).f11, _412_v37);
      r1 = ((_414_v39).minus(new BigNumber((_461_v57).cardinality()))).multipliedBy(new BigNumber((_dafny.Seq.update((_403_v33)[_module.__default.safeIndex(new BigNumber(772), new BigNumber((_403_v33).length))], _module.__default.safeIndex(new BigNumber((_462_v58).length), new BigNumber(((_403_v33)[_module.__default.safeIndex(new BigNumber(772), new BigNumber((_403_v33).length))]).length)), _406_v35)).length));
      r2 = _412_v37;
      return [r0, r1, r2];
    }
    m7(p0, p1, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = _dafny.Seq.UnicodeFromString("");
      let r2 = _dafny.ZERO;
      let _463_v0;
      let _nw80 = Array((new BigNumber(13)).toNumber()).fill(false);
      _463_v0 = _nw80;
      let _464_v1;
      _464_v1 = _dafny.Seq.of((_this).f12);
      let _465_v2;
      _465_v2 = _dafny.Map.Empty.slice().updateUnsafe(_463_v0,_464_v1);
      _465_v2 = ((_465_v2).Merge(_465_v2)).Merge((_465_v2).Merge(_465_v2));
      if (p1) {
        let _466_v3;
        _466_v3 = _dafny.Seq.UnicodeFromString("jmjjfn");
        r1 = _dafny.Seq.Concat(_466_v3, _466_v3);
        let _467_v4;
        let _nw81 = Array((new BigNumber(6)).toNumber());
        _nw81[(_dafny.ZERO).toNumber()] = _463_v0;
        _nw81[(_dafny.ONE).toNumber()] = _463_v0;
        _nw81[(new BigNumber(2)).toNumber()] = _463_v0;
        _nw81[(new BigNumber(3)).toNumber()] = _463_v0;
        _nw81[(new BigNumber(4)).toNumber()] = _463_v0;
        _nw81[(new BigNumber(5)).toNumber()] = _463_v0;
        _467_v4 = _nw81;
        let _index62 = _module.__default.safeIndex(new BigNumber(881), new BigNumber((_467_v4).length));
        (_467_v4)[_index62] = _463_v0;
        let _468_v5;
        _468_v5 = _dafny.Map.Empty.slice().updateUnsafe(((_this).f12).plus((_this).f12),_463_v0);
        let _index63 = _module.__default.safeIndex(new BigNumber(881), new BigNumber((_467_v4).length));
        (_467_v4)[_index63] = (((_468_v5).contains((_this).f12)) ? ((_468_v5).get((_this).f12)) : (_463_v0));
        let _469_v6;
        _469_v6 = _dafny.Map.Empty.slice().updateUnsafe(!(p0),p1);
        let _470_v7;
        _470_v7 = _dafny.Set.fromElements(_469_v6, _469_v6, _469_v6);
        _470_v7 = _470_v7;
        let _471_v8;
        _471_v8 = true;
        let _472_v9;
        _472_v9 = _dafny.Seq.of(_469_v6);
        _471_v8 = _module.__default.fm23(_dafny.Seq.Create(_module.__default.abs(new BigNumber(140)), function (_473_i0) {
          return new _dafny.CodePoint('p'.codePointAt(0));
        }), new BigNumber((_472_v9).length), (_this).f12, p0, globalState);
        _471_v8 = ((_this).f12).isLessThan(((p0) ? ((_this).f12) : (new BigNumber(941))));
      } else {
        let _474_v10;
        _474_v10 = _dafny.Set.fromElements((_this).f12);
        (globalState).f2 = _module.__default.fm24((_474_v10).Difference(_474_v10), globalState);
        let _475_v11;
        let _nw82 = new _module.C0();
        _nw82.__ctor((_this).f12);
        _475_v11 = _nw82;
        let _476_v12;
        _476_v12 = _dafny.Seq.of(_475_v11);
        let _477_v13;
        let _nw83 = Array((new BigNumber(6)).toNumber());
        _nw83[(_dafny.ZERO).toNumber()] = new BigNumber((_476_v12).length);
        _nw83[(_dafny.ONE).toNumber()] = _module.__default.safeModuloInt((_475_v11).f13, (_475_v11).f13);
        _nw83[(new BigNumber(2)).toNumber()] = ((_this).f12).minus((_dafny.ZERO).minus((_475_v11).f13));
        _nw83[(new BigNumber(3)).toNumber()] = new BigNumber((_dafny.Seq.update(_dafny.Seq.Concat(_464_v1, _dafny.Seq.of((_475_v11).f13, new BigNumber((_464_v1).length), (_this).f12)), _module.__default.safeIndex((_475_v11).f13, new BigNumber((_dafny.Seq.Concat(_464_v1, _dafny.Seq.of((_475_v11).f13, new BigNumber((_464_v1).length), (_this).f12))).length)), (_475_v11).f13)).length);
        _nw83[(new BigNumber(4)).toNumber()] = (_this).f12;
        _nw83[(new BigNumber(5)).toNumber()] = (_this).f12;
        _477_v13 = _nw83;
        let _index64 = _module.__default.safeIndex(new BigNumber(352), new BigNumber((_477_v13).length));
        (_477_v13)[_index64] = (_this).f12;
        let _478_v14;
        _478_v14 = _module.D3.create_DC7((_this).f12, (new BigNumber(316)).multipliedBy((_this).f12));
        let _479_v15;
        _479_v15 = _dafny.MultiSet.fromElements((_475_v11).f13, (_this).f12);
        let _480_v16;
        _480_v16 = _dafny.Map.Empty.slice().updateUnsafe(_479_v15,new BigNumber(-719));
        let _481_v17;
        _481_v17 = _module.D2.create_DC3(_480_v16);
        let _482_v18;
        _482_v18 = _module.D2.create_DC3((_481_v17).dtor_cf5);
        let _483_v19;
        _483_v19 = _module.D4.create_DC9((_this).f12, new _dafny.CodePoint('e'.codePointAt(0)), _464_v1, (_475_v11).f13, (_this).f11);
        let _484_v20;
        _484_v20 = _dafny.Seq.of(_483_v19);
        let _485_v21;
        _485_v21 = _dafny.Seq.of(_484_v20);
        let _index65 = _module.__default.safeIndex(new BigNumber(352), new BigNumber((_477_v13).length));
        let _rhs54 = _module.__default.safeModuloInt((_475_v11).f13, (_475_v11).f13);
        let _rhs55 = new BigNumber((_485_v21).length);
        let _rhs56 = _478_v14;
        let _rhs57 = _481_v17;
        let _lhs32 = _477_v13;
        let _lhs33 = _module.__default.safeIndex(new BigNumber(352), new BigNumber((_477_v13).length));
        r0 = _rhs54;
        _lhs32[_lhs33] = _rhs55;
        _478_v14 = _rhs56;
        _481_v17 = _rhs57;
        let _486_v22;
        let _487_v23;
        let _488_v24;
        let _out23;
        let _out24;
        let _out25;
        let _outcollector9 = (_this).m6(globalState);
        _out23 = _outcollector9[0];
        _out24 = _outcollector9[1];
        _out25 = _outcollector9[2];
        _486_v22 = _out23;
        _487_v23 = _out24;
        _488_v24 = _out25;
        _488_v24 = !(((p1) ? ((_this).f11) : (true)));
        let _489_v25;
        let _nw84 = Array((new BigNumber(7)).toNumber()).fill([]);
        _489_v25 = _nw84;
        let _index66 = _module.__default.safeIndex(new BigNumber(966), new BigNumber((_489_v25).length));
        (_489_v25)[_index66] = _477_v13;
        let _index67 = _module.__default.safeIndex(new BigNumber(251), new BigNumber((_463_v0).length));
        (_463_v0)[_index67] = !(!(false));
        let _index68 = _module.__default.safeIndex(new BigNumber(352), new BigNumber((_477_v13).length));
        let _index69 = _module.__default.safeIndex(new BigNumber(966), new BigNumber((_489_v25).length));
        let _index70 = _module.__default.safeIndex(new BigNumber(251), new BigNumber((_463_v0).length));
        let _rhs58 = new BigNumber(365);
        let _rhs59 = _477_v13;
        let _rhs60 = p1;
        let _lhs34 = _477_v13;
        let _lhs35 = _module.__default.safeIndex(new BigNumber(352), new BigNumber((_477_v13).length));
        let _lhs36 = _489_v25;
        let _lhs37 = _module.__default.safeIndex(new BigNumber(966), new BigNumber((_489_v25).length));
        let _lhs38 = _463_v0;
        let _lhs39 = _module.__default.safeIndex(new BigNumber(251), new BigNumber((_463_v0).length));
        _lhs34[_lhs35] = _rhs58;
        _lhs36[_lhs37] = _rhs59;
        _lhs38[_lhs39] = _rhs60;
      }
      let _490_v26;
      _490_v26 = _dafny.Seq.UnicodeFromString("vtttnk");
      let _491_v28;
      let _init10 = function (_492_i1) {
        return _module.D2.create_DC3(function () {
  let _coll18 = new _dafny.Map();
  for (const _compr_18 of (_dafny.MultiSet.fromElements(_dafny.MultiSet.fromElements(new BigNumber(91)), _dafny.MultiSet.fromElements((_this).f12))).Elements) {
    let _493_v27 = _compr_18;
    if ((_dafny.MultiSet.fromElements(_dafny.MultiSet.fromElements(new BigNumber(91)), _dafny.MultiSet.fromElements((_this).f12))).contains(_493_v27)) {
      _coll18.push([_493_v27,(_this).f12]);
    }
  }
  return _coll18;
}());
      };
      let _nw85 = Array((new BigNumber(24)).toNumber());
      for (let _i0_10 = 0; _i0_10 < new BigNumber(_nw85.length); _i0_10++) {
        _nw85[_i0_10] = _init10(new BigNumber(_i0_10));
      }
      _491_v28 = _nw85;
      let _494_v29;
      _494_v29 = _dafny.Map.Empty.slice().updateUnsafe(_491_v28,_490_v26);
      let _495_v30;
      _495_v30 = new _dafny.CodePoint('u'.codePointAt(0));
      r1 = _dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.Concat(_module.__default.fm25(!((_this).f11), (_this).f12, globalState), _490_v26), (((_494_v29).contains(_491_v28)) ? ((_494_v29).get(_491_v28)) : (_dafny.Seq.UnicodeFromString("i")))), _module.__default.safeIndex((_this).f12, new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Concat(_module.__default.fm25(!((_this).f11), (_this).f12, globalState), _490_v26), (((_494_v29).contains(_491_v28)) ? ((_494_v29).get(_491_v28)) : (_dafny.Seq.UnicodeFromString("i"))))).length)), _495_v30);
      (globalState).f2 = new BigNumber(238);
      let _496_v31;
      _496_v31 = _dafny.Map.Empty.slice().updateUnsafe(p1,_463_v0);
      let _497_v32;
      _497_v32 = _dafny.Map.Empty.slice().updateUnsafe((_this).f11,(((_496_v31).contains((_this).f11)) ? ((_496_v31).get((_this).f11)) : (_463_v0)));
      _496_v31 = _497_v32;
      let _498_v33;
      _498_v33 = _dafny.Set.fromElements(new BigNumber(616), (_this).f12);
      let _499_v35;
      _499_v35 = _module.D1.create_DC2(((_this).f12).plus(new BigNumber(611)), !((_498_v33).equals(function () {
  let _coll19 = new _dafny.Set();
  for (const _compr_19 of _dafny.IntegerRange(new BigNumber(835), new BigNumber(567))) {
    let _500_v34 = _compr_19;
    if (((new BigNumber(835)).isLessThanOrEqualTo(_500_v34)) && ((_500_v34).isLessThan(new BigNumber(567)))) {
      _coll19.add((_500_v34).minus(new BigNumber(97)));
    }
  }
  return _coll19;
}())), (_this).f12);
      let _source9 = _499_v35;
      if (_source9.is_DC2) {
        let _501___mcc_h0 = (_source9).cf2;
        let _502___mcc_h1 = (_source9).cf3;
        let _503___mcc_h2 = (_source9).cf4;
        let _504_cf4 = _503___mcc_h2;
        let _505_cf3 = _502___mcc_h1;
        let _506_cf2 = _501___mcc_h0;
        _505_cf3 = (_this).f11;
        let _507_v36;
        let _nw86 = new _module.C0();
        _nw86.__ctor(new BigNumber(-580));
        _507_v36 = _nw86;
        let _508_v37;
        _508_v37 = _dafny.Set.fromElements(_495_v30, _495_v30);
        let _509_v38;
        _509_v38 = _module.D7.create_DC13(_508_v37);
        _508_v37 = (_508_v37).Union((_509_v38).dtor_cf25);
        let _510_v39;
        _510_v39 = _dafny.Map.Empty.slice().updateUnsafe((_this).f12,new BigNumber((_490_v26).length));
        let _511_v40;
        _511_v40 = _dafny.MultiSet.fromElements(_510_v39);
        let _512_v41;
        _512_v41 = _module.D2.create_DC4((((_511_v40).contains(_module.__default.fm28(_506_cf2, globalState))) ? ((_511_v40).get(_module.__default.fm28(_506_cf2, globalState))) : (_506_cf2)));
        let _513_v42;
        _513_v42 = _dafny.Seq.of((_this).f11, (_this).f11);
        let _514_v43;
        _514_v43 = _dafny.Map.Empty.slice().updateUnsafe(_module.D7.create_DC13(_508_v37),new BigNumber(-23));
        _512_v41 = (((_513_v42)[_module.__default.safeIndex(new BigNumber((_514_v43).length), new BigNumber((_513_v42).length))]) ? (((p1) ? (_512_v41) : (_module.D2.create_DC4((_this).f12)))) : (_512_v41));
      } else {
        let _515___mcc_h3 = (_source9).cf1;
        let _516_cf1 = _515___mcc_h3;
        if (p0) {
          let _517_v44;
          _517_v44 = _dafny.Set.fromElements(_dafny.Seq.UnicodeFromString("c"), _490_v26);
          let _518_v45;
          _518_v45 = _module.D6.create_DC12((_this).f11, _module.__default.fm26((_this).f12, (_this).f12, (_this).f11, globalState));
          _490_v26 = _dafny.Seq.Concat(_dafny.Seq.update(_dafny.Seq.UnicodeFromString("xramcle"), _module.__default.safeIndex(new BigNumber((_517_v44).length), new BigNumber((_dafny.Seq.UnicodeFromString("xramcle")).length)), (_518_v45).dtor_cf24), _490_v26);
          let _519_v46;
          let _nw87 = new _module.C0();
          _nw87.__ctor((_this).f12);
          _519_v46 = _nw87;
          let _520_v47;
          _520_v47 = false;
          let _521_v48;
          _521_v48 = _dafny.Seq.of(_519_v46, _519_v46, _519_v46);
          let _522_v49;
          _522_v49 = _dafny.Map.Empty.slice().updateUnsafe((_519_v46).f13,_490_v26);
          let _index71 = _module.__default.safeIndex(new BigNumber(451), new BigNumber((_463_v0).length));
          (_463_v0)[_index71] = p1;
          let _523_v50;
          _523_v50 = _module.D8.create_DC15(_463_v0);
          let _524_v51;
          _524_v51 = _dafny.Map.Empty.slice().updateUnsafe((_523_v50).dtor_cf27,p1);
          let _index72 = _module.__default.safeIndex(new BigNumber(451), new BigNumber((_463_v0).length));
          let _rhs61 = _520_v47;
          let _rhs62 = _dafny.Seq.Concat(_521_v48, _dafny.Seq.update(_521_v48, _module.__default.safeIndex((_this).f12, new BigNumber((_521_v48).length)), (_521_v48)[_module.__default.safeIndex((_519_v46).f13, new BigNumber((_521_v48).length))]));
          let _rhs63 = _dafny.Map.Empty.slice().updateUnsafe((_this).f12,_module.__default.fm25((((_524_v51).contains(_463_v0)) ? ((_524_v51).get(_463_v0)) : (p0)), new BigNumber(185), globalState));
          let _rhs64 = true;
          let _lhs40 = _463_v0;
          let _lhs41 = _module.__default.safeIndex(new BigNumber(451), new BigNumber((_463_v0).length));
          _520_v47 = _rhs61;
          _521_v48 = _rhs62;
          _522_v49 = _rhs63;
          _lhs40[_lhs41] = _rhs64;
          r0 = ((_this).f12).plus(new BigNumber((_dafny.Seq.Concat(_490_v26, _490_v26)).length));
          let _525_v52;
          _525_v52 = _dafny.MultiSet.fromElements((_519_v46).f13);
          let _526_v53;
          _526_v53 = _dafny.Seq.of((_525_v52).IsDisjointFrom(_dafny.MultiSet.fromElements((_519_v46).f13)), p1, p0, (_this).f11);
          _526_v53 = ((false) ? (_526_v53) : (_dafny.Seq.of(p1, _520_v47, p1)));
        } else {
          let _527_v54;
          _527_v54 = _dafny.Map.Empty.slice().updateUnsafe((_this).f12,_dafny.Seq.IsProperPrefixOf(_490_v26, _490_v26));
          let _528_v55;
          _528_v55 = _dafny.MultiSet.fromElements(_464_v1, _464_v1, _dafny.Seq.of(new BigNumber(-84)));
          _527_v54 = (_527_v54).update((((_528_v55).contains(_464_v1)) ? ((_528_v55).get(_464_v1)) : ((_this).f12)), (_this).f11);
          let _529_v56;
          let _530_v57;
          let _531_v58;
          let _out26;
          let _out27;
          let _out28;
          let _outcollector10 = _module.__default.m0((_this).f12, globalState);
          _out26 = _outcollector10[0];
          _out27 = _outcollector10[1];
          _out28 = _outcollector10[2];
          _529_v56 = _out26;
          _530_v57 = _out27;
          _531_v58 = _out28;
          _529_v56 = (((_530_v57) ? (_498_v33) : (_498_v33))).IsProperSubsetOf(_498_v33);
          _463_v0 = _463_v0;
          let _532_v59;
          let _533_v60;
          let _534_v61;
          let _out29;
          let _out30;
          let _out31;
          let _outcollector11 = _module.__default.m0((_this).f12, globalState);
          _out29 = _outcollector11[0];
          _out30 = _outcollector11[1];
          _out31 = _outcollector11[2];
          _532_v59 = _out29;
          _533_v60 = _out30;
          _534_v61 = _out31;
        }
        let _535_v62;
        _535_v62 = _module.D8.create_DC15(_463_v0);
        let _pat_let_tv14 = _463_v0;
        let _source10 = function (_pat_let4_0) {
          return function (_536_dt__update__tmp_h0) {
            return function (_pat_let5_0) {
              return function (_537_dt__update_hcf27_h0) {
                return _module.D8.create_DC15(_537_dt__update_hcf27_h0);
              }(_pat_let5_0);
            }(_pat_let_tv14);
          }(_pat_let4_0);
        }(_535_v62);
        if (_source10.is_DC16) {
          let _538___mcc_h4 = (_source10).cf28;
          let _539___mcc_h5 = (_source10).cf29;
          let _540___mcc_h6 = (_source10).cf30;
          let _541___mcc_h7 = (_source10).cf31;
          let _542_cf31 = _541___mcc_h7;
          let _543_cf30 = _540___mcc_h6;
          let _544_cf29 = _539___mcc_h5;
          let _545_cf28 = _538___mcc_h4;
          let _index73 = _module.__default.safeIndex(new BigNumber(406), new BigNumber((_463_v0).length));
          (_463_v0)[_index73] = (_498_v33).IsSubsetOf(_498_v33);
          let _index74 = _module.__default.safeIndex(new BigNumber(406), new BigNumber((_463_v0).length));
          (_463_v0)[_index74] = false;
          let _546_v63;
          _546_v63 = _dafny.MultiSet.fromElements(_542_cf31, new BigNumber((_490_v26).length));
          (globalState).f2 = ((_545_cf28).plus(new BigNumber((_546_v63).cardinality()))).multipliedBy(_542_cf31);
          let _547_v64;
          _547_v64 = _dafny.Set.fromElements(p0);
          let _548_v65;
          let _nw88 = Array((new BigNumber(2)).toNumber());
          _nw88[(_dafny.ZERO).toNumber()] = _542_cf31;
          _nw88[(_dafny.ONE).toNumber()] = _module.__default.safeModuloInt(_545_cf28, new BigNumber((_547_v64).length));
          _548_v65 = _nw88;
          let _index75 = _module.__default.safeIndex(new BigNumber(202), new BigNumber((_548_v65).length));
          (_548_v65)[_index75] = _545_cf28;
          let _index76 = _module.__default.safeIndex(new BigNumber(202), new BigNumber((_548_v65).length));
          (_548_v65)[_index76] = new BigNumber(-366);
          let _549_v66;
          _549_v66 = _module.D8.create_DC16(_542_cf31, p0, (_this).f11, _545_cf28);
          _544_cf29 = (_549_v66).dtor_cf29;
        } else if (_source10.is_DC15) {
          let _550___mcc_h8 = (_source10).cf27;
          let _551_cf27 = _550___mcc_h8;
          let _552_v67;
          let _nw89 = new _module.C1();
          _nw89.__ctor((_464_v1)[_module.__default.safeIndex((_this).f12, new BigNumber((_464_v1).length))], (_this).f3);
          _552_v67 = _nw89;
          _552_v67 = _552_v67;
          let _553_v68;
          let _554_v69;
          let _555_v70;
          let _out32;
          let _out33;
          let _out34;
          let _outcollector12 = (_this).m6(globalState);
          _out32 = _outcollector12[0];
          _out33 = _outcollector12[1];
          _out34 = _outcollector12[2];
          _553_v68 = _out32;
          _554_v69 = _out33;
          _555_v70 = _out34;
          r1 = _dafny.Seq.Concat(_490_v26, _dafny.Seq.Concat(_490_v26, _490_v26));
          r1 = _490_v26;
        } else {
          let _556___mcc_h9 = (_source10).cf32;
          let _557_cf32 = _556___mcc_h9;
          let _558_v71;
          _558_v71 = false;
          let _559_v72;
          let _nw90 = new _module.C0();
          _nw90.__ctor((_this).f12);
          _559_v72 = _nw90;
          let _560_v73;
          _560_v73 = _module.D4.create_DC8(_559_v72);
          let _rhs65 = p0;
          let _rhs66 = p0;
          let _rhs67 = _560_v73;
          _558_v71 = _rhs65;
          _558_v71 = _rhs66;
          _560_v73 = _rhs67;
          let _561_v74;
          _561_v74 = _dafny.MultiSet.fromElements(_558_v71);
          let _562_v75;
          let _nw91 = Array((new BigNumber(23)).toNumber()).fill(_dafny.ZERO);
          _562_v75 = _nw91;
          let _index77 = _module.__default.safeIndex(new BigNumber(685), new BigNumber((_562_v75).length));
          (_562_v75)[_index77] = (_this).f12;
          let _563_v76;
          _563_v76 = _dafny.Map.Empty.slice().updateUnsafe(_558_v71,_561_v74);
          let _564_v77;
          _564_v77 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("eeosvtnfb"),(((_563_v76).contains((_this).f11)) ? ((_563_v76).get((_this).f11)) : (_561_v74)));
          let _index78 = _module.__default.safeIndex(new BigNumber(685), new BigNumber((_562_v75).length));
          let _rhs68 = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Concat(_490_v26, _490_v26), _dafny.Seq.update(_dafny.Seq.Concat(_490_v26, _490_v26), _module.__default.safeIndex((_559_v72).f13, new BigNumber((_dafny.Seq.Concat(_490_v26, _490_v26)).length)), _495_v30))).length);
          let _rhs69 = _module.__default.safeModuloInt((_dafny.ZERO).minus((_559_v72).f13), (_559_v72).f13);
          let _rhs70 = (((_564_v77).contains(((_module.__default.fm23(_490_v26, (_559_v72).f13, (_559_v72).f13, _558_v71, globalState)) ? (_490_v26) : (_dafny.Seq.update(_490_v26, _module.__default.safeIndex((_module.__default.fm30((_559_v72).f13, (_this).f12, _dafny.Seq.UnicodeFromString("hkitllpy"), globalState)).dtor_cf7, new BigNumber((_490_v26).length)), new _dafny.CodePoint('f'.codePointAt(0))))))) ? ((_564_v77).get(((_module.__default.fm23(_490_v26, (_559_v72).f13, (_559_v72).f13, _558_v71, globalState)) ? (_490_v26) : (_dafny.Seq.update(_490_v26, _module.__default.safeIndex((_module.__default.fm30((_559_v72).f13, (_this).f12, _dafny.Seq.UnicodeFromString("hkitllpy"), globalState)).dtor_cf7, new BigNumber((_490_v26).length)), new _dafny.CodePoint('f'.codePointAt(0))))))) : (_561_v74));
          let _rhs71 = (_this).f12;
          let _lhs42 = globalState;
          let _lhs43 = globalState;
          let _lhs44 = _562_v75;
          let _lhs45 = _module.__default.safeIndex(new BigNumber(685), new BigNumber((_562_v75).length));
          _lhs42.f2 = _rhs68;
          _lhs43.f2 = _rhs69;
          _561_v74 = _rhs70;
          _lhs44[_lhs45] = _rhs71;
          let _565_v78;
          let _nw92 = new _module.C0();
          _nw92.__ctor((_464_v1)[_module.__default.safeIndex((_559_v72).f13, new BigNumber((_464_v1).length))]);
          _565_v78 = _nw92;
          let _566_v79;
          _566_v79 = _dafny.Seq.of(!(_558_v71));
          let _567_v80;
          _567_v80 = _566_v79;
          let _568_v81;
          _568_v81 = _dafny.Map.Empty.slice().updateUnsafe(_567_v80,p1);
          _568_v81 = (_568_v81).update(_567_v80, (_this).f11);
        }
        let _569_v82;
        _569_v82 = _dafny.Set.fromElements(false);
        _569_v82 = (_module.__default.fm31(p1, (_this).f12, (_this).f11, globalState)).Difference(_569_v82);
        _463_v0 = _463_v0;
      }
      r0 = (_this).f12;
      let _570_v83;
      _570_v83 = _dafny.Map.Empty.slice().updateUnsafe((_this).f12,p0);
      r1 = _dafny.Seq.Concat(_module.__default.fm25((((_570_v83).contains((_this).f12)) ? ((_570_v83).get((_this).f12)) : (p0)), new BigNumber((_module.__default.fm32(true, p1, (_this).f12, (_this).f11, globalState)).length), globalState), _490_v26);
      r2 = new BigNumber((_490_v26).length);
      return [r0, r1, r2];
    }
    get f11() {
      let _this = this;
      return _this._f11;
    };
    get f12() {
      let _this = this;
      return _this._f12;
    };
  };

  $module.C3 = class C3 {
    constructor () {
      this._tname = "_module.C3";
      this._f9 = _dafny.MultiSet.Empty;
      this._f10 = false;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f9, f10) {
      let _this = this;
      (_this)._f9 = f9;
      (_this)._f10 = f10;
      return;
    }
    fm19(p0, p1, globalState) {
      let _this = this;
      if (!(false)) {
        return new BigNumber((_dafny.Seq.of(new BigNumber(550), (_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(-925), new BigNumber(484), new BigNumber((_dafny.Set.fromElements(true)).length))).cardinality()), new BigNumber(-771))).length))))).length);
      } else {
        return (_dafny.ZERO).minus(new BigNumber(-312));
      }
    };
    fm20(p0, p1, p2, p3, globalState) {
      let _this = this;
      return new BigNumber(483);
    };
    m5(p0, p1, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = false;
      let r2 = new _dafny.CodePoint('D'.codePointAt(0));
      let r3 = false;
      let _571_v0;
      let _nw93 = Array((new BigNumber(5)).toNumber()).fill(_dafny.Map.Empty);
      _571_v0 = _nw93;
      let _572_v1;
      _572_v1 = _dafny.Map.Empty.slice().updateUnsafe(p1,p1);
      let _573_v2;
      let _nw94 = Array((new BigNumber(3)).toNumber());
      _nw94[(_dafny.ZERO).toNumber()] = p0;
      _nw94[(_dafny.ONE).toNumber()] = p0;
      _nw94[(new BigNumber(2)).toNumber()] = new BigNumber((_572_v1).length);
      _573_v2 = _nw94;
      let _574_v3;
      _574_v3 = _dafny.Map.Empty.slice().updateUnsafe(_573_v2,_573_v2);
      let _index79 = _module.__default.safeIndex(new BigNumber(803), new BigNumber((_571_v0).length));
      (_571_v0)[_index79] = (_574_v3).Merge(_574_v3);
      let _index80 = _module.__default.safeIndex(new BigNumber(803), new BigNumber((_571_v0).length));
      (_571_v0)[_index80] = _574_v3;
      let _575_v5;
      _575_v5 = _dafny.Seq.of((_this).f9);
      let _576_v6;
      _576_v6 = _module.D2.create_DC3(function () {
  let _coll20 = new _dafny.Map();
  for (const _compr_20 of (_575_v5).Elements) {
    let _577_v4 = _compr_20;
    if (_dafny.Seq.contains(_575_v5, _577_v4)) {
      _coll20.push([_577_v4,new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(p0, p0))).cardinality())]);
    }
  }
  return _coll20;
}());
      let _source11 = _576_v6;
      if (_source11.is_DC4) {
        let _578___mcc_h0 = (_source11).cf6;
        let _579_cf6 = _578___mcc_h0;
        let _580_v7;
        _580_v7 = _dafny.Map.Empty.slice().updateUnsafe((p1) && (p1),(p0).minus(p0));
        _580_v7 = (_580_v7).update(p1, _579_cf6);
        let _581_v8;
        _581_v8 = _dafny.Seq.UnicodeFromString("lfvoqiom");
        let _582_v9;
        _582_v9 = _dafny.Set.fromElements(p0, p0, p0);
        let _583_v10;
        _583_v10 = _dafny.Seq.of(_dafny.Set.fromElements((_dafny.ZERO).minus(new BigNumber((_581_v8).length)), p0, new BigNumber(30)), _582_v9);
        let _584_v11;
        _584_v11 = new _dafny.CodePoint('k'.codePointAt(0));
        let _585_v12;
        _585_v12 = _dafny.Map.Empty.slice().updateUnsafe(p0,_573_v2);
        let _586_v13;
        _586_v13 = _dafny.Seq.of(_579_cf6);
        let _587_v14;
        _587_v14 = _dafny.Seq.of((_this).f10, p1);
        let _588_v15;
        _588_v15 = _module.D2.create_DC5(p0, (_this).f10, p1, (_this).f10, (_this).f10);
        let _589_v17;
        let _nw95 = Array((new BigNumber(27)).toNumber());
        _nw95[(_dafny.ZERO).toNumber()] = _module.__default.fm21(new BigNumber((_583_v10).length), _584_v11, _module.__default.fm21((_dafny.ZERO).minus(new BigNumber(((_585_v12).update((_this).fm19(_584_v11, (_586_v13)[_module.__default.safeIndex(new BigNumber((_587_v14).length), new BigNumber((_586_v13).length))], globalState), _573_v2)).length)), _584_v11, p1, (_this).f10, globalState), p1, globalState);
        _nw95[(_dafny.ONE).toNumber()] = (_582_v9).IsSubsetOf(_582_v9);
        _nw95[(new BigNumber(2)).toNumber()] = ((_this).f9).IsProperSubsetOf(_dafny.MultiSet.FromArray(_586_v13));
        _nw95[(new BigNumber(3)).toNumber()] = (_this).f10;
        _nw95[(new BigNumber(4)).toNumber()] = p1;
        _nw95[(new BigNumber(5)).toNumber()] = _module.__default.fm21(new BigNumber((_581_v8).length), _584_v11, (_this).f10, false, globalState);
        _nw95[(new BigNumber(6)).toNumber()] = p1;
        _nw95[(new BigNumber(7)).toNumber()] = (_this).f10;
        _nw95[(new BigNumber(8)).toNumber()] = ((false) ? (p1) : ((_this).f10));
        _nw95[(new BigNumber(9)).toNumber()] = p1;
        _nw95[(new BigNumber(10)).toNumber()] = p1;
        _nw95[(new BigNumber(11)).toNumber()] = (_588_v15).dtor_cf11;
        _nw95[(new BigNumber(12)).toNumber()] = (((_this).f10) ? (true) : ((_this).f10));
        _nw95[(new BigNumber(13)).toNumber()] = !((_this).f10);
        _nw95[(new BigNumber(14)).toNumber()] = (_this).f10;
        _nw95[(new BigNumber(15)).toNumber()] = p1;
        _nw95[(new BigNumber(16)).toNumber()] = false;
        _nw95[(new BigNumber(17)).toNumber()] = (_this).f10;
        _nw95[(new BigNumber(18)).toNumber()] = ((_this).f10) === (!((_this).f10));
        _nw95[(new BigNumber(19)).toNumber()] = p1;
        _nw95[(new BigNumber(20)).toNumber()] = (p1) && ((_this).f10);
        _nw95[(new BigNumber(21)).toNumber()] = p1;
        _nw95[(new BigNumber(22)).toNumber()] = false;
        _nw95[(new BigNumber(23)).toNumber()] = (_dafny.Set.fromElements(p0, _579_cf6)).IsProperSubsetOf(function () {
          let _coll21 = new _dafny.Set();
          for (const _compr_21 of (_586_v13).Elements) {
            let _590_v16 = _compr_21;
            if (_dafny.Seq.contains(_586_v13, _590_v16)) {
              _coll21.add(_module.__default.safeModuloInt(_590_v16, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(994)), function (_591_i0) {
                return new BigNumber(121);
              })).length)));
            }
          }
          return _coll21;
        }());
        _nw95[(new BigNumber(24)).toNumber()] = p1;
        _nw95[(new BigNumber(25)).toNumber()] = _dafny.areEqual(_586_v13, _586_v13);
        _nw95[(new BigNumber(26)).toNumber()] = _module.__default.fm21(_579_cf6, _584_v11, (((_572_v1).contains(true)) ? ((_572_v1).get(true)) : ((_this).f10)), p1, globalState);
        _589_v17 = _nw95;
        let _index81 = _module.__default.safeIndex(new BigNumber(725), new BigNumber((_589_v17).length));
        (_589_v17)[_index81] = p1;
        let _index82 = _module.__default.safeIndex(new BigNumber(725), new BigNumber((_589_v17).length));
        (_589_v17)[_index82] = (_this).f10;
        _573_v2 = _573_v2;
        _589_v17 = _589_v17;
      } else if (_source11.is_DC5) {
        let _592___mcc_h1 = (_source11).cf7;
        let _593___mcc_h2 = (_source11).cf8;
        let _594___mcc_h3 = (_source11).cf9;
        let _595___mcc_h4 = (_source11).cf10;
        let _596___mcc_h5 = (_source11).cf11;
        let _597_cf11 = _596___mcc_h5;
        let _598_cf10 = _595___mcc_h4;
        let _599_cf9 = _594___mcc_h3;
        let _600_cf8 = _593___mcc_h2;
        let _601_cf7 = _592___mcc_h1;
        _597_cf11 = !(p0).isEqualTo((_601_cf7).minus(_601_cf7));
        let _602_v18;
        _602_v18 = _dafny.Seq.UnicodeFromString("vsovhibx");
        let _603_v19;
        _603_v19 = _dafny.Seq.of(new BigNumber(270), new BigNumber((_602_v18).length));
        let _604_v20;
        _604_v20 = new _dafny.CodePoint('q'.codePointAt(0));
        let _605_v21;
        _605_v21 = _dafny.Seq.of(_600_cf8);
        r1 = ((!(_dafny.Map.Empty.slice().updateUnsafe(_576_v6,p1)).contains(_module.__default.fm22(_603_v19, _601_cf7, _604_v20, globalState))) ? (_module.__default.fm21(p0, _604_v20, (_605_v21)[_module.__default.safeIndex(_601_cf7, new BigNumber((_605_v21).length))], _598_cf10, globalState)) : (_module.__default.fm21(_601_cf7, _604_v20, (_this).f10, _598_cf10, globalState)));
        let _606_v22;
        _606_v22 = _dafny.Map.Empty.slice().updateUnsafe(_601_cf7,_573_v2);
        _606_v22 = ((_606_v22).Merge((_606_v22).update(new BigNumber(963), _573_v2))).Merge(_dafny.Map.Empty.slice().updateUnsafe(p0,_573_v2));
        let _607_v23;
        _607_v23 = _dafny.Map.Empty.slice().updateUnsafe(_598_cf10,_dafny.Map.Empty.slice().updateUnsafe(_597_cf11,new BigNumber(-126)));
        let _608_v24;
        let _nw96 = Array((new BigNumber(4)).toNumber()).fill(_dafny.Map.Empty);
        _608_v24 = _nw96;
        let _609_v25;
        let _nw97 = new _module.C1();
        _nw97.__ctor((_dafny.ZERO).minus(new BigNumber((_607_v23).length)), _608_v24);
        _609_v25 = _nw97;
      } else {
        let _610___mcc_h6 = (_source11).cf5;
        let _611_cf5 = _610___mcc_h6;
        let _612_v26;
        _612_v26 = _dafny.Seq.UnicodeFromString("lnqpg");
        let _613_v27;
        _613_v27 = new _dafny.CodePoint('b'.codePointAt(0));
        let _614_v28;
        _614_v28 = _module.D4.create_DC9(new BigNumber(991), _613_v27, _dafny.Seq.Create(_module.__default.abs(new BigNumber(119)), function (_615_i1) {
  return new BigNumber(141);
}), p0, (_this).f10);
        let _616_v29;
        _616_v29 = _dafny.Seq.of(p0, new BigNumber(-848));
        let _617_v30;
        _617_v30 = _dafny.Seq.of(new BigNumber((_dafny.Seq.Concat(_612_v26, _module.__default.fm25((_614_v28).dtor_cf20, new BigNumber((_616_v29).length), globalState))).length));
        (globalState).f2 = (_616_v29)[_module.__default.safeIndex(p0, new BigNumber((_616_v29).length))];
        (globalState).f2 = p0;
        let _618_v31;
        let _nw98 = Array((new BigNumber(27)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
        _618_v31 = _nw98;
        let _index83 = _module.__default.safeIndex(new BigNumber(842), new BigNumber((_618_v31).length));
        (_618_v31)[_index83] = _module.__default.fm25((_this).f10, new BigNumber(444), globalState);
        let _index84 = _module.__default.safeIndex(new BigNumber(842), new BigNumber((_618_v31).length));
        (_618_v31)[_index84] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(641)), ((_619_v27) => function (_620_i2) {
          return _619_v27;
        })(_613_v27));
        r3 = p1;
      }
      let _621_v32;
      let _nw99 = Array((new BigNumber(24)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
      _621_v32 = _nw99;
      let _622_v33;
      _622_v33 = _dafny.Map.Empty.slice().updateUnsafe(_621_v32,p1);
      r0 = (((_622_v33).contains(_621_v32)) ? ((_622_v33).get(_621_v32)) : ((_this).f10));
      let _623_v34;
      let _init11 = function (_624_i3) {
        return _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('r'.codePointAt(0)),_dafny.Seq.UnicodeFromString("rvkrqon"));
      };
      let _nw100 = Array((new BigNumber(6)).toNumber());
      for (let _i0_11 = 0; _i0_11 < new BigNumber(_nw100.length); _i0_11++) {
        _nw100[_i0_11] = _init11(new BigNumber(_i0_11));
      }
      _623_v34 = _nw100;
      let _625_v35;
      _625_v35 = new _dafny.CodePoint('t'.codePointAt(0));
      let _626_v36;
      _626_v36 = _dafny.Seq.UnicodeFromString("aewdkgg");
      let _627_v37;
      _627_v37 = _dafny.Map.Empty.slice().updateUnsafe(_625_v35,_module.__default.fm25(false, new BigNumber((_626_v36).length), globalState));
      let _index85 = _module.__default.safeIndex(new BigNumber(975), new BigNumber((_623_v34).length));
      (_623_v34)[_index85] = _627_v37;
      let _628_v38;
      _628_v38 = _dafny.Map.Empty.slice().updateUnsafe((_this).f9,(_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of((_this).f10, p1, !((_this).f10), (_this).f10, p1)).length)));
      let _pat_let_tv15 = _627_v37;
      let _pat_let_tv16 = _625_v35;
      let _pat_let_tv17 = _626_v36;
      let _pat_let_tv18 = _627_v37;
      let _pat_let_tv19 = _627_v37;
      let _index86 = _module.__default.safeIndex(new BigNumber(975), new BigNumber((_623_v34).length));
      (_623_v34)[_index86] = function (_source12) {
        if (_source12.is_DC4) {
          let _629___mcc_h7 = (_source12).cf6;
          let _630_cf6 = _629___mcc_h7;
          return _pat_let_tv15;
        } else if (_source12.is_DC5) {
          let _631___mcc_h8 = (_source12).cf7;
          let _632___mcc_h9 = (_source12).cf8;
          let _633___mcc_h10 = (_source12).cf9;
          let _634___mcc_h11 = (_source12).cf10;
          let _635___mcc_h12 = (_source12).cf11;
          let _636_cf11 = _635___mcc_h12;
          let _637_cf10 = _634___mcc_h11;
          let _638_cf9 = _633___mcc_h10;
          let _639_cf8 = _632___mcc_h9;
          let _640_cf7 = _631___mcc_h8;
          return (_dafny.Map.Empty.slice().updateUnsafe(_pat_let_tv16,_pat_let_tv17)).Merge(_pat_let_tv18);
        } else {
          let _641___mcc_h13 = (_source12).cf5;
          let _642_cf5 = _641___mcc_h13;
          return _pat_let_tv19;
        }
      }(_module.D2.create_DC3(_628_v38));
      let _643_v40;
      let _init12 = ((_644_p0) => function (_645_i4) {
        return function () {
          let _coll22 = new _dafny.Map();
          for (const _compr_22 of (_dafny.Set.fromElements(_644_p0)).Elements) {
            let _646_v39 = _compr_22;
            if ((_dafny.Set.fromElements(_644_p0)).contains(_646_v39)) {
              _coll22.push([(_646_v39).minus(_644_p0),(_this).f10]);
            }
          }
          return _coll22;
        }();
      })(p0);
      let _nw101 = Array((new BigNumber(15)).toNumber());
      for (let _i0_12 = 0; _i0_12 < new BigNumber(_nw101.length); _i0_12++) {
        _nw101[_i0_12] = _init12(new BigNumber(_i0_12));
      }
      _643_v40 = _nw101;
      let _647_v41;
      let _nw102 = new _module.C2();
      _nw102.__ctor((_this).f10, p0, _643_v40);
      _647_v41 = _nw102;
      let _648_v42;
      let _nw103 = Array((new BigNumber(10)).toNumber()).fill([]);
      _648_v42 = _nw103;
      _648_v42 = _648_v42;
      r0 = p1;
      r1 = (_647_v41).f11;
      let _649_v43;
      let _init13 = ((_650_p0, _651_v41) => function (_652_i5) {
        return _dafny.Seq.of(_650_p0, new BigNumber((_dafny.Seq.of((_651_v41).f11, (_this).f10)).length), (_651_v41).f12);
      })(p0, _647_v41);
      let _nw104 = Array((new BigNumber(22)).toNumber());
      for (let _i0_13 = 0; _i0_13 < new BigNumber(_nw104.length); _i0_13++) {
        _nw104[_i0_13] = _init13(new BigNumber(_i0_13));
      }
      _649_v43 = _nw104;
      let _653_v44;
      _653_v44 = _dafny.Map.Empty.slice().updateUnsafe(_649_v43,(_647_v41).f12);
      let _654_v45;
      _654_v45 = _dafny.Map.Empty.slice().updateUnsafe(p1,(_647_v41).f12);
      let _655_v46;
      _655_v46 = _dafny.Seq.of(p1, (_647_v41).f11, (_647_v41).f11, (_this).f10, (_647_v41).f11);
      let _656_v47;
      _656_v47 = _dafny.Seq.of(new BigNumber((_655_v46).length), p0);
      let _657_v48;
      _657_v48 = _module.D4.create_DC9((((_653_v44).contains(_649_v43)) ? ((_653_v44).get(_649_v43)) : ((((_654_v45).contains((_647_v41).f11)) ? ((_654_v45).get((_647_v41).f11)) : (new BigNumber(273))))), _625_v35, _656_v47, p0, (_647_v41).f11);
      r2 = (_657_v48).dtor_cf17;
      r3 = ((_this).fm19(_625_v35, (_647_v41).f12, globalState)).isLessThanOrEqualTo(new BigNumber(636));
      return [r0, r1, r2, r3];
    }
    get f9() {
      let _this = this;
      return _this._f9;
    };
    get f10() {
      let _this = this;
      return _this._f10;
    };
  };

  $module.C4 = class C4 {
    constructor () {
      this._tname = "_module.C4";
      this._f3 = [];
      this._f7 = false;
      this._f8 = _dafny.MultiSet.Empty;
    }
    _parentTraits() {
      return [_module.T1, _module.T0];
    }
    get f3() {
      let _this = this;
      return _this._f3;
    };
    __ctor(f7, f8, f3) {
      let _this = this;
      (_this)._f7 = f7;
      (_this)._f8 = f8;
      (_this)._f3 = f3;
      return;
    }
    fm13(p0, p1, p2, p3, globalState) {
      let _this = this;
      return (((true) ? (_dafny.Set.fromElements(new BigNumber(196), new BigNumber((_dafny.Seq.UnicodeFromString("yapel")).length), new BigNumber((_dafny.Seq.of(_dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("ruu")).length)))).length))) : (_dafny.Set.fromElements(new BigNumber(532), new BigNumber(424))))).Intersect(_dafny.Set.fromElements(new BigNumber(566)));
    };
    fm14(p0, p1, globalState) {
      let _this = this;
      return _dafny.Seq.UnicodeFromString("aljqjwu");
    };
    m1(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = new _dafny.CodePoint('D'.codePointAt(0));
      let _658_v0;
      _658_v0 = new BigNumber(671);
      let _659_v1;
      _659_v1 = _module.D1.create_DC2(_658_v0, p1, _658_v0);
      let _660_v2;
      let _nw105 = Array((new BigNumber(10)).toNumber());
      _nw105[(_dafny.ZERO).toNumber()] = (_this).f7;
      _nw105[(_dafny.ONE).toNumber()] = (_658_v0).isLessThanOrEqualTo(_658_v0);
      _nw105[(new BigNumber(2)).toNumber()] = (_659_v1).dtor_cf3;
      _nw105[(new BigNumber(3)).toNumber()] = !(p1);
      _nw105[(new BigNumber(4)).toNumber()] = (_658_v0).isLessThan(new BigNumber((p0).length));
      _nw105[(new BigNumber(5)).toNumber()] = p1;
      _nw105[(new BigNumber(6)).toNumber()] = (_659_v1).dtor_cf3;
      _nw105[(new BigNumber(7)).toNumber()] = (_658_v0).isLessThanOrEqualTo(_658_v0);
      _nw105[(new BigNumber(8)).toNumber()] = (_this).f7;
      _nw105[(new BigNumber(9)).toNumber()] = (_this).f7;
      _660_v2 = _nw105;
      let _661_v3;
      let _nw106 = Array((new BigNumber(22)).toNumber()).fill(_dafny.ZERO);
      _661_v3 = _nw106;
      let _index87 = _module.__default.safeIndex(new BigNumber(414), new BigNumber((_661_v3).length));
      (_661_v3)[_index87] = _658_v0;
      let _662_v6;
      _662_v6 = _dafny.Set.fromElements(_658_v0, new BigNumber((p2).length));
      let _663_v7;
      _663_v7 = _dafny.Map.Empty.slice().updateUnsafe(function () {
        let _coll23 = new _dafny.Map();
        for (const _compr_23 of (_662_v6).Elements) {
          let _664_v5 = _compr_23;
          if ((_662_v6).contains(_664_v5)) {
            _coll23.push([(_664_v5).multipliedBy(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('t'.codePointAt(0)),(_this).f7)).length)),true]);
          }
        }
        return _coll23;
      }(),(_this).f7);
      let _665_v8;
      _665_v8 = _dafny.Map.Empty.slice().updateUnsafe(_658_v0,p1);
      let _index88 = _module.__default.safeIndex(new BigNumber(414), new BigNumber((_661_v3).length));
      let _rhs72 = _660_v2;
      let _rhs73 = _658_v0;
      let _rhs74 = new BigNumber(((function () {
        let _coll24 = new _dafny.Map();
        for (const _compr_24 of (_663_v7).Keys.Elements) {
          let _666_v4 = _compr_24;
          if ((_663_v7).contains(_666_v4)) {
            _coll24.push([_666_v4,_658_v0]);
          }
        }
        return _coll24;
      }()).update(_665_v8, (_658_v0).multipliedBy(_658_v0))).length);
      let _lhs46 = _661_v3;
      let _lhs47 = _module.__default.safeIndex(new BigNumber(414), new BigNumber((_661_v3).length));
      let _lhs48 = globalState;
      _660_v2 = _rhs72;
      _lhs46[_lhs47] = _rhs73;
      _lhs48.f2 = _rhs74;
      let _667_i0;
      _667_i0 = _dafny.ZERO;
      L2: {
        while (p1) {
          C2: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_667_i0)) {
              break L2;
            }
            _667_i0 = (_667_i0).plus(_dafny.ONE);
            let _668_v9;
            _668_v9 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm15(!(p1), true, globalState),_658_v0);
            let _669_v10;
            _669_v10 = _module.D2.create_DC3(_668_v9);
            let _670_v11;
            _670_v11 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_658_v0),new BigNumber(226));
            let _671_v12;
            _671_v12 = _dafny.Seq.of(_660_v2);
            let _672_v13;
            let _nw107 = Array((new BigNumber(8)).toNumber());
            _nw107[(_dafny.ZERO).toNumber()] = _module.D2.create_DC3(_668_v9);
            _nw107[(_dafny.ONE).toNumber()] = _669_v10;
            _nw107[(new BigNumber(2)).toNumber()] = _module.D2.create_DC3((_dafny.Map.Empty.slice().updateUnsafe((_this).f8,new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(538)), ((_673_v3, _674_p2, _675_v0) => function (_676_i1) {
  return new BigNumber(((_dafny.MultiSet.fromElements((_673_v3)[_module.__default.safeIndex(new BigNumber(414), new BigNumber((_673_v3).length))], new BigNumber(-947))).update(new BigNumber((_674_p2).length), _module.__default.abs(_675_v0))).cardinality());
})(_661_v3, p2, _658_v0))).length))).update((_this).f8, (((_670_v11).contains(new BigNumber((_671_v12).length))) ? ((_670_v11).get(new BigNumber((_671_v12).length))) : (_658_v0))));
            _nw107[(new BigNumber(3)).toNumber()] = _669_v10;
            _nw107[(new BigNumber(4)).toNumber()] = _669_v10;
            _nw107[(new BigNumber(5)).toNumber()] = _669_v10;
            _nw107[(new BigNumber(6)).toNumber()] = _669_v10;
            _nw107[(new BigNumber(7)).toNumber()] = _669_v10;
            _672_v13 = _nw107;
            let _index89 = _module.__default.safeIndex(new BigNumber(780), new BigNumber((_672_v13).length));
            (_672_v13)[_index89] = _669_v10;
            let _index90 = _module.__default.safeIndex(new BigNumber(780), new BigNumber((_672_v13).length));
            (_672_v13)[_index90] = _669_v10;
            let _677_v14;
            _677_v14 = _dafny.Seq.UnicodeFromString("c");
            _677_v14 = _dafny.Seq.Concat(_677_v14, _dafny.Seq.Concat(p2, p2));
            if ((_this).f7) {
              let _678_v15;
              _678_v15 = false;
              _678_v15 = ((_662_v6).Intersect(function () {
                let _coll25 = new _dafny.Set();
                for (const _compr_25 of (_662_v6).Elements) {
                  let _679_v16 = _compr_25;
                  if ((_662_v6).contains(_679_v16)) {
                    _coll25.add((_679_v16).minus(new BigNumber(713)));
                  }
                }
                return _coll25;
              }())).IsDisjointFrom(_662_v6);
              _678_v15 = p1;
              let _680_v17;
              _680_v17 = _dafny.Seq.of(p1, p1);
              let _681_v18;
              _681_v18 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_658_v0),_680_v17);
              _681_v18 = (_681_v18).update((_661_v3)[_module.__default.safeIndex(new BigNumber(414), new BigNumber((_661_v3).length))], _dafny.Seq.Concat(_dafny.Seq.of((_this).f7, _678_v15), _680_v17));
              _678_v15 = _678_v15;
              let _index91 = _module.__default.safeIndex(new BigNumber(414), new BigNumber((_661_v3).length));
              (_661_v3)[_index91] = _658_v0;
            } else {
              let _682_v19;
              _682_v19 = _module.D3.create_DC6(new _dafny.CodePoint('l'.codePointAt(0)));
              _658_v0 = (p0)[_module.__default.safeIndex(new BigNumber((_module.__default.fm16((_661_v3)[_module.__default.safeIndex(new BigNumber(414), new BigNumber((_661_v3).length))], (_682_v19).dtor_cf12, (_661_v3)[_module.__default.safeIndex(new BigNumber(414), new BigNumber((_661_v3).length))], globalState)).length), new BigNumber((p0).length))];
              let _683_v20;
              _683_v20 = _dafny.Seq.of(p2, _dafny.Seq.of(new _dafny.CodePoint('r'.codePointAt(0))));
              (globalState).f2 = (_module.__default.fm17(_dafny.Seq.of(_677_v14), (_661_v3)[_module.__default.safeIndex(new BigNumber(414), new BigNumber((_661_v3).length))], new BigNumber((p0).length), (_661_v3)[_module.__default.safeIndex(new BigNumber(414), new BigNumber((_661_v3).length))], globalState)).plus(_module.__default.fm17(_683_v20, _658_v0, (_661_v3)[_module.__default.safeIndex(new BigNumber(414), new BigNumber((_661_v3).length))], _658_v0, globalState));
              let _684_v21;
              let _nw108 = Array((new BigNumber(22)).toNumber()).fill(_module.D2.Default());
              _684_v21 = _nw108;
              _684_v21 = _684_v21;
              let _685_v22;
              _685_v22 = _dafny.Map.Empty.slice().updateUnsafe(false,true);
              let _686_v23;
              _686_v23 = _dafny.Seq.of((_661_v3)[_module.__default.safeIndex(new BigNumber(414), new BigNumber((_661_v3).length))], (_661_v3)[_module.__default.safeIndex(new BigNumber(414), new BigNumber((_661_v3).length))], (new BigNumber((_685_v22).length)).multipliedBy((_661_v3)[_module.__default.safeIndex(new BigNumber(414), new BigNumber((_661_v3).length))]), (_dafny.ZERO).minus(new BigNumber(-212)), _module.__default.safeModuloInt(new BigNumber(571), _658_v0));
              let _687_v24;
              _687_v24 = _dafny.Set.fromElements(_660_v2);
              let _rhs75 = (_dafny.ZERO).minus(_module.__default.safeModuloInt((_dafny.ZERO).minus(new BigNumber((_677_v14).length)), _658_v0));
              let _rhs76 = _658_v0;
              let _rhs77 = _module.__default.fm18(new BigNumber((_687_v24).length), new BigNumber((p0).length), globalState);
              let _rhs78 = (_661_v3)[_module.__default.safeIndex(new BigNumber(414), new BigNumber((_661_v3).length))];
              let _rhs79 = p0;
              let _lhs49 = globalState;
              let _lhs50 = globalState;
              _658_v0 = _rhs75;
              _lhs49.f2 = _rhs76;
              _659_v1 = _rhs77;
              _lhs50.f2 = _rhs78;
              _686_v23 = _rhs79;
              let _688_v25;
              let _nw109 = Array((new BigNumber(16)).toNumber()).fill([]);
              _688_v25 = _nw109;
              let _index92 = _module.__default.safeIndex(new BigNumber(306), new BigNumber((_688_v25).length));
              (_688_v25)[_index92] = _661_v3;
              let _index93 = _module.__default.safeIndex(new BigNumber(306), new BigNumber((_688_v25).length));
              (_688_v25)[_index93] = _661_v3;
            }
            let _689_v26;
            _689_v26 = _dafny.Set.fromElements(true);
            let _690_v27;
            _690_v27 = _dafny.Map.Empty.slice().updateUnsafe(_689_v26,new BigNumber((_677_v14).length));
            _690_v27 = (_690_v27).update(_689_v26, new BigNumber(680));
          }
        }
      }
      let _691_i2;
      _691_i2 = _dafny.ZERO;
      L3: {
        while (p1) {
          C3: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_691_i2)) {
              break L3;
            }
            _691_i2 = (_691_i2).plus(_dafny.ONE);
            (globalState).f2 = (_dafny.ZERO).minus(new BigNumber(((((p1) ? (_665_v8) : ((_665_v8).update((_661_v3)[_module.__default.safeIndex(new BigNumber(414), new BigNumber((_661_v3).length))], (_this).f7)))).Merge(_665_v8)).length));
            let _692_v28;
            let _nw110 = new _module.C0();
            _nw110.__ctor(((p1) ? (_module.__default.fm17(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-339)), ((_693_p2) => function (_694_i3) {
              return _693_p2;
            })(p2)), (_661_v3)[_module.__default.safeIndex(new BigNumber(414), new BigNumber((_661_v3).length))], (_661_v3)[_module.__default.safeIndex(new BigNumber(414), new BigNumber((_661_v3).length))], _658_v0, globalState)) : ((_661_v3)[_module.__default.safeIndex(new BigNumber(414), new BigNumber((_661_v3).length))])));
            _692_v28 = _nw110;
            let _695_v29;
            let _nw111 = new _module.C3();
            _nw111.__ctor((_this).f8, (_this).f7);
            _695_v29 = _nw111;
            (globalState).f2 = _658_v0;
          }
        }
      }
      let _696_v30;
      _696_v30 = false;
      _696_v30 = p1;
      let _697_i4;
      _697_i4 = _dafny.ZERO;
      L4: {
        while (!_dafny.areEqual(p0, p0)) {
          C4: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_697_i4)) {
              break L4;
            }
            _697_i4 = (_697_i4).plus(_dafny.ONE);
            let _698_v31;
            _698_v31 = new _dafny.CodePoint('x'.codePointAt(0));
            let _rhs80 = ((_696_v30) ? (_698_v31) : ((p2)[_module.__default.safeIndex(_658_v0, new BigNumber((p2).length))]));
            let _rhs81 = (_658_v0).isLessThanOrEqualTo(_module.__default.safeModuloInt((_dafny.ZERO).minus(_658_v0), _658_v0));
            let _rhs82 = (((_661_v3)[_module.__default.safeIndex(new BigNumber(414), new BigNumber((_661_v3).length))]).multipliedBy(_658_v0)).isEqualTo(((p1) ? (_658_v0) : (_658_v0)));
            r1 = _rhs80;
            _696_v30 = _rhs81;
            _696_v30 = _rhs82;
            let _699_v32;
            let _init14 = ((_700_v30, _701_p1) => function (_702_i5) {
              return (_dafny.Set.fromElements(_700_v30, !((_this).f7), _701_p1)).Difference(_dafny.Set.fromElements(_701_p1));
            })(_696_v30, p1);
            let _nw112 = Array((new BigNumber(16)).toNumber());
            for (let _i0_14 = 0; _i0_14 < new BigNumber(_nw112.length); _i0_14++) {
              _nw112[_i0_14] = _init14(new BigNumber(_i0_14));
            }
            _699_v32 = _nw112;
            let _index94 = _module.__default.safeIndex(new BigNumber(670), new BigNumber((_699_v32).length));
            (_699_v32)[_index94] = _dafny.Set.fromElements(true, p1, (_this).f7, !(_696_v30));
            let _703_v33;
            _703_v33 = _dafny.Set.fromElements((_this).f7);
            let _index95 = _module.__default.safeIndex(new BigNumber(670), new BigNumber((_699_v32).length));
            (_699_v32)[_index95] = _703_v33;
            _696_v30 = (_658_v0).isLessThan((_661_v3)[_module.__default.safeIndex(new BigNumber(414), new BigNumber((_661_v3).length))]);
            let _704_v34;
            _704_v34 = _module.D2.create_DC4(_module.__default.safeModuloInt(_658_v0, _658_v0));
            let _source13 = _704_v34;
            if (_source13.is_DC4) {
              let _705___mcc_h0 = (_source13).cf6;
              let _706_cf6 = _705___mcc_h0;
              _698_v31 = _698_v31;
              _661_v3 = _661_v3;
              let _707_v35;
              _707_v35 = _dafny.Map.Empty.slice().updateUnsafe((_661_v3)[_module.__default.safeIndex(new BigNumber(414), new BigNumber((_661_v3).length))],new BigNumber(169));
              let _708_v36;
              _708_v36 = _dafny.Map.Empty.slice().updateUnsafe(((((_707_v35).contains(_658_v0)) ? ((_707_v35).get(_658_v0)) : (_module.__default.fm24(_662_v6, globalState)))).plus((_661_v3)[_module.__default.safeIndex(new BigNumber(414), new BigNumber((_661_v3).length))]),(_706_cf6).multipliedBy(_706_cf6));
              let _709_v37;
              _709_v37 = _dafny.MultiSet.fromElements((_this).fm13(_658_v0, new BigNumber(495), p1, (_dafny.ZERO).minus(_706_cf6), globalState));
              _707_v35 = (_707_v35).update(_706_cf6, (((_709_v37).contains(_dafny.Set.fromElements(_706_cf6))) ? ((_709_v37).get(_dafny.Set.fromElements(_706_cf6))) : (new BigNumber((p2).length))));
              _696_v30 = _module.__default.fm23(((!(p1)) ? (p2) : (p2)), _658_v0, _706_cf6, _696_v30, globalState);
            } else if (_source13.is_DC5) {
              let _710___mcc_h1 = (_source13).cf7;
              let _711___mcc_h2 = (_source13).cf8;
              let _712___mcc_h3 = (_source13).cf9;
              let _713___mcc_h4 = (_source13).cf10;
              let _714___mcc_h5 = (_source13).cf11;
              let _715_cf11 = _714___mcc_h5;
              let _716_cf10 = _713___mcc_h4;
              let _717_cf9 = _712___mcc_h3;
              let _718_cf8 = _711___mcc_h2;
              let _719_cf7 = _710___mcc_h1;
              let _720_v38;
              let _init15 = ((_721_v31) => function (_722_i6) {
                return _721_v31;
              })(_698_v31);
              let _nw113 = Array((new BigNumber(21)).toNumber());
              for (let _i0_15 = 0; _i0_15 < new BigNumber(_nw113.length); _i0_15++) {
                _nw113[_i0_15] = _init15(new BigNumber(_i0_15));
              }
              _720_v38 = _nw113;
              let _723_v39;
              _723_v39 = _dafny.Seq.of(_720_v38, _720_v38);
              _658_v0 = ((_661_v3)[_module.__default.safeIndex(new BigNumber(414), new BigNumber((_661_v3).length))]).minus(new BigNumber((_723_v39).length));
              let _724_v40;
              let _nw114 = new _module.C0();
              _nw114.__ctor(_719_cf7);
              _724_v40 = _nw114;
              let _725_v41;
              _725_v41 = _dafny.Seq.of(_718_cf8);
              let _726_v42;
              _726_v42 = _dafny.MultiSet.fromElements(_dafny.Seq.Concat(_725_v41, _725_v41), _dafny.Seq.update(_725_v41, _module.__default.safeIndex(new BigNumber(720), new BigNumber((_725_v41).length)), !(true)), _dafny.Seq.of(_718_cf8), _725_v41);
              let _727_v43;
              _727_v43 = _dafny.Map.Empty.slice().updateUnsafe(_716_cf10,_module.__default.fm33(_716_cf10, _dafny.Seq.of((_dafny.ZERO).minus((_724_v40).f13)), _717_cf9, _698_v31, globalState));
              _726_v42 = (((_727_v43).contains(!(_696_v30))) ? ((_727_v43).get(!(_696_v30))) : (_726_v42));
              let _728_v44;
              let _init16 = ((_729_cf11) => function (_730_i7) {
                return _dafny.Seq.of(_729_cf11, _729_cf11);
              })(_715_cf11);
              let _nw115 = Array((new BigNumber(2)).toNumber());
              for (let _i0_16 = 0; _i0_16 < new BigNumber(_nw115.length); _i0_16++) {
                _nw115[_i0_16] = _init16(new BigNumber(_i0_16));
              }
              _728_v44 = _nw115;
              let _index96 = _module.__default.safeIndex(new BigNumber(426), new BigNumber((_728_v44).length));
              (_728_v44)[_index96] = _725_v41;
              let _index97 = _module.__default.safeIndex(new BigNumber(426), new BigNumber((_728_v44).length));
              (_728_v44)[_index97] = _725_v41;
            } else {
              let _731___mcc_h6 = (_source13).cf5;
              let _732_cf5 = _731___mcc_h6;
              let _733_v45;
              let _nw116 = new _module.C2();
              _nw116.__ctor((_this).f7, new BigNumber(296), ((p1) ? ((_this).f3) : ((_this).f3)));
              _733_v45 = _nw116;
              let _734_v46;
              _734_v46 = _dafny.MultiSet.fromElements(p1, !((_733_v45).f11), true);
              let _735_v47;
              _735_v47 = _dafny.Map.Empty.slice().updateUnsafe(false,_661_v3);
              let _736_v48;
              _736_v48 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((p2).length),_661_v3);
              let _737_v49;
              let _nw117 = Array((new BigNumber(18)).toNumber());
              _nw117[(_dafny.ZERO).toNumber()] = _661_v3;
              _nw117[(_dafny.ONE).toNumber()] = _661_v3;
              _nw117[(new BigNumber(2)).toNumber()] = (((_735_v47).contains((_733_v45).f11)) ? ((_735_v47).get((_733_v45).f11)) : (_661_v3));
              _nw117[(new BigNumber(3)).toNumber()] = _661_v3;
              _nw117[(new BigNumber(4)).toNumber()] = _661_v3;
              _nw117[(new BigNumber(5)).toNumber()] = _661_v3;
              _nw117[(new BigNumber(6)).toNumber()] = _661_v3;
              _nw117[(new BigNumber(7)).toNumber()] = _661_v3;
              _nw117[(new BigNumber(8)).toNumber()] = _661_v3;
              _nw117[(new BigNumber(9)).toNumber()] = _661_v3;
              _nw117[(new BigNumber(10)).toNumber()] = _661_v3;
              _nw117[(new BigNumber(11)).toNumber()] = _661_v3;
              _nw117[(new BigNumber(12)).toNumber()] = _661_v3;
              _nw117[(new BigNumber(13)).toNumber()] = _661_v3;
              _nw117[(new BigNumber(14)).toNumber()] = _661_v3;
              _nw117[(new BigNumber(15)).toNumber()] = _661_v3;
              _nw117[(new BigNumber(16)).toNumber()] = _661_v3;
              _nw117[(new BigNumber(17)).toNumber()] = (((_736_v48).contains((_dafny.ZERO).minus((_733_v45).f12))) ? ((_736_v48).get((_dafny.ZERO).minus((_733_v45).f12))) : (_661_v3));
              _737_v49 = _nw117;
              let _index98 = _module.__default.safeIndex(new BigNumber(914), new BigNumber((_737_v49).length));
              (_737_v49)[_index98] = _661_v3;
              let _738_v50;
              _738_v50 = _dafny.Map.Empty.slice().updateUnsafe((_this).f7,(_dafny.MultiSet.fromElements(p1)).Intersect(_734_v46));
              let _index99 = _module.__default.safeIndex(new BigNumber(914), new BigNumber((_737_v49).length));
              let _rhs83 = (((_738_v50).contains(p1)) ? ((_738_v50).get(p1)) : (_734_v46));
              let _rhs84 = p1;
              let _rhs85 = (((_736_v48).contains(((_661_v3)[_module.__default.safeIndex(new BigNumber(414), new BigNumber((_661_v3).length))]).multipliedBy((_733_v45).f12))) ? ((_736_v48).get(((_661_v3)[_module.__default.safeIndex(new BigNumber(414), new BigNumber((_661_v3).length))]).multipliedBy((_733_v45).f12))) : (_661_v3));
              let _lhs51 = _737_v49;
              let _lhs52 = _module.__default.safeIndex(new BigNumber(914), new BigNumber((_737_v49).length));
              _734_v46 = _rhs83;
              _696_v30 = _rhs84;
              _lhs51[_lhs52] = _rhs85;
              let _739_v51;
              _739_v51 = _dafny.Map.Empty.slice().updateUnsafe(_658_v0,(_661_v3)[_module.__default.safeIndex(new BigNumber(414), new BigNumber((_661_v3).length))]);
              let _740_v52;
              _740_v52 = _dafny.Map.Empty.slice().updateUnsafe(_739_v51,((_this).f8).Intersect(_dafny.MultiSet.fromElements((_661_v3)[_module.__default.safeIndex(new BigNumber(414), new BigNumber((_661_v3).length))])));
              _740_v52 = (_740_v52).update(((_739_v51).update((_733_v45).f12, (_661_v3)[_module.__default.safeIndex(new BigNumber(414), new BigNumber((_661_v3).length))])).update(new BigNumber(((_this).f8).cardinality()), (_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.fromElements((_this).f7)).cardinality()))), (_this).f8);
              let _rhs86 = true;
              let _rhs87 = (_dafny.ZERO).minus((_module.__default.safeDivisionInt((_733_v45).f12, (_661_v3)[_module.__default.safeIndex(new BigNumber(414), new BigNumber((_661_v3).length))])).multipliedBy(_658_v0));
              let _rhs88 = (_733_v45).f12;
              let _rhs89 = (_this).f7;
              let _rhs90 = ((_733_v45).f12).isLessThan(_module.__default.safeModuloInt(_658_v0, _658_v0));
              let _lhs53 = globalState;
              _696_v30 = _rhs86;
              r0 = _rhs87;
              _lhs53.f2 = _rhs88;
              _696_v30 = _rhs89;
              _696_v30 = _rhs90;
            }
          }
        }
      }
      _696_v30 = ((_module.D2.create_DC5(_658_v0, _696_v30, (_this).f7, !(_696_v30), _696_v30)).dtor_cf7).isEqualTo(new BigNumber((p2).length));
      r0 = _658_v0;
      let _741_v53;
      _741_v53 = new _dafny.CodePoint('u'.codePointAt(0));
      r1 = _741_v53;
      return [r0, r1];
    }
    get f7() {
      let _this = this;
      return _this._f7;
    };
    get f8() {
      let _this = this;
      return _this._f8;
    };
  };

  $module.C5 = class C5 {
    constructor () {
      this._tname = "_module.C5";
      this._f6 = new _dafny.CodePoint('D'.codePointAt(0));
    }
    _parentTraits() {
      return [];
    }
    __ctor(f6) {
      let _this = this;
      (_this)._f6 = f6;
      return;
    }
    fm7(p0, p1, p2, globalState) {
      let _this = this;
      return (((false) ? (new BigNumber(976)) : (new BigNumber(852)))).isLessThanOrEqualTo(new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(!(false),new BigNumber((_dafny.MultiSet.fromElements((_module.D1.create_DC2(new BigNumber(840), true, new BigNumber(-480))).dtor_cf3, true, false)).cardinality()))).Merge(_dafny.Map.Empty.slice().updateUnsafe(!(false),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.Seq.of(false, false))).length)))).length));
    };
    m4(p0, p1, globalState) {
      let _this = this;
      let _742_v0;
      _742_v0 = false;
      if (!(_742_v0)) {
        let _743_v1;
        _743_v1 = _dafny.Seq.UnicodeFromString("otnenv");
        _742_v0 = !_dafny.areEqual(_743_v1, _dafny.Seq.UnicodeFromString("sypxlsr"));
        let _744_v2;
        let _nw118 = Array((new BigNumber(11)).toNumber()).fill(_dafny.Map.Empty);
        _744_v2 = _nw118;
        let _index100 = _module.__default.safeIndex(new BigNumber(439), new BigNumber((_744_v2).length));
        (_744_v2)[_index100] = _dafny.Map.Empty.slice().updateUnsafe(p1,p1);
        let _745_v3;
        _745_v3 = _dafny.Map.Empty.slice().updateUnsafe(p0,_742_v0);
        let _746_v4;
        _746_v4 = _dafny.MultiSet.fromElements(p0, p0);
        let _747_v5;
        _747_v5 = _dafny.Map.Empty.slice().updateUnsafe(_746_v4,p1);
        let _748_v6;
        _748_v6 = _module.D2.create_DC3(_747_v5);
        let _749_v7;
        _749_v7 = _dafny.Seq.of(new BigNumber(((_748_v6).dtor_cf5).length), p0);
        let _750_v8;
        _750_v8 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((p0).multipliedBy(new BigNumber((_745_v3).length))),new BigNumber((_749_v7).length));
        let _index101 = _module.__default.safeIndex(new BigNumber(439), new BigNumber((_744_v2).length));
        (_744_v2)[_index101] = _750_v8;
        let _751_v9;
        _751_v9 = _dafny.Map.Empty.slice().updateUnsafe(!(_742_v0),new BigNumber(-76));
        let _752_v10;
        _752_v10 = _dafny.MultiSet.fromElements(!(_742_v0));
        _751_v9 = (_751_v9).update(!(_742_v0), new BigNumber((_752_v10).cardinality()));
        let _753_v11;
        let _init17 = ((_754_v3, _755_p1, _756_v0) => function (_757_i0) {
          return _dafny.Set.fromElements(new BigNumber(((_754_v3).update(_755_p1, _756_v0)).length));
        })(_745_v3, p1, _742_v0);
        let _nw119 = Array((new BigNumber(2)).toNumber());
        for (let _i0_17 = 0; _i0_17 < new BigNumber(_nw119.length); _i0_17++) {
          _nw119[_i0_17] = _init17(new BigNumber(_i0_17));
        }
        _753_v11 = _nw119;
        let _758_v12;
        let _init18 = ((_759_v0) => function (_760_i1) {
          return _759_v0;
        })(_742_v0);
        let _nw120 = Array((new BigNumber(22)).toNumber());
        for (let _i0_18 = 0; _i0_18 < new BigNumber(_nw120.length); _i0_18++) {
          _nw120[_i0_18] = _init18(new BigNumber(_i0_18));
        }
        _758_v12 = _nw120;
        let _rhs91 = !(((!(_742_v0)) ? (_742_v0) : (_742_v0)));
        let _rhs92 = _753_v11;
        let _rhs93 = !(_742_v0) || (_742_v0);
        let _rhs94 = _742_v0;
        let _rhs95 = _758_v12;
        _742_v0 = _rhs91;
        _753_v11 = _rhs92;
        _742_v0 = _rhs93;
        _742_v0 = _rhs94;
        _758_v12 = _rhs95;
        _743_v1 = _dafny.Seq.Concat(_dafny.Seq.Concat(_743_v1, _module.__default.fm8(p1, globalState)), _743_v1);
      } else {
        _742_v0 = _742_v0;
        let _761_v13;
        _761_v13 = _dafny.Map.Empty.slice().updateUnsafe(_742_v0,_dafny.Seq.UnicodeFromString("uqv"));
        let _762_v14;
        _762_v14 = _dafny.Seq.UnicodeFromString("keousjicc");
        let _763_v15;
        _763_v15 = _dafny.Seq.of((_dafny.ZERO).minus(new BigNumber((_761_v13).length)), new BigNumber((_762_v14).length), p0, p0, (_dafny.ZERO).minus(p1));
        let _764_v16;
        _764_v16 = _dafny.Map.Empty.slice().updateUnsafe(((false) ? (p0) : (p0)),_763_v15);
        let _765_v17;
        let _init19 = ((_766_p0, _767_v0, _768_p1) => function (_769_i2) {
          return _module.D1.create_DC2(_766_p0, _767_v0, _768_p1);
        })(p0, _742_v0, p1);
        let _nw121 = Array((new BigNumber(20)).toNumber());
        for (let _i0_19 = 0; _i0_19 < new BigNumber(_nw121.length); _i0_19++) {
          _nw121[_i0_19] = _init19(new BigNumber(_i0_19));
        }
        _765_v17 = _nw121;
        let _770_v18;
        _770_v18 = _module.D1.create_DC2(p0, true, p1);
        let _index102 = _module.__default.safeIndex(new BigNumber(445), new BigNumber((_765_v17).length));
        (_765_v17)[_index102] = _770_v18;
        let _771_v19;
        _771_v19 = _dafny.Map.Empty.slice().updateUnsafe(p0,_762_v14);
        let _772_v20;
        _772_v20 = _dafny.Set.fromElements((_this).fm7(p0, _742_v0, _771_v19, globalState), false);
        let _773_v21;
        _773_v21 = _dafny.Map.Empty.slice().updateUnsafe(_742_v0,p1);
        let _774_v22;
        _774_v22 = _dafny.Map.Empty.slice().updateUnsafe(_773_v21,_dafny.Seq.UnicodeFromString("grdlrs"));
        let _775_v23;
        let _nw122 = Array((new BigNumber(4)).toNumber()).fill(false);
        _775_v23 = _nw122;
        let _776_v24;
        _776_v24 = _dafny.Map.Empty.slice().updateUnsafe(_775_v23,_742_v0);
        let _777_v25;
        let _nw123 = Array((new BigNumber(25)).toNumber());
        _nw123[(_dafny.ZERO).toNumber()] = (_dafny.ZERO).minus(p1);
        _nw123[(_dafny.ONE).toNumber()] = p1;
        _nw123[(new BigNumber(2)).toNumber()] = p0;
        _nw123[(new BigNumber(3)).toNumber()] = p0;
        _nw123[(new BigNumber(4)).toNumber()] = p1;
        _nw123[(new BigNumber(5)).toNumber()] = p1;
        _nw123[(new BigNumber(6)).toNumber()] = (_dafny.ZERO).minus(new BigNumber((_772_v20).length));
        _nw123[(new BigNumber(7)).toNumber()] = new BigNumber(519);
        _nw123[(new BigNumber(8)).toNumber()] = new BigNumber(118);
        _nw123[(new BigNumber(9)).toNumber()] = p0;
        _nw123[(new BigNumber(10)).toNumber()] = (_763_v15)[_module.__default.safeIndex(p1, new BigNumber((_763_v15).length))];
        _nw123[(new BigNumber(11)).toNumber()] = p1;
        _nw123[(new BigNumber(12)).toNumber()] = p0;
        _nw123[(new BigNumber(13)).toNumber()] = _module.__default.fm9((_dafny.ZERO).minus(p1), globalState);
        _nw123[(new BigNumber(14)).toNumber()] = p0;
        _nw123[(new BigNumber(15)).toNumber()] = p0;
        _nw123[(new BigNumber(16)).toNumber()] = (p1).multipliedBy((_dafny.ZERO).minus(p1));
        _nw123[(new BigNumber(17)).toNumber()] = new BigNumber(((((_774_v22).contains(_773_v21)) ? ((_774_v22).get(_773_v21)) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(220)), function (_778_i3) {
          return (_this).f6;
        })))).length);
        _nw123[(new BigNumber(18)).toNumber()] = (((_773_v21).contains(_742_v0)) ? ((_773_v21).get(_742_v0)) : (p1));
        _nw123[(new BigNumber(19)).toNumber()] = _module.__default.safeDivisionInt(_module.__default.fm9(p0, globalState), new BigNumber((_762_v14).length));
        _nw123[(new BigNumber(20)).toNumber()] = p1;
        _nw123[(new BigNumber(21)).toNumber()] = _module.__default.safeDivisionInt((_dafny.ZERO).minus(p0), p0);
        _nw123[(new BigNumber(22)).toNumber()] = new BigNumber(((_776_v24).Merge(_dafny.Map.Empty.slice().updateUnsafe(_775_v23,_742_v0))).length);
        _nw123[(new BigNumber(23)).toNumber()] = new BigNumber((_dafny.Seq.UnicodeFromString("ngqqsqvhs")).length);
        _nw123[(new BigNumber(24)).toNumber()] = p1;
        _777_v25 = _nw123;
        let _779_v26;
        _779_v26 = _dafny.Seq.of(_742_v0, _742_v0, _742_v0);
        let _780_v27;
        _780_v27 = _dafny.Seq.of(_module.__default.fm10(_742_v0, globalState), _779_v26);
        let _781_v28;
        _781_v28 = (_780_v27)[_module.__default.safeIndex(new BigNumber(62), new BigNumber((_780_v27).length))];
        let _index103 = _module.__default.safeIndex(new BigNumber(351), new BigNumber((_777_v25).length));
        (_777_v25)[_index103] = new BigNumber(((_781_v28)).length);
        let _782_v29;
        _782_v29 = _dafny.MultiSet.fromElements(p0, (new BigNumber((_763_v15).length)).multipliedBy(p0));
        let _783_v30;
        _783_v30 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_773_v21).length),p1);
        let _784_v31;
        _784_v31 = _dafny.MultiSet.fromElements(_module.__default.fm12(_742_v0, p0, _772_v20, _783_v30, globalState));
        let _index104 = _module.__default.safeIndex(new BigNumber(445), new BigNumber((_765_v17).length));
        let _index105 = _module.__default.safeIndex(new BigNumber(351), new BigNumber((_777_v25).length));
        let _rhs96 = (_764_v16).Merge(_764_v16);
        let _rhs97 = _770_v18;
        let _rhs98 = p1;
        let _rhs99 = _module.__default.fm11(_742_v0, _742_v0, _module.D2.create_DC5(p0, _742_v0, _742_v0, _742_v0, _742_v0), globalState);
        let _rhs100 = (_dafny.ZERO).minus(((((_784_v31).contains(_763_v15)) ? ((_784_v31).get(_763_v15)) : (new BigNumber((_779_v26).length)))).plus(p1));
        let _lhs54 = _765_v17;
        let _lhs55 = _module.__default.safeIndex(new BigNumber(445), new BigNumber((_765_v17).length));
        let _lhs56 = _777_v25;
        let _lhs57 = _module.__default.safeIndex(new BigNumber(351), new BigNumber((_777_v25).length));
        let _lhs58 = globalState;
        _764_v16 = _rhs96;
        _lhs54[_lhs55] = _rhs97;
        _lhs56[_lhs57] = _rhs98;
        _782_v29 = _rhs99;
        _lhs58.f2 = _rhs100;
        if (_742_v0) {
          let _index106 = _module.__default.safeIndex(new BigNumber(351), new BigNumber((_777_v25).length));
          (_777_v25)[_index106] = (_777_v25)[_module.__default.safeIndex(new BigNumber(351), new BigNumber((_777_v25).length))];
          let _index107 = _module.__default.safeIndex(new BigNumber(86), new BigNumber((_775_v23).length));
          (_775_v23)[_index107] = _742_v0;
          let _785_v32;
          _785_v32 = _dafny.Map.Empty.slice().updateUnsafe(_742_v0,_742_v0);
          let _786_v33;
          _786_v33 = _dafny.Set.fromElements((_dafny.ZERO).minus((_777_v25)[_module.__default.safeIndex(new BigNumber(351), new BigNumber((_777_v25).length))]), new BigNumber((_783_v30).length));
          let _787_v34;
          _787_v34 = _dafny.Seq.of(_773_v21);
          let _788_v35;
          _788_v35 = _module.D2.create_DC5((_777_v25)[_module.__default.safeIndex(new BigNumber(351), new BigNumber((_777_v25).length))], _742_v0, _742_v0, (((_785_v32).contains(_742_v0)) ? ((_785_v32).get(_742_v0)) : (_742_v0)), (_this).fm7(p0, _742_v0, _771_v19, globalState));
          let _789_v36;
          _789_v36 = _dafny.Map.Empty.slice().updateUnsafe(_775_v23,_788_v35);
          let _790_v37;
          _790_v37 = _dafny.Map.Empty.slice().updateUnsafe(_789_v36,_786_v33);
          let _index108 = _module.__default.safeIndex(new BigNumber(86), new BigNumber((_775_v23).length));
          let _rhs101 = !(true);
          let _rhs102 = (((_777_v25)[_module.__default.safeIndex(new BigNumber(351), new BigNumber((_777_v25).length))]).plus(p0)).plus((p1).plus(p0));
          let _rhs103 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(285)), function (_791_i4) {
            return (_this).f6;
          });
          let _rhs104 = (_785_v32).update((_773_v21).equals((_787_v34)[_module.__default.safeIndex(p1, new BigNumber((_787_v34).length))]), _dafny.Seq.contains(_762_v14, (_this).f6));
          let _rhs105 = (((_790_v37).contains(_789_v36)) ? ((_790_v37).get(_789_v36)) : (_786_v33));
          let _lhs59 = _775_v23;
          let _lhs60 = _module.__default.safeIndex(new BigNumber(86), new BigNumber((_775_v23).length));
          let _lhs61 = globalState;
          _lhs59[_lhs60] = _rhs101;
          _lhs61.f2 = _rhs102;
          _762_v14 = _rhs103;
          _785_v32 = _rhs104;
          _786_v33 = _rhs105;
          let _792_v38;
          _792_v38 = _dafny.Map.Empty.slice().updateUnsafe((_777_v25)[_module.__default.safeIndex(new BigNumber(351), new BigNumber((_777_v25).length))],(_775_v23)[_module.__default.safeIndex(new BigNumber(86), new BigNumber((_775_v23).length))]);
          let _793_v41;
          _793_v41 = _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('o'.codePointAt(0)),(_775_v23)[_module.__default.safeIndex(new BigNumber(86), new BigNumber((_775_v23).length))]);
          let _794_v42;
          let _nw124 = Array((new BigNumber(14)).toNumber());
          _nw124[(_dafny.ZERO).toNumber()] = _792_v38;
          _nw124[(_dafny.ONE).toNumber()] = _792_v38;
          _nw124[(new BigNumber(2)).toNumber()] = function () {
            let _coll26 = new _dafny.Map();
            for (const _compr_26 of _dafny.IntegerRange(new BigNumber(305), new BigNumber(57))) {
              let _795_v39 = _compr_26;
              if (((new BigNumber(305)).isLessThanOrEqualTo(_795_v39)) && ((_795_v39).isLessThan(new BigNumber(57)))) {
                _coll26.push([(_795_v39).minus(p1),_742_v0]);
              }
            }
            return _coll26;
          }();
          _nw124[(new BigNumber(3)).toNumber()] = _792_v38;
          _nw124[(new BigNumber(4)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe((_777_v25)[_module.__default.safeIndex(new BigNumber(351), new BigNumber((_777_v25).length))],(_775_v23)[_module.__default.safeIndex(new BigNumber(86), new BigNumber((_775_v23).length))]);
          _nw124[(new BigNumber(5)).toNumber()] = _792_v38;
          _nw124[(new BigNumber(6)).toNumber()] = _792_v38;
          _nw124[(new BigNumber(7)).toNumber()] = function () {
            let _coll27 = new _dafny.Map();
            for (const _compr_27 of (_792_v38).Keys.Elements) {
              let _796_v40 = _compr_27;
              if ((_792_v38).contains(_796_v40)) {
                _coll27.push([_module.__default.safeModuloInt(_796_v40, p0),(_775_v23)[_module.__default.safeIndex(new BigNumber(86), new BigNumber((_775_v23).length))]]);
              }
            }
            return _coll27;
          }();
          _nw124[(new BigNumber(8)).toNumber()] = _module.__default.fm34((_775_v23)[_module.__default.safeIndex(new BigNumber(86), new BigNumber((_775_v23).length))], new _dafny.CodePoint('k'.codePointAt(0)), _763_v15, globalState);
          _nw124[(new BigNumber(9)).toNumber()] = _792_v38;
          _nw124[(new BigNumber(10)).toNumber()] = _792_v38;
          _nw124[(new BigNumber(11)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_783_v30).length),(((_793_v41).contains((_this).f6)) ? ((_793_v41).get((_this).f6)) : ((((_785_v32).contains((_775_v23)[_module.__default.safeIndex(new BigNumber(86), new BigNumber((_775_v23).length))])) ? ((_785_v32).get((_775_v23)[_module.__default.safeIndex(new BigNumber(86), new BigNumber((_775_v23).length))])) : (_742_v0)))));
          _nw124[(new BigNumber(12)).toNumber()] = _792_v38;
          _nw124[(new BigNumber(13)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe((_777_v25)[_module.__default.safeIndex(new BigNumber(351), new BigNumber((_777_v25).length))],_742_v0);
          _794_v42 = _nw124;
          let _797_v43;
          let _nw125 = new _module.C2();
          _nw125.__ctor(true, (_777_v25)[_module.__default.safeIndex(new BigNumber(351), new BigNumber((_777_v25).length))], _794_v42);
          _797_v43 = _nw125;
          let _798_v44;
          _798_v44 = _dafny.Map.Empty.slice().updateUnsafe((_this).fm7((_777_v25)[_module.__default.safeIndex(new BigNumber(351), new BigNumber((_777_v25).length))], _742_v0, _771_v19, globalState),_797_v43);
          let _799_v45;
          _799_v45 = _dafny.Seq.of(_797_v43);
          _798_v44 = (_798_v44).update(true, (_799_v45)[_module.__default.safeIndex((_777_v25)[_module.__default.safeIndex(new BigNumber(351), new BigNumber((_777_v25).length))], new BigNumber((_799_v45).length))]);
          let _800_v46;
          _800_v46 = _module.D3.create_DC7(new BigNumber((_762_v14).length), p0);
          let _801_v47;
          _801_v47 = _dafny.Map.Empty.slice().updateUnsafe(_800_v46,_777_v25);
          let _802_v48;
          _802_v48 = _dafny.Map.Empty.slice().updateUnsafe(_742_v0,_777_v25);
          let _803_v49;
          let _nw126 = Array((new BigNumber(17)).toNumber());
          _nw126[(_dafny.ZERO).toNumber()] = (((_801_v47).contains(_800_v46)) ? ((_801_v47).get(_800_v46)) : (_777_v25));
          _nw126[(_dafny.ONE).toNumber()] = _777_v25;
          _nw126[(new BigNumber(2)).toNumber()] = _777_v25;
          _nw126[(new BigNumber(3)).toNumber()] = _777_v25;
          _nw126[(new BigNumber(4)).toNumber()] = _777_v25;
          _nw126[(new BigNumber(5)).toNumber()] = _777_v25;
          _nw126[(new BigNumber(6)).toNumber()] = _777_v25;
          _nw126[(new BigNumber(7)).toNumber()] = (((_802_v48).contains((_775_v23)[_module.__default.safeIndex(new BigNumber(86), new BigNumber((_775_v23).length))])) ? ((_802_v48).get((_775_v23)[_module.__default.safeIndex(new BigNumber(86), new BigNumber((_775_v23).length))])) : (_777_v25));
          _nw126[(new BigNumber(8)).toNumber()] = _777_v25;
          _nw126[(new BigNumber(9)).toNumber()] = _777_v25;
          _nw126[(new BigNumber(10)).toNumber()] = _777_v25;
          _nw126[(new BigNumber(11)).toNumber()] = _777_v25;
          _nw126[(new BigNumber(12)).toNumber()] = _777_v25;
          _nw126[(new BigNumber(13)).toNumber()] = _777_v25;
          _nw126[(new BigNumber(14)).toNumber()] = _777_v25;
          _nw126[(new BigNumber(15)).toNumber()] = _777_v25;
          _nw126[(new BigNumber(16)).toNumber()] = _777_v25;
          _803_v49 = _nw126;
          let _nw127 = Array((new BigNumber(9)).toNumber()).fill([]);
          _803_v49 = _nw127;
          _785_v32 = (_785_v32).update((_775_v23)[_module.__default.safeIndex(new BigNumber(86), new BigNumber((_775_v23).length))], (_775_v23)[_module.__default.safeIndex(new BigNumber(86), new BigNumber((_775_v23).length))]);
        } else {
          let _804_v50;
          _804_v50 = _module.D3.create_DC7((_777_v25)[_module.__default.safeIndex(new BigNumber(351), new BigNumber((_777_v25).length))], p0);
          let _805_v51;
          _805_v51 = _dafny.Seq.of(_804_v50, _804_v50, _804_v50);
          let _index109 = _module.__default.safeIndex(new BigNumber(140), new BigNumber((_775_v23).length));
          (_775_v23)[_index109] = _742_v0;
          let _806_v52;
          _806_v52 = _module.D4.create_DC9(p0, (_this).f6, _763_v15, new BigNumber((_762_v14).length), true);
          let _807_v53;
          _807_v53 = _dafny.Map.Empty.slice().updateUnsafe(_806_v52,_775_v23);
          let _index110 = _module.__default.safeIndex(new BigNumber(140), new BigNumber((_775_v23).length));
          let _rhs106 = _805_v51;
          let _rhs107 = _dafny.Seq.update(_762_v14, _module.__default.safeIndex((_777_v25)[_module.__default.safeIndex(new BigNumber(351), new BigNumber((_777_v25).length))], new BigNumber((_762_v14).length)), (_this).f6);
          let _rhs108 = _module.__default.fm17(_dafny.Seq.of(_762_v14), p1, _module.__default.safeModuloInt(p0, new BigNumber(-351)), (_777_v25)[_module.__default.safeIndex(new BigNumber(351), new BigNumber((_777_v25).length))], globalState);
          let _rhs109 = (((_807_v53).contains(_806_v52)) ? ((_807_v53).get(_806_v52)) : (_775_v23));
          let _rhs110 = (((_779_v26)[_module.__default.safeIndex((_777_v25)[_module.__default.safeIndex(new BigNumber(351), new BigNumber((_777_v25).length))], new BigNumber((_779_v26).length))]) ? (_742_v0) : (!(((_777_v25)[_module.__default.safeIndex(new BigNumber(351), new BigNumber((_777_v25).length))]).isLessThan(new BigNumber((_762_v14).length)))));
          let _lhs62 = globalState;
          let _lhs63 = _775_v23;
          let _lhs64 = _module.__default.safeIndex(new BigNumber(140), new BigNumber((_775_v23).length));
          _805_v51 = _rhs106;
          _762_v14 = _rhs107;
          _lhs62.f2 = _rhs108;
          _775_v23 = _rhs109;
          _lhs63[_lhs64] = _rhs110;
          let _index111 = _module.__default.safeIndex(new BigNumber(140), new BigNumber((_775_v23).length));
          (_775_v23)[_index111] = _module.__default.fm23(_dafny.Seq.Concat(_762_v14, _dafny.Seq.update(_762_v14, _module.__default.safeIndex(new BigNumber((_763_v15).length), new BigNumber((_762_v14).length)), (_this).f6)), p1, p1, !(false), globalState);
          let _808_v54;
          _808_v54 = _dafny.Seq.of(_781_v28);
          _742_v0 = _dafny.Seq.IsPrefixOf(_808_v54, _808_v54);
          let _index112 = _module.__default.safeIndex(new BigNumber(351), new BigNumber((_777_v25).length));
          (_777_v25)[_index112] = p0;
          _773_v21 = _773_v21;
        }
        let _index113 = _module.__default.safeIndex(new BigNumber(351), new BigNumber((_777_v25).length));
        let _rhs111 = new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(356)), function (_809_i5) {
          return (_this).f6;
        })).length);
        let _rhs112 = _dafny.Map.Empty.slice().updateUnsafe(true,_dafny.Seq.Concat(_762_v14, _dafny.Seq.Create(_module.__default.abs(new BigNumber(121)), function (_810_i6) {
          return (_this).f6;
        })));
        let _lhs65 = _777_v25;
        let _lhs66 = _module.__default.safeIndex(new BigNumber(351), new BigNumber((_777_v25).length));
        _lhs65[_lhs66] = _rhs111;
        _761_v13 = _rhs112;
        let _811_v55;
        let _nw128 = Array((new BigNumber(2)).toNumber()).fill(_dafny.Map.Empty);
        _811_v55 = _nw128;
        let _812_v56;
        let _nw129 = new _module.C2();
        _nw129.__ctor((_779_v26)[_module.__default.safeIndex(new BigNumber((_762_v14).length), new BigNumber((_779_v26).length))], new BigNumber((_779_v26).length), _811_v55);
        _812_v56 = _nw129;
      }
      let _813_v57;
      _813_v57 = _module.D6.create_DC11(_dafny.Seq.UnicodeFromString("x"));
      let _source14 = _813_v57;
      if (_source14.is_DC12) {
        let _814___mcc_h0 = (_source14).cf23;
        let _815___mcc_h1 = (_source14).cf24;
        let _816_cf24 = _815___mcc_h1;
        let _817_cf23 = _814___mcc_h0;
        if (_817_cf23) {
          let _818_v58;
          _818_v58 = _dafny.Seq.UnicodeFromString("wgyx");
          let _819_v59;
          let _nw130 = Array((new BigNumber(6)).toNumber());
          _nw130[(_dafny.ZERO).toNumber()] = _818_v58;
          _nw130[(_dafny.ONE).toNumber()] = _818_v58;
          _nw130[(new BigNumber(2)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(536)), ((_820_cf24) => function (_821_i7) {
            return _820_cf24;
          })(_816_cf24));
          _nw130[(new BigNumber(3)).toNumber()] = _818_v58;
          _nw130[(new BigNumber(4)).toNumber()] = _dafny.Seq.UnicodeFromString("ef");
          _nw130[(new BigNumber(5)).toNumber()] = _818_v58;
          _819_v59 = _nw130;
          let _822_v60;
          _822_v60 = _module.D9.create_DC18(_819_v59);
          let _823_v61;
          let _nw131 = Array((new BigNumber(18)).toNumber());
          _nw131[(_dafny.ZERO).toNumber()] = _819_v59;
          _nw131[(_dafny.ONE).toNumber()] = _819_v59;
          _nw131[(new BigNumber(2)).toNumber()] = _819_v59;
          _nw131[(new BigNumber(3)).toNumber()] = _819_v59;
          _nw131[(new BigNumber(4)).toNumber()] = _819_v59;
          _nw131[(new BigNumber(5)).toNumber()] = _819_v59;
          _nw131[(new BigNumber(6)).toNumber()] = _819_v59;
          _nw131[(new BigNumber(7)).toNumber()] = _819_v59;
          _nw131[(new BigNumber(8)).toNumber()] = _819_v59;
          _nw131[(new BigNumber(9)).toNumber()] = _819_v59;
          _nw131[(new BigNumber(10)).toNumber()] = _819_v59;
          _nw131[(new BigNumber(11)).toNumber()] = (_822_v60).dtor_cf33;
          _nw131[(new BigNumber(12)).toNumber()] = _819_v59;
          _nw131[(new BigNumber(13)).toNumber()] = _819_v59;
          _nw131[(new BigNumber(14)).toNumber()] = _819_v59;
          _nw131[(new BigNumber(15)).toNumber()] = ((_817_cf23) ? (_819_v59) : (_819_v59));
          _nw131[(new BigNumber(16)).toNumber()] = _819_v59;
          _nw131[(new BigNumber(17)).toNumber()] = _819_v59;
          _823_v61 = _nw131;
          let _index114 = _module.__default.safeIndex(new BigNumber(959), new BigNumber((_823_v61).length));
          (_823_v61)[_index114] = _819_v59;
          let _index115 = _module.__default.safeIndex(new BigNumber(959), new BigNumber((_823_v61).length));
          (_823_v61)[_index115] = _819_v59;
          let _824_v62;
          let _nw132 = Array((new BigNumber(8)).toNumber()).fill(false);
          _824_v62 = _nw132;
          _824_v62 = _824_v62;
          _817_cf23 = !(_module.__default.fm21((p1).multipliedBy(new BigNumber(73)), new _dafny.CodePoint('i'.codePointAt(0)), _dafny.Seq.IsProperPrefixOf(_818_v58, _818_v58), _817_cf23, globalState));
          let _825_v63;
          _825_v63 = _dafny.Set.fromElements(!(_742_v0));
          let _826_v64;
          _826_v64 = _dafny.Map.Empty.slice().updateUnsafe(_825_v63,_818_v58);
          _826_v64 = (_826_v64).update(_825_v63, _dafny.Seq.Create(_module.__default.abs(new BigNumber(981)), function (_827_i8) {
            return (_this).f6;
          }));
          let _828_v66;
          _828_v66 = _module.D2.create_DC5(new BigNumber((function () {
  let _coll28 = new _dafny.Map();
  for (const _compr_28 of _dafny.IntegerRange(new BigNumber(345), new BigNumber(-615))) {
    let _829_v65 = _compr_28;
    if (((new BigNumber(345)).isLessThanOrEqualTo(_829_v65)) && ((_829_v65).isLessThan(new BigNumber(-615)))) {
      _coll28.push([(_829_v65).multipliedBy(p0),p1]);
    }
  }
  return _coll28;
}()).length), _742_v0, _817_cf23, _817_cf23, _817_cf23);
          let _pat_let_tv20 = _817_cf23;
          let _pat_let_tv21 = _817_cf23;
          _828_v66 = function (_pat_let6_0) {
            return function (_830_dt__update__tmp_h0) {
              return function (_pat_let7_0) {
                return function (_831_dt__update_hcf8_h0) {
                  return function (_pat_let8_0) {
                    return function (_832_dt__update_hcf10_h0) {
                      return _module.D2.create_DC5((_830_dt__update__tmp_h0).dtor_cf7, _831_dt__update_hcf8_h0, (_830_dt__update__tmp_h0).dtor_cf9, _832_dt__update_hcf10_h0, (_830_dt__update__tmp_h0).dtor_cf11);
                    }(_pat_let8_0);
                  }(_pat_let_tv21);
                }(_pat_let7_0);
              }(_pat_let_tv20);
            }(_pat_let6_0);
          }(_828_v66);
        } else {
          let _833_v67;
          _833_v67 = _dafny.Map.Empty.slice().updateUnsafe((_742_v0) && (_817_cf23),p1);
          _833_v67 = (_833_v67).update(_817_cf23, p0);
          let _834_v68;
          _834_v68 = _dafny.Seq.of(new BigNumber(291), p0);
          let _835_v69;
          _835_v69 = _dafny.Seq.UnicodeFromString("ulgwrp");
          let _836_v70;
          let _nw133 = Array((new BigNumber(7)).toNumber()).fill(_dafny.Map.Empty);
          _836_v70 = _nw133;
          let _837_v71;
          let _nw134 = new _module.C2();
          _nw134.__ctor(true, p1, _836_v70);
          _837_v71 = _nw134;
          let _838_v72;
          _838_v72 = _dafny.Set.fromElements(p0);
          let _839_v73;
          let _nw135 = new _module.C0();
          _nw135.__ctor(p0);
          _839_v73 = _nw135;
          let _840_v74;
          _840_v74 = _dafny.MultiSet.fromElements(_839_v73, _839_v73);
          let _841_v75;
          let _nw136 = Array((new BigNumber(15)).toNumber());
          _nw136[(_dafny.ZERO).toNumber()] = p0;
          _nw136[(_dafny.ONE).toNumber()] = new BigNumber((_835_v69).length);
          _nw136[(new BigNumber(2)).toNumber()] = p0;
          _nw136[(new BigNumber(3)).toNumber()] = p1;
          _nw136[(new BigNumber(4)).toNumber()] = p1;
          _nw136[(new BigNumber(5)).toNumber()] = p1;
          _nw136[(new BigNumber(6)).toNumber()] = new BigNumber((_834_v68).length);
          _nw136[(new BigNumber(7)).toNumber()] = new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_837_v71,_817_cf23)).length);
          _nw136[(new BigNumber(8)).toNumber()] = new BigNumber((_838_v72).length);
          _nw136[(new BigNumber(9)).toNumber()] = p1;
          _nw136[(new BigNumber(10)).toNumber()] = p1;
          _nw136[(new BigNumber(11)).toNumber()] = p1;
          _nw136[(new BigNumber(12)).toNumber()] = new BigNumber((_834_v68).length);
          _nw136[(new BigNumber(13)).toNumber()] = new BigNumber((_840_v74).cardinality());
          _nw136[(new BigNumber(14)).toNumber()] = new BigNumber(-423);
          _841_v75 = _nw136;
          let _842_v76;
          _842_v76 = _dafny.Map.Empty.slice().updateUnsafe(_841_v75,_742_v0);
          let _843_v77;
          _843_v77 = _module.D8.create_DC16((_834_v68)[_module.__default.safeIndex(p1, new BigNumber((_834_v68).length))], (((_842_v76).contains(_841_v75)) ? ((_842_v76).get(_841_v75)) : (_817_cf23)), _817_cf23, p1);
          _742_v0 = (_843_v77).dtor_cf29;
          let _844_v78;
          _844_v78 = _dafny.Seq.of(_835_v69);
          _835_v69 = _dafny.Seq.Concat(_835_v69, _dafny.Seq.update(_dafny.Seq.Concat(_835_v69, _dafny.Seq.Create(_module.__default.abs(new BigNumber(13)), function (_845_i9) {
            return (_this).f6;
          })), _module.__default.safeIndex((_834_v68)[_module.__default.safeIndex(_module.__default.fm17(_844_v78, p0, new BigNumber((_835_v69).length), new BigNumber((_838_v72).length), globalState), new BigNumber((_834_v68).length))], new BigNumber((_dafny.Seq.Concat(_835_v69, _dafny.Seq.Create(_module.__default.abs(new BigNumber(13)), function (_846_i9) {
            return (_this).f6;
          }))).length)), _816_cf24));
          let _index116 = _module.__default.safeIndex(new BigNumber(309), new BigNumber((_841_v75).length));
          (_841_v75)[_index116] = p0;
          let _847_v79;
          _847_v79 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm35(globalState),_817_cf23);
          let _848_v80;
          _848_v80 = _dafny.Set.fromElements(_835_v69);
          let _index117 = _module.__default.safeIndex(new BigNumber(309), new BigNumber((_841_v75).length));
          let _rhs113 = p1;
          let _rhs114 = _813_v57;
          let _rhs115 = p1;
          let _rhs116 = (((_847_v79).contains(_848_v80)) ? ((_847_v79).get(_848_v80)) : (!(_742_v0)));
          let _rhs117 = _838_v72;
          let _lhs67 = globalState;
          let _lhs68 = _841_v75;
          let _lhs69 = _module.__default.safeIndex(new BigNumber(309), new BigNumber((_841_v75).length));
          _lhs67.f2 = _rhs113;
          _813_v57 = _rhs114;
          _lhs68[_lhs69] = _rhs115;
          _742_v0 = _rhs116;
          _838_v72 = _rhs117;
          let _849_v81;
          _849_v81 = _dafny.Map.Empty.slice().updateUnsafe(_841_v75,p0);
          _849_v81 = (_849_v81).update(_841_v75, new BigNumber(657));
        }
        (globalState).f2 = p1;
        let _850_v82;
        _850_v82 = _dafny.Seq.UnicodeFromString("a");
        _850_v82 = _dafny.Seq.Concat(_850_v82, ((_742_v0) ? (_850_v82) : (_850_v82)));
        if (_module.__default.fm23(_dafny.Seq.Create(_module.__default.abs(new BigNumber(608)), ((_851_cf24) => function (_852_i10) {
          return _851_cf24;
        })(_816_cf24)), new BigNumber(296), _module.__default.safeModuloInt(p0, new BigNumber((_850_v82).length)), _817_cf23, globalState)) {
          let _853_v83;
          let _init20 = ((_854_p0) => function (_855_i11) {
            return (_855_i11).minus(_854_p0);
          })(p0);
          let _nw137 = Array((new BigNumber(4)).toNumber());
          for (let _i0_20 = 0; _i0_20 < new BigNumber(_nw137.length); _i0_20++) {
            _nw137[_i0_20] = _init20(new BigNumber(_i0_20));
          }
          _853_v83 = _nw137;
          let _index118 = _module.__default.safeIndex(new BigNumber(915), new BigNumber((_853_v83).length));
          (_853_v83)[_index118] = p0;
          let _856_v84;
          _856_v84 = _dafny.Seq.of(_850_v82, _dafny.Seq.Create(_module.__default.abs(new BigNumber(775)), ((_857_cf24) => function (_858_i12) {
            return _857_cf24;
          })(_816_cf24)));
          let _859_v85;
          _859_v85 = _dafny.Seq.of((_856_v84)[_module.__default.safeIndex(p0, new BigNumber((_856_v84).length))], _850_v82, _module.__default.fm25(_742_v0, p0, globalState));
          let _860_v86;
          _860_v86 = _dafny.Map.Empty.slice().updateUnsafe(p0,_850_v82);
          let _index119 = _module.__default.safeIndex(new BigNumber(915), new BigNumber((_853_v83).length));
          let _rhs118 = _module.__default.fm17(((_742_v0) ? (_859_v85) : (_859_v85)), p1, new BigNumber(82), p0, globalState);
          let _rhs119 = (_this).fm7(p0, _817_cf23, (_860_v86).update(p0, _850_v82), globalState);
          let _lhs70 = _853_v83;
          let _lhs71 = _module.__default.safeIndex(new BigNumber(915), new BigNumber((_853_v83).length));
          _lhs70[_lhs71] = _rhs118;
          _817_cf23 = _rhs119;
          let _861_v87;
          let _init21 = ((_862_v0, _863_cf23) => function (_864_i13) {
            return _dafny.Seq.Concat(_dafny.Seq.of(_862_v0, _863_cf23), _dafny.Seq.of(!(true)));
          })(_742_v0, _817_cf23);
          let _nw138 = Array((new BigNumber(15)).toNumber());
          for (let _i0_21 = 0; _i0_21 < new BigNumber(_nw138.length); _i0_21++) {
            _nw138[_i0_21] = _init21(new BigNumber(_i0_21));
          }
          _861_v87 = _nw138;
          let _865_v88;
          _865_v88 = _dafny.Seq.of(_817_cf23);
          let _index120 = _module.__default.safeIndex(new BigNumber(45), new BigNumber((_861_v87).length));
          (_861_v87)[_index120] = _dafny.Seq.Concat(_865_v88, _865_v88);
          let _index121 = _module.__default.safeIndex(new BigNumber(45), new BigNumber((_861_v87).length));
          (_861_v87)[_index121] = _dafny.Seq.Concat(_dafny.Seq.of(_817_cf23), _865_v88);
          let _866_v89;
          _866_v89 = _dafny.Seq.of(((_853_v83)[_module.__default.safeIndex(new BigNumber(915), new BigNumber((_853_v83).length))]).plus(p0));
          let _867_v90;
          _867_v90 = _module.D4.create_DC9((_dafny.ZERO).minus((_853_v83)[_module.__default.safeIndex(new BigNumber(915), new BigNumber((_853_v83).length))]), (_this).f6, _866_v89, p0, _817_cf23);
          let _868_v91;
          _868_v91 = _dafny.Set.fromElements(_742_v0, _817_cf23);
          let _869_v92;
          _869_v92 = _dafny.Map.Empty.slice().updateUnsafe(p0,p0);
          _866_v89 = ((_817_cf23) ? (_866_v89) : (_dafny.Seq.Concat((_867_v90).dtor_cf18, _module.__default.fm12(true, p0, _868_v91, _869_v92, globalState))));
          let _870_v93;
          _870_v93 = _dafny.Map.Empty.slice().updateUnsafe(false,p0);
          _870_v93 = (_870_v93).update(_817_cf23, (_dafny.ZERO).minus(p1));
          _742_v0 = ((_dafny.Seq.IsPrefixOf(_dafny.Seq.Create(_module.__default.abs(new BigNumber(900)), ((_871_v83) => function (_872_i14) {
            return (_871_v83)[_module.__default.safeIndex(new BigNumber(915), new BigNumber((_871_v83).length))];
          })(_853_v83)), _dafny.Seq.of(p0, p0))) ? ((new BigNumber(834)).isEqualTo(p1)) : (_742_v0));
        } else {
          let _873_v94;
          let _init22 = ((_874_p1) => function (_875_i15) {
            return (_875_i15).plus(_874_p1);
          })(p1);
          let _nw139 = Array((new BigNumber(26)).toNumber());
          for (let _i0_22 = 0; _i0_22 < new BigNumber(_nw139.length); _i0_22++) {
            _nw139[_i0_22] = _init22(new BigNumber(_i0_22));
          }
          _873_v94 = _nw139;
          let _index122 = _module.__default.safeIndex(new BigNumber(457), new BigNumber((_873_v94).length));
          (_873_v94)[_index122] = p1;
          let _index123 = _module.__default.safeIndex(new BigNumber(457), new BigNumber((_873_v94).length));
          (_873_v94)[_index123] = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.update(_850_v82, _module.__default.safeIndex(p1, new BigNumber((_850_v82).length)), _816_cf24), _dafny.Seq.UnicodeFromString("q")), _dafny.Seq.UnicodeFromString("lojxgqr"))).length);
          let _876_v95;
          let _nw140 = new _module.C0();
          _nw140.__ctor(p0);
          _876_v95 = _nw140;
          let _877_v96;
          _877_v96 = _module.D4.create_DC8(_876_v95);
          let _878_v97;
          _878_v97 = _dafny.MultiSet.fromElements(_877_v96, _877_v96);
          let _879_v98;
          _879_v98 = _module.D7.create_DC14((((_878_v97).contains(_877_v96)) ? ((_878_v97).get(_877_v96)) : (p0)));
          let _pat_let_tv22 = _873_v94;
          let _pat_let_tv23 = _873_v94;
          _879_v98 = function (_pat_let9_0) {
            return function (_880_dt__update__tmp_h1) {
              return function (_pat_let10_0) {
                return function (_881_dt__update_hcf26_h0) {
                  return _module.D7.create_DC14(_881_dt__update_hcf26_h0);
                }(_pat_let10_0);
              }((_pat_let_tv23)[_module.__default.safeIndex(new BigNumber(457), new BigNumber((_pat_let_tv22).length))]);
            }(_pat_let9_0);
          }(_module.__default.fm36(_742_v0, globalState));
          let _882_v99;
          _882_v99 = _873_v94;
          let _883_v100;
          _883_v100 = _dafny.MultiSet.fromElements((_882_v99), _873_v94, _873_v94);
          _883_v100 = _dafny.MultiSet.fromElements(_873_v94, ((_742_v0) ? (_873_v94) : (_873_v94)), _873_v94, _873_v94, _873_v94);
          _883_v100 = _883_v100;
          _817_cf23 = _817_cf23;
        }
      } else {
        let _884___mcc_h2 = (_source14).cf22;
        let _885_cf22 = _884___mcc_h2;
        let _886_v101;
        _886_v101 = _dafny.Seq.of(_885_cf22, _885_cf22);
        let _887_v102;
        _887_v102 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.contains((_813_v57).dtor_cf22, (_this).f6),_module.__default.fm17(_886_v101, p1, new BigNumber(803), p0, globalState));
        _887_v102 = (_887_v102).update(_742_v0, ((_742_v0) ? (p1) : (new BigNumber((_885_cf22).length))));
        let _888_v103;
        _888_v103 = _dafny.Seq.of(p0, new BigNumber(192));
        let _889_v104;
        _889_v104 = _module.D2.create_DC4((_888_v103)[_module.__default.safeIndex(p0, new BigNumber((_888_v103).length))]);
        let _890_v105;
        _890_v105 = _dafny.Seq.of(_889_v104);
        _890_v105 = _dafny.Seq.Concat(_dafny.Seq.of(_889_v104), _dafny.Seq.Concat(_890_v105, _890_v105));
        let _891_v106;
        let _nw141 = Array((new BigNumber(3)).toNumber()).fill(false);
        _891_v106 = _nw141;
        let _892_v107;
        _892_v107 = _dafny.Seq.of(_891_v106);
        let _893_v108;
        _893_v108 = _dafny.Seq.of((_892_v107)[_module.__default.safeIndex(p1, new BigNumber((_892_v107).length))], _891_v106);
        let _894_v109;
        _894_v109 = _module.D8.create_DC15((_893_v108)[_module.__default.safeIndex(p0, new BigNumber((_893_v108).length))]);
        let _source15 = _894_v109;
        if (_source15.is_DC16) {
          let _895___mcc_h3 = (_source15).cf28;
          let _896___mcc_h4 = (_source15).cf29;
          let _897___mcc_h5 = (_source15).cf30;
          let _898___mcc_h6 = (_source15).cf31;
          let _899_cf31 = _898___mcc_h6;
          let _900_cf30 = _897___mcc_h5;
          let _901_cf29 = _896___mcc_h4;
          let _902_cf28 = _895___mcc_h3;
          _742_v0 = _742_v0;
          _900_cf30 = _900_cf30;
          let _903_v110;
          let _nw142 = Array((new BigNumber(8)).toNumber()).fill(_dafny.ZERO);
          _903_v110 = _nw142;
          let _904_v111;
          _904_v111 = _dafny.MultiSet.fromElements(p1);
          let _index124 = _module.__default.safeIndex(new BigNumber(150), new BigNumber((_903_v110).length));
          (_903_v110)[_index124] = _module.__default.safeModuloInt((((_904_v111).contains(_899_cf31)) ? ((_904_v111).get(_899_cf31)) : (new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_901_cf29,p0)).length))), (_889_v104).dtor_cf6);
          let _index125 = _module.__default.safeIndex(new BigNumber(150), new BigNumber((_903_v110).length));
          (_903_v110)[_index125] = _902_cf28;
          let _905_v112;
          _905_v112 = new _dafny.CodePoint('m'.codePointAt(0));
          _905_v112 = _905_v112;
        } else if (_source15.is_DC15) {
          let _906___mcc_h7 = (_source15).cf27;
          let _907_cf27 = _906___mcc_h7;
          let _908_v113;
          _908_v113 = _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(_742_v0,_742_v0));
          let _909_v114;
          _909_v114 = _dafny.Map.Empty.slice().updateUnsafe(p1,(_908_v113)[_module.__default.safeIndex(p0, new BigNumber((_908_v113).length))]);
          let _910_v115;
          _910_v115 = _dafny.Set.fromElements(p1);
          let _911_v116;
          _911_v116 = _dafny.Map.Empty.slice().updateUnsafe(_742_v0,_742_v0);
          let _912_v117;
          _912_v117 = _dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe(_742_v0,_742_v0), (((_909_v114).contains(_module.__default.fm24(_910_v115, globalState))) ? ((_909_v114).get(_module.__default.fm24(_910_v115, globalState))) : (_911_v116)));
          let _913_v118;
          _913_v118 = _dafny.Seq.of(_912_v117);
          let _914_v119;
          _914_v119 = _dafny.Seq.of(_742_v0, _742_v0, _742_v0, _742_v0);
          _742_v0 = (((_913_v118)[_module.__default.safeIndex(new BigNumber((_914_v119).length), new BigNumber((_913_v118).length))]).Union(_912_v117)).IsSubsetOf(_dafny.MultiSet.FromArray(_dafny.Seq.Concat(_908_v113, _module.__default.fm37(_module.D1.create_DC2(p1, false, p0), _742_v0, globalState))));
          _742_v0 = _742_v0;
          (globalState).f2 = new BigNumber(14);
          let _915_v121;
          _915_v121 = _dafny.Map.Empty.slice().updateUnsafe(_885_cf22,new BigNumber(746));
          let _916_v122;
          _916_v122 = _dafny.MultiSet.fromElements(p1, new BigNumber((function () {
            let _coll29 = new _dafny.Map();
            for (const _compr_29 of (_915_v121).Keys.Elements) {
              let _917_v120 = _compr_29;
              if ((_915_v121).contains(_917_v120)) {
                _coll29.push([_917_v120,_742_v0]);
              }
            }
            return _coll29;
          }()).length), p1);
          let _918_v123;
          let _nw143 = new _module.C3();
          _nw143.__ctor(_916_v122, _742_v0);
          _918_v123 = _nw143;
          let _nw144 = new _module.C3();
          _nw144.__ctor((_918_v123).f9, true);
          _918_v123 = _nw144;
        } else {
          let _919___mcc_h8 = (_source15).cf32;
          let _920_cf32 = _919___mcc_h8;
          (globalState).f2 = p0;
          let _921_v124;
          _921_v124 = _dafny.MultiSet.fromElements(p1, p1);
          let _922_v125;
          let _nw145 = Array((new BigNumber(21)).toNumber()).fill(_dafny.Map.Empty);
          _922_v125 = _nw145;
          let _923_v126;
          let _nw146 = new _module.C4();
          _nw146.__ctor(_742_v0, _921_v124, _922_v125);
          _923_v126 = _nw146;
          let _924_v127;
          let _init23 = function (_925_i16) {
            return _module.__default.safeDivisionInt(_925_i16, new BigNumber(778));
          };
          let _nw147 = Array((new BigNumber(28)).toNumber());
          for (let _i0_23 = 0; _i0_23 < new BigNumber(_nw147.length); _i0_23++) {
            _nw147[_i0_23] = _init23(new BigNumber(_i0_23));
          }
          _924_v127 = _nw147;
          let _926_v128;
          _926_v128 = _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('e'.codePointAt(0)),_924_v127);
          let _927_v129;
          _927_v129 = _924_v127;
          let _928_v130;
          _928_v130 = _dafny.Map.Empty.slice().updateUnsafe(_742_v0,_924_v127);
          let _929_v131;
          let _nw148 = Array((new BigNumber(17)).toNumber());
          _nw148[(_dafny.ZERO).toNumber()] = ((_742_v0) ? (_924_v127) : (_924_v127));
          _nw148[(_dafny.ONE).toNumber()] = _924_v127;
          _nw148[(new BigNumber(2)).toNumber()] = _924_v127;
          _nw148[(new BigNumber(3)).toNumber()] = _924_v127;
          _nw148[(new BigNumber(4)).toNumber()] = _924_v127;
          _nw148[(new BigNumber(5)).toNumber()] = (((_926_v128).contains((_this).f6)) ? ((_926_v128).get((_this).f6)) : (_924_v127));
          _nw148[(new BigNumber(6)).toNumber()] = _924_v127;
          _nw148[(new BigNumber(7)).toNumber()] = _924_v127;
          _nw148[(new BigNumber(8)).toNumber()] = (_927_v129);
          _nw148[(new BigNumber(9)).toNumber()] = _924_v127;
          _nw148[(new BigNumber(10)).toNumber()] = _924_v127;
          _nw148[(new BigNumber(11)).toNumber()] = _924_v127;
          _nw148[(new BigNumber(12)).toNumber()] = _924_v127;
          _nw148[(new BigNumber(13)).toNumber()] = _924_v127;
          _nw148[(new BigNumber(14)).toNumber()] = (((_928_v130).contains((_923_v126).f7)) ? ((_928_v130).get((_923_v126).f7)) : (_924_v127));
          _nw148[(new BigNumber(15)).toNumber()] = _924_v127;
          _nw148[(new BigNumber(16)).toNumber()] = _924_v127;
          _929_v131 = _nw148;
          let _index126 = _module.__default.safeIndex(new BigNumber(536), new BigNumber((_929_v131).length));
          (_929_v131)[_index126] = _924_v127;
          let _index127 = _module.__default.safeIndex(new BigNumber(536), new BigNumber((_929_v131).length));
          (_929_v131)[_index127] = _924_v127;
          let _rhs120 = _742_v0;
          let _rhs121 = new BigNumber(((((p0).isEqualTo(new BigNumber(439))) ? (_dafny.Seq.Concat(_885_cf22, _885_cf22)) : (_885_cf22))).length);
          let _lhs72 = globalState;
          _742_v0 = _rhs120;
          _lhs72.f2 = _rhs121;
        }
        let _930_v132;
        let _nw149 = Array((new BigNumber(9)).toNumber()).fill(_dafny.ZERO);
        _930_v132 = _nw149;
        let _931_v133;
        _931_v133 = _dafny.MultiSet.fromElements(p1, p0, p0, p1);
        let _932_v134;
        _932_v134 = _dafny.Map.Empty.slice().updateUnsafe(!(_742_v0),_742_v0);
        let _index128 = _module.__default.safeIndex(new BigNumber(50), new BigNumber((_930_v132).length));
        (_930_v132)[_index128] = (((_931_v133).contains(_module.__default.fm9(new BigNumber((_885_cf22).length), globalState))) ? ((_931_v133).get(_module.__default.fm9(new BigNumber((_885_cf22).length), globalState))) : ((_dafny.ZERO).minus(new BigNumber((_932_v134).length))));
        let _933_v135;
        _933_v135 = _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('b'.codePointAt(0)),_dafny.Seq.Concat(_885_cf22, _885_cf22));
        let _934_v136;
        _934_v136 = _dafny.Map.Empty.slice().updateUnsafe(p1,_930_v132);
        let _935_v138;
        _935_v138 = _dafny.Set.fromElements((_this).f6);
        let _index129 = _module.__default.safeIndex(new BigNumber(50), new BigNumber((_930_v132).length));
        let _rhs122 = ((!(true)) ? (p0) : (p0));
        let _rhs123 = ((_742_v0) ? (_933_v135) : ((function () {
          let _coll30 = new _dafny.Map();
          for (const _compr_30 of (_935_v138).Elements) {
            let _936_v137 = _compr_30;
            if ((_935_v138).contains(_936_v137)) {
              _coll30.push([_936_v137,_885_cf22]);
            }
          }
          return _coll30;
        }()).update(new _dafny.CodePoint('f'.codePointAt(0)), _885_cf22)));
        let _rhs124 = _module.__default.safeModuloInt((p1).plus(p1), p1);
        let _rhs125 = p1;
        let _rhs126 = _dafny.Map.Empty.slice().updateUnsafe(p1,_930_v132);
        let _lhs73 = _930_v132;
        let _lhs74 = _module.__default.safeIndex(new BigNumber(50), new BigNumber((_930_v132).length));
        let _lhs75 = globalState;
        let _lhs76 = globalState;
        _lhs73[_lhs74] = _rhs122;
        _933_v135 = _rhs123;
        _lhs75.f2 = _rhs124;
        _lhs76.f2 = _rhs125;
        _934_v136 = _rhs126;
      }
      let _937_v139;
      _937_v139 = _dafny.Seq.UnicodeFromString("f");
      let _938_v140;
      _938_v140 = _dafny.Seq.of(p0, p0);
      let _939_v141;
      _939_v141 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_938_v140).length),!(_742_v0));
      let _940_v142;
      let _nw150 = Array((new BigNumber(11)).toNumber());
      _nw150[(_dafny.ZERO).toNumber()] = _939_v141;
      _nw150[(_dafny.ONE).toNumber()] = _939_v141;
      _nw150[(new BigNumber(2)).toNumber()] = _939_v141;
      _nw150[(new BigNumber(3)).toNumber()] = _939_v141;
      _nw150[(new BigNumber(4)).toNumber()] = _939_v141;
      _nw150[(new BigNumber(5)).toNumber()] = _939_v141;
      _nw150[(new BigNumber(6)).toNumber()] = _939_v141;
      _nw150[(new BigNumber(7)).toNumber()] = _939_v141;
      _nw150[(new BigNumber(8)).toNumber()] = _939_v141;
      _nw150[(new BigNumber(9)).toNumber()] = _939_v141;
      _nw150[(new BigNumber(10)).toNumber()] = _939_v141;
      _940_v142 = _nw150;
      let _941_v143;
      let _nw151 = new _module.C2();
      _nw151.__ctor(_742_v0, new BigNumber((_937_v139).length), _940_v142);
      _941_v143 = _nw151;
      let _942_v144;
      _942_v144 = _dafny.Map.Empty.slice().updateUnsafe(p1,_941_v143);
      let _943_v145;
      _943_v145 = _module.D2.create_DC5(p1, false, _742_v0, _742_v0, _742_v0);
      let _source16 = _module.D2.create_DC5(new BigNumber((_942_v144).length), (_943_v145).dtor_cf8, _742_v0, _742_v0, !_dafny.Seq.contains(_937_v139, (_this).f6));
      if (_source16.is_DC4) {
        let _944___mcc_h9 = (_source16).cf6;
        let _945_cf6 = _944___mcc_h9;
        let _946_v146;
        let _nw152 = new _module.C2();
        _nw152.__ctor(false, new BigNumber((((!(_742_v0)) ? (_938_v140) : (_938_v140))).length), (_941_v143).f3);
        _946_v146 = _nw152;
        let _947_v147;
        let _nw153 = Array((new BigNumber(26)).toNumber());
        _nw153[(_dafny.ZERO).toNumber()] = (_946_v146).f11;
        _nw153[(_dafny.ONE).toNumber()] = (_946_v146).f11;
        _nw153[(new BigNumber(2)).toNumber()] = _742_v0;
        _nw153[(new BigNumber(3)).toNumber()] = _742_v0;
        _nw153[(new BigNumber(4)).toNumber()] = (_946_v146).f11;
        _nw153[(new BigNumber(5)).toNumber()] = (_946_v146).f11;
        _nw153[(new BigNumber(6)).toNumber()] = _module.__default.fm23(_module.__default.fm25(_module.__default.fm23(_937_v139, new BigNumber(924), (_946_v146).f12, _742_v0, globalState), p1, globalState), new BigNumber(379), _945_cf6, (_946_v146).f11, globalState);
        _nw153[(new BigNumber(7)).toNumber()] = !(_742_v0);
        _nw153[(new BigNumber(8)).toNumber()] = _742_v0;
        _nw153[(new BigNumber(9)).toNumber()] = _742_v0;
        _nw153[(new BigNumber(10)).toNumber()] = _742_v0;
        _nw153[(new BigNumber(11)).toNumber()] = _742_v0;
        _nw153[(new BigNumber(12)).toNumber()] = _742_v0;
        _nw153[(new BigNumber(13)).toNumber()] = false;
        _nw153[(new BigNumber(14)).toNumber()] = _742_v0;
        _nw153[(new BigNumber(15)).toNumber()] = (_946_v146).f11;
        _nw153[(new BigNumber(16)).toNumber()] = false;
        _nw153[(new BigNumber(17)).toNumber()] = true;
        _nw153[(new BigNumber(18)).toNumber()] = (_946_v146).f11;
        _nw153[(new BigNumber(19)).toNumber()] = !((_946_v146).f11);
        _nw153[(new BigNumber(20)).toNumber()] = _742_v0;
        _nw153[(new BigNumber(21)).toNumber()] = (_946_v146).f11;
        _nw153[(new BigNumber(22)).toNumber()] = (_946_v146).f11;
        _nw153[(new BigNumber(23)).toNumber()] = (_946_v146).f11;
        _nw153[(new BigNumber(24)).toNumber()] = (_946_v146).f11;
        _nw153[(new BigNumber(25)).toNumber()] = _742_v0;
        _947_v147 = _nw153;
        let _948_v148;
        _948_v148 = _dafny.MultiSet.fromElements(new BigNumber((_939_v141).length));
        let _949_v149;
        let _nw154 = new _module.C3();
        _nw154.__ctor(_948_v148, _742_v0);
        _949_v149 = _nw154;
        let _950_v150;
        _950_v150 = _dafny.Map.Empty.slice().updateUnsafe(_947_v147,_949_v149);
        _742_v0 = ((_742_v0) ? (!(_950_v150).contains(_947_v147)) : ((_949_v149).f10));
        let _951_v151;
        _951_v151 = _dafny.MultiSet.fromElements((_949_v149).f10, (_949_v149).f10, (_946_v146).f11);
        let _952_v152;
        _952_v152 = _dafny.Seq.of((_951_v151).IsDisjointFrom(_951_v151), (_949_v149).f10, true, (_946_v146).f11, (_949_v149).f10);
        if ((_952_v152)[_module.__default.safeIndex(p0, new BigNumber((_952_v152).length))]) {
          let _953_v154;
          _953_v154 = _dafny.Map.Empty.slice().updateUnsafe((_946_v146).f12,(_dafny.ZERO).minus((_946_v146).f12));
          let _954_v155;
          _954_v155 = _dafny.Set.fromElements(new BigNumber((function () {
            let _coll31 = new _dafny.Map();
            for (const _compr_31 of (_953_v154).Keys.Elements) {
              let _955_v153 = _compr_31;
              if ((_953_v154).contains(_955_v153)) {
                _coll31.push([_module.__default.safeModuloInt(_955_v153, p1),(_946_v146).f11]);
              }
            }
            return _coll31;
          }()).length));
          (globalState).f2 = (_dafny.ZERO).minus(_module.__default.fm24(((true) ? (_954_v155) : (_954_v155)), globalState));
          let _956_v156;
          _956_v156 = _dafny.MultiSet.fromElements(_947_v147, _947_v147);
          let _957_v157;
          _957_v157 = _dafny.Seq.of(_956_v156);
          let _958_v158;
          _958_v158 = _dafny.Seq.of(_947_v147, _947_v147, _947_v147);
          _956_v156 = ((_957_v157)[_module.__default.safeIndex(_945_cf6, new BigNumber((_957_v157).length))]).Difference((_dafny.MultiSet.fromElements(_947_v147, (_958_v158)[_module.__default.safeIndex(new BigNumber(-888), new BigNumber((_958_v158).length))], _947_v147)).Union(_956_v156));
          let _959_v159;
          _959_v159 = _dafny.Seq.of((_949_v149).f10);
          (globalState).f2 = new BigNumber(((_959_v159)).length);
          _945_cf6 = (_946_v146).f12;
          _937_v139 = _module.__default.fm8(new BigNumber(826), globalState);
        } else {
          (globalState).f2 = (_dafny.ZERO).minus((_946_v146).f12);
          let _960_v160;
          _960_v160 = _dafny.Map.Empty.slice().updateUnsafe((_949_v149).f10,_dafny.areEqual(_937_v139, (_813_v57).dtor_cf22));
          _742_v0 = (((_960_v160).contains((_946_v146).f11)) ? ((_960_v160).get((_946_v146).f11)) : (!(!(((((_939_v141).contains((_dafny.ZERO).minus((_946_v146).f12))) ? ((_939_v141).get((_dafny.ZERO).minus((_946_v146).f12))) : ((_949_v149).f10))) && ((_949_v149).f10)))));
          let _961_v161;
          let _nw155 = Array((new BigNumber(18)).toNumber()).fill(_dafny.ZERO);
          _961_v161 = _nw155;
          _961_v161 = _961_v161;
          _960_v160 = (_960_v160).update(_dafny.Seq.contains(_937_v139, (_this).f6), (_946_v146).f11);
          let _962_v162;
          _962_v162 = _dafny.Map.Empty.slice().updateUnsafe((_946_v146).f12,p0);
          let _963_v164;
          _963_v164 = _962_v162;
          let _964_v165;
          let _nw156 = Array((new BigNumber(5)).toNumber());
          _nw156[(_dafny.ZERO).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe((_946_v146).f12,new BigNumber((_937_v139).length))).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(57),p1));
          _nw156[(_dafny.ONE).toNumber()] = _962_v162;
          _nw156[(new BigNumber(2)).toNumber()] = function () {
            let _coll32 = new _dafny.Map();
            for (const _compr_32 of _dafny.IntegerRange(new BigNumber(572), new BigNumber(990))) {
              let _965_v163 = _compr_32;
              if (((new BigNumber(572)).isLessThanOrEqualTo(_965_v163)) && ((_965_v163).isLessThan(new BigNumber(990)))) {
                _coll32.push([(_965_v163).multipliedBy(p1),p0]);
              }
            }
            return _coll32;
          }();
          _nw156[(new BigNumber(3)).toNumber()] = (_962_v162).Merge((_963_v164));
          _nw156[(new BigNumber(4)).toNumber()] = _962_v162;
          _964_v165 = _nw156;
          let _index130 = _module.__default.safeIndex(new BigNumber(473), new BigNumber((_964_v165).length));
          (_964_v165)[_index130] = _962_v162;
          let _index131 = _module.__default.safeIndex(new BigNumber(473), new BigNumber((_964_v165).length));
          (_964_v165)[_index131] = (_962_v162).update(p1, (((_949_v149).f10) ? (new BigNumber((_952_v152).length)) : (new BigNumber(-935))));
        }
        if ((_949_v149).f10) {
          let _966_v166;
          let _967_v167;
          let _out35;
          let _out36;
          let _outcollector13 = (_946_v146).m1(_938_v140, ((_946_v146).f11) || ((_946_v146).f11), _937_v139, globalState);
          _out35 = _outcollector13[0];
          _out36 = _outcollector13[1];
          _966_v166 = _out35;
          _967_v167 = _out36;
          (globalState).f2 = (p0).minus((_946_v146).f12);
          let _968_v168;
          let _969_v169;
          let _970_v170;
          let _971_v171;
          let _out37;
          let _out38;
          let _out39;
          let _out40;
          let _outcollector14 = (_949_v149).m5(p1, (_949_v149).f10, globalState);
          _out37 = _outcollector14[0];
          _out38 = _outcollector14[1];
          _out39 = _outcollector14[2];
          _out40 = _outcollector14[3];
          _968_v168 = _out37;
          _969_v169 = _out38;
          _970_v170 = _out39;
          _971_v171 = _out40;
          let _972_v172;
          _972_v172 = _module.D9.create_DC19(_941_v143, (_dafny.ZERO).minus((((_948_v148).contains((_946_v146).f12)) ? ((_948_v148).get((_946_v146).f12)) : (new BigNumber((_951_v151).cardinality())))), _module.__default.fm9(new BigNumber((_938_v140).length), globalState));
          _972_v172 = _972_v172;
          let _973_v173;
          let _nw157 = Array((new BigNumber(2)).toNumber());
          _nw157[(_dafny.ZERO).toNumber()] = _951_v151;
          _nw157[(_dafny.ONE).toNumber()] = _951_v151;
          _973_v173 = _nw157;
          let _974_v174;
          let _nw158 = new _module.C0();
          _nw158.__ctor(p1);
          _974_v174 = _nw158;
          let _975_v175;
          _975_v175 = _dafny.Set.fromElements((_946_v146).f11);
          let _rhs127 = ((_971_v171) ? (_967_v167) : (_970_v170));
          let _rhs128 = ((_975_v175).Union(_975_v175)).equals(_975_v175);
          let _rhs129 = _973_v173;
          let _rhs130 = _974_v174;
          _967_v167 = _rhs127;
          _968_v168 = _rhs128;
          _973_v173 = _rhs129;
          _974_v174 = _rhs130;
        } else {
          let _976_v176;
          _976_v176 = _dafny.Set.fromElements((_946_v146).f12, (_938_v140)[_module.__default.safeIndex((_dafny.ZERO).minus((_dafny.ZERO).minus(p1)), new BigNumber((_938_v140).length))]);
          let _977_v177;
          _977_v177 = _dafny.Map.Empty.slice().updateUnsafe((_946_v146).f11,_976_v176);
          let _978_v178;
          _978_v178 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.areEqual(_938_v140, _938_v140),(_module.__default.fm38(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_949_v149).f10,(_952_v152)[_module.__default.safeIndex(new BigNumber(492), new BigNumber((_952_v152).length))])).length), (_952_v152)[_module.__default.safeIndex(p1, new BigNumber((_952_v152).length))], (_949_v149).f10, (_949_v149).f10, globalState)).Union((((_977_v177).contains((_949_v149).f10)) ? ((_977_v177).get((_949_v149).f10)) : (_dafny.Set.fromElements((_946_v146).f12, p1)))));
          _977_v177 = (_977_v177).update(!(!(_976_v176).contains(_945_cf6)), ((false) ? (function () {
            let _coll33 = new _dafny.Set();
            for (const _compr_33 of _dafny.IntegerRange(new BigNumber(121), new BigNumber(74))) {
              let _979_v179 = _compr_33;
              if (((new BigNumber(121)).isLessThanOrEqualTo(_979_v179)) && ((_979_v179).isLessThan(new BigNumber(74)))) {
                _coll33.add(_module.__default.safeModuloInt(_979_v179, new BigNumber((function () {
                  let _coll34 = new _dafny.Map();
                  for (const _compr_34 of _dafny.IntegerRange(new BigNumber(124), new BigNumber(935))) {
                    let _980_v180 = _compr_34;
                    if (((new BigNumber(124)).isLessThanOrEqualTo(_980_v180)) && ((_980_v180).isLessThan(new BigNumber(935)))) {
                      _coll34.push([(_980_v180).minus(_945_cf6),p1]);
                    }
                  }
                  return _coll34;
                }()).length)));
              }
            }
            return _coll33;
          }()) : (_976_v176)));
          let _981_v181;
          _981_v181 = _dafny.Set.fromElements((_this).f6);
          let _982_v182;
          _982_v182 = _dafny.Set.fromElements((_949_v149).f10, (_981_v181).IsDisjointFrom(_981_v181), false);
          _982_v182 = _982_v182;
          (globalState).f2 = _945_cf6;
          _976_v176 = _976_v176;
          let _983_v183;
          _983_v183 = _982_v182;
          _983_v183 = (_dafny.Set.fromElements((_949_v149).f10)).Difference(_982_v182);
        }
      } else if (_source16.is_DC5) {
        let _984___mcc_h10 = (_source16).cf7;
        let _985___mcc_h11 = (_source16).cf8;
        let _986___mcc_h12 = (_source16).cf9;
        let _987___mcc_h13 = (_source16).cf10;
        let _988___mcc_h14 = (_source16).cf11;
        let _989_cf11 = _988___mcc_h14;
        let _990_cf10 = _987___mcc_h13;
        let _991_cf9 = _986___mcc_h12;
        let _992_cf8 = _985___mcc_h11;
        let _993_cf7 = _984___mcc_h10;
        let _994_v184;
        let _nw159 = Array((new BigNumber(16)).toNumber()).fill(_dafny.MultiSet.Empty);
        _994_v184 = _nw159;
        let _995_v185;
        let _nw160 = Array((new BigNumber(11)).toNumber()).fill(_dafny.ZERO);
        _995_v185 = _nw160;
        let _index132 = _module.__default.safeIndex(new BigNumber(389), new BigNumber((_994_v184).length));
        (_994_v184)[_index132] = (_dafny.MultiSet.fromElements(_995_v185, _995_v185, _995_v185)).update(_995_v185, _module.__default.abs(p1));
        let _996_v186;
        _996_v186 = _dafny.MultiSet.fromElements(_995_v185, _995_v185, _995_v185, _995_v185);
        let _index133 = _module.__default.safeIndex(new BigNumber(389), new BigNumber((_994_v184).length));
        (_994_v184)[_index133] = ((_992_cf8) ? ((_996_v186).Difference(_996_v186)) : ((_996_v186).Union(_996_v186)));
        let _997_v187;
        _997_v187 = _dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("polxqwadi"));
        _997_v187 = ((_997_v187).Union(_997_v187)).update(_dafny.Seq.UnicodeFromString("uiiokhs"), _module.__default.abs(new BigNumber(-980)));
        _937_v139 = (_813_v57).dtor_cf22;
        let _998_v188;
        let _init24 = ((_999_cf8) => function (_1000_i17) {
          return _dafny.Set.fromElements(_999_cf8);
        })(_992_cf8);
        let _nw161 = Array((new BigNumber(5)).toNumber());
        for (let _i0_24 = 0; _i0_24 < new BigNumber(_nw161.length); _i0_24++) {
          _nw161[_i0_24] = _init24(new BigNumber(_i0_24));
        }
        _998_v188 = _nw161;
        let _index134 = _module.__default.safeIndex(new BigNumber(560), new BigNumber((_998_v188).length));
        (_998_v188)[_index134] = _module.__default.fm16(p0, new _dafny.CodePoint('l'.codePointAt(0)), _993_cf7, globalState);
        let _1001_v189;
        _1001_v189 = _dafny.Set.fromElements(p1);
        let _1002_v190;
        _1002_v190 = _dafny.Seq.of(!(false));
        let _1003_v191;
        _1003_v191 = _dafny.Set.fromElements((_1001_v189).IsSubsetOf(_dafny.Set.fromElements(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(544)), ((_1004_p1) => function (_1005_i18) {
          return _1004_p1;
        })(p1))).length))), _990_cf10, (_1002_v190)[_module.__default.safeIndex(p0, new BigNumber((_1002_v190).length))]);
        let _index135 = _module.__default.safeIndex(new BigNumber(560), new BigNumber((_998_v188).length));
        (_998_v188)[_index135] = _1003_v191;
      } else {
        let _1006___mcc_h15 = (_source16).cf5;
        let _1007_cf5 = _1006___mcc_h15;
        let _1008_v192;
        let _nw162 = new _module.C2();
        _nw162.__ctor(_742_v0, _module.__default.safeDivisionInt(p0, new BigNumber(130)), _940_v142);
        _1008_v192 = _nw162;
        if ((_1008_v192).f11) {
          _742_v0 = _742_v0;
          let _1009_v193;
          let _nw163 = Array((new BigNumber(20)).toNumber());
          _nw163[(_dafny.ZERO).toNumber()] = false;
          _nw163[(_dafny.ONE).toNumber()] = (_1008_v192).f11;
          _nw163[(new BigNumber(2)).toNumber()] = _742_v0;
          _nw163[(new BigNumber(3)).toNumber()] = _742_v0;
          _nw163[(new BigNumber(4)).toNumber()] = (_1008_v192).f11;
          _nw163[(new BigNumber(5)).toNumber()] = (_1008_v192).f11;
          _nw163[(new BigNumber(6)).toNumber()] = (_1008_v192).f11;
          _nw163[(new BigNumber(7)).toNumber()] = (_1008_v192).f11;
          _nw163[(new BigNumber(8)).toNumber()] = (_1008_v192).f11;
          _nw163[(new BigNumber(9)).toNumber()] = (_1008_v192).f11;
          _nw163[(new BigNumber(10)).toNumber()] = false;
          _nw163[(new BigNumber(11)).toNumber()] = (_1008_v192).f11;
          _nw163[(new BigNumber(12)).toNumber()] = (_1008_v192).f11;
          _nw163[(new BigNumber(13)).toNumber()] = (_1008_v192).f11;
          _nw163[(new BigNumber(14)).toNumber()] = true;
          _nw163[(new BigNumber(15)).toNumber()] = _742_v0;
          _nw163[(new BigNumber(16)).toNumber()] = _742_v0;
          _nw163[(new BigNumber(17)).toNumber()] = _742_v0;
          _nw163[(new BigNumber(18)).toNumber()] = (_1008_v192).f11;
          _nw163[(new BigNumber(19)).toNumber()] = (_1008_v192).f11;
          _1009_v193 = _nw163;
          let _1010_v194;
          _1010_v194 = _dafny.Map.Empty.slice().updateUnsafe(_module.D8.create_DC15(_1009_v193),_module.D9.create_DC19(_941_v143, p0, (_1008_v192).f12));
          _939_v141 = (_939_v141).update((_dafny.ZERO).minus((new BigNumber((_1010_v194).length)).plus((_1008_v192).f12)), (_1008_v192).f11);
          let _1011_v195;
          _1011_v195 = _module.D8.create_DC15(_1009_v193);
          let _pat_let_tv24 = _1009_v193;
          _1011_v195 = function (_pat_let11_0) {
            return function (_1012_dt__update__tmp_h3) {
              return function (_pat_let12_0) {
                return function (_1013_dt__update_hcf27_h0) {
                  return _module.D8.create_DC15(_1013_dt__update_hcf27_h0);
                }(_pat_let12_0);
              }(_pat_let_tv24);
            }(_pat_let11_0);
          }(_1011_v195);
          _742_v0 = (p0).isLessThanOrEqualTo((_1008_v192).f12);
          _742_v0 = (((new BigNumber((_937_v139).length)).isLessThanOrEqualTo((_dafny.ZERO).minus((_938_v140)[_module.__default.safeIndex(p0, new BigNumber((_938_v140).length))]))) ? ((_1008_v192).f11) : (((_1008_v192).f11) === (_module.__default.fm23(_937_v139, new BigNumber((_dafny.MultiSet.fromElements((_1008_v192).f12)).cardinality()), p1, (_1008_v192).f11, globalState))));
        } else {
          let _rhs131 = !(!(!((_1008_v192).f11)));
          let _rhs132 = _937_v139;
          _742_v0 = _rhs131;
          _937_v139 = _rhs132;
          (globalState).f2 = (new BigNumber(((_dafny.Set.fromElements(new BigNumber(36), new BigNumber((_937_v139).length))).Intersect(function () {
            let _coll35 = new _dafny.Set();
            for (const _compr_35 of (_dafny.Seq.of(p1)).Elements) {
              let _1014_v196 = _compr_35;
              if (_dafny.Seq.contains(_dafny.Seq.of(p1), _1014_v196)) {
                _coll35.add((_1014_v196).multipliedBy(new BigNumber(489)));
              }
            }
            return _coll35;
          }())).length)).minus(p1);
          _742_v0 = _742_v0;
          let _1015_v197;
          _1015_v197 = _dafny.Map.Empty.slice().updateUnsafe((_1008_v192).f12,new BigNumber((_939_v141).length));
          (globalState).f2 = _module.__default.safeDivisionInt((((_1015_v197).contains((_1008_v192).f12)) ? ((_1015_v197).get((_1008_v192).f12)) : ((_1008_v192).f12)), p1);
          let _1016_v198;
          _1016_v198 = new _dafny.CodePoint('e'.codePointAt(0));
          _1016_v198 = (_this).f6;
        }
        (globalState).f2 = (_dafny.ZERO).minus(((_1008_v192).f12).plus(new BigNumber((_dafny.Set.fromElements(false)).length)));
        (globalState).f2 = ((p1).plus(p1)).multipliedBy(((_dafny.ZERO).minus(new BigNumber(-916))).minus(p0));
      }
      let _1017_v199;
      let _nw164 = Array((new BigNumber(7)).toNumber()).fill(false);
      _1017_v199 = _nw164;
      let _index136 = _module.__default.safeIndex(new BigNumber(840), new BigNumber((_1017_v199).length));
      (_1017_v199)[_index136] = _dafny.Seq.IsProperPrefixOf(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-512)), function (_1018_i19) {
        return (_this).f6;
      }), _937_v139);
      let _1019_v200;
      _1019_v200 = _dafny.Map.Empty.slice().updateUnsafe(p1,p0);
      let _1020_v201;
      _1020_v201 = _module.D3.create_DC6((_this).f6);
      let _1021_v202;
      let _nw165 = Array((new BigNumber(11)).toNumber()).fill(_dafny.ZERO);
      _1021_v202 = _nw165;
      let _index137 = _module.__default.safeIndex(new BigNumber(52), new BigNumber((_1021_v202).length));
      (_1021_v202)[_index137] = p1;
      let _index138 = _module.__default.safeIndex(new BigNumber(840), new BigNumber((_1017_v199).length));
      let _index139 = _module.__default.safeIndex(new BigNumber(52), new BigNumber((_1021_v202).length));
      let _rhs133 = _742_v0;
      let _rhs134 = _1019_v200;
      let _rhs135 = ((p0).minus(p1)).minus(((_dafny.ZERO).minus(p1)).minus(new BigNumber(468)));
      let _rhs136 = _1020_v201;
      let _rhs137 = p1;
      let _lhs77 = _1017_v199;
      let _lhs78 = _module.__default.safeIndex(new BigNumber(840), new BigNumber((_1017_v199).length));
      let _lhs79 = globalState;
      let _lhs80 = _1021_v202;
      let _lhs81 = _module.__default.safeIndex(new BigNumber(52), new BigNumber((_1021_v202).length));
      _lhs77[_lhs78] = _rhs133;
      _1019_v200 = _rhs134;
      _lhs79.f2 = _rhs135;
      _1020_v201 = _rhs136;
      _lhs80[_lhs81] = _rhs137;
      let _index140 = _module.__default.safeIndex(new BigNumber(840), new BigNumber((_1017_v199).length));
      (_1017_v199)[_index140] = !(true);
      let _1022_v203;
      _1022_v203 = _dafny.MultiSet.fromElements(p0);
      let _1023_v204;
      let _nw166 = new _module.C3();
      _nw166.__ctor(_1022_v203, _742_v0);
      _1023_v204 = _nw166;
      let _1024_v205;
      _1024_v205 = _dafny.Map.Empty.slice().updateUnsafe(_1023_v204,(_this).f6);
      let _index141 = _module.__default.safeIndex(new BigNumber(840), new BigNumber((_1017_v199).length));
      (_1017_v199)[_index141] = !(_1024_v205).equals(_dafny.Map.Empty.slice().updateUnsafe(_1023_v204,(_module.D3.create_DC6((_this).f6)).dtor_cf12));
      return;
    }
    get f6() {
      let _this = this;
      return _this._f6;
    };
  };

  $module.C6 = class C6 {
    constructor () {
      this._tname = "_module.C6";
      this.f5 = _dafny.ZERO;
      this._f4 = _dafny.Map.Empty;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f4, f5) {
      let _this = this;
      (_this)._f4 = f4;
      (_this).f5 = f5;
      return;
    }
    m2(p0, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = false;
      let _1025_i0;
      _1025_i0 = _dafny.ZERO;
      L5: {
        while (p0) {
          C5: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1025_i0)) {
              break L5;
            }
            _1025_i0 = (_1025_i0).plus(_dafny.ONE);
            let _1026_v0;
            let _init25 = ((_1027_p0) => function (_1028_i1) {
              return _module.D1.create_DC2(_this.f5, _1027_p0, _this.f5);
            })(p0);
            let _nw167 = Array((new BigNumber(10)).toNumber());
            for (let _i0_25 = 0; _i0_25 < new BigNumber(_nw167.length); _i0_25++) {
              _nw167[_i0_25] = _init25(new BigNumber(_i0_25));
            }
            _1026_v0 = _nw167;
            let _1029_v1;
            _1029_v1 = _module.D1.create_DC2(_this.f5, p0, _this.f5);
            let _index142 = _module.__default.safeIndex(new BigNumber(620), new BigNumber((_1026_v0).length));
            (_1026_v0)[_index142] = _1029_v1;
            let _index143 = _module.__default.safeIndex(new BigNumber(620), new BigNumber((_1026_v0).length));
            (_1026_v0)[_index143] = _1029_v1;
            let _1030_v2;
            _1030_v2 = new _dafny.CodePoint('x'.codePointAt(0));
            let _1031_v3;
            let _nw168 = new _module.C5();
            _nw168.__ctor(_1030_v2);
            _1031_v3 = _nw168;
            let _1032_v4;
            _1032_v4 = _dafny.Map.Empty.slice().updateUnsafe(p0,false);
            let _1033_v5;
            _1033_v5 = _dafny.Seq.of((_dafny.ZERO).minus(_this.f5), _this.f5);
            let _1034_v6;
            _1034_v6 = _dafny.Map.Empty.slice().updateUnsafe(_1032_v4,_1033_v5);
            _1034_v6 = (_1034_v6).update(_module.__default.fm39(false, globalState), _1033_v5);
            let _1035_v7;
            let _nw169 = Array((new BigNumber(21)).toNumber()).fill(false);
            _1035_v7 = _nw169;
            let _1036_v8;
            _1036_v8 = _dafny.Set.fromElements(_1035_v7, _1035_v7, _1035_v7, _1035_v7, _1035_v7);
            let _1037_v9;
            _1037_v9 = _dafny.Seq.UnicodeFromString("eueyat");
            let _1038_v10;
            _1038_v10 = _module.D6.create_DC11(_dafny.Seq.UnicodeFromString("kuxdxnku"));
            let _1039_v11;
            _1039_v11 = _dafny.Map.Empty.slice().updateUnsafe(_1036_v8,new BigNumber((_dafny.Seq.Concat(_1037_v9, (_1038_v10).dtor_cf22)).length));
            let _1040_v12;
            _1040_v12 = _dafny.Seq.of(_1031_v3, _1031_v3);
            let _1041_v13;
            _1041_v13 = _dafny.MultiSet.fromElements(new BigNumber((_1040_v12).length), _this.f5);
            let _1042_v14;
            _1042_v14 = _dafny.Seq.of(p0);
            _1039_v11 = (_1039_v11).update((_1036_v8).Difference(_1036_v8), ((p0) ? (new BigNumber((_1041_v13).cardinality())) : (new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_1033_v5)[_module.__default.safeIndex(new BigNumber((_1042_v14).length), new BigNumber((_1033_v5).length))],_1032_v4)).length))));
          }
        }
      }
      let _1043_v15;
      _1043_v15 = _dafny.MultiSet.fromElements(true, !(p0), p0, p0);
      let _1044_v16;
      _1044_v16 = new _dafny.CodePoint('x'.codePointAt(0));
      let _1045_v17;
      _1045_v17 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(884),_module.__default.fm21(_this.f5, _1044_v16, p0, false, globalState));
      let _1046_v18;
      _1046_v18 = _dafny.Seq.UnicodeFromString("tbwc");
      let _1047_v19;
      _1047_v19 = _dafny.MultiSet.fromElements((new BigNumber((_1043_v15).cardinality())).multipliedBy(_this.f5), _this.f5, _module.__default.safeModuloInt(new BigNumber((_1045_v17).length), new BigNumber((_1046_v18).length)));
      let _1048_v20;
      let _nw170 = Array((new BigNumber(19)).toNumber()).fill(_module.D2.Default());
      _1048_v20 = _nw170;
      let _1049_v21;
      _1049_v21 = _dafny.Seq.of(_1048_v20);
      let _rhs138 = (_1047_v19).Intersect(_1047_v19);
      let _rhs139 = _this.f5;
      let _rhs140 = _dafny.Seq.of(_1048_v20);
      let _lhs82 = _this;
      _1047_v19 = _rhs138;
      _lhs82.f5 = _rhs139;
      _1049_v21 = _rhs140;
      let _1050_v22;
      let _nw171 = Array((new BigNumber(12)).toNumber()).fill([]);
      _1050_v22 = _nw171;
      let _1051_v23;
      let _nw172 = Array((new BigNumber(28)).toNumber()).fill(false);
      _1051_v23 = _nw172;
      let _index144 = _module.__default.safeIndex(new BigNumber(678), new BigNumber((_1050_v22).length));
      (_1050_v22)[_index144] = _1051_v23;
      let _index145 = _module.__default.safeIndex(new BigNumber(678), new BigNumber((_1050_v22).length));
      (_1050_v22)[_index145] = _1051_v23;
      let _hi4 = _this.f5;
      for (let _1052_i2 = ((_dafny.ZERO).minus(_this.f5)).minus(_this.f5); _1052_i2.isLessThan(_hi4); _1052_i2 = _1052_i2.plus(_dafny.ONE)) {
        r1 = p0;
        _1043_v15 = _1043_v15;
        let _1053_v24;
        let _init26 = ((_1054_v19) => function (_1055_i3) {
          return (_1054_v19).Difference(_1054_v19);
        })(_1047_v19);
        let _nw173 = Array((new BigNumber(2)).toNumber());
        for (let _i0_26 = 0; _i0_26 < new BigNumber(_nw173.length); _i0_26++) {
          _nw173[_i0_26] = _init26(new BigNumber(_i0_26));
        }
        _1053_v24 = _nw173;
        _1053_v24 = _1053_v24;
        let _1056_v25;
        _1056_v25 = _dafny.Seq.of(new BigNumber(((_this).f4).length), _1052_i2, _1052_i2, _this.f5, new BigNumber((_1046_v18).length));
        let _1057_v26;
        _1057_v26 = _dafny.Set.fromElements(false, p0);
        let _1058_v27;
        _1058_v27 = _dafny.Map.Empty.slice().updateUnsafe(_this.f5,_this.f5);
        let _1059_v28;
        let _nw174 = Array((new BigNumber(21)).toNumber());
        _nw174[(_dafny.ZERO).toNumber()] = _dafny.Seq.Concat(_1056_v25, _module.__default.fm12(p0, new BigNumber((_1056_v25).length), _1057_v26, _1058_v27, globalState));
        _nw174[(_dafny.ONE).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(699)), ((_1060_i2) => function (_1061_i4) {
          return _1060_i2;
        })(_1052_i2));
        _nw174[(new BigNumber(2)).toNumber()] = _module.__default.fm12(p0, new BigNumber((_1046_v18).length), _1057_v26, _1058_v27, globalState);
        _nw174[(new BigNumber(3)).toNumber()] = _dafny.Seq.Concat(_1056_v25, _1056_v25);
        _nw174[(new BigNumber(4)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(240)), ((_1062_i2) => function (_1063_i5) {
          return _1062_i2;
        })(_1052_i2));
        _nw174[(new BigNumber(5)).toNumber()] = _dafny.Seq.of(_this.f5);
        _nw174[(new BigNumber(6)).toNumber()] = _1056_v25;
        _nw174[(new BigNumber(7)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.update(_dafny.Seq.of(_this.f5, _1052_i2), _module.__default.safeIndex(_1052_i2, new BigNumber((_dafny.Seq.of(_this.f5, _1052_i2)).length)), _this.f5), _1056_v25);
        _nw174[(new BigNumber(8)).toNumber()] = _1056_v25;
        _nw174[(new BigNumber(9)).toNumber()] = _1056_v25;
        _nw174[(new BigNumber(10)).toNumber()] = _1056_v25;
        _nw174[(new BigNumber(11)).toNumber()] = _1056_v25;
        _nw174[(new BigNumber(12)).toNumber()] = _1056_v25;
        _nw174[(new BigNumber(13)).toNumber()] = _1056_v25;
        _nw174[(new BigNumber(14)).toNumber()] = _1056_v25;
        _nw174[(new BigNumber(15)).toNumber()] = _dafny.Seq.of(_1052_i2);
        _nw174[(new BigNumber(16)).toNumber()] = _dafny.Seq.of(new BigNumber((_1056_v25).length), _1052_i2, _1052_i2, _1052_i2);
        _nw174[(new BigNumber(17)).toNumber()] = _1056_v25;
        _nw174[(new BigNumber(18)).toNumber()] = _1056_v25;
        _nw174[(new BigNumber(19)).toNumber()] = _dafny.Seq.Concat(_1056_v25, _dafny.Seq.of(new BigNumber(729)));
        _nw174[(new BigNumber(20)).toNumber()] = _dafny.Seq.Concat(_1056_v25, _1056_v25);
        _1059_v28 = _nw174;
        let _1064_v29;
        _1064_v29 = _dafny.Map.Empty.slice().updateUnsafe((_1056_v25)[_module.__default.safeIndex(_1052_i2, new BigNumber((_1056_v25).length))],_1059_v28);
        _1059_v28 = (((_1064_v29).contains(_this.f5)) ? ((_1064_v29).get(_this.f5)) : (_1059_v28));
      }
      let _1065_i6;
      _1065_i6 = _dafny.ZERO;
      L6: {
        while (!(p0)) {
          C6: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1065_i6)) {
              break L6;
            }
            _1065_i6 = (_1065_i6).plus(_dafny.ONE);
            let _1066_v30;
            _1066_v30 = _dafny.Map.Empty.slice().updateUnsafe(p0,_1051_v23);
            let _1067_v31;
            let _nw175 = Array((new BigNumber(3)).toNumber()).fill(_dafny.Map.Empty);
            _1067_v31 = _nw175;
            let _1068_v32;
            let _nw176 = new _module.C2();
            _nw176.__ctor(_module.__default.fm21(_this.f5, _1044_v16, p0, p0, globalState), _this.f5, _1067_v31);
            _1068_v32 = _nw176;
            let _1069_v33;
            _1069_v33 = _dafny.Seq.of(new BigNumber(((_1066_v30).update(p0, (_1050_v22)[_module.__default.safeIndex(new BigNumber(678), new BigNumber((_1050_v22).length))])).length), new BigNumber((_dafny.Seq.of(_1068_v32)).length), _this.f5, (_1068_v32).f12);
            let _1070_v34;
            _1070_v34 = _dafny.Seq.of(_dafny.MultiSet.FromArray(_1069_v33));
            let _1071_v35;
            let _nw177 = new _module.C3();
            _nw177.__ctor((_dafny.MultiSet.fromElements(new BigNumber((_1047_v19).cardinality()), _this.f5)).Union((_1070_v34)[_module.__default.safeIndex((_1068_v32).f12, new BigNumber((_1070_v34).length))]), p0);
            _1071_v35 = _nw177;
            let _1072_v36;
            _1072_v36 = _dafny.Seq.of((_1068_v32).f11);
            let _1073_v37;
            _1073_v37 = _module.D4.create_DC9(new BigNumber((_dafny.MultiSet.FromArray(_1072_v36)).cardinality()), _1044_v16, _dafny.Seq.of(new BigNumber((_dafny.Seq.of((_1068_v32).f12, _this.f5)).length)), _this.f5, p0);
            _1044_v16 = (_1073_v37).dtor_cf17;
            let _1074_v38;
            _1074_v38 = _dafny.Seq.of(_module.__default.fm40((_1071_v35).f10, (_1068_v32).f12, _1044_v16, (_1068_v32).f11, globalState), (_dafny.MultiSet.fromElements((_1068_v32).f11)).update((_1068_v32).f11, _module.__default.abs(_this.f5)), _1043_v15);
            let _1075_v39;
            _1075_v39 = _dafny.Set.fromElements((_1068_v32).f12);
            let _1076_v40;
            _1076_v40 = _dafny.Seq.of(_1075_v39, _1075_v39, _1075_v39, _1075_v39);
            let _1077_v41;
            let _nw178 = Array((new BigNumber(5)).toNumber());
            _nw178[(_dafny.ZERO).toNumber()] = (_dafny.ZERO).minus((_1068_v32).f12);
            _nw178[(_dafny.ONE).toNumber()] = new BigNumber(((_1043_v15).Intersect((_1074_v38)[_module.__default.safeIndex(new BigNumber(308), new BigNumber((_1074_v38).length))])).cardinality());
            _nw178[(new BigNumber(2)).toNumber()] = new BigNumber((_1076_v40).length);
            _nw178[(new BigNumber(3)).toNumber()] = _this.f5;
            _nw178[(new BigNumber(4)).toNumber()] = (_1068_v32).f12;
            _1077_v41 = _nw178;
            let _index146 = _module.__default.safeIndex(new BigNumber(999), new BigNumber((_1077_v41).length));
            (_1077_v41)[_index146] = _this.f5;
            let _1078_v42;
            _1078_v42 = _dafny.Set.fromElements(p0);
            let _index147 = _module.__default.safeIndex(new BigNumber(999), new BigNumber((_1077_v41).length));
            let _rhs141 = (_1068_v32).f12;
            let _rhs142 = new BigNumber((((_1078_v42).Intersect(_1078_v42)).Difference(_dafny.Set.fromElements(true, false))).length);
            let _rhs143 = new BigNumber((_dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(748)), ((_1079_v32, _1080_v17) => function (_1081_i7) {
              return ((_1079_v32).f12).multipliedBy(new BigNumber((_1080_v17).length));
            })(_1068_v32, _1045_v17)), _module.__default.safeIndex(_this.f5, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(748)), ((_1082_v32, _1083_v17) => function (_1084_i7) {
              return ((_1082_v32).f12).multipliedBy(new BigNumber((_1083_v17).length));
            })(_1068_v32, _1045_v17))).length)), (_dafny.ZERO).minus(((_dafny.ZERO).minus(_this.f5)).multipliedBy((_1068_v32).f12)))).length);
            let _rhs144 = (_1071_v35).f10;
            let _lhs83 = globalState;
            let _lhs84 = _this;
            let _lhs85 = _1077_v41;
            let _lhs86 = _module.__default.safeIndex(new BigNumber(999), new BigNumber((_1077_v41).length));
            _lhs83.f2 = _rhs141;
            _lhs84.f5 = _rhs142;
            _lhs85[_lhs86] = _rhs143;
            r1 = _rhs144;
            let _index148 = _module.__default.safeIndex(new BigNumber(848), new BigNumber((_1051_v23).length));
            (_1051_v23)[_index148] = !(_this.f5).isEqualTo((_1077_v41)[_module.__default.safeIndex(new BigNumber(999), new BigNumber((_1077_v41).length))]);
            let _index149 = _module.__default.safeIndex(new BigNumber(848), new BigNumber((_1051_v23).length));
            (_1051_v23)[_index149] = (((_this.f5).isLessThan((_1068_v32).f12)) ? ((_1068_v32).f11) : (_dafny.Seq.contains(_1046_v18, _1044_v16)));
          }
        }
      }
      let _1085_v43;
      _1085_v43 = _dafny.Seq.of((_this.f5).plus(_this.f5));
      let _1086_v44;
      _1086_v44 = _dafny.Map.Empty.slice().updateUnsafe(_this.f5,_1085_v43);
      _1085_v43 = _dafny.Seq.Concat(_dafny.Seq.Concat(_1085_v43, (((_1086_v44).contains(_this.f5)) ? ((_1086_v44).get(_this.f5)) : (_1085_v43))), _1085_v43);
      let _1087_v45;
      _1087_v45 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1046_v18).length),_1044_v16);
      let _1088_v46;
      _1088_v46 = _dafny.MultiSet.fromElements(_1087_v45);
      r0 = _module.__default.fm9(new BigNumber(((_1088_v46).update(_1087_v45, _module.__default.abs(_this.f5))).cardinality()), globalState);
      r1 = (new BigNumber((_dafny.Seq.Concat(_1046_v18, _1046_v18)).length)).isLessThanOrEqualTo(_this.f5);
      return [r0, r1];
    }
    m3(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = _dafny.ZERO;
      let r2 = _dafny.Set.Empty;
      let _1089_v0;
      let _nw179 = Array((new BigNumber(18)).toNumber()).fill(false);
      _1089_v0 = _nw179;
      let _1090_v1;
      _1090_v1 = _module.D8.create_DC15(_1089_v0);
      let _source17 = _1090_v1;
      if (_source17.is_DC16) {
        let _1091___mcc_h0 = (_source17).cf28;
        let _1092___mcc_h1 = (_source17).cf29;
        let _1093___mcc_h2 = (_source17).cf30;
        let _1094___mcc_h3 = (_source17).cf31;
        let _1095_cf31 = _1094___mcc_h3;
        let _1096_cf30 = _1093___mcc_h2;
        let _1097_cf29 = _1092___mcc_h1;
        let _1098_cf28 = _1091___mcc_h0;
        let _1099_v2;
        let _nw180 = Array((new BigNumber(28)).toNumber()).fill(_dafny.ZERO);
        _1099_v2 = _nw180;
        let _1100_v3;
        _1100_v3 = _dafny.Set.fromElements(_1097_cf29);
        let _index150 = _module.__default.safeIndex(new BigNumber(584), new BigNumber((_1099_v2).length));
        (_1099_v2)[_index150] = new BigNumber((((_1096_cf30) ? (_1100_v3) : (_1100_v3))).length);
        let _index151 = _module.__default.safeIndex(new BigNumber(584), new BigNumber((_1099_v2).length));
        (_1099_v2)[_index151] = _1095_cf31;
        _1098_cf28 = _1098_cf28;
        let _1101_v4;
        _1101_v4 = new _dafny.CodePoint('x'.codePointAt(0));
        let _1102_v5;
        _1102_v5 = _module.D3.create_DC6(_1101_v4);
        let _1103_v6;
        let _nw181 = new _module.C5();
        _nw181.__ctor((_1102_v5).dtor_cf12);
        _1103_v6 = _nw181;
        let _1104_v7;
        _1104_v7 = _dafny.MultiSet.fromElements((_1099_v2)[_module.__default.safeIndex(new BigNumber(584), new BigNumber((_1099_v2).length))], p1);
        let _index152 = _module.__default.safeIndex(new BigNumber(584), new BigNumber((_1099_v2).length));
        (_1099_v2)[_index152] = _module.__default.safeModuloInt(p1, (_this.f5).multipliedBy((((_1104_v7).contains(_this.f5)) ? ((_1104_v7).get(_this.f5)) : (new BigNumber((_dafny.Seq.UnicodeFromString("m")).length)))));
      } else if (_source17.is_DC15) {
        let _1105___mcc_h4 = (_source17).cf27;
        let _1106_cf27 = _1105___mcc_h4;
        let _1107_v8;
        _1107_v8 = true;
        _1107_v8 = (_module.__default.fm41(_this.f5, _this.f5, globalState)).dtor_cf30;
        let _1108_v9;
        _1108_v9 = _dafny.Map.Empty.slice().updateUnsafe(p3,p2);
        let _1109_v10;
        _1109_v10 = _dafny.Map.Empty.slice().updateUnsafe(p1,p3);
        let _1110_v11;
        let _init27 = ((_1111_v9) => function (_1112_i0) {
          return _1111_v9;
        })(_1108_v9);
        let _nw182 = Array((new BigNumber(26)).toNumber());
        for (let _i0_27 = 0; _i0_27 < new BigNumber(_nw182.length); _i0_27++) {
          _nw182[_i0_27] = _init27(new BigNumber(_i0_27));
        }
        _1110_v11 = _nw182;
        let _1113_v12;
        let _nw183 = new _module.C2();
        _nw183.__ctor((((_1108_v9).contains(((((_this).f4).contains(!(_1107_v8))) ? (((_this).f4).get(!(_1107_v8))) : (p1)))) ? ((_1108_v9).get(((((_this).f4).contains(!(_1107_v8))) ? (((_this).f4).get(!(_1107_v8))) : (p1)))) : (p2)), new BigNumber((_1109_v10).length), _1110_v11);
        _1113_v12 = _nw183;
        let _1114_v13;
        _1114_v13 = new _dafny.CodePoint('l'.codePointAt(0));
        let _1115_v14;
        _1115_v14 = _dafny.Seq.of(_1114_v13);
        let _1116_v15;
        _1116_v15 = _dafny.Seq.of(_dafny.Seq.of(_1114_v13), _1115_v14);
        (globalState).f2 = (_dafny.ZERO).minus(_module.__default.safeModuloInt((_dafny.ZERO).minus(_module.__default.fm17(_1116_v15, (_dafny.ZERO).minus(p1), new BigNumber((_1108_v9).length), p1, globalState)), p1));
        let _1117_v16;
        let _nw184 = Array((new BigNumber(5)).toNumber());
        _nw184[(_dafny.ZERO).toNumber()] = (_1113_v12).f12;
        _nw184[(_dafny.ONE).toNumber()] = p3;
        _nw184[(new BigNumber(2)).toNumber()] = ((_1107_v8) ? (p3) : ((_dafny.ZERO).minus(new BigNumber((_1115_v14).length))));
        _nw184[(new BigNumber(3)).toNumber()] = (p3).plus(p1);
        _nw184[(new BigNumber(4)).toNumber()] = p3;
        _1117_v16 = _nw184;
        let _init28 = function (_1118_i1) {
          return (_1118_i1).plus(_this.f5);
        };
        let _nw185 = Array((new BigNumber(21)).toNumber());
        for (let _i0_28 = 0; _i0_28 < new BigNumber(_nw185.length); _i0_28++) {
          _nw185[_i0_28] = _init28(new BigNumber(_i0_28));
        }
        _1117_v16 = _nw185;
      } else {
        let _1119___mcc_h5 = (_source17).cf32;
        let _1120_cf32 = _1119___mcc_h5;
        let _1121_v17;
        let _nw186 = new _module.C5();
        _nw186.__ctor(_module.__default.fm26(_this.f5, (_dafny.ZERO).minus((_dafny.ZERO).minus(((((_this).f4).contains(p2)) ? (((_this).f4).get(p2)) : (p1)))), p2, globalState));
        _1121_v17 = _nw186;
        let _1122_v18;
        _1122_v18 = _dafny.Map.Empty.slice().updateUnsafe(p3,(_1121_v17).f6);
        let _1123_v19;
        _1123_v19 = _dafny.Seq.of(p1, new BigNumber(-934), new BigNumber((_1122_v18).length));
        let _1124_v20;
        _1124_v20 = _dafny.Set.fromElements(_dafny.Seq.Concat(_dafny.Seq.of(_this.f5, p3, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(467)), function (_1125_i2) {
          return new BigNumber(228);
        })).length), p1), _1123_v19));
        _1124_v20 = _1124_v20;
        let _1126_v21;
        let _nw187 = Array((new BigNumber(29)).toNumber()).fill(_dafny.MultiSet.Empty);
        _1126_v21 = _nw187;
        let _1127_v22;
        _1127_v22 = _dafny.Seq.of(_1126_v21);
        let _1128_v23;
        let _nw188 = Array((new BigNumber(18)).toNumber());
        _nw188[(_dafny.ZERO).toNumber()] = _1126_v21;
        _nw188[(_dafny.ONE).toNumber()] = _1126_v21;
        _nw188[(new BigNumber(2)).toNumber()] = _1126_v21;
        _nw188[(new BigNumber(3)).toNumber()] = _1126_v21;
        _nw188[(new BigNumber(4)).toNumber()] = _1126_v21;
        _nw188[(new BigNumber(5)).toNumber()] = _1126_v21;
        _nw188[(new BigNumber(6)).toNumber()] = _1126_v21;
        _nw188[(new BigNumber(7)).toNumber()] = _1126_v21;
        _nw188[(new BigNumber(8)).toNumber()] = ((p2) ? (_1126_v21) : (_1126_v21));
        _nw188[(new BigNumber(9)).toNumber()] = (_1127_v22)[_module.__default.safeIndex(p3, new BigNumber((_1127_v22).length))];
        _nw188[(new BigNumber(10)).toNumber()] = _1126_v21;
        _nw188[(new BigNumber(11)).toNumber()] = _1126_v21;
        _nw188[(new BigNumber(12)).toNumber()] = _1126_v21;
        _nw188[(new BigNumber(13)).toNumber()] = (_1127_v22)[_module.__default.safeIndex(_this.f5, new BigNumber((_1127_v22).length))];
        _nw188[(new BigNumber(14)).toNumber()] = _1126_v21;
        _nw188[(new BigNumber(15)).toNumber()] = ((true) ? (_1126_v21) : (_1126_v21));
        _nw188[(new BigNumber(16)).toNumber()] = _1126_v21;
        _nw188[(new BigNumber(17)).toNumber()] = _1126_v21;
        _1128_v23 = _nw188;
        let _index153 = _module.__default.safeIndex(new BigNumber(460), new BigNumber((_1128_v23).length));
        (_1128_v23)[_index153] = _1126_v21;
        let _index154 = _module.__default.safeIndex(new BigNumber(460), new BigNumber((_1128_v23).length));
        (_1128_v23)[_index154] = _1126_v21;
        let _1129_v24;
        _1129_v24 = _dafny.Map.Empty.slice().updateUnsafe(p3,p2);
        if (((p2) || (p2)) === ((((_1129_v24).contains(_this.f5)) ? ((_1129_v24).get(_this.f5)) : (p2)))) {
          let _1130_v25;
          _1130_v25 = false;
          _1130_v25 = p2;
          let _1131_v26;
          let _nw189 = new _module.C0();
          _nw189.__ctor((_this.f5).minus(_this.f5));
          _1131_v26 = _nw189;
          let _1132_v27;
          _1132_v27 = _dafny.Seq.of(_module.__default.fm25(false, new BigNumber((_1123_v19).length), globalState));
          let _1133_v28;
          _1133_v28 = _dafny.Set.fromElements(p1);
          let _rhs145 = _module.__default.fm17(_1132_v27, (_1131_v26).f13, _module.__default.fm9(new BigNumber((_1133_v28).length), globalState), (_dafny.ZERO).minus(p1), globalState);
          let _rhs146 = _1131_v26;
          let _lhs87 = globalState;
          _lhs87.f2 = _rhs145;
          _1131_v26 = _rhs146;
          _1130_v25 = _1130_v25;
          let _1134_v29;
          _1134_v29 = _dafny.Map.Empty.slice().updateUnsafe(_1130_v25,_1130_v25);
          let _1135_v30;
          _1135_v30 = _dafny.Set.fromElements((((_1134_v29).contains(_1130_v25)) ? ((_1134_v29).get(_1130_v25)) : (_1130_v25)));
          _1130_v25 = !(_1135_v30).contains(p2);
          let _1136_v31;
          _1136_v31 = _dafny.Seq.UnicodeFromString("myiqun");
          let _1137_v32;
          _1137_v32 = _module.D6.create_DC11(_1136_v31);
          let _1138_v33;
          _1138_v33 = _dafny.Seq.of(_1137_v32);
          let _1139_v34;
          let _nw190 = Array((new BigNumber(20)).toNumber());
          _1139_v34 = _nw190;
          let _1140_v35;
          let _nw191 = Array((new BigNumber(28)).toNumber()).fill(_dafny.ZERO);
          _1140_v35 = _nw191;
          let _1141_v36;
          _1141_v36 = _dafny.MultiSet.fromElements(_1140_v35, _1140_v35);
          let _1142_v37;
          _1142_v37 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(_1138_v33,_1139_v34),_1141_v36);
          let _1143_v38;
          _1143_v38 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(_1137_v32),_1139_v34);
          _1142_v37 = (_1142_v37).update(_1143_v38, _1141_v36);
        } else {
          let _1144_v39;
          let _init29 = ((_1145_p1, _1146_p2) => function (_1147_i3) {
            return (_1147_i3).multipliedBy(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_1145_p1,_dafny.Seq.of(_1146_p2, _1146_p2))).length));
          })(p1, p2);
          let _nw192 = Array((new BigNumber(26)).toNumber());
          for (let _i0_29 = 0; _i0_29 < new BigNumber(_nw192.length); _i0_29++) {
            _nw192[_i0_29] = _init29(new BigNumber(_i0_29));
          }
          _1144_v39 = _nw192;
          let _index155 = _module.__default.safeIndex(new BigNumber(999), new BigNumber((_1144_v39).length));
          (_1144_v39)[_index155] = _this.f5;
          let _index156 = _module.__default.safeIndex(new BigNumber(999), new BigNumber((_1144_v39).length));
          (_1144_v39)[_index156] = _module.__default.safeDivisionInt(new BigNumber((_dafny.Seq.update(_dafny.Seq.of(p2, p2), _module.__default.safeIndex(p3, new BigNumber((_dafny.Seq.of(p2, p2)).length)), p2)).length), new BigNumber((_dafny.Seq.of(_1089_v0)).length));
          let _1148_v40;
          let _1149_v41;
          let _out41;
          let _out42;
          let _outcollector15 = (_this).m2(p2, globalState);
          _out41 = _outcollector15[0];
          _out42 = _outcollector15[1];
          _1148_v40 = _out41;
          _1149_v41 = _out42;
          let _1150_v42;
          _1150_v42 = _dafny.Seq.UnicodeFromString("yxohwdduo");
          let _1151_v43;
          _1151_v43 = _dafny.Map.Empty.slice().updateUnsafe((_1121_v17).f6,(_dafny.ZERO).minus(new BigNumber((_1150_v42).length)));
          _1151_v43 = (_1151_v43).update((_1121_v17).f6, new BigNumber(435));
          _1149_v41 = _module.__default.fm21((_1144_v39)[_module.__default.safeIndex(new BigNumber(999), new BigNumber((_1144_v39).length))], (_1121_v17).f6, _1149_v41, ((_1149_v41) ? (false) : (_1149_v41)), globalState);
          let _1152_v44;
          _1152_v44 = _dafny.Map.Empty.slice().updateUnsafe(p2,p1);
          _1152_v44 = (_1152_v44).update(!((_1144_v39)[_module.__default.safeIndex(new BigNumber(999), new BigNumber((_1144_v39).length))]).isEqualTo((_1144_v39)[_module.__default.safeIndex(new BigNumber(999), new BigNumber((_1144_v39).length))]), (_1148_v40).plus((_1144_v39)[_module.__default.safeIndex(new BigNumber(999), new BigNumber((_1144_v39).length))]));
        }
      }
      if (p2) {
        let _1153_v45;
        _1153_v45 = _dafny.Map.Empty.slice().updateUnsafe((p3).multipliedBy(p3),_1089_v0);
        let _1154_v47;
        _1154_v47 = _dafny.MultiSet.fromElements((_dafny.ZERO).minus(_module.__default.fm24(function () {
          let _coll36 = new _dafny.Set();
          for (const _compr_36 of (_dafny.MultiSet.fromElements(new BigNumber(-165), new BigNumber(396))).Elements) {
            let _1155_v46 = _compr_36;
            if ((_dafny.MultiSet.fromElements(new BigNumber(-165), new BigNumber(396))).contains(_1155_v46)) {
              _coll36.add((_1155_v46).multipliedBy(new BigNumber(-395)));
            }
          }
          return _coll36;
        }(), globalState)), p1);
        let _1156_v48;
        _1156_v48 = new _dafny.CodePoint('d'.codePointAt(0));
        let _1157_v49;
        _1157_v49 = _dafny.Seq.of(p3);
        let _1158_v50;
        _1158_v50 = _dafny.Seq.of(_1157_v49);
        let _1159_v51;
        _1159_v51 = _dafny.Seq.of(new BigNumber((_1158_v50).length), p3);
        let _1160_v52;
        _1160_v52 = _module.D4.create_DC9(new BigNumber((_1154_v47).cardinality()), _1156_v48, _1157_v49, _this.f5, p2);
        _1153_v45 = (_1153_v45).update((_1160_v52).dtor_cf19, _1089_v0);
        let _1161_v53;
        _1161_v53 = false;
        let _1162_v54;
        _1162_v54 = _module.D8.create_DC16(new BigNumber((_1159_v51).length), _1161_v53, false, _this.f5);
        let _rhs147 = (_1162_v54).dtor_cf30;
        let _rhs148 = p2;
        let _rhs149 = p3;
        let _rhs150 = p1;
        let _lhs88 = _this;
        _1161_v53 = _rhs147;
        _1161_v53 = _rhs148;
        r1 = _rhs149;
        _lhs88.f5 = _rhs150;
        let _1163_v55;
        _1163_v55 = _dafny.Seq.of(_1161_v53);
        let _1164_v56;
        _1164_v56 = _dafny.Map.Empty.slice().updateUnsafe(p2,_1163_v55);
        let _1165_v57;
        _1165_v57 = _dafny.Seq.UnicodeFromString("isekkdev");
        let _1166_v58;
        _1166_v58 = _dafny.MultiSet.fromElements(!(_1161_v53));
        let _1167_v59;
        _1167_v59 = _dafny.Seq.of(_1165_v57, _1165_v57, _dafny.Seq.update(_dafny.Seq.update(_1165_v57, _module.__default.safeIndex(_this.f5, new BigNumber((_1165_v57).length)), _1156_v48), _module.__default.safeIndex((_dafny.ZERO).minus(new BigNumber((_1166_v58).cardinality())), new BigNumber((_dafny.Seq.update(_1165_v57, _module.__default.safeIndex(_this.f5, new BigNumber((_1165_v57).length)), _1156_v48)).length)), _1156_v48));
        let _1168_v60;
        _1168_v60 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1154_v47).cardinality()),_1161_v53);
        let _1169_v61;
        _1169_v61 = _dafny.Seq.of(_1168_v60, _1168_v60);
        let _1170_v63;
        _1170_v63 = _module.D6.create_DC11(_dafny.Seq.UnicodeFromString("wwxjc"));
        let _1171_v65;
        let _nw193 = Array((new BigNumber(13)).toNumber()).fill(_dafny.Map.Empty);
        _1171_v65 = _nw193;
        let _1172_v66;
        let _nw194 = new _module.C4();
        _nw194.__ctor(_1161_v53, _1154_v47, _1171_v65);
        _1172_v66 = _nw194;
        let _1173_v67;
        _1173_v67 = _module.D9.create_DC19(_1172_v66, new BigNumber(80), (_dafny.ZERO).minus(_this.f5));
        let _1174_v68;
        _1174_v68 = _dafny.Map.Empty.slice().updateUnsafe((_1157_v49)[_module.__default.safeIndex(new BigNumber(395), new BigNumber((_1157_v49).length))],_1156_v48);
        let _1175_v69;
        let _nw195 = Array((new BigNumber(27)).toNumber());
        _nw195[(_dafny.ZERO).toNumber()] = p1;
        _nw195[(_dafny.ONE).toNumber()] = new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(p2,_1163_v55)).Merge(_1164_v56)).length);
        _nw195[(new BigNumber(2)).toNumber()] = (p3).plus(_this.f5);
        _nw195[(new BigNumber(3)).toNumber()] = p3;
        _nw195[(new BigNumber(4)).toNumber()] = new BigNumber((_1167_v59).length);
        _nw195[(new BigNumber(5)).toNumber()] = p1;
        _nw195[(new BigNumber(6)).toNumber()] = _module.__default.fm24(function () {
          let _coll37 = new _dafny.Set();
          for (const _compr_37 of ((_1169_v61)[_module.__default.safeIndex(p1, new BigNumber((_1169_v61).length))]).Keys.Elements) {
            let _1176_v62 = _compr_37;
            if (((_1169_v61)[_module.__default.safeIndex(p1, new BigNumber((_1169_v61).length))]).contains(_1176_v62)) {
              _coll37.add((_1176_v62).multipliedBy((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(new BigNumber(559), new BigNumber(505))).length))));
            }
          }
          return _coll37;
        }(), globalState);
        _nw195[(new BigNumber(7)).toNumber()] = (((_1154_v47).contains(p1)) ? ((_1154_v47).get(p1)) : ((_module.D2.create_DC4(p3)).dtor_cf6));
        _nw195[(new BigNumber(8)).toNumber()] = (new BigNumber((_1165_v57).length)).minus((_module.D7.create_DC14(p1)).dtor_cf26);
        _nw195[(new BigNumber(9)).toNumber()] = _module.__default.safeModuloInt(new BigNumber(((_1170_v63).dtor_cf22).length), (_1159_v51)[_module.__default.safeIndex(p1, new BigNumber((_1159_v51).length))]);
        _nw195[(new BigNumber(10)).toNumber()] = _module.__default.fm24(function () {
          let _coll38 = new _dafny.Set();
          for (const _compr_38 of _dafny.IntegerRange(new BigNumber(190), new BigNumber(666))) {
            let _1177_v64 = _compr_38;
            if (((new BigNumber(190)).isLessThanOrEqualTo(_1177_v64)) && ((_1177_v64).isLessThan(new BigNumber(666)))) {
              _coll38.add((_1177_v64).minus(new BigNumber((_1166_v58).cardinality())));
            }
          }
          return _coll38;
        }(), globalState);
        _nw195[(new BigNumber(11)).toNumber()] = _module.__default.safeDivisionInt(new BigNumber((_1159_v51).length), p1);
        _nw195[(new BigNumber(12)).toNumber()] = p1;
        _nw195[(new BigNumber(13)).toNumber()] = p3;
        _nw195[(new BigNumber(14)).toNumber()] = _this.f5;
        _nw195[(new BigNumber(15)).toNumber()] = _this.f5;
        _nw195[(new BigNumber(16)).toNumber()] = (p1).minus(_this.f5);
        _nw195[(new BigNumber(17)).toNumber()] = p3;
        _nw195[(new BigNumber(18)).toNumber()] = _this.f5;
        _nw195[(new BigNumber(19)).toNumber()] = (_1173_v67).dtor_cf35;
        _nw195[(new BigNumber(20)).toNumber()] = (p1).minus(_this.f5);
        _nw195[(new BigNumber(21)).toNumber()] = _this.f5;
        _nw195[(new BigNumber(22)).toNumber()] = _this.f5;
        _nw195[(new BigNumber(23)).toNumber()] = new BigNumber((_dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(225),new _dafny.CodePoint('l'.codePointAt(0))), _1174_v68, _1174_v68, _1174_v68)).length);
        _nw195[(new BigNumber(24)).toNumber()] = p1;
        _nw195[(new BigNumber(25)).toNumber()] = p3;
        _nw195[(new BigNumber(26)).toNumber()] = p3;
        _1175_v69 = _nw195;
        let _index157 = _module.__default.safeIndex(new BigNumber(963), new BigNumber((_1175_v69).length));
        (_1175_v69)[_index157] = p3;
        let _1178_v70;
        let _nw196 = Array((new BigNumber(22)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
        _1178_v70 = _nw196;
        let _index158 = _module.__default.safeIndex(new BigNumber(31), new BigNumber((_1089_v0).length));
        (_1089_v0)[_index158] = ((p2) ? (p2) : (_1161_v53));
        let _1179_v72;
        _1179_v72 = _dafny.Map.Empty.slice().updateUnsafe(_1165_v57,_1165_v57);
        let _index159 = _module.__default.safeIndex(new BigNumber(963), new BigNumber((_1175_v69).length));
        let _index160 = _module.__default.safeIndex(new BigNumber(31), new BigNumber((_1089_v0).length));
        let _rhs151 = _module.__default.fm17(_1167_v59, new BigNumber((function () {
          let _coll39 = new _dafny.Map();
          for (const _compr_39 of (_1179_v72).Keys.Elements) {
            let _1180_v71 = _compr_39;
            if ((_1179_v72).contains(_1180_v71)) {
              _coll39.push([_1180_v71,_1165_v57]);
            }
          }
          return _coll39;
        }()).length), p1, _this.f5, globalState);
        let _rhs152 = _1178_v70;
        let _rhs153 = _module.__default.fm21(((((_this).f4).contains((_1162_v54).dtor_cf30)) ? (((_this).f4).get((_1162_v54).dtor_cf30)) : (_this.f5)), _1156_v48, p2, p2, globalState);
        let _lhs89 = _1175_v69;
        let _lhs90 = _module.__default.safeIndex(new BigNumber(963), new BigNumber((_1175_v69).length));
        let _lhs91 = _1089_v0;
        let _lhs92 = _module.__default.safeIndex(new BigNumber(31), new BigNumber((_1089_v0).length));
        _lhs89[_lhs90] = _rhs151;
        _1178_v70 = _rhs152;
        _lhs91[_lhs92] = _rhs153;
        let _1181_v73;
        _1181_v73 = _dafny.MultiSet.fromElements(_1163_v55);
        let _1182_v74;
        _1182_v74 = _dafny.Map.Empty.slice().updateUnsafe(_1161_v53,p2);
        let _1183_v75;
        _1183_v75 = _1175_v69;
        let _1184_v76;
        _1184_v76 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm8((((_1181_v73).contains(_1163_v55)) ? ((_1181_v73).get(_1163_v55)) : (new BigNumber((_1182_v74).length))), globalState),_1183_v75);
        _1184_v76 = (_1184_v76).update(_1165_v57, _1183_v75);
        let _1185_v77;
        _1185_v77 = _dafny.Map.Empty.slice().updateUnsafe(p3,_this.f5);
        (globalState).f2 = _module.__default.safeDivisionInt(p3, ((p2) ? (new BigNumber((_1185_v77).length)) : (new BigNumber((_1165_v57).length))));
      } else {
        let _1186_v78;
        _1186_v78 = true;
        _1186_v78 = (true) || (p2);
        let _1187_v79;
        _1187_v79 = _dafny.Set.fromElements(_this.f5);
        if (!((_module.__default.fm38(new BigNumber(773), true, !(p2), _1186_v78, globalState)).Union(_1187_v79)).contains(p3)) {
          let _1188_v80;
          _1188_v80 = new _dafny.CodePoint('v'.codePointAt(0));
          _1188_v80 = _module.__default.fm26(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_1186_v78,_1186_v78)).length), new BigNumber((_1187_v79).length), p2, globalState);
          let _1189_v81;
          let _nw197 = new _module.C5();
          _nw197.__ctor(_1188_v80);
          _1189_v81 = _nw197;
          _1188_v80 = (_1189_v81).f6;
          r0 = p1;
          let _1190_v82;
          _1190_v82 = _dafny.Seq.UnicodeFromString("cvmfvuww");
          let _1191_v83;
          _1191_v83 = _dafny.Seq.of(p2, _1186_v78, _1186_v78);
          let _1192_v84;
          _1192_v84 = _dafny.Map.Empty.slice().updateUnsafe(_1190_v82,_1191_v83);
          let _1193_v85;
          _1193_v85 = _dafny.Map.Empty.slice().updateUnsafe(p1,_dafny.Seq.UnicodeFromString("oaxbduxb"));
          _1192_v84 = (_1192_v84).update(_dafny.Seq.Concat(_dafny.Seq.update((((_1193_v85).contains(new BigNumber(229))) ? ((_1193_v85).get(new BigNumber(229))) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(-136)), ((_1194_v81) => function (_1195_i4) {
            return (_1194_v81).f6;
          })(_1189_v81)))), _module.__default.safeIndex(_this.f5, new BigNumber(((((_1193_v85).contains(new BigNumber(229))) ? ((_1193_v85).get(new BigNumber(229))) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(-136)), ((_1196_v81) => function (_1197_i4) {
            return (_1196_v81).f6;
          })(_1189_v81))))).length)), _1188_v80), _1190_v82), _1191_v83);
        } else {
          let _1198_v86;
          _1198_v86 = _dafny.MultiSet.fromElements(p3);
          let _1199_v87;
          let _nw198 = new _module.C3();
          _nw198.__ctor((_1198_v86).Difference(_1198_v86), !(false));
          _1199_v87 = _nw198;
          let _1200_v88;
          _1200_v88 = _dafny.Seq.UnicodeFromString("tplrn");
          _1186_v78 = !(_dafny.Seq.IsProperPrefixOf(_dafny.Seq.UnicodeFromString("jjpsq"), _dafny.Seq.Concat(_1200_v88, _dafny.Seq.UnicodeFromString("diviu"))));
          let _1201_v89;
          _1201_v89 = _dafny.Map.Empty.slice().updateUnsafe(_1186_v78,_1186_v78);
          _1201_v89 = (_1201_v89).update(true, p2);
          let _1202_v90;
          let _init30 = ((_1203_p3) => function (_1204_i5) {
            return (_1204_i5).multipliedBy(_1203_p3);
          })(p3);
          let _nw199 = Array((new BigNumber(20)).toNumber());
          for (let _i0_30 = 0; _i0_30 < new BigNumber(_nw199.length); _i0_30++) {
            _nw199[_i0_30] = _init30(new BigNumber(_i0_30));
          }
          _1202_v90 = _nw199;
          let _index161 = _module.__default.safeIndex(new BigNumber(949), new BigNumber((_1202_v90).length));
          (_1202_v90)[_index161] = p1;
          let _1205_v91;
          _1205_v91 = new _dafny.CodePoint('h'.codePointAt(0));
          let _index162 = _module.__default.safeIndex(new BigNumber(949), new BigNumber((_1202_v90).length));
          (_1202_v90)[_index162] = (_1199_v87).fm19(_1205_v91, new BigNumber(272), globalState);
          (globalState).f2 = _module.__default.safeDivisionInt(_module.__default.safeDivisionInt(new BigNumber(74), (_dafny.ZERO).minus((((_1198_v86).contains(_this.f5)) ? ((_1198_v86).get(_this.f5)) : (_this.f5)))), p1);
        }
        (_this).f5 = _module.__default.fm24(_1187_v79, globalState);
        let _1206_v92;
        let _1207_v93;
        let _out43;
        let _out44;
        let _outcollector16 = (_this).m2(p2, globalState);
        _out43 = _outcollector16[0];
        _out44 = _outcollector16[1];
        _1206_v92 = _out43;
        _1207_v93 = _out44;
        let _1208_v94;
        _1208_v94 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-249),_1207_v93);
        _1208_v94 = (_1208_v94).update(new BigNumber((_dafny.Seq.UnicodeFromString("vpbxb")).length), p2);
      }
      (_this).f5 = (p1).multipliedBy(p3);
      let _1209_i6;
      _1209_i6 = _dafny.ZERO;
      L7: {
        while (p2) {
          C7: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1209_i6)) {
              break L7;
            }
            _1209_i6 = (_1209_i6).plus(_dafny.ONE);
            let _1210_v95;
            _1210_v95 = _dafny.MultiSet.fromElements(p3);
            let _1211_v96;
            _1211_v96 = _dafny.Seq.of(_1210_v95, _1210_v95, _1210_v95);
            (globalState).f2 = _module.__default.safeModuloInt((_dafny.ZERO).minus(_module.__default.safeModuloInt(p1, new BigNumber((_1211_v96).length))), _this.f5);
            let _1212_v97;
            let _nw200 = Array((new BigNumber(7)).toNumber());
            _1212_v97 = _nw200;
            _1212_v97 = _1212_v97;
            let _1213_v98;
            _1213_v98 = false;
            let _1214_v99;
            _1214_v99 = _dafny.Seq.UnicodeFromString("vm");
            let _1215_v100;
            _1215_v100 = _dafny.Map.Empty.slice().updateUnsafe(_this.f5,_dafny.Seq.IsProperPrefixOf(_1214_v99, _1214_v99));
            let _1216_v101;
            _1216_v101 = _dafny.Set.fromElements(_1089_v0);
            let _rhs154 = !((((_1215_v100).contains((new BigNumber((_1216_v101).length)).minus(_this.f5))) ? ((_1215_v100).get((new BigNumber((_1216_v101).length)).minus(_this.f5))) : (((p2) ? (p2) : (false)))));
            let _rhs155 = !(_1213_v98);
            _1213_v98 = _rhs154;
            _1213_v98 = _rhs155;
            (globalState).f2 = p3;
          }
        }
      }
      let _1217_v102;
      _1217_v102 = _dafny.Seq.UnicodeFromString("tp");
      let _1218_v103;
      _1218_v103 = _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('q'.codePointAt(0)),_dafny.Seq.Concat(_1217_v102, _1217_v102));
      let _1219_v104;
      _1219_v104 = new _dafny.CodePoint('b'.codePointAt(0));
      _1218_v103 = (_1218_v103).update(_1219_v104, _dafny.Seq.UnicodeFromString("eltjfofh"));
      let _1220_v105;
      _1220_v105 = _dafny.Set.fromElements(false);
      let _1221_v106;
      _1221_v106 = _dafny.Map.Empty.slice().updateUnsafe(p2,false);
      let _1222_v107;
      _1222_v107 = _dafny.Map.Empty.slice().updateUnsafe(p3,new BigNumber((_1221_v106).length));
      let _1223_v108;
      _1223_v108 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.Set.fromElements(p2, p2, p2)).IsProperSubsetOf(_1220_v105),new BigNumber(((_1222_v107).update(p1, _this.f5)).length));
      _1223_v108 = (_1223_v108).update(p2, _this.f5);
      r0 = new BigNumber(883);
      r1 = (_dafny.ZERO).minus(p1);
      let _1224_v109;
      _1224_v109 = _module.D2.create_DC5(p3, p2, false, p2, p2);
      r2 = ((!((_1224_v109).dtor_cf10)) ? (function () {
        let _coll40 = new _dafny.Set();
        for (const _compr_40 of (_module.__default.fm42(false, p3, p2, globalState)).Keys.Elements) {
          let _1225_v110 = _compr_40;
          if ((_module.__default.fm42(false, p3, p2, globalState)).contains(_1225_v110)) {
            _coll40.add(_1225_v110);
          }
        }
        return _coll40;
      }()) : (_module.__default.fm35(globalState)));
      return [r0, r1, r2];
    }
    get f4() {
      let _this = this;
      return _this._f4;
    };
  };

  $module.C7 = class C7 {
    constructor () {
      this._tname = "_module.C7";
      this._f3 = [];
    }
    _parentTraits() {
      return [_module.T0, _module.T1];
    }
    get f3() {
      let _this = this;
      return _this._f3;
    };
    __ctor(f3) {
      let _this = this;
      (_this)._f3 = f3;
      return;
    }
    fm0(p0, globalState) {
      let _this = this;
      return (_module.__default.safeDivisionInt(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!(false),new BigNumber((_dafny.Seq.of(new BigNumber(297))).length))).length), new BigNumber((_dafny.Seq.UnicodeFromString("fuwqkmw")).length))).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(679)), function (_1226_i0) {
        return new _dafny.CodePoint('j'.codePointAt(0));
      })).length));
    };
    fm1(p0, p1, p2, p3, globalState) {
      let _this = this;
      return new BigNumber(133);
    };
    m1(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = new _dafny.CodePoint('D'.codePointAt(0));
      let _1227_v0;
      _1227_v0 = new BigNumber(-681);
      r0 = ((p1) ? (_1227_v0) : (_1227_v0));
      if (p1) {
        let _1228_v1;
        _1228_v1 = _dafny.Seq.of(p1);
        let _1229_v2;
        _1229_v2 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.update(_1228_v1, _module.__default.safeIndex(_1227_v0, new BigNumber((_1228_v1).length)), p1),_dafny.Seq.update(p0, _module.__default.safeIndex(_1227_v0, new BigNumber((p0).length)), new BigNumber((p2).length)));
        let _1230_v3;
        _1230_v3 = new _dafny.CodePoint('t'.codePointAt(0));
        let _1231_v4;
        _1231_v4 = _dafny.Map.Empty.slice().updateUnsafe(p1,(_1228_v1)[_module.__default.safeIndex(new BigNumber(303), new BigNumber((_1228_v1).length))]);
        let _1232_v5;
        _1232_v5 = _dafny.Seq.of(_1231_v4, _1231_v4);
        let _1233_v6;
        _1233_v6 = _module.D1.create_DC1(_1232_v5);
        let _1234_v7;
        _1234_v7 = _dafny.MultiSet.fromElements(p1);
        let _1235_v8;
        let _nw201 = Array((new BigNumber(29)).toNumber());
        _nw201[(_dafny.ZERO).toNumber()] = (_1229_v2).contains(_1228_v1);
        _nw201[(_dafny.ONE).toNumber()] = p1;
        _nw201[(new BigNumber(2)).toNumber()] = p1;
        _nw201[(new BigNumber(3)).toNumber()] = (p1) || (false);
        _nw201[(new BigNumber(4)).toNumber()] = !(p1) || (p1);
        _nw201[(new BigNumber(5)).toNumber()] = p1;
        _nw201[(new BigNumber(6)).toNumber()] = p1;
        _nw201[(new BigNumber(7)).toNumber()] = (_1227_v0).isLessThan(_1227_v0);
        _nw201[(new BigNumber(8)).toNumber()] = p1;
        _nw201[(new BigNumber(9)).toNumber()] = p1;
        _nw201[(new BigNumber(10)).toNumber()] = !(p1) || (p1);
        _nw201[(new BigNumber(11)).toNumber()] = p1;
        _nw201[(new BigNumber(12)).toNumber()] = (_1227_v0).isLessThan(_1227_v0);
        _nw201[(new BigNumber(13)).toNumber()] = !(_module.__default.fm2(_1230_v3, (_1233_v6).dtor_cf1, p1, globalState));
        _nw201[(new BigNumber(14)).toNumber()] = p1;
        _nw201[(new BigNumber(15)).toNumber()] = (new BigNumber(-190)).isLessThanOrEqualTo(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_1227_v0,true)).length));
        _nw201[(new BigNumber(16)).toNumber()] = p1;
        _nw201[(new BigNumber(17)).toNumber()] = _module.__default.fm2(_1230_v3, _module.__default.fm3(globalState), false, globalState);
        _nw201[(new BigNumber(18)).toNumber()] = (_1227_v0).isLessThanOrEqualTo(new BigNumber(-359));
        _nw201[(new BigNumber(19)).toNumber()] = p1;
        _nw201[(new BigNumber(20)).toNumber()] = p1;
        _nw201[(new BigNumber(21)).toNumber()] = p1;
        _nw201[(new BigNumber(22)).toNumber()] = p1;
        _nw201[(new BigNumber(23)).toNumber()] = p1;
        _nw201[(new BigNumber(24)).toNumber()] = p1;
        _nw201[(new BigNumber(25)).toNumber()] = !((_1234_v7).contains(true));
        _nw201[(new BigNumber(26)).toNumber()] = p1;
        _nw201[(new BigNumber(27)).toNumber()] = p1;
        _nw201[(new BigNumber(28)).toNumber()] = p1;
        _1235_v8 = _nw201;
        let _index163 = _module.__default.safeIndex(new BigNumber(667), new BigNumber((_1235_v8).length));
        (_1235_v8)[_index163] = p1;
        let _index164 = _module.__default.safeIndex(new BigNumber(667), new BigNumber((_1235_v8).length));
        (_1235_v8)[_index164] = p1;
        let _1236_v9;
        let _init31 = ((_1237_v0, _1238_v1, _1239_v8, _1240_p1) => function (_1241_i0) {
          return (_dafny.Map.Empty.slice().updateUnsafe(_module.D1.create_DC2(new BigNumber(-156), false, _1237_v0),_1238_v1)).Merge(_dafny.Map.Empty.slice().updateUnsafe(_module.D1.create_DC2(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_1239_v8)[_module.__default.safeIndex(new BigNumber(667), new BigNumber((_1239_v8).length))],_1237_v0)).length), _1240_p1, _1237_v0),_1238_v1));
        })(_1227_v0, _1228_v1, _1235_v8, p1);
        let _nw202 = Array((new BigNumber(15)).toNumber());
        for (let _i0_31 = 0; _i0_31 < new BigNumber(_nw202.length); _i0_31++) {
          _nw202[_i0_31] = _init31(new BigNumber(_i0_31));
        }
        _1236_v9 = _nw202;
        let _1242_v10;
        _1242_v10 = _dafny.MultiSet.fromElements(_1227_v0);
        let _1243_v11;
        _1243_v11 = _module.D1.create_DC2(new BigNumber((_1242_v10).cardinality()), (_1235_v8)[_module.__default.safeIndex(new BigNumber(667), new BigNumber((_1235_v8).length))], _1227_v0);
        let _1244_v12;
        _1244_v12 = _dafny.Map.Empty.slice().updateUnsafe(_1243_v11,_1228_v1);
        let _index165 = _module.__default.safeIndex(new BigNumber(119), new BigNumber((_1236_v9).length));
        (_1236_v9)[_index165] = (_1244_v12).update(_1243_v11, _1228_v1);
        let _1245_v13;
        _1245_v13 = _dafny.Set.fromElements(p1);
        let _index166 = _module.__default.safeIndex(new BigNumber(119), new BigNumber((_1236_v9).length));
        (_1236_v9)[_index166] = ((_1244_v12).update(_module.D1.create_DC2(_1227_v0, (_1235_v8)[_module.__default.safeIndex(new BigNumber(667), new BigNumber((_1235_v8).length))], new BigNumber((_1245_v13).length)), _1228_v1)).Merge(_dafny.Map.Empty.slice().updateUnsafe(_1243_v11,_1228_v1));
        let _index167 = _module.__default.safeIndex(new BigNumber(667), new BigNumber((_1235_v8).length));
        (_1235_v8)[_index167] = (_1243_v11).dtor_cf3;
        _1242_v10 = _1242_v10;
        if (_module.__default.fm2(_1230_v3, _1232_v5, p1, globalState)) {
          _1233_v6 = _module.__default.fm4(p1, _1230_v3, _1227_v0, _module.__default.fm5(_1227_v0, globalState), globalState);
          let _1246_v14;
          let _init32 = ((_1247_p2) => function (_1248_i1) {
            return _1247_p2;
          })(p2);
          let _nw203 = Array((new BigNumber(29)).toNumber());
          for (let _i0_32 = 0; _i0_32 < new BigNumber(_nw203.length); _i0_32++) {
            _nw203[_i0_32] = _init32(new BigNumber(_i0_32));
          }
          _1246_v14 = _nw203;
          let _index168 = _module.__default.safeIndex(new BigNumber(515), new BigNumber((_1246_v14).length));
          (_1246_v14)[_index168] = p2;
          let _index169 = _module.__default.safeIndex(new BigNumber(515), new BigNumber((_1246_v14).length));
          (_1246_v14)[_index169] = p2;
          let _index170 = _module.__default.safeIndex(new BigNumber(667), new BigNumber((_1235_v8).length));
          (_1235_v8)[_index170] = (_1235_v8)[_module.__default.safeIndex(new BigNumber(667), new BigNumber((_1235_v8).length))];
          let _1249_v15;
          _1249_v15 = _dafny.Map.Empty.slice().updateUnsafe(p2,new _dafny.CodePoint('n'.codePointAt(0)));
          _1249_v15 = (_1249_v15).update(p2, _module.__default.fm5(new BigNumber((_1228_v1).length), globalState));
          r0 = _1227_v0;
        } else {
          let _index171 = _module.__default.safeIndex(new BigNumber(667), new BigNumber((_1235_v8).length));
          (_1235_v8)[_index171] = (_1235_v8)[_module.__default.safeIndex(new BigNumber(667), new BigNumber((_1235_v8).length))];
          let _1250_v16;
          let _init33 = ((_1251_v11) => function (_1252_i2) {
            return _1251_v11;
          })(_1243_v11);
          let _nw204 = Array((new BigNumber(10)).toNumber());
          for (let _i0_33 = 0; _i0_33 < new BigNumber(_nw204.length); _i0_33++) {
            _nw204[_i0_33] = _init33(new BigNumber(_i0_33));
          }
          _1250_v16 = _nw204;
          let _1253_v17;
          _1253_v17 = _dafny.Seq.of(_1235_v8, _1235_v8, _1235_v8, _1235_v8);
          let _1254_v18;
          _1254_v18 = _dafny.Map.Empty.slice().updateUnsafe(_1250_v16,_dafny.Seq.Concat(_1253_v17, _1253_v17));
          _1254_v18 = (_1254_v18).update(_1250_v16, _dafny.Seq.of(_1235_v8, _1235_v8));
          r0 = new BigNumber(597);
          let _1255_v19;
          _1255_v19 = _dafny.Seq.of(_1228_v1);
          let _1256_v20;
          _1256_v20 = _1228_v1;
          let _1257_v21;
          _1257_v21 = _dafny.Map.Empty.slice().updateUnsafe((_1243_v11).dtor_cf3,_dafny.Seq.Concat((_1255_v19)[_module.__default.safeIndex(_1227_v0, new BigNumber((_1255_v19).length))], _module.__default.fm6(new BigNumber((p2).length), (_this).fm1((_1228_v1)[_module.__default.safeIndex(new BigNumber(-124), new BigNumber((_1228_v1).length))], _1230_v3, _1256_v20, (_1235_v8)[_module.__default.safeIndex(new BigNumber(667), new BigNumber((_1235_v8).length))], globalState), _1230_v3, globalState)));
          _1257_v21 = (_1257_v21).update((_1235_v8)[_module.__default.safeIndex(new BigNumber(667), new BigNumber((_1235_v8).length))], _1228_v1);
          let _1258_v22;
          let _init34 = ((_1259_p2) => function (_1260_i3) {
            return (_1260_i3).plus(new BigNumber((_1259_p2).length));
          })(p2);
          let _nw205 = Array((new BigNumber(20)).toNumber());
          for (let _i0_34 = 0; _i0_34 < new BigNumber(_nw205.length); _i0_34++) {
            _nw205[_i0_34] = _init34(new BigNumber(_i0_34));
          }
          _1258_v22 = _nw205;
          let _1261_v23;
          _1261_v23 = _dafny.Seq.UnicodeFromString("bgdm");
          let _index172 = _module.__default.safeIndex(new BigNumber(667), new BigNumber((_1235_v8).length));
          let _rhs156 = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("khi"), _1261_v23)).length);
          let _rhs157 = _1258_v22;
          let _rhs158 = !((_module.__default.fm2(_1230_v3, _1232_v5, !(!(true)), globalState)) && ((_1235_v8)[_module.__default.safeIndex(new BigNumber(667), new BigNumber((_1235_v8).length))])) || ((_1227_v0).isEqualTo(_1227_v0));
          let _rhs159 = p2;
          let _lhs93 = globalState;
          let _lhs94 = _1235_v8;
          let _lhs95 = _module.__default.safeIndex(new BigNumber(667), new BigNumber((_1235_v8).length));
          _lhs93.f2 = _rhs156;
          _1258_v22 = _rhs157;
          _lhs94[_lhs95] = _rhs158;
          _1261_v23 = _rhs159;
        }
      } else {
        (globalState).f2 = _1227_v0;
        let _1262_v24;
        _1262_v24 = false;
        _1262_v24 = _1262_v24;
        let _1263_v25;
        let _nw206 = Array((new BigNumber(11)).toNumber()).fill(false);
        _1263_v25 = _nw206;
        let _1264_v26;
        _1264_v26 = new _dafny.CodePoint('i'.codePointAt(0));
        let _1265_v27;
        _1265_v27 = _dafny.Map.Empty.slice().updateUnsafe(_1262_v24,p1);
        let _1266_v28;
        _1266_v28 = _dafny.Seq.of(_1265_v27);
        let _index173 = _module.__default.safeIndex(new BigNumber(845), new BigNumber((_1263_v25).length));
        (_1263_v25)[_index173] = _module.__default.fm2(_1264_v26, _1266_v28, _1262_v24, globalState);
        let _index174 = _module.__default.safeIndex(new BigNumber(845), new BigNumber((_1263_v25).length));
        (_1263_v25)[_index174] = _1262_v24;
        _1265_v27 = (_1265_v27).Merge(_1265_v27);
        let _1267_v29;
        _1267_v29 = _dafny.MultiSet.fromElements(_1227_v0, _1227_v0, _1227_v0);
        let _1268_v30;
        _1268_v30 = _dafny.Map.Empty.slice().updateUnsafe(_1227_v0,_1227_v0);
        let _1269_v31;
        _1269_v31 = _dafny.MultiSet.fromElements(p1);
        let _1270_v32;
        let _nw207 = Array((new BigNumber(18)).toNumber());
        _nw207[(_dafny.ZERO).toNumber()] = new BigNumber((_1267_v29).cardinality());
        _nw207[(_dafny.ONE).toNumber()] = _1227_v0;
        _nw207[(new BigNumber(2)).toNumber()] = _1227_v0;
        _nw207[(new BigNumber(3)).toNumber()] = _1227_v0;
        _nw207[(new BigNumber(4)).toNumber()] = _1227_v0;
        _nw207[(new BigNumber(5)).toNumber()] = _1227_v0;
        _nw207[(new BigNumber(6)).toNumber()] = _1227_v0;
        _nw207[(new BigNumber(7)).toNumber()] = new BigNumber(297);
        _nw207[(new BigNumber(8)).toNumber()] = _1227_v0;
        _nw207[(new BigNumber(9)).toNumber()] = (_module.D1.create_DC2(new BigNumber(591), (_1263_v25)[_module.__default.safeIndex(new BigNumber(845), new BigNumber((_1263_v25).length))], _1227_v0)).dtor_cf4;
        _nw207[(new BigNumber(10)).toNumber()] = _1227_v0;
        _nw207[(new BigNumber(11)).toNumber()] = new BigNumber((_1268_v30).length);
        _nw207[(new BigNumber(12)).toNumber()] = _1227_v0;
        _nw207[(new BigNumber(13)).toNumber()] = _1227_v0;
        _nw207[(new BigNumber(14)).toNumber()] = _1227_v0;
        _nw207[(new BigNumber(15)).toNumber()] = new BigNumber((p2).length);
        _nw207[(new BigNumber(16)).toNumber()] = _1227_v0;
        _nw207[(new BigNumber(17)).toNumber()] = (((_1269_v31).contains(p1)) ? ((_1269_v31).get(p1)) : (_1227_v0));
        _1270_v32 = _nw207;
        let _1271_v33;
        _1271_v33 = _dafny.Map.Empty.slice().updateUnsafe((new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-718)), function (_1272_i4) {
          return new _dafny.CodePoint('l'.codePointAt(0));
        })).length)).isLessThanOrEqualTo(new BigNumber((_dafny.Set.fromElements(p1, (_1263_v25)[_module.__default.safeIndex(new BigNumber(845), new BigNumber((_1263_v25).length))])).length)),_1270_v32);
        _1271_v33 = (_1271_v33).Merge(_1271_v33);
      }
      let _1273_v34;
      let _nw208 = Array((new BigNumber(2)).toNumber());
      _nw208[(_dafny.ZERO).toNumber()] = p1;
      _nw208[(_dafny.ONE).toNumber()] = p1;
      _1273_v34 = _nw208;
      let _index175 = _module.__default.safeIndex(new BigNumber(649), new BigNumber((_1273_v34).length));
      (_1273_v34)[_index175] = !(false) || (p1);
      let _1274_v35;
      _1274_v35 = _dafny.Map.Empty.slice().updateUnsafe(p1,true);
      let _1275_v36;
      _1275_v36 = _dafny.Seq.of(_1274_v35);
      let _index176 = _module.__default.safeIndex(new BigNumber(649), new BigNumber((_1273_v34).length));
      (_1273_v34)[_index176] = !(!(p1)) || (_module.__default.fm2(_module.__default.fm5(_1227_v0, globalState), _1275_v36, p1, globalState));
      let _1276_v37;
      let _nw209 = new _module.C0();
      _nw209.__ctor(_1227_v0);
      _1276_v37 = _nw209;
      let _1277_v38;
      _1277_v38 = _dafny.MultiSet.fromElements(_1227_v0, new BigNumber((p2).length), (_1276_v37).f13);
      let _hi5 = _1227_v0;
      for (let _1278_i5 = _module.__default.safeDivisionInt(new BigNumber((_1277_v38).cardinality()), _1227_v0); _1278_i5.isLessThan(_hi5); _1278_i5 = _1278_i5.plus(_dafny.ONE)) {
        _1274_v35 = (_1274_v35).update(p1, true);
        if (p1) {
          let _1279_v39;
          _1279_v39 = _dafny.Seq.of(new _dafny.CodePoint('l'.codePointAt(0)));
          _1279_v39 = p2;
          let _1280_v40;
          _1280_v40 = _dafny.Seq.of(_1277_v38, _1277_v38, _1277_v38);
          let _1281_v41;
          _1281_v41 = _dafny.Set.fromElements((_1273_v34)[_module.__default.safeIndex(new BigNumber(649), new BigNumber((_1273_v34).length))]);
          let _1282_v42;
          _1282_v42 = _dafny.Seq.of((_1273_v34)[_module.__default.safeIndex(new BigNumber(649), new BigNumber((_1273_v34).length))]);
          let _1283_v43;
          let _nw210 = Array((new BigNumber(16)).toNumber());
          _nw210[(_dafny.ZERO).toNumber()] = _dafny.MultiSet.fromElements((_1276_v37).f13);
          _nw210[(_dafny.ONE).toNumber()] = _1277_v38;
          _nw210[(new BigNumber(2)).toNumber()] = _1277_v38;
          _nw210[(new BigNumber(3)).toNumber()] = (_1280_v40)[_module.__default.safeIndex(new BigNumber(384), new BigNumber((_1280_v40).length))];
          _nw210[(new BigNumber(4)).toNumber()] = _1277_v38;
          _nw210[(new BigNumber(5)).toNumber()] = _1277_v38;
          _nw210[(new BigNumber(6)).toNumber()] = _dafny.MultiSet.fromElements((_1276_v37).f13);
          _nw210[(new BigNumber(7)).toNumber()] = _dafny.MultiSet.fromElements((_1276_v37).f13, new BigNumber(907));
          _nw210[(new BigNumber(8)).toNumber()] = _1277_v38;
          _nw210[(new BigNumber(9)).toNumber()] = _dafny.MultiSet.fromElements(_1278_i5, (_1276_v37).f13);
          _nw210[(new BigNumber(10)).toNumber()] = _1277_v38;
          _nw210[(new BigNumber(11)).toNumber()] = _dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("ffaybi")).length), new BigNumber((_1281_v41).length), _1278_i5, (_dafny.ZERO).minus((_1276_v37).f13));
          _nw210[(new BigNumber(12)).toNumber()] = _1277_v38;
          _nw210[(new BigNumber(13)).toNumber()] = _dafny.MultiSet.fromElements(new BigNumber((_1282_v42).length));
          _nw210[(new BigNumber(14)).toNumber()] = _1277_v38;
          _nw210[(new BigNumber(15)).toNumber()] = _1277_v38;
          _1283_v43 = _nw210;
          let _1284_v44;
          _1284_v44 = _dafny.Map.Empty.slice().updateUnsafe(_1283_v43,_1227_v0);
          _1284_v44 = (_1284_v44).update(_1283_v43, (_1276_v37).f13);
          let _index177 = _module.__default.safeIndex(new BigNumber(649), new BigNumber((_1273_v34).length));
          (_1273_v34)[_index177] = !(p1);
          let _index178 = _module.__default.safeIndex(new BigNumber(649), new BigNumber((_1273_v34).length));
          (_1273_v34)[_index178] = false;
          (globalState).f2 = new BigNumber(804);
        } else {
          let _1285_v45;
          _1285_v45 = _dafny.MultiSet.fromElements(p1, (_1273_v34)[_module.__default.safeIndex(new BigNumber(649), new BigNumber((_1273_v34).length))]);
          let _1286_v46;
          _1286_v46 = _dafny.Seq.of((_1273_v34)[_module.__default.safeIndex(new BigNumber(649), new BigNumber((_1273_v34).length))], p1, p1);
          let _1287_v47;
          _1287_v47 = _dafny.MultiSet.fromElements((_1285_v45).Difference(_dafny.MultiSet.FromArray(_1286_v46)));
          _1287_v47 = (((p1) ? (_1287_v47) : (_1287_v47))).Union(_dafny.MultiSet.FromArray(_dafny.Seq.of(_dafny.MultiSet.fromElements(p1), _dafny.MultiSet.fromElements(p1, _module.__default.fm23(p2, (_1276_v37).f13, (_1276_v37).f13, false, globalState)), _1285_v45, _1285_v45)));
          let _1288_v48;
          _1288_v48 = _module.D1.create_DC2(new BigNumber((p2).length), p1, (_1276_v37).f13);
          let _index179 = _module.__default.safeIndex(new BigNumber(649), new BigNumber((_1273_v34).length));
          (_1273_v34)[_index179] = ((_1276_v37).f13).isLessThanOrEqualTo((_1288_v48).dtor_cf4);
          let _1289_v49;
          let _nw211 = Array((new BigNumber(24)).toNumber()).fill(_dafny.Map.Empty);
          _1289_v49 = _nw211;
          let _index180 = _module.__default.safeIndex(new BigNumber(303), new BigNumber((_1289_v49).length));
          (_1289_v49)[_index180] = _module.__default.fm43(_dafny.MultiSet.FromArray(p0), _1278_i5, _1278_i5, globalState);
          let _1290_v50;
          _1290_v50 = _dafny.Map.Empty.slice().updateUnsafe(p2,(_dafny.ZERO).minus((_1276_v37).f13));
          let _1291_v51;
          _1291_v51 = _dafny.Map.Empty.slice().updateUnsafe(_1285_v45,_dafny.MultiSet.fromElements((_1276_v37).f13));
          let _1292_v52;
          _1292_v52 = _dafny.Seq.of(_1290_v50);
          let _1293_v53;
          _1293_v53 = _dafny.Set.fromElements((_1273_v34)[_module.__default.safeIndex(new BigNumber(649), new BigNumber((_1273_v34).length))], p1);
          let _1294_v54;
          _1294_v54 = _dafny.Seq.of(_1290_v50, (_1292_v52)[_module.__default.safeIndex(new BigNumber((_1293_v53).length), new BigNumber((_1292_v52).length))]);
          let _1295_v56;
          _1295_v56 = _dafny.Map.Empty.slice().updateUnsafe(p2,(_1273_v34)[_module.__default.safeIndex(new BigNumber(649), new BigNumber((_1273_v34).length))]);
          let _index181 = _module.__default.safeIndex(new BigNumber(303), new BigNumber((_1289_v49).length));
          let _index182 = _module.__default.safeIndex(new BigNumber(649), new BigNumber((_1273_v34).length));
          let _rhs160 = (_1291_v51).Merge(_1291_v51);
          let _rhs161 = ((_1290_v50).Merge((_1294_v54)[_module.__default.safeIndex(_1278_i5, new BigNumber((_1294_v54).length))])).Merge(function () {
            let _coll41 = new _dafny.Map();
            for (const _compr_41 of (_1295_v56).Keys.Elements) {
              let _1296_v55 = _compr_41;
              if ((_1295_v56).contains(_1296_v55)) {
                _coll41.push([_1296_v55,(_1276_v37).f13]);
              }
            }
            return _coll41;
          }());
          let _rhs162 = (_1278_i5).isLessThan((_1276_v37).f13);
          let _lhs96 = _1289_v49;
          let _lhs97 = _module.__default.safeIndex(new BigNumber(303), new BigNumber((_1289_v49).length));
          let _lhs98 = _1273_v34;
          let _lhs99 = _module.__default.safeIndex(new BigNumber(649), new BigNumber((_1273_v34).length));
          _lhs96[_lhs97] = _rhs160;
          _1290_v50 = _rhs161;
          _lhs98[_lhs99] = _rhs162;
          _1227_v0 = (_dafny.ZERO).minus((_dafny.ZERO).minus(_module.__default.safeModuloInt((_1276_v37).f13, _module.__default.safeDivisionInt(_1278_i5, new BigNumber(803)))));
          let _1297_v57;
          _1297_v57 = _dafny.Map.Empty.slice().updateUnsafe(_1278_i5,p1);
          let _1298_v58;
          _1298_v58 = _dafny.Map.Empty.slice().updateUnsafe(_1297_v57,new BigNumber(-541));
          let _index183 = _module.__default.safeIndex(new BigNumber(649), new BigNumber((_1273_v34).length));
          (_1273_v34)[_index183] = !(_1298_v58).contains(_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_1278_i5),p1));
        }
        let _1299_v59;
        _1299_v59 = _module.D1.create_DC2((_1276_v37).f13, (_1273_v34)[_module.__default.safeIndex(new BigNumber(649), new BigNumber((_1273_v34).length))], _1278_i5);
        let _index184 = _module.__default.safeIndex(new BigNumber(649), new BigNumber((_1273_v34).length));
        (_1273_v34)[_index184] = (_1299_v59).dtor_cf3;
        let _1300_v60;
        _1300_v60 = _module.D2.create_DC4(new BigNumber((_1277_v38).cardinality()));
        _1300_v60 = _1300_v60;
      }
      if (p1) {
        let _1301_v61;
        _1301_v61 = _dafny.Seq.UnicodeFromString("hyoeph");
        _1301_v61 = _1301_v61;
        let _1302_v62;
        _1302_v62 = new _dafny.CodePoint('c'.codePointAt(0));
        r1 = _1302_v62;
        let _1303_v63;
        let _nw212 = Array((new BigNumber(8)).toNumber());
        _nw212[(_dafny.ZERO).toNumber()] = (_1276_v37).f13;
        _nw212[(_dafny.ONE).toNumber()] = (_1276_v37).f13;
        _nw212[(new BigNumber(2)).toNumber()] = _1227_v0;
        _nw212[(new BigNumber(3)).toNumber()] = (_1276_v37).f13;
        _nw212[(new BigNumber(4)).toNumber()] = new BigNumber(159);
        _nw212[(new BigNumber(5)).toNumber()] = (p0)[_module.__default.safeIndex(_1227_v0, new BigNumber((p0).length))];
        _nw212[(new BigNumber(6)).toNumber()] = (_1276_v37).f13;
        _nw212[(new BigNumber(7)).toNumber()] = new BigNumber((_dafny.MultiSet.FromArray(p2)).cardinality());
        _1303_v63 = _nw212;
        let _1304_v64;
        _1304_v64 = _1303_v63;
        let _1305_v65;
        _1305_v65 = _dafny.Map.Empty.slice().updateUnsafe((_1276_v37).f13,_1304_v64);
        let _1306_v66;
        _1306_v66 = _dafny.Set.fromElements(p1);
        let _1307_v67;
        _1307_v67 = _dafny.Map.Empty.slice().updateUnsafe(_1227_v0,(_1276_v37).f13);
        _1305_v65 = (_1305_v65).update(new BigNumber((_dafny.Seq.Concat(_module.__default.fm12(!(p1), _1227_v0, _1306_v66, _1307_v67, globalState), p0)).length), _1304_v64);
        let _1308_v68;
        let _nw213 = new _module.C2();
        _nw213.__ctor((false) && ((_1273_v34)[_module.__default.safeIndex(new BigNumber(649), new BigNumber((_1273_v34).length))]), _module.__default.safeModuloInt(_1227_v0, (_1276_v37).f13), (_this).f3);
        _1308_v68 = _nw213;
        let _1309_v69;
        _1309_v69 = _dafny.Map.Empty.slice().updateUnsafe(p1,_1308_v68);
        _1308_v68 = (((_1309_v69).contains((_1308_v68).f11)) ? ((_1309_v69).get((_1308_v68).f11)) : (_1308_v68));
        let _1310_v70;
        _1310_v70 = _module.D4.create_DC8(_1276_v37);
        let _pat_let_tv25 = _1276_v37;
        _1310_v70 = function (_pat_let13_0) {
          return function (_1311_dt__update__tmp_h0) {
            return function (_pat_let14_0) {
              return function (_1312_dt__update_hcf15_h0) {
                return _module.D4.create_DC8(_1312_dt__update_hcf15_h0);
              }(_pat_let14_0);
            }(_pat_let_tv25);
          }(_pat_let13_0);
        }(_1310_v70);
      } else {
        let _1313_v71;
        let _nw214 = new _module.C2();
        _nw214.__ctor(p1, (_1276_v37).f13, (_this).f3);
        _1313_v71 = _nw214;
        let _1314_v72;
        let _nw215 = Array((new BigNumber(2)).toNumber());
        _nw215[(_dafny.ZERO).toNumber()] = _1313_v71;
        _nw215[(_dafny.ONE).toNumber()] = _1313_v71;
        _1314_v72 = _nw215;
        let _index185 = _module.__default.safeIndex(new BigNumber(721), new BigNumber((_1314_v72).length));
        (_1314_v72)[_index185] = _1313_v71;
        let _index186 = _module.__default.safeIndex(new BigNumber(721), new BigNumber((_1314_v72).length));
        (_1314_v72)[_index186] = _1313_v71;
        let _1315_v73;
        _1315_v73 = _dafny.MultiSet.fromElements((_1273_v34)[_module.__default.safeIndex(new BigNumber(649), new BigNumber((_1273_v34).length))], (_1273_v34)[_module.__default.safeIndex(new BigNumber(649), new BigNumber((_1273_v34).length))]);
        let _1316_v74;
        _1316_v74 = _dafny.Seq.of((((_1315_v73).contains((_1273_v34)[_module.__default.safeIndex(new BigNumber(649), new BigNumber((_1273_v34).length))])) ? ((_1315_v73).get((_1273_v34)[_module.__default.safeIndex(new BigNumber(649), new BigNumber((_1273_v34).length))])) : ((_1276_v37).f13)), (_dafny.ZERO).minus((_1276_v37).f13), _1227_v0, (_1276_v37).f13, (_1276_v37).f13);
        _1316_v74 = _1316_v74;
        let _1317_v75;
        let _nw216 = Array((new BigNumber(8)).toNumber()).fill(_dafny.Map.Empty);
        _1317_v75 = _nw216;
        let _nw217 = Array((new BigNumber(20)).toNumber()).fill(_dafny.Map.Empty);
        _1317_v75 = _nw217;
        _1315_v73 = _1315_v73;
        let _1318_v76;
        _1318_v76 = _dafny.Map.Empty.slice().updateUnsafe((_1276_v37).f13,(_1276_v37).f13);
        let _1319_v77;
        _1319_v77 = _module.D1.create_DC2((((_1318_v76).contains((_1276_v37).f13)) ? ((_1318_v76).get((_1276_v37).f13)) : (_1227_v0)), _dafny.Seq.IsPrefixOf(p2, p2), _1227_v0);
        _1319_v77 = _1319_v77;
      }
      r0 = _1227_v0;
      let _1320_v78;
      _1320_v78 = new _dafny.CodePoint('n'.codePointAt(0));
      r1 = _1320_v78;
      return [r0, r1];
    }
  };
  return $module;
})(); // end of module _module
_dafny.HandleHaltExceptions(() => _module.__default.Main(_dafny.UnicodeFromMainArguments(require('process').argv)));
